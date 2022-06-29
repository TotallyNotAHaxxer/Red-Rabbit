#!/usr/bin/perl
#
# testserver.pl - copyright 2006 rain forest puppy
#
# This is a specialized web server emulator used to feed test
# cases to a web client
#

use IO::Socket;
use IO::Select;
$|++;

my $testcasedir = shift || '.';
chdir($testcasedir);

my $insocket = IO::Socket::INET->new(LocalPort => 8088, LocalHost => '127.0.0.1',
					Listen => 20, Proto  => 'tcp', Reuse  => 1);
die $! unless $insocket;

my $readers = IO::Select->new() or die $!;
$readers->add($insocket);

%clients = (); # headersend, content-length, data, # reqs on this socket
%crawler = (); # requests made to crawler handler

while(1){
	my @ready = $readers->can_read;
	foreach my $handle (@ready){
		my $fno = fileno($handle);
		if($fno == fileno($insocket)){
			my $newclient = $handle->accept();
			$clients{ fileno($newclient) } = [-1,-1,'',0];
			$newclient->autoflush(1);
			$readers->add($newclient);
			next;
		}

		my $data;
		my $res = sysread($handle, $data, 2048);
		if(!defined $res || $res==0){
			display_error("Client closed socket before request")
				if($clients{$fno}->[3] == 0);
			$readers->remove($fno);
			eval { $handle->close; };
			delete $clients{$fno};
			next;
		}

		$clients{$fno}->[2] .= $data;
	
		if(precheck_request($handle)){
			handle_request($handle);
		}
	}
}


# figure out if we have enough data to process the request
sub precheck_request {
	my $handle = shift;
	my $fno = fileno($handle);

	return 0 if(!defined $clients{$fno});

	# first, see if we have any data
	return 0 if(length($clients{$fno}->[2])==0);

	# find the end of the headers
	my $headersend = -1;
	if( $clients{$fno}->[0] == -1){
		if( $clients{$fno}->[2] =~ /(\r{0,1}\n\r{0,1}\n)/ ){
			my $endtoken = $1;
			$headersend = index( $clients{$fno}->[2], $endtoken );
			if($headersend == -1){ # weird...shutdown to prevent looping
				shutdown_request($handle);
				return 0;
			}
			$headersend += length($endtoken);
		} else {
			return 0;
		}
		$clients{$fno}->[0] = $headersend;
	}

	# do we have content data (dictated by content-length)
	my $contentlen = 0;
	if($clients{$fno}->[1] == -1){
		if( substr( $clients{$fno}->[2], 0, $headersend ) =~
				/\ncontent-length:[ \t]*(.+?)\r{0,1}\n/i){
			$contentlen = $1;
		} else {
			# no valid content-length, and we have all headers, so
			# we should be good to go
			$clients{$fno}->[1] = 0;
			return 1;
		}
		$clients{$fno}->[1] = $contentlen;
	}

	# do we have enough data to satisfy content length?
	if( (length( $clients{$fno}->[2] ) - $clients{$fno}->[0]) >= 
			$clients{$fno}->[1] ){
		return 1;
	}

	# we're not quite there yet
	return 0;
}


# process the request (assumes enough data based on precheck_request)
sub handle_request {
	my $handle = shift;
	my $fno = fileno($handle);

	if(!defined $clients{$fno}){
		display_error("Integrity error; client doesn't exist in state table");
		shutdown_request($handle);
		return;
	}
	
	$clients{$fno}->[3]++;

	my ($requestline, $url);

	if($clients{$fno}->[2]=~m/^([a-z0-9]+[ \t]+([^ \t\r\n]+)(.*?)\n)/i ){
		$requestline = $1;
		$url = $2;
	} else {
		display_error("Bad HTTP request format", $clients{$fno}->[2]);
		shutdown_request($handle);
		return;
	}

	if( $url=~m#^/AUTH/(.+?)/# ){
		my $answer = '';
		my $a = $1;
		$a =~ tr/+/ /;
		if($clients{$fno}->[2] !~ /Authorization: $a/){
			$answer = <<EOT;
HTTP/1.0 401 Authentication Required
WWW-Authenticate: Basic
Connection: close
Content-Length: 0

EOT
		}
		else
		{
			$answer = <<EOT;
HTTP/1.0 200 OK
Connection: close

You successfully authenticated
EOT
		}

		syswrite($handle, $answer, length($answer));
		shutdown_request($handle);
		return;
	}


	if( $url=~m#^/TESTCASE/(.+?)/# ){
		run_testcase($handle, $1);
		return;
	}
	
	if( $url=~m#^/CUSTOM/(.+)$# ){
		run_custom($handle, $1);
		return;
	}

	if( $url=~m#^/CRAWLERRESET/# ){
		%crawler = ();
		my $answer = <<EOT;
HTTP/1.0 200 OK
Status: crawler reset
Connection: close

Crawler was reset.
EOT
		syswrite($handle, $answer, length($answer));
		shutdown_request($handle);
		return;
	}

	if( $url=~m#^/CRAWLERSTART/# ){
		$crawler{ $url }++;
		my $answer = <<EOT;
HTTP/1.0 200 OK
Status: crawler start
Connection: close

EOT
		my $answer2 = <<EOT;
<a href="/CRAWLER/A/">A</a>
<a href="/CRAWLER/B/">B</a>
<a href="/CRAWLER/C/">C</a>
<a href="/CRAWLER/D/">D</a>
<a href="/CRAWLER/E/">E</a>
EOT
		$answer .= $answer2 if($requestline!~/^HEAD /);
		syswrite($handle, $answer, length($answer));
		shutdown_request($handle);
		return;
	}

	if( $url=~m#^/CRAWLER/([0-9]+)/# ){
		$crawler{ $url }++;
		my $x = $1 +1;
		my $answer = <<EOT;
HTTP/1.0 200 OK
Status: crawler continuation
Connection: close

EOT
		my $answer2 = <<EOT;
<a href="/CRAWLER/$x/">$x</a>
EOT
		$answer .= $answer2 if($requestline!~/^HEAD /);
		syswrite($handle, $answer, length($answer));
		shutdown_request($handle);
		return;
	}
	
	if( $url=~m#^/CRAWLER/([A-Z]+)/# ){
		$crawler{ $url }++;
		my $x = $1;
		my $answer = <<EOT;
HTTP/1.0 200 OK
Status: crawler continuation
Connection: close

EOT
		my $answer2 = '';
		if(length($x) >= 4){
			$answer2 .= 'No more links';
		} else {
			$x .= $x;
			$x2 = substr($x,0,1);
			$answer2 .= "<a href=\"/CRAWLER/$x/\">$x</a>\n";
			$answer2 .= "<a href=\"/CRAWLER/$x2/\">$x2</a>\n";
			$answer2 .= "<a href=\"/CRAWLER/$x/\">$x</a>\n";
		}
		$answer .= $answer2 if($requestline !~ /^HEAD /);
		syswrite($handle, $answer, length($answer));
		shutdown_request($handle);
		return;
	}

	if( $url=~m#^/CRAWLERRESULT/# ){
		my $answer = <<EOT;
HTTP/1.0 200 OK
Status: crawler result
Connection: close

EOT
		foreach(sort keys %crawler){
			$answer .= $_ . ': ' . $crawler{$_} . "\n";
		}
		syswrite($handle, $answer, length($answer));
		shutdown_request($handle);
		return;
	}

	if( $url=~m#^/SHUTDOWN/# ){
		exit;
	}
	
	# if we get here, it's not a testcase or custom test; so error out	
	display_error("Non-testcase request", $clients{$fno}->[2]);
	my $answer = <<EOT;
HTTP/1.0 403 Forbidden
Status: non-testcase request
Connection: close

You requested a forbidden resource.
EOT
	
	syswrite($handle, $answer, length($answer));
	shutdown_request($handle);
}

sub run_testcase {
	my $handle = shift;
	my $case = shift;

	$case=~tr/-_a-z0-9A-Z//cd;
	
	if(!-e $case.'.case'){
		my $answer = <<EOT;
HTTP/1.0 404 Test case not found
Status: missing test case
Connection: close

The specified test case was not found.
EOT
		syswrite($handle, $answer, length($answer));
		shutdown_request($handle);
		return;
	}

	if(!open(IN,"<$case".'.case')){
		my $answer = <<EOT;
HTTP/1.0 500 Problem opening test case
Status: test case error
Connection: close

There was a problem opening the test case.
EOT
		syswrite($handle, $answer, length($answer));
		shutdown_request($handle);
		return;
	}

	binmode(IN); # Stupid Windows
	my $target = new_default_attribs();

	my $headers = '';
	my $data = '';

	# first line might be custom attributes
	my $attribs = <IN>;
	if($attribs !~ m#^HTTP/#){
		parse_custom_attribs($attribs, $target);
	} else {
		$headers .= $attribs;
	}

	my $inheaders=1;
	while(<IN>){
		if($inheaders) {
			$headers .= $_;
		} else {
			$data .= $_;
		}
		if($inheaders && ($_ eq "\x0a" || $_ eq "\x0d\x0a")){
			$inheaders--;
		}
	}
	close(IN);

	_run_ex($handle, $target, $headers, $data);
}

sub _run_ex {
	my $handle = shift;
	my $target = shift;
	my $headers = shift;
	my $data = shift;

	if($target->{stallinitial}>0){
		sleep($target->{stallinitial});
	}	

	if($target->{'100continue'} > 0){
		my $msg = <<EOT;
HTTP/1.0 100 Continue
Content-length: 0
Status: still working...

EOT
		for(1..$target->{'100continue'}){
			if(!syswrite($handle, $msg, length($msg))){
				display_error("Problem sending 100 Continue messages");
				shutdown_request($handle);
				return;
			}
			sleep(1);
		}
	}
	
	if(!syswrite($handle, $headers, length($headers))){
		display_error("Problem sending HTTP response headers");
		shutdown_request($handle);
		return;
	}

	if($target->{stallmidstream}>0){
		sleep($target->{stallmidstream});
	}

	if(!syswrite($handle, $data, length($data))){
		display_error("Problem sending response data");
		shutdown_request($handle);
		return;
	}

	if($target->{stallend}>0){
		sleep($target->{stallend});
	}

	if($target->{keepalive}==1){
		if($target->{forceclose}==1){
			shutdown_request($handle);
		} else {
			reset_request($handle);
		}
	} else {
		if($target->{forceopen}==1){
			reset_request($handle);
		} else {
			shutdown_request($handle);
		}
	}
}

sub run_custom {
	my $handle = shift;
	my $attribs = shift;

	my $target = new_default_attribs();
	parse_custom_attribs($attribs, $target);

	my $headers = "HTTP/" . $target->{httpversion} .
		" " . $target->{httpresponse} . " Response" .
		$target->{eol};
	
	if($target->{keepalive}==1){
		$headers .= "Connection: keep-alive" . $target->{eol};
	} else {
		$headers .= "Connection: close" . $target->{eol};
	}

	if($target->{chunked}==1){
		$headers .= "Transfer-encoding: chunked" . $target->{eol};
	} else {
		$headers .= "Content-length: " . $target->{contentlength} .
			$target->{eol};
	}

	$headers .= $target->{eol};

	my $data = $target->{content};

	if($target->{contentusepost}==1 || $target->{contentreflect}==1){
		$data = '';
		my $fno = fileno($handle);
		if($target->{contentreflect}==1){
			$data .= substr($clients{$fno}->[2], 0,
				$clients{$fno}->[0]);
		}
		$data .= substr( $clients{$fno}->[2],
				$clients{$fno}->[0], $clients{$fno}->[1] );
	}
	
	if($target->{chunked}==1){
		$data = sprintf("%02x",length($data)) . $target->{eol}
			. $data . $target->{eol} . '0' .
			$target->{eol};
	}	

	_run_ex($handle, $target, $headers, $data);

}

sub parse_custom_attribs {
	my $attribstr = shift;
	my $attribhr = shift;

	my @attr = split(/\//, $attribstr);
	foreach (@attr){
		next if($_ eq '');
		my ($k,$v) = split(/=/, $_, 2);
		if($v=~tr/%//){
			$v =~ s/%([a-f0-9][a-f0-9])/pack("C",hex($1))/eig;
		}
		if(defined $attribhr->{$k}){
			$attribhr->{$k} = $v; 
		}
	}
}

sub new_default_attribs {
	my %attribs = (

		httpversion => '1.0',	# http version to use
		httpresponse => 200,	# http response code
		eol => "\x0d\x0a",	# EOL
		keepalive => 0,		# use keepalives?
		chunked => 0,		# use chunked output?
		content => 'abcde',	# content to return
		contentusepost => 0,	# get content from posted data
		contentreflect => 0,	# get content from full request
		contentlength => 5,	# value of content-length header
		forceclose => 0,	# close the connection no matter what
		forceopen => 0,		# keep the connection open no matter what
		'100continue' => 0,	# number of 100 Continue responses to inject
		stallinitial => 0,	# wait specified number of seconds before responding
		stallmidstream => 0,	# wait specified number of seconds during response
		stallend => 0		# wait specified number of seconds after response 
					# (but before connection close)
	
	);
	
	return \%attribs;
}

sub reset_request {
	my $handle = shift;
	my $fno = fileno($handle);
	
	substr($clients{$fno}->[2], 0, $clients{$fno}->[0], '');
	$clients{$fno}->[0] = -1;
	if($clients{$fno}->[1] > 0){
		substr($clients{$fno}->[2], 0, $clients{$fno}->[1], '');
	}
	$clients{$fno}->[1] = -1;
	
	# we're supposed to be lenient about trailing CRLFs...
	if(length($clients{$fno}->[2])>0){
		$clients{$fno}->[2] =~ s/^[\r\n]+//;
	}
}

sub shutdown_request {
	my $handle = shift;
	my $fno = fileno($handle);
	
#	sleep(1);
	eval { $handle->close; };
	$readers->remove($fno);
	delete $clients{$fno};
}

sub display_error {
	print STDERR '-' x 76, "\n";
	print STDERR "Error: ", shift, "\n";
	if( defined $_[0] ){
		print STDERR shift, "\n";
	}
}

