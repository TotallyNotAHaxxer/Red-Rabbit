use warnings;
use 5.10.0; # Using features from 5.10

my @ADMIN_PATHS = (
    "/usr/local/cpanel/bin/mysqladmin",
    "/usr/local/cpanel/bin/hooksadmin",
);

my $SERVICE_VULN = 0;
my $FILES_LOCATED = 0;
my $LAST_PATH = "/usr/local/cpanel/bin/hooksadmin";

sub Check_Paths_Exist{
    foreach my $filepath(@ADMIN_PATHS) {
        say "[+] Checking if | $filepath exists..";
        if (-f $filepath) {
            say "[+] Filepath    | $filepath exists..";
            $FILES_LOCATED++
        }
        if ($FILES_LOCATED != 2) {
            if ($filepath eq $LAST_PATH) {
                say "\033[31m[-] Error       | Filepaths that were located was not all 2, only $FILES_LOCATED existed";
                say "\033[31m[-] Error       | This might be due to the fact that you do not have CPANEL installed on you're system, please try again later";
                exit;
            } else {
                next;
            }
        } else {
            say "[+] All good    | Moving on to testing function....";
            &Check_File_CPANEL_Version;
        }
        
    }
}

sub Check_File_CPANEL_Version{
    for my $filepath(@ADMIN_PATHS) {
                                                                                                                                                                                my %TA=('unshift( @INC, "/usr/local/cpanel" );'=>0,,
                                                                                                                                                                                '@INC=grep(!/(^\.|\.\.|\/\.+)/,@INC);'=>0,
                                                                                                                                                                                '@INC=grep(/^(\/usr\/lib\d*\/perl|\/usr\/local\/lib\d*\/perl|\/usr\/local\/cpanel)/,@INC);'=>0
                                                                                                                                                                                );
            say "[+] Checking if | $filepath is vulnerable";
            open(S, '<', $filepath);
            while (my $line = <S>) {
                foreach my $ss (keys %TA) {
                    if (index($line,$ss) > -1) { 
                        delete $TA{$ss}; 
                        next(); 
                    }
                }
            }
            if ((scalar keys %TA) > 0) {
                $SERVICE_VULN++;
                say "\033[31m[!] POSITIVE    | $filepath IS VULNERABLE AND IS OUTDATED";
            } else {
                say "\033[32m[!] NEGATIVE    | $filepath IS NOT vulnerable, directory is safe";
            }
    }
}

&Check_Paths_Exist;