#!/usr/bin/perl
# ArkAngeL43 rewrite author:
        # there wasnt much to change but there was a few bugs where some would 
        # not scan, spready out code as well a bit to run fix
        # added time function

        # creds to the og author who wrote this, the script was published off github however it seems as if the user that posted it doesnt exist anymore, but it was written in 2010 so i would imagine so
#
#
# for the ones reading this code try to find me the actuall developer of this im quite intriguied i really couldnt find him  

$skip="TTOU TTIN TSTP STOP CONT CHLD STKFLT ALRM PIPE USR2 SEGV USR1 KILL FPE BUS IOT ABRT TRAP ILL QUIT INT HUP _DYNAMIC _GLOBAL_OFFSET_TABLE_ --";
# (2/4): script signals.
$SIG{'INT'}=\&dataexit;
$SIG{'TSTP'}=\&dataexit;
# (3/4): script routines.
sub out{print STDERR"[*] @_";}
sub outr{print STDERR"@_";}
sub outq{print STDERR"[!] @_";exit(-1);}
sub isvalid{$char=substr(shift,0,1);if(ord($char)>64&&ord($char)<91 || ord($char)>47&&ord($char)<58 || ord($char)==45 || ord($char)==95){return(1);}return(0);}

sub readbinary{
 out("$0(3): getenv() BINARY SCANNER \n");
 # color codes dont work with outq
 open(BINARY,shift)||outq("ERR: FATAL: -< could not open binary.\n");out(" Opened binary successfully \n");
 @months = qw( Jan Feb Mar Apr May Jun Jul Aug Sep Oct Nov Dec );
 @days = qw(Sun Mon Tue Wed Thu Fri Sat Sun);
 @read=<BINARY>;close(BINARY);$i=0;$tokens=@read;out("Scanning binary($tokens): ");while($read[$i]){

  @tmpread=split(chr(0),$read[$i]);$tokens=@tmpread;$j=-1;while($j<$tokens){

   $j++;$k=0;while(isvalid(substr($tmpread[$j],$k,1))&&length($tmpread[$j])>1){

    if($k+1==length($tmpread[$j])){

     $m=0;@s=split(/ /,$skip);$l=0;while($s[$l]){if($s[$l]eq$tmpread[$j]){$m++;}$l++;}

     @s=split(/,/,$result);$l=0;while($s[$l]){if($s[$l]eq$tmpread[$j]||$s[$l]eq" $tmpread[$j]"){$m++;}$l++;}

     if(!$m&&substr($tmpread[$j],0,3)ne"SIG"&&substr($tmpread[$j],0,2)ne"__"&&substr($tmpread[$j],length($tmpread[$j])-2,2)ne"__"){

      if(!$result){$result=$tmpread[$j];}

      else{$result="$result, $tmpread[$j]";}
     }
    }
    $k++;
   }
  }
  $i++;outr(".");
 }

 ($sec,$min,$hour,$mday,$mon,$year,$wday,$yday,$isdst) = localtime();
 print "$mday $months[$mon] $days[$wday]\n";
 print("[*] Finished scan at -> $hour:$min:$sec\n");
}

sub data{
 if($result){out("typical getenv() possibilities: $result.\n");}
 else{out("no typical getenv() possibilities found.\n");}
}

sub dataexit{outr("cut!\n");data;outq("cut run, finished.\n");}
# (4/4): script init.
if(!$ARGV[0]){outq("syntax: $0 </path/to/binary>\n");}
if(!-f$ARGV[0]){outq("error, binary not found.\n");}
readbinary($ARGV[0]);data;out("clean run, finished.\n");exit(0);
