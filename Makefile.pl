# THIS FILE WAS NOT AUTO GENERATED THIS SCRIPT WAS MADE AS AN INSTALLER DO NOT TOUCH THIS FILE UNLESS YOU ARE A DEVELOPER OF THE OFFICAL RED RABBIT PROJECT
#
# PROJECT  - THE RED RABBIT PROJECT
# AUTHOR   - ArkAngeL43
# LOCATION - https://www.github.com/ArkAngeL43/Red-Rabbit-V4 
# FILENAME - Makefile.pl
# FILERUN  - Makefile.pl
# 
#
#
#
#
# Updated: Thur, May 19th 2022 
#
# New support for arch linux systems, the perl INC paths are not the same as they were before for debian systems, so given the change i the author decided to 
# Allow support for all PERL INC 5.34 PATHS
#
#
#
#
#
#
#
#
#
#
#
#
#
#
system("clear");
use Cwd;
use File::Copy;
$| = 1;
binmode STDOUT, ":utf8";
my $dotdotdotdotdotdotdotdotdotdot = ".";
print $camel;
for (1..1961) {
    sleep 0.9;
    print " $dotdotdotdotdotdotdotdotdotdot";
}
print "\n";
sleep(2);
my $SCRIPT_VERSION               = "V 1.0";
my $SCRIPT_NAME                  = "Red Rabbit";
my $SCRIPT_AUTHOR                = "ArkAngeL43";
my $WORKING_DIR                  = Cwd::cwd();
my $BRUTE_MOD                    = "modules/perl/mod/BRUTE.pm";
my $OFSL_MOD                     = "modules/perl/mod/OFSL.pm";
my $TABLE_MOD                    = "modules/perl/mod/TABLE.pm";
my $IMAGE_MOD                    = "/usr/share/perl5/Image";
my $PERL_MODULE_FILE_DST_DEBIAN  = "/usr/share/perl5";
my $PERL_MODULE_FILE_DST_ARCH    = "/usr/lib/perl5/5.34/site_perl";
my $OPTIONS               = shift;
my $OPERATING_SYSTEM_FORM = shift;
my $counter = 0;
#=HEAD1_SET_VALUES_MODULES_NAMES_AND_ARRAYS

my @LOCATE_FILES;

my %MODULES = (
	'English'           => 0,
	'Image::ExifTool'   => 0,
    'strict'            => 0,
	'warnings'          => 0,
	'Getopt::Std'       => 0,
	'Cwd'               => 0,
	'English'           => 0,
	'POSIX'             => 0,
	'File::Basename'    => 0,
	'Image::ExifTool'   => 0,
	'Data::HexDump'     => 0,
	'String::CRC32'     => 0,
	'HTTP::Request'     => 0,
	'LWP::UserAgent'    => 0,
	'LWP::UserAgent'    => 0,
	'Net::Netmask'      => 0,
	'Net::Traceroute'   => 0,
	'Net::Telnet'       => 0,
	'GD'                => 0
);

my @SYS_COMMANDS_PERL = (
	"sudo cpan install Text::Table",
	"sudo cpan install Tk",
	"sudo cpan install Tk::NoteBook",
	"sudo cpan install Tk::ROText",
	"sudo cpan install Tk::widgets",
	"sudo cpan install Getopt::Std",
	"sudo cpan install Data::HexDump",
	"sudo cpan install Image::ExifTool",
	"sudo cpan install HTTP::Request",
	"sudo cpan install LWP::UserAgent",
	"sudo cpan install English",
	"sudo cpan install MIME::Base64",
	"sudo cpan install HTTP::Tiny",
	"sudo cpan install Fcntl",
	"sudo cpan install Net::Netmask",
	"sudo cpan install Net::Traceroute",
	"sudo cpan install Net::Telnet",
	"sudo cpan install YAML",
	"sudo cpan install JSON",

);

my @SYS_COMMANDS = (
	"go get -u github.com/spf13/pflag",
    "go get -u github.com/google/gopacket/layers",
	"go get -u github.com/google/gopacket",
	"go get -u github.com/google/gopacket/pcap",
	"go get -u golang.org/x/crypto/ssh",
	"go get -u github.com/go-sql-driver/mysql",
	"go get -u github.com/jlaffaye/ftp",
	"go get -u github.com/lib/pq",
	"go get -u gopkg.in/mgo.v2",
	"go get -u github.com/PuerkitoBio/goquery",
	"go get -u github.com/atotto/clipboard",
	"go get -u github.com/steelx/extractlinks",
	"go get -u golang.org/x/net/html",
    "go get -u github.com/flopp/go-staticmaps",
	"go get -u github.com/fogleman/gg",
	"go get -u github.com/golang/geo/s2",
	"go get -u github.com/rwcarlsen/goexif/exif",
);

my @MODS = (
	"modules/ruby/color.rb",
	"modules/ruby/fake-ap.rb",
	"modules/ruby/file_parser.rb",
	"modules/ruby/ftp.rb",
	"modules/ruby/net.rb",
	"modules/ruby/shod.rb",
	"modules/perl/mod/BRUTE.pm",
	"modules/perl/mod/OFSL.pm",
	"modules/perl/mod/TABLE.pm",	
);


my @make_commands_CPP = (
	"sudo go build main.go",
	"sudo make -C modules/go-main/80211",
);

my @clean_source_code = (
	"sudo rm main.go",
	"sudo rm modules/go-main/80211/interfaces.cpp",
	"sudo rm modules/go-main/80211/main.cpp",
	"sudo rm modules/go-main/80211/mod/goul/goul.go",
);

my @PODS = (
	"modules/perl/pod/ipath.pod",
	"modules/perl/pod/main-all.pod",
	"modules/perl/pod/notes.pod",
	"modules/perl/pod/OFS.pod",
	"modules/perl/pod/table.pod",
	"modules/perl/pod/tests.pod",
	"modules/perl/pod/type.pod",
);

my @DEBIAN_FILEPATHS = (
	"/usr/share/perl5/OFSL.pm",
	"/usr/share/perl5/BRUTE.pm",
	"/usr/share/perl5/TABLE.pm",
);

my @DEBIAN_FILEPATHS_TO_SEND = (
	"/usr/share/perl5",
);

my @ARCH_FILEPATHS_TO_SEND = (
	"/usr/lib/perl5/5.34/site_perl",
);

my @NAMES = (
		"text/banners/rr6.txt" => 0, 
		"text/banners/screen-logo-small.txt",
		"text/banners/shark.txt",
		"text/banners/team-logo.txt",
		"text/help/advance.txt",
		"text/help/command-flags.txt",
		"text/menus/verticle_menu.txt",
		"text/menus/landscape_menu.txt",
		"examples/pcap/aaa.pcap",
		"examples/pcap/IMAP - Authenticate CRAM-MD5.cap",
		"examples/pcap/ospf-md5_key-1234.pcap",
		"examples/pcap/SMB - NTLMSSP (Windows 10).pcapng",
		"examples/pcap/ARP - Broadcast.pcap",
		"examples/pcap/IMAP - Authenticate CRAM-MD5.pcapng",
		"examples/pcap/ospf-simple.cap",
		"examples/pcap/SMTP - Auth Cram MD5.pcap",
		"examples/pcap/ARP - Broadcast.pcapng",
		"examples/pcap/IMAP - Authenticate Digest-MD5.pcap",
		"examples/pcap/ospf simple password authentication.cap",
		"examples/pcap/SMTP - Auth Cram MD5.pcapng",
		"examples/pcap/Asterisk_ZFONE_XLITE.pcap",
		"examples/pcap/IMAP - Authenticate Digest-MD5.pcapng",
		"examples/pcap/OSPF_with_MD5_auth.cap",
		"examples/pcap/SMTP - Auth Login.pcapng",
		"examples/pcap/dot11.pcapng",
		"examples/pcap/IMAP - Authenticate Plain (Base64).pcap",
		"examples/pcap/POP3.pcap",
		"examples/pcap/SMTP - Auth Plain.pcap",
		"examples/pcap/DTMFsipinfo.pcap",
		"examples/pcap/IMAP - Authenticate Plain (Base64).pcapng",
		"examples/pcap/POP3.pcapng",
		"examples/pcap/SMTP - Auth Plain.pcapng",
		"examples/pcap/FAX-Call-t38-CA-TDM-SIP-FB-1.pcap",
		"examples/pcap/IMAP - Authenticate XYMPKI.pcap",
		"examples/pcap/SIP CALL.pcap",
		"examples/pcap/smtp-login.pcap",
		"examples/pcap/Ftp.pcap",
		"examples/pcap/IMAP - Authenticate XYMPKI.pcapng",
		"examples/pcap/SIP_DTMF2.pcap",
		"examples/pcap/smtp.pcap",
		"examples/pcap/Ftp.pcapng",
		"examples/pcap/IMAP - Login (Plaintext) 1.pcap",
		"examples/pcap/sip-rtp-dvi4.pcap",
		"examples/pcap/TCP - 5 Packets.pcap",
		"examples/pcap/h223-over-rtp.cap",
		"examples/pcap/IMAP - Login (Plaintext) 1.pcapng",
		"examples/pcap/sip-rtp-g711.pcap",
		"examples/pcap/TCP - 5 Packets.pcapng",
		"examples/pcap/h263-over-rtp.pcap",
		"examples/pcap/IMAP - Login (Plaintext) 2.pcap",
		"examples/pcap/sip-rtp-g722.pcap",
		"examples/pcap/TCP - File Downloads.pcap",
		"examples/pcap/HTTP - Basic Authentication.pcap",
		"examples/pcap/IMAP - Login (Plaintext) 2.pcapng",
		"examples/pcap/sip-rtp-g726.pcap",
		"examples/pcap/TCP - File Downloads.pcapng",
		"examples/pcap/examples/pcap/HTTP - Basic Authentication.pcapng",
		"examples/pcap/IMAP - Login (Plaintext) 3.pcap",
		"examples/pcap/sip-rtp-g729a.pcap",
		"examples/pcap/TCP - Network.pcap",
		"examples/pcap/examples/pcap/HTTP - Digest Authentication.pcap",
		"examples/pcap/IMAP - Login (Plaintext) 3.pcapng",
		"examples/pcap/sip-rtp-gsm.pcap",
		"examples/pcap/TCP - Network.pcapng",
		"examples/pcap/HTTP - Digest Authentication.pcapng",
		"examples/pcap/Kerberos-816.pcap",
		"examples/pcap/sip-rtp-ilbc.pcap",
		"examples/pcap/TCP - Tcpreplay.pcap",
		"examples/pcap/examples/pcap/HTTP - Digest-MD5.pcap",
		"examples/pcap/Kerberos-816.pcapng",
		"examples/pcap/sip-rtp-l16.pcap",
		"examples/pcap/TCP - Tcpreplay.pcapng",
		"examples/pcap/HTTP - Digest-MD5.pcapng",
		"examples/pcap/Kerberos - v5 TCP.pcap",
		"examples/pcap/sip-rtp-lpc.pcap",
		"examples/pcap/Telnet - Char Mode2.pcap",
		"examples/pcap/HTTP - JPG Download.pcap",
		"examples/pcap/Kerberos - v5 TCP.pcapng", "sip-rtp_only.pcap",
		"examples/pcap/Telnet - Char Mode2.pcapng",
		"examples/pcap/HTTP - JPG Download.pcapng",
		"examples/pcap/Kerberos v5 UDP 2.pcap", "sip-rtp-opus.pcap",
		"examples/pcap/Telnet - Char Mode.pcap",
		"examples/pcap/examples/pcap/HTTP - NTLM GSSAPI.pcap",
		"examples/pcap/Kerberos v5 UDP 2.pcapng",
		"examples/pcap/sip-rtp-speex.pcap",
		"examples/pcap/Telnet - Char Mode.pcapng",
		"examples/pcap/HTTP - NTLM GSSAPI.pcapng",
		"examples/pcap/Kerberos v5 - UDP 3.pcap",
		"examples/pcap/SMB - NTLM cifs_SessionSetupAndX_NTLM_Plain.pcap",
		"examples/pcap/Telnet - Line Mode.pcap",
		"examples/pcap/HTTP - NTLM.pcap",
		"examples/pcap/Kerberos v5 - UDP 3.pcapng",
		"examples/pcap/SMB - NTLM cifs_SessionSetupAndX_NTLM_Plain.pcapng",
		"examples/pcap/Telnet - Line Mode.pcapng",
		"examples/pcap/HTTP - NTLM.pcapng",
		"examples/pcap/Kerberos - v5 UDP.pcap",
		"examples/pcap/SMB - NTLMSSP Single Session (Windows 10).pcap",
		"examples/pcap/Telnet.pcap",
		"examples/pcap/HTTP - PDF file download.pcap", "examples/pcap/Kerberos - v5 UDP.pcapng",
		"examples/pcap/SMB - NTLMSSP Single Session (Windows 10).pcapng", "examples/pcap/Telnet.pcapng",
		"examples/pcap/HTTP - PDF file download.pcapng", "examples/pcap/MagicJack+_short_call.pcap",
		"examples/pcap/SMB - NTLMSSP (smb3 aes 128 ccm).pcap",
		"examples/pcap/udp_1_packet.pcap",
		"examples/pcap/HTTP - Small File.pcap", "examples/pcap/metasploit-sip-invite-spoof.pcap",
		"examples/pcap/SMB - NTLMSSP (smb3 aes 128 ccm).pcapng",
		"examples/pcap/udp_1_packet.pcapng",
		"examples/pcap/HTTP - Small File.pcapng",
		"examples/pcap/only_sip_sdp_rtp.pcap", 
		"examples/pcap/SMB - NTLMSSP (Windows 10).pcap",
);





main($OPTIONS);

sub end_install {
	foreach my $make_in(@make_commands_CPP) {
		print "[+] Operation: \033[37m(SYSTEM|OS)\033[39m Running -> $make_in\n";
		system($make_in);
	}
	foreach my $source(@clean_source_code) {
		print "[+] Operation: \033[37m(SYSTEM|OS)\033[39m Removing source code -> $source\n";
		system($source);
	}
}

sub command_install {
	print("[+] Operation: \033[35m(Setting)\033[39m Working Directory -> $WORKING_DIR\n");
	print("[+] Operation: \033[35m(Setting)\033[39m Operating System  -> $OPERATING_SYSTEM_FORM\n");
	sleep(1);
	if ($OPERATING_SYSTEM_FORM eq "arch") {
		my @ARCH_FILEPATHS_TO_SEND = (
			"/usr/lib/perl5/5.34/site_perl",
		);
		my @modules = (
			$OFSL_MOD,
			$TABLE_MOD,
			$BRUTE_MOD,
		);
		foreach my $modpath(@ARCH_FILEPATHS_TO_SEND) {
			foreach my $modfile(@modules) {
				print("[+] Operation: \033[35m(Setting)\033[39m Sending Module    -> $modfile\n");
				system("cp $modfile -R $modpath");
				print("[+] Operation: \033[35m(Setting)\033[39m Modules New Path  -> $modpath\n");				
			}
		}
	} else {
		if ($OPERATING_SYSTEM_FORM eq "debian") {
			my @modules = (
				$OFSL_MOD,
				$TABLE_MOD,
				$BRUTE_MOD,
			);
			foreach my $sendpath(@DEBIAN_FILEPATHS_TO_SEND) {
				foreach my $module(@modules) {
					print("[+] Operation: \033[35m(Setting)\033[39m Sending Module    -> $module\n");
					print("[+] Operation: \033[35m(Setting)\033[39m Module Filepath   -> $sendpath\n");
					system("cp $module -R $sendpath");
				}
			}
		}
	}
	print "[+] Operation: \033[37m(SYSTEM|OS)\033[39m Running GOLANG module installs\n";
	sleep(1);
	foreach my $gomod(@SYS_COMMANDS) {
		print "[+] Operation: \033[37m(SYSTEM|OS)\033[39m Running -> $gomod\n";
		system("$gomod");
		$counter++;
	}
	print "[+] Operation: \033[37m(SYSTEM|OS)\033[39m Installed -> $counter GO libs\n";
	print "[+] Operation: \033[37m(SYSTEM|OS)\033[39m Running all perl installs\n";
	sleep(1);
	foreach my $perlmod(@SYS_COMMANDS_PERL) {
		print "[+] Operation: \033[37m(SYSTEM|OS)\033[39m Running -> $perlmod\n";
		system($perlmod);		
	}
	&end_install;
}



# i was way to lazy to make a dev version, so i copied this one from the other command check i made
sub command_check_dev {
	print("[+] Operation: PREFORMING DEVELOPMENT INSTALL.....");
	print("[+] Operation: \033[35m(Setting)\033[39m Working Directory -> $WORKING_DIR\n");
	print("[+] Operation: \033[35m(Setting)\033[39m Operating System  -> $OPERATING_SYSTEM_FORM\n");
	sleep(1);
	if ($OPERATING_SYSTEM_FORM eq "arch") {
		my @ARCH_FILEPATHS_TO_SEND = (
			"/usr/lib/perl5/5.34/site_perl",
		);
		my @modules = (
			$OFSL_MOD,
			$TABLE_MOD,
			$BRUTE_MOD,
		);
		foreach my $modpath(@ARCH_FILEPATHS_TO_SEND) {
			foreach my $modfile(@modules) {
				print("[+] Operation: \033[35m(Setting)\033[39m Sending Module    -> $modfile\n");
				system("cp $modfile -R $modpath");
				print("[+] Operation: \033[35m(Setting)\033[39m Modules New Path  -> $modpath\n");				
			}
		}
	} else {
		if ($OPERATING_SYSTEM_FORM eq "debian") {
			my @modules = (
				$OFSL_MOD,
				$TABLE_MOD,
				$BRUTE_MOD,
			);
			foreach my $sendpath(@DEBIAN_FILEPATHS_TO_SEND) {
				foreach my $module(@modules) {
					print("[+] Operation: \033[35m(Setting)\033[39m Sending Module    -> $module\n");
					print("[+] Operation: \033[35m(Setting)\033[39m Module Filepath   -> $sendpath\n");
					system("cp $module -R $sendpath");
				}
			}
		}
	}
}



sub command_check {
	sleep(1);
	print "[+] Operation: \033[35m(Setting)\033[39m Checking all modules and filepaths for needed files\n";
	foreach (keys %MODULES){
		eval "use $_;";
		if(!$@){
			my $name = $MODULES{$_}++;
			print "[+] Operation: \033[35m(Testing)\033[39m Found module $_\n";
			print "[+] Operation: \033[35m(Testing)\033[39m TEST OK++ \n";
		} else {
			print "[+] Operation: \033[35m(Errorzz)\033[39m Could not find module [$_] FATAL Missing??????????\n";
			print "[+] Operation: \033[35m(Install)\033[39m Installing missing module...\n";
			system("sudo cpan install $_");
			print "\n";
		}
	}
	if ($OPERATING_SYSTEM_FORM eq "arch") {
		my @ARCH_FILEPATHS_TO_SEND = (
			"/usr/lib/perl5/5.34/site_perl",
		);
		my @modules = (
			$OFSL_MOD,
			$TABLE_MOD,
			$BRUTE_MOD,
		);
		foreach my $modpath(@ARCH_FILEPATHS_TO_SEND) {
			my @new_path = (
				$modpath."/BRUTE.pm",
				$modpath."/TABLE.pm",
				$modpath."/OFSL.pm",
			);
			foreach my $modfile(@modules) {
				foreach my $newpath(@new_path) {
					if (-f $newpath) {
							print "[+] Operation: \033[35m(Testing)\033[39m Found module $newpath\n";
							print "[+] Operation: \033[35m(Testing)\033[39m TEST OK++ \n";
					} else {
						print "[+] Operation: \033[35m(Errorzz)\033[39m Could not find module [$newpath] FATAL Missing??????????\n";
						print "[+] Operation: \033[35m(Install)\033[39m Installing missing module...\n";
						&command_install;
					}

				}
			}
		}
	} else {
		if ($OPERATING_SYSTEM_FORM eq "debian") {
			my @modules = (
				$OFSL_MOD,
				$TABLE_MOD,
				$BRUTE_MOD,
			);
			foreach my $sendpath(@DEBIAN_FILEPATHS_TO_SEND) {
				my @newpaths = (
					$sendpath."/OFSL.pm",
					$sendpath."/BRUTE.pm",
					$sendpath."/TABLE.pm",
				);
				foreach my $module(@modules) {
					foreach my $checkpath(@newpaths) {
						if (-f $checkpath) {
								print "[+] Operation: \033[35m(Testing)\033[39m Found module $checkpath\n";
								print "[+] Operation: \033[35m(Testing)\033[39m TEST OK++ \n";
						} else {
							print "[+] Operation: \033[35m(Errorzz)\033[39m Could not find module [$checkpath] FATAL Missing??????????\n";
							print "[+] Operation: \033[35m(Install)\033[39m Installing missing module...\n";
							&command_install;
						}
					}
				}
			}
		}
	}
	exit();
}

sub command_SUPER {
	print "Checking all.....";
}

sub menu {
	print "MAKEFILE: <RR6> ERROR: No arguments defined, please use the following settings\n";
	print "ARGV: install    | installs all files and directories\n";
	print "ARGV: clean      | will clean and erase the RR6 directory\n";
	print "ARGV: check      | will check and run all modules to see if installed\n";
	print "ARGV: checkS     | Will SUPER check all files and modules and install modules\n";
	print "################################################################################\n";
	print "Usage Install < Arch linux >   : sudo perl Makefile.pl install arch \n";
	print "Usage Install < Debian Linux > : sudo perl Makefile.pl install debain\n";
	print "Usage Check   < Arch linux >   : sudo perl Makefile.pl check arch \n";
	print "Usage Check   < Debian linux > : sudo perl Makefile.pl check debian \n";
	print "Usage Checks  < Arch linux >   : sudo perl Makefile.pl checkS arch\n";
	print "Usage Checks  < Debian linux > : sudo perl Makefile.pl checkS debian\n";
	exit(0);
}

sub main {
	my $commands___options = shift;
	if ( $commands___options eq 'dev_install')   { command_check_dev;                 }
	if ( $commands___options eq '')              { menu;            				  }
	if ( $commands___options eq 'install')       { command_install; 				  }
	if ( $commands___options eq 'check')         { command_check;   				  }
	if ($OPERATING_SYSTEM_FORM != "debian" or $OPERATING_SYSTEM_FORM != "arch" ) {print "[-] WARN: Your system is not debian or arch, can not continue";menu;}
	if ($OPERATING_SYSTEM_FORM eq "") {print "[-] WARN: Operating system ARGUMENT 2 was empty, please specify an OS";menu;}
	#
	#
	#
	#
	#
	#
	#
	#
	#
	#
	#
	# MAJOR BREAK NEEDED
	#
	#
	#
	#
	#
	#
	#
	#
	#
	#
	#
	if ( $commands___options eq 'checkS') {		
			foreach my $output_files (@NAMES) {
				open(OP, $output_files);
				print "\033[0;39m\033[34m+\033[35mPerl MakeFile \033[34m+: \033[39m \033[36m(Running)\033[39m Searching for Example files and required banners....\n";
				while(<OP>) {
					print "\033[0;39m\033[34m+\033[35mPerl MakeFile \033[34m+: \033[39m \033[36m(WorkDir)\033[39m \033[4;34mFOUND FILE $output_files\n\033[4;39m";
					print "\033[0;39m\033[34m+\033[35mPerl MakeFile \033[34m+: \033[39m \033[34m(Testing)\033[39m FILE OK.....\n" or die ("ERROR: Could not open up file this is a FATAL ISSUE! THESE FILES ARE NEEDED");
				}
				close(OP);
			}
			print "\033[0;39m\033[34m+\033[35mPerl MakeFile \033[34m+: \033[39m \033[37m(USER WARNING!!!!!) Searching for .mod, .in, .1in, .pod, .mod, .pm, .mod. etc files\n\033[39m";
			sleep(1);
			foreach my $Modded_files(@MODS) {
				print "\033[0;39m\033[34m+\033[35mPerl MakeFile \033[34m+: \033[39m \033[36m(Running)\033[39m Searching for required modules....\n";
				open(MOD, $Modded_files);
				print "\033[0;39m\033[34m+\033[35mPerl MakeFile \033[34m+: \033[39m \033[36m(WorkDir)\033[39m \033[4;34mFOUND MOD $Modded_files\n\033[4;39m" or die("[-] Operation: \033[34m(Testing)\033[39m FILE IS NOT OK!!!!!! FILENAME $Modded_files might be missing or corrupt");
				print "\033[0;39m\033[34m+\033[35mPerl MakeFile \033[34m+: \033[39m \033[34m(Testing)\033[39m FILE OK.....\n" or die("[-] Operation: \033[34m(Testing)\033[39m FILE IS NOT OK!!!!!! FILENAME $Modded_files might be missing or corrupt");
				close(MOD);
			}
			print "\033[0;39m\033[34m+\033[35mPerl MakeFile \033[34m+: \033[39m (OUTPUT WARNING) !!!!!!! SEARCHING FOR ALL .pod FILES FOR MANUAL MAKEFILE CONFIGURATION AND SCRIPT SETTINGS AND NOTES.....\n\033[39m";
			sleep(1);
			foreach my $PODS (@PODS) {
				print "\033[0;39m\033[34m+\033[35mPerl MakeFile \033[34m+: \033[39m \033[36m(RUNNING)\033[39m Searching for required POD files\n";
				open(POD, $PODS);
				print "\033[0;39m\033[34m+\033[35mPerl MakeFile \033[34m+: \033[39m(\033[4;33mTesting\033[0;39m)\033[39m FOUND POD FILE $PODS\n" or die("MAKE ERROR: Could not locate perl mod or pod files CORRUPTION ERROR 225");
				close(POD);
			} 
			open(FH, "sub/config-files.txt") or die $!;
			while(<FH>) {
				print "\033[0;39m\033[34m+\033[35mPerl MakeFile\033[34m+: (Testing)\033[39m \033[37m [ Checking filenames ] \n";
				opendir(DIR,$_);
				my @dir = readdir(DIR);
				close DIR;
				if (-d $dir . "/" . $directory ){
					print "\033[0;39m\033[34m+\033[35mPerl MakeFile\033[34m+: \033[39m Searching Files.....\n";
					print "\033[0;39m\033[34m+\033[35mPerl MakeFile\033[34m+: \033[39m Found $_\n";
				} else {
					print "\033[0;39m\033[34m+\033[35mPerl MakeFile \033[34m+: \033[39m Skipping..\n";
				}
			} 
			command_install;
			command_check;
	}
}