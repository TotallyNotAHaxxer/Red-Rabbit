#!/usr/bin/perl
#
#Author => ArkAngeL43
# Simple PHP BB version checker
use LWP::UserAgent;

# set useragent 
$useragent = LWP::UserAgent->new;


# website
$site = $ARGV[0];
#
# get the changelog
print "\033[31m[+] Trying -> $site\n";
$res = $useragent->get($site.'/docs/CHANGELOG.html');
if($res->is_success) {
    @ver=$res->content=~/<li><a href="(.*?)\">Changes since (.*?)<\/a><\/li>/i;
    $ver[1]=~/(\d+)\.(\d+)\.(\d+)/;
    $version=$1.'.'.$2.'.'.scalar $3+1;
    print 'Version: '.$version if $version; 
} else {
        print "\n\033[38m[!] Version: NULL: Could not fetch version using useragent -> $useragent on site -> $site\n";
}

