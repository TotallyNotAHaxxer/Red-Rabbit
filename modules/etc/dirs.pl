#!/usr/bin/perl
use strict;
use LWP::Simple qw($ua get);


$ua = LWP::UserAgent->new;
$ua->agent('Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.3) Gecko/20070309 Firefox/2.0.0.9');
my @found;


my $url = $ARGV[0];
my $dirlist = $ARGV[1];
my $results = $ARGV[2];

open(ifile, "<$dirlist") || die "Couldn't open file\n";
my @dirs =<ifile>;
close("ifile");
&search;


sub search {
	foreach my $dir(@dirs){
	print "$url/$dir";
	my $response = $ua->get("$url$dir");
	# test if the url is real with regex str 
	#
	print "\033[34m[*] STAT: Verified URL       -> $url$dir \n" if $url =~ /^(?:(?:https?|s?ftp))/i;
	print "\033[32m[!] STAT: Processing Request -> $url$dir \n";
	#
	if($response->status_line !~ m/^404/){
		push(@found,"$url/$dir");
		print "\033[38m[+] URL came back true => $url$dir";
	} else {
		print "\033[31m[!] URL came back FALSE For testinf and response => $response -> $url$dir \n"
	}
}
}



