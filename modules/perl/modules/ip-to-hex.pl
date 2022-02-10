#!/usr/bin/perl
use strict;
my $ip = shift;

my @octets = split(/\./, $ip);

print "ip->$ip->CONV-> 0x";
foreach my $octet(@octets) {
   $octet =~ s/$octet/sprintf("%X",$octet)/eg;
   print $octet
}
print "\n"; 