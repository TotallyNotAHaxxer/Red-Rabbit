#!/usr/bin/perl
$|++;

require '../LW2.pm';

%CONFIG = (
	proxy_host => '',
	proxy_auth => 0,
	proxy_user => '',
	proxy_pass => ''
);

@NORMALHOSTS = (
	'http://www.apache.org/',	# latest apache
	'http://www.microsoft.com/'	# IIS
);

@SSLHOSTS = (
	'https://www.microsoft.com/',	# IIS SSL
	'https://www.redhat.com/',	# Apache SSL
	'https://www.sun.com'		# Sun java web server SSL
);

@ERRORHOSTS = (
	'http://www.wiretrip.net:81/', 	# closed HTTP port
	'http://www.google.com:82/', 	# filtered HTTP port
	'http://non-exist.wiretrip.net'	# non-exist host
);

@ERRORHOSTSSSL = (
	'https://www.wiretrip.net:81/',  # closed HTTPS port
	'https://www.google.com:82/',    # filtered HTTPS port
	'https://non-exist.wiretrip.net' # non-exist host
);


print STDOUT <<EOT;




Welcome to the Libwhisker2 connection tester!

The goal of this script is to make HTTP and HTTPS connections to various
(public) web servers, in an effort to assess how your particular
platform performs during the connection attempts.  After the connection
testing is done, you will be presented with a collection of data and
given the chance to auto-submit the data to the Libwhisker maintainers.
Nothing will be sent without your express consent and review.

Note: in order to use this script, you will need the ability to make web
requests to the internet (via proxy is ok).

Let's get started!

------------------------------------------------------------------------

EOT


print STDOUT "Do you need to use a proxy server to access the internet?\n";
$res = yesno();

if($res eq 'y'){
	print STDOUT <<EOT;

Please enter in the address of your proxy server, in the form of
server:port.  So, for example, if you need to use the proxy server at
DNS address proxy.mynetwork.com on port 8080, you should enter
proxy.mynetwork.com:8080.  If you need to use the proxy server at IP
address 192.168.2.14 on port 8000, then enter in 192.168.2.14:8000.  Both
an address and a port number are required.

EOT

	$temp = '';
	do {
		$temp = getinfo('proxy server address');
		if($temp!~/^[^:]+:[0-9]+$/){
			print STDOUT "Bad proxy address format [expecting server:port]\n";
			$temp = '';
		}
	} while($temp eq '');

	$CONFIG{proxy_host} = $temp;
		
	print STDOUT "\nDo you need to authenticate to your proxy server ?\n";
	$res = yesno();

	if($res eq 'y'){
		print STDOUT <<EOT;

Please select the proxy authentication type:

0 - None
1 - Basic
2 - NTLM
3 - Other

EOT
		$res = 0;
		do {
			$res = getinfo('auth type');
		} while($res !~ m/^[0-3]$/);

		if($res == 3){
			print STDOUT <<EOT;

I'm sorry, but Libwhisker2 currently does not support proxy auth types
other than Basic or NTLM.  Since you specified you need something else,
we will be unable to continue.  Please submit a feature addition request
to add the type of proxy authentication you need.

Goodbye.
EOT
			exit;
		}

		$CONFIG{proxy_auth} = $res;
		if($res > 0){
			print STDOUT "\nPlease enter in proxy username\n";
			$CONFIG{proxy_user} = getinfo('proxy username');
			print STDOUT "\nPlease enter in proxy password\n";
			$CONFIG{proxy_pass} = getinfo('proxy password');
		}
	}
}

print STDOUT <<EOT;

------------------------------------------------------------------------

This connection script will make requests against a configured set of
web sites.  You also have the option of making a request to a
user-specified web site as well, in order to test how libwhisker may
interact with that particular server & server configuration.

Would you like to specify an additional user-supplied web site?

EOT

$res = yesno();

if($res eq 'y'){

	print STDOUT <<EOT;


Please enter in your target test server, in the format of:

	http[s]://address[:port]/

For example, to make access www.mycompany.com via HTTPS/SSL, you would
enter https://www.mycompany.com/

To access web01.mycompany.com via normal HTTP on port 8080, you would
enter http://web01.mycompany.com:8080/

Ultimately you should enter in the same full URL that you would enter
into a web browser in order to view the site.

EOT

if(!LW2::ssl_is_available()){
	print STDOUT <<EOT;
NOTICE!!  The Perl SSL modules are not installed, so https URLs will be ignored

EOT
}

	$temp = '';
	do {
		$temp = getinfo('web server address');
		if($temp !~ m#^http[s]*://.+#){
			print STDOUT "Error: address not in expected format\n";
			$temp = '';
		}
	} while ($temp eq '');

	if($temp =~ m#^https://#){
		push @SSLHOSTS, $temp;
	} else {
		push @NORMALHOSTS, $temp;
	}
}

###################################
# check ssl library options

@x = LW2::ssl_is_available();
if($x[0] == 1 && $x[1] eq 'Net::SSLeay'){
	eval "use Net::SSL";
	if ( !$@ ){
		# Net::SSL is also available; use it?
		print STDOUT <<EOT;

------------------------------------------------------------------------

Libwhisker uses Net::SSLeay by default; however, it appears your system
also has Net::SSL installed too.  We suggest you run this script twice,
once with each SSL module (i.e. answer 'no' the first time and 'yes'
the second time, below).

Would you like to use Net::SSL instead of Net::SSLeay?

EOT
		$res = yesno();
		if($res eq 'y'){
			$LW2::LW_SSL_LIB = 2;
			$LW2::_SSL_LIBRARY = 'Net::SSL';
		}
	}
}


# at this point, we should have all the data...so let's recap

print STDOUT <<EOT;

------------------------------------------------------------------------

Current configuration:

EOT

if($CONFIG{proxy_host} eq ''){
	print STDOUT "\tUse proxy: No\n";
} else {
	print STDOUT "\tUse proxy: Yes\n";
	if($CONFIG{proxy_auth} > 0){
		print STDOUT "\tProxy auth type: Basic\n"
			if($CONFIG{proxy_auth} == 1);
		print STDOUT "\tProxy auth type: NTLM\n"
			if($CONFIG{proxy_auth} == 2);
		print STDOUT "\tProxy username: ", $CONFIG{proxy_user}, "\n";
		print STDOUT "\tProxy password: <not displayed>\n";
	}
}

print STDOUT "\nTarget web sites:\n";

@HOSTS = ();
push @HOSTS, @NORMALHOSTS;
push @HOSTS, @SSLHOSTS	if(LW2::ssl_is_available());
push @HOSTS, @ERRORHOSTS;
push @HOSTS, @ERRORHOSTSSSL  if(LW2::ssl_is_available());

foreach(@HOSTS){
	print "\t", $_, "\n";
}

print STDOUT "\n\nDoes this look correct?\n";
$res = yesno();
if($res eq 'n'){
	print STDOUT <<EOT;

You indicated something didn't look correct.  Please start the script over
to select the correct values.

Goodbye.

EOT
	exit;
}

# if we get here, we're good to go!
@DATA = ();

print STDOUT <<EOT;

------------------------------------------------------------------------

Starting connection test.  This may take several minutes; please be
patient.

EOT

@T = ();
push @T, @NORMALHOSTS;
push @T, @ERRORHOSTS;
foreach $host (@T){
	print "Testing '$host'...\n";

	# first in nonblocking mode
	$LW2::LW_NONBLOCK_CONNECT = 1;
	connect_test($host);
	LW2::http_reset();
	
	# then force blocking mode
	$LW2::LW_NONBLOCK_CONNECT = 0;
	connect_test($host);
	LW2::http_reset();
}

@T = ();
push @T, @SSLHOSTS;
push @T, @ERRORHOSTSSSL;

foreach $host (@T){
	print "Testing '$host'...\n";

	# first in nonblocking mode, non SSL KA mode
	$LW2::LW_NONBLOCK_CONNECT = 1;
	$LW2::LW_SSL_KEEPALIVE = 0;
	connect_test($host);
	LW2::http_reset();

	# next in nonblocking, SSL KA mode
	$LW2::LW_NONBLOCK_CONNECT = 1;
	$LW2::LW_SSL_KEEPALIVE = 1;
	connect_test($host);
	LW2::http_reset();
	
	# then force blocking mode, non SSL KA
	$LW2::LW_NONBLOCK_CONNECT = 0;
	$LW2::LW_SSL_KEEPALIVE = 0;
	connect_test($host);
	LW2::http_reset();

	# last force blocking mode, SSL KA mode
	$LW2::LW_NONBLOCK_CONNECT = 0;
	$LW2::LW_SSL_KEEPALIVE = 1;
	connect_test($host);
	LW2::http_reset();
}



print STDOUT <<EOT;

---------------------------------------------------------------------

Test complete.  Here is the result data:

EOT

$asv = '';
$wos = '';
if($^O =~ /win32/i || $^O =~ /cygwin/i ){
	eval {
		eval "use Win32";
		if(!$@){
			$asv = Win32::BuildNumber(); 
			$wos = Win32::GetOSName();
		}
	}
}


$PLATFORM = "Perl $] (". sprintf("%vd", $^V). " $asv) on $^O $wos\n";
if(LW2::ssl_is_available()){
	@x = LW2::ssl_is_available();
	$PLATFORM .= "Using SSL: $x[1] $x[2]\n";
}

print STDOUT $PLATFORM, join("\n", @DATA), "\n";

print STDOUT <<EOT;

---------------------------------------------------------------------

Hit 'y' to submit this result data to wiretrip.net (libwhisker's
maintainer).  Hit 'n' to skip the automated submission.

NOTE: only the data displayed above is submitted; no other identifing
information is included (although your IP address is obviously
involved in the HTTP communication process).  The information is purely
for compiling result data to better enable libwhisker to work on more
platforms, and will only be used for development purposes.

EOT

print STDOUT "Would you like to automatically submit these results?\n";
$res = yesno();

if($res eq 'y'){
	$REQ = LW2::http_new_request();
	$RESP = LW2::http_new_response();
	LW2::uri_split('http://www.wiretrip.net/rfp/libwhisker-data.asp',
		$REQ);
	$REQ->{whisker}->{method} = 'POST';
	$REQ->{whisker}->{data} = 'data=' . LW2::uri_escape($PLATFORM . join("\n", @DATA));
	
	if($CONFIG{proxy_host} ne ''){
		my @x = split(/:/, $CONFIG{proxy_host});
		$REQ->{whisker}->{proxy_host} = $x[0];
		$REQ->{whisker}->{proxy_port} = $x[1];

		if($CONFIG{proxy_auth} == 1){
			LW2::auth_set('proxy-basic', $REQ,
				$CONFIG{proxy_user},
				$CONFIG{proxy_pass});
		}

		if($CONFIG{proxy_auth} == 2){
			LW2::auth_set('proxy-ntlm', $REQ,
				$CONFIG{proxy_user},
				$CONFIG{proxy_pass});
		}
	}
	
	LW2::http_fixup_request($REQ);
	if(LW2::http_do_request($REQ,$RESP)){
		print STDOUT <<EOT;

!! There was an error during the automated submission process.  The
error is below.

$RESP->{whisker}->{error}

EOT
		exit;
	}

	print LW2::dump('resp', $RESP), "\n";

	print STDOUT <<EOT;

Your results have been submitted.  Thank you for helping to support
libwhisker development.

EOT

} else {
	print STDOUT <<EOT;

Since you opted to not automatically submit your result data, it is not
going to benefit the libwhisker development effort.  Please reconsider
re-running this script and allowing the automated data submission in
order to help aid in libwhisker development and platform support.

EOT

}


print STDOUT "Thanks for choosing libwhisker2!  - rfp\n\n";

exit;

#######################################################################

sub getinfo {
	my $prompt = shift;
	my $x = '';
	do {
		print STDOUT "Enter $prompt: ";
		$x = <STDIN>;
		$x=~tr/\r\n//d;
	} while ($x eq '');
	return $x;
}

sub yesno {
	my $x = '';
	do {
		print STDOUT "Choice [y/n]: ";
		$x = <STDIN>;
		$x=~tr/YyNn//cd;
	} while ($x eq '');
	return lc($x);
}

sub connect_test {
	my $host = shift;

	my $REQ = LW2::http_new_request();
	my $RESP = LW2::http_new_response();

	$REQ->{whisker}->{host} = '';
	$REQ->{whisker}->{port} = 0;

	LW2::uri_split($host, $REQ);

	if($REQ->{whisker}->{host} eq '' || $REQ->{whisker}->{port} == 0){
		print STDERR "Error: bad host '$host'\n";
		return 0;
	}

	if($REQ->{whisker}->{ssl}==1 && !LW2::ssl_is_available()){
		print STDERR "Skipping SSL site '$host'\n";
		return 0;
	}

	if($CONFIG{proxy_host} ne ''){
		my @x = split(/:/, $CONFIG{proxy_host});
		$REQ->{whisker}->{proxy_host} = $x[0];
		$REQ->{whisker}->{proxy_port} = $x[1];

		if($CONFIG{proxy_auth} == 1){
			LW2::auth_set('proxy-basic', $REQ,
				$CONFIG{proxy_user},
				$CONFIG{proxy_pass});
		}

		if($CONFIG{proxy_auth} == 2){
			LW2::auth_set('proxy-ntlm', $REQ,
				$CONFIG{proxy_user},
				$CONFIG{proxy_pass});
		}
	}

	$REQ->{whisker}->{retry} = 0;
	$REQ->{whisker}->{timeout} = 13;

	LW2::http_fixup_request($REQ);

	my %data = (
		times => '',
		nb => $LW2::LW_NONBLOCK_CONNECT,
		sslka => $LW2::LW_SSL_KEEPALIVE,
		reqs => '',
		syns => '',
		sstate => ''
	);

	my $loop = 0;
	do {

		my $bef_nb = $LW2::LW_NONBLOCK_CONNECT;
		my $bef_ssl = $LW2::LW_SSL_KEEPALIVE;
	
		my $start = time;
		my $res = LW2::http_do_request($REQ,$RESP);
		my $diff = (time - $start);

		$RESP->{whisker}->{stats_reqs} ||= 0;
		$RESP->{whisker}->{stats_syns} ||= 0;
		$RESP->{whisker}->{socket_state} ||= 0;

		$data{times} .= ",$diff";
		$data{reqs} .= ','.$RESP->{whisker}->{stats_reqs};
		$data{syns} .= ','.$RESP->{whisker}->{stats_syns};
		$data{sstate} .= ','.$RESP->{whisker}->{socket_state};
		$data{nb} .= ','.$LW2::LW_NONBLOCK_CONNECT;

		if($res){
			push @DATA, "$host";
			push @DATA, "\t" . make_data_line(\%data);
			push @DATA, "\tError: " .
				 $RESP->{whisker}->{error};
			return 0;	
		}

		if($RESP->{whisker}->{code} != 200 && (
			$RESP->{whisker}->{code} < 300 ||
			$RESP->{whisker}->{code} > 305) ){
			push @DATA, "$host";
			push @DATA, "\t" . make_data_line(\%data);
			push @DATA, "\tNon-200/30x response";
			return 0;
		}

		$loop++;

	} while($loop < 4);

	push @DATA, "$host";
	push @DATA, "\t" . make_data_line(\%data);

	return 1;
}

sub make_data_line {
	my $hr = shift;

	$hr->{reqs} =~ s/^,//;
	$hr->{syns} =~ s/^,//;
	$hr->{sstate} =~ s/^,//;
	$hr->{times} =~ s/^,//;

	my $l = 'time '.$hr->{times}.'/ska '.$hr->{sslka}.
		'/nb '.$hr->{nb}.'/syn '.$hr->{syns}.'/req '.
		$hr->{reqs}.'/sock '.$hr->{sstate};
	
	return $l;
}
