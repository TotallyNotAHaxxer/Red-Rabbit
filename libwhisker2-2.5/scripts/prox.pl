#!/usr/bin/perl
#
# prox.pl - copyright rfp 2001,2002
#
# This is a general proxy used to troubleshoot protocol problems.  Really
# designed to handle HTTP and other ASCII single-connection protocols.

use IO::Socket;
use IO::Select;
$|++;

my $insocket = IO::Socket::INET->new(LocalPort => 81, Listen => 20,
					Proto  => 'tcp', Reuse  => 1);
die $! unless $insocket;

my $readers = IO::Select->new() or die $!;

while(1){

	my $incoming = $insocket->accept();
	my $outgoing = IO::Socket::INET->new("localhost:80") or die $!;

	$incoming->autoflush(1);
	$outgoing->autoflush(1);

	$readers->add($incoming);	
	$readers->add($outgoing);	

	$check=1;
	while($check){
		my @ready = $readers->can_read;
		for my $handle (@ready){

			my $data;
			my $c = sysread($handle,$data,2048);

			if($c > 0){
				if($handle eq $incoming){
					syswrite($outgoing,$data);
					$dir = '>';
				} else {
					syswrite($incoming,$data);
					$dir = '<';
				}
				$data=~s/\r/\[r\]/g;
				$data=~s/\n/\[n\]\n/g;
				print STDOUT $dir x 75, "\n", $data, "\n";
			} else {
				$incoming->close;
				$outgoing->close;
				print '<> closed ','<>'x32,"\n";
				$check=0;
				break;
			}
		}
	}

	$readers->remove($incoming);
	$readers->remove($outgoing);
}
