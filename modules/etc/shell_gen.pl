#!/usr/bin/perl

print "\n[!] Please input the shellcode -=> ";
$x1=<>;
my $data = "$x1";
chomp($data);
my @values = split(undef,$data);

print ("\033[31m[+] NEW SHELLCODE ------------....\n\033[39m");

foreach my $val (@values) {
  chomp($val);
  print '\x';
  print unpack(H8,"$val");
}

print "\n";
exit 0;