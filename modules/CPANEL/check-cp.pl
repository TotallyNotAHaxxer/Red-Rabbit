#!/usr/bin/perl
# cpanel11 - expcheck.pl          Copyright(c) 2006 cPanel Inc.
#                                 All rights Reserved.
# copyright@cpanel.net            http://www.cpanel.net
#
# this a rewrite holding better more standardized code rewritten by ArkAngeL43
# credits to CPANEL for the idea of this scanner
use feature 'say';

my $notok=0;

my @paths = (
    "/usr/local/cpanel/bin/mysqladmin",
    "/usr/local/cpanel/bin/hooksadmin",
);


sub check() {
    for my $path (@paths) {
    my %TA=('unshift( @INC, "/usr/local/cpanel" );'=>0,,'@INC=grep(!/(^\.|\.\.|\/\.+)/,@INC);'=>0,'@INC=grep(/^(\/usr\/lib\d*\/perl|\/usr\/local\/lib\d*\/perl|\/usr\/local\/cpanel)/,@INC);'=>0);
    say "\033[36mChecking Path -> $path\n";
    if(! -e $path) {
        say "\033[31m[*] WARNING: FATAL: STAT: CHECK:.....[..-> Script or path not installed STAT (OK)_Finished scan" . "\n"; next();
    } 
    open(SM, '<', $path);
    #
    while(my $l = <SM>) {
        foreach my $a (keys %TA) {
            if (index($l,$a) > -1) { delete $TA{$a}; next(); }
        }
    }
    if ((scalar keys %TA) > 0) {

        $notok++;
        say "unsafe";
    } else {
        say "safe";
    say "PATH-> $path";
    }
    say "..Done\n";
    }
    if ($notok) {
        say "\033[31m[!] Warning: FATAL: Tests have all failed to pass, this might be due to the fact your using a verion of CPANEL that is vulnerable";
        say "\033[31m[!] Warning: FATAL: It is highly suggested you navigate to the Home directory of RR5 and run the following command";
        say "\033[31m[!] `perl sec095206.pl` to update your CONF file and version of CPANEL";
        say "\033[31m[!] This file will automatically update to the newest version as of Sunday Feb 6 2022, if you feel this script is older";
        say "\033[31m[!] Please navigate to the following URL -> http://layer1.cpanel.net/installer/sec092506.pl to download the right one";
        say "\033[31m[!] \n~ArkAngeL43";
    } else {
        say "\033[32m[*] System has been patched or does not utilize a version of CPANEL that is vulnerable";

    }
}


check();