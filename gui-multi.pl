#!/usr/bin/perl -w
use strict;
use Tk;
use Tk::NoteBook;
use Tk::ROText;
# nyan cat theme
use Tk::widgets qw(Photo Animation);
use OFSL;
use IO::Socket;


my $base_text_color         = "red";
my $back                    = "black";
my $Back_Page_Color         = "black";
my $base_text_color_backing = "red";




my $mw = MainWindow->new(
    -title=>"RR6 Tabbed Enviroment (Perl5)",
    -background=>"black",
    -foreground=>"light blue"
);

$mw->geometry("900x900+100+300");
my $book = $mw->NoteBook(
    -ipadx=>0,
    -foreground=>$base_text_color,
    -background=>$back,
    -backpagecolor=>$back,
    -inactivebackground=>$Back_Page_Color
)->grid(
    -row=>0,
    -column=>0,
    -sticky=>"w",
    -pady=>10
);


my $tab1       = $book->add("Sheet 1", -label=>"Test port status");
my $tab2       = $book->add("Sheet 2", -label=>"OUI Lookup");
my $tab3       = $book->add("Sheet 3", -label=>"Host Ping ICMP");
my $tab4       = $book->add("Sheet 4", -label=>"Host Ping SYN");
my $tab5       = $book->add("Sheet 5", -label=>"Host Ping UDP");
my $tab6       = $book->add("Sheet 6", -label=>"Host Ping TCP");
my $tab7       = $book->add("Sheet 7", -label=>"GIF Injection");
my $tab8       = $book->add("Sheet 8", -label=>"BMP Injection");
my $tab9       = $book->add("Sheet 9", -label=>"WEBP Injection");
my $tab10      = $book->add("Sheet 10", -label=>"JPG Injection");
my $file       = 'gui-conf/nyan.gif';
my $img        = $mw->Animation('-format' => 'gif',-file => $file);
my $lab        = $mw->Label(-image => $img);
my @months = qw( Jan Feb Mar Apr May Jun Jul Aug Sep Oct Nov Dec );
my @days = qw(Sun Mon Tue Wed Thu Fri Sat Sun);
my  ($sec,$min,$hour,$mday,$mon,$year,$wday,$yday,$isdst) = localtime();



# settings
my ($mac,$port,$host,$CIDR,$image,$payload,$location) = ""x7;

my $oPad = $mw->Scrolled(
    'ROText',
    -scrollbars=>"e",
    -width=>"65",
    -height=>"60",
    -foreground=>"blue",
    -background=>"black")->grid(
        -row=>1,
        -columnspan=>5,
        -pady=>5,
        -column=>0);






# label for image entry [ JPG ]
$tab10->Label(
    -text=> "Image")->grid(
        -row=>0,
        -column=>0,
        -sticky=>"w");

# image entry [ JPG ]
$tab10->Entry(
    -textvariable=>\$image)->grid(
        -row=>0,
        -column=>1,
        -sticky=>"w",
        -pady=>5);

# payload entry label [ JPG ] 
$tab10->Label(
    -text=> "Payload")->grid(
        -row=>1,
        -column=>2,
        -sticky=>"w");

# payload entry [ JPG ] 
$tab10->Entry(
    -textvariable=>\$payload)->grid(
        -row=>1,
        -column=>1,
        -sticky=>"w",
        -pady=>5);

# inject image [ JPG ] button
$tab10->Button(
    -text=>"Inject Image",
    -command=>\&jpgjection,)->grid(
        -row=>3,
        -column=>0,
        -sticky=>"we",
        -pady=>5);









# label for image entry [ WEBP ]
$tab9->Label(
    -text=> "Image")->grid(
        -row=>0,
        -column=>0,
        -sticky=>"w");

# image entry [ WEBP ]
$tab9->Entry(
    -textvariable=>\$image)->grid(
        -row=>0,
        -column=>1,
        -sticky=>"w",
        -pady=>5);

# payload entry label [ WEBP ] 
$tab9->Label(
    -text=> "Payload")->grid(
        -row=>1,
        -column=>2,
        -sticky=>"w");

# payload entry [ WEBP ] 
$tab9->Entry(
    -textvariable=>\$payload)->grid(
        -row=>1,
        -column=>1,
        -sticky=>"w",
        -pady=>5);

# inject image [ WEBP ] button
$tab9->Button(
    -text=>"Inject Image",
    -command=>\&webpjection,)->grid(
        -row=>3,
        -column=>0,
        -sticky=>"we",
        -pady=>5);










# label for image entry [ BMP ]
$tab8->Label(
    -text=> "Image")->grid(
        -row=>0,
        -column=>0,
        -sticky=>"w");

# image entry [ BMP ]
$tab8->Entry(
    -textvariable=>\$image)->grid(
        -row=>0,
        -column=>1,
        -sticky=>"w",
        -pady=>5);

# payload entry label [ BMP ] 
$tab8->Label(
    -text=> "Payload")->grid(
        -row=>1,
        -column=>2,
        -sticky=>"w");

# payload entry [ BMP ] 
$tab8->Entry(
    -textvariable=>\$payload)->grid(
        -row=>1,
        -column=>1,
        -sticky=>"w",
        -pady=>5);

# inject image [ BMP ] button
$tab8->Button(
    -text=>"Inject Image",
    -command=>\&bmpjection,)->grid(
        -row=>3,
        -column=>0,
        -sticky=>"we",
        -pady=>5);









# label for image entry
$tab7->Label(
    -text=> "Image")->grid(
        -row=>0,
        -column=>0,
        -sticky=>"w");

# image entry 
$tab7->Entry(
    -textvariable=>\$image)->grid(
        -row=>0,
        -column=>1,
        -sticky=>"w",
        -pady=>5);

# payload entry label
$tab7->Label(
    -text=> "Payload")->grid(
        -row=>1,
        -column=>2,
        -sticky=>"w");

# payload entry
$tab7->Entry(
    -textvariable=>\$payload)->grid(
        -row=>1,
        -column=>1,
        -sticky=>"w",
        -pady=>5);

# inject image button
$tab7->Button(
    -text=>"Inject Image",
    -command=>\&gifjection,)->grid(
        -row=>3,
        -column=>0,
        -sticky=>"we",
        -pady=>5);










# label for perl port check
$tab1->Label(
    -text=>"Port")->grid(
        -row=>0,
        -column=>0,
        -sticky=>"w",
        -pady=>5);


# entry for port checker (port)
$tab1->Entry(
    -textvariable=>\$port)->grid(
        -row=>0,
        -column=>1,
        -sticky=>"w");

$tab1->Label(
    -text=>"Host")->grid(
        -row=>1,
        -column=>0,
        -sticky=>"w",
        -pady=>5);


# entry for port checker (host)
$tab1->Entry(
    -textvariable=>\$host)->grid(
        -row=>1,
        -column=>1,
        -sticky=>"w");


$tab1->Button(
    -text=>"Test Port",
    -command=>\&port_s_check,)->grid(
        -row=>3,
        -column=>0,
        -sticky=>"we",
        -pady=>5);


# label for oui lookup
$tab2->Label(
    -text=>"MAC Address ")->grid(
        -row=>0,
        -column=>0,
        -sticky=>"w",
        -pady=>5);

# entry for oui lookup
$tab2->Entry(
    -textvariable=>\$mac)->grid(
        -row=>0,
        -column=>1,
        -sticky=>"w");



$tab2->Button(
    -text=>"Scan",
    -command=>\&lookup_oui,)->grid(
        -row=>3,
        -column=>0,
        -sticky=>"we",
        -pady=>5);














# label for ICMP ping
$tab3->Label(
    -text=>"Network Range -> ")->grid(
        -row=>0,
        -column=>0);

# entry for ICMP ping
$tab3->Entry(
    -textvariable=>\$CIDR)->grid(
        -row=>0,
        -column=>1,
        -padx=>5,
        -pady=>5);

# ICMP ping activation 
$tab3->Button(
    -text=>"Scan",
    -command=>\&hosts_ICMP,
    )->grid(
        -row=>2,
        -column=>0,
        -pady=>5,
        -sticky=>"ew"
);










# label for host ping SYN
$tab4->Label(
        -text=>"Network Range -> ")->grid(
        -row=>0,
        -column=>0);

# entry for host ping SYN
$tab4->Entry(
    -textvariable=>\$CIDR)->grid(
        -row=>0,
        -column=>1,
        -padx=>5,
        -pady=>5);

#button for host ping SYN
$tab4->Button(
    -text=>"Scan",
    -command=>\&hosts_SYN,
    )->grid(
        -row=>2,
        -column=>0,
        -pady=>5,
        -sticky=>"ew"
);









# label for host ping UDP
$tab5->Label(
        -text=>"Network Range -> ")->grid(
        -row=>0,
        -column=>0);

# entry for host ping UDP
$tab5->Entry(
    -textvariable=>\$CIDR)->grid(
        -row=>0,
        -column=>1,
        -padx=>5,
        -pady=>5);

#button for host ping UDP
$tab5->Button(
    -text=>"Scan",
    -command=>\&hosts_UDP,
    )->grid(
        -row=>2,
        -column=>0,
        -pady=>5,
        -sticky=>"ew"
);












# label for host ping TCP
$tab6->Label(
        -text=>"Network Range -> ")->grid(
        -row=>0,
        -column=>0);

# entry for host ping TCP
$tab6->Entry(
    -textvariable=>\$CIDR)->grid(
        -row=>0,
        -column=>1,
        -padx=>5,
        -pady=>5);

#button for host ping TCP
$tab6->Button(
    -text=>"Scan",
    -command=>\&hosts_TCP,
    )->grid(
        -row=>2,
        -column=>0,
        -pady=>5,
        -sticky=>"ew"
);

# exit sub caller
$mw->Button(
    -text=>"Exit",
    -command=>sub{ exit; },
    -background=>"red"
    )->grid(
        -row=>0,
        -column=>-0,
        -pady=>1,
        -sticky=>"e",
);

$mw->Button(
    -text=>'Spawn the CAT',
    -command=>\&anime,
    -background=>"blue",
    )->grid(
        -row=>1,
        -column=>-0,
        -pady=>1,
        -sticky=>"e"
);


$oPad->insert("end","!!!!!!!!!!!!!!!!!!!!!!dont spawn the cat!!!!!!!!!!!!!!!!!!!!!\n");
$oPad->insert("end","-------------------------------------------------------------\n");
$oPad->insert("end","\n");


MainLoop();


# routines and calls

#FC:58:FA:D2:5C:83
sub oui_lookup {
  $oPad->insert("end","\n");
  my $mac_address = shift;
  $mac_address =~ s/:/-/g;
  my $oui = substr $mac_address, 0, 8;
  open (my $oui_file, '<', "config/oui.txt") or die $!;
  while (my $line = <$oui_file>)
  {
    if($line =~ /$oui/i)
    {
      my ($address, $manufacturer_name) = split /\t+/, $line;
      return "\nAddr -> $mac_address \nManufacturer -> $manufacturer_name";
      last;
    }
  }
  return "Unknown";
}


sub hosts_ICMP {
    $oPad->insert("end","\n");
    $oPad->insert("end","Scanning...\n");
    my @hosts = `sudo perl r6.pl -o discover_icmp -t $CIDR`;
    $oPad->insert("end",@hosts);
    return;
}

sub hosts_TCP {
    $oPad->insert("end","\n");
    $oPad->insert("end","Scanning...\n");
    my @hosts = `sudo perl r6.pl -o discover_tcp -t $CIDR`;
    $oPad->insert("end",@hosts);
    return;
}

sub hosts_SYN {
    $oPad->insert("end","\n");
    $oPad->insert("end","Scanning...\n");
    my @hosts = `sudo perl r6.pl -o discover_syn -t $CIDR`;
    $oPad->insert("end",@hosts);
    return;
}

sub hosts_UDP {
    $oPad->insert("end","\n");
    $oPad->insert("end","Scanning...\n");
    my @hosts = `sudo perl r6.pl -o discover_udp -t $CIDR`;
    $oPad->insert("end",@hosts);
    return;
}

sub lookup_oui(){
    $oPad->insert("end","\n");
    my @hosts = oui_lookup($mac);
    $oPad->insert("end",@hosts) if(scalar(@hosts)>0);
    return;
}

sub port_s_check(){
    $oPad->insert("end","\n");
    my @ret = `sudo perl r6.pl -o slashc -e $port -r $host`;
    $oPad->insert("end","Checking Host | $host:$port | \n");
    $oPad->insert("end",@ret) if(scalar(@ret)>0);
    return;
}

sub gifjection(){
    $oPad->insert("end","\n");
    my @ret = `sudo perl r6.pl -o gif -f $image -p $payload`;
    $oPad->insert("end","Injecting Payload -> | $payload | \n");
    $oPad->insert("end","Injecting Image   -> | $image   | \n");
    $oPad->insert("end",@ret) if(scalar(@ret)>0);
    return;
}

sub webpjection(){
    $oPad->insert("end","\n");
    my @ret = `sudo perl r6.pl -o webp -f $image -p $payload`;
    $oPad->insert("end","Injecting Payload -> | $payload | \n");
    $oPad->insert("end","Injecting Image   -> | $image   | \n");
    $oPad->insert("end",@ret) if(scalar(@ret)>0);
    return;
}

sub jpgjection(){
    $oPad->insert("end","\n");
    my @ret = `sudo perl r6.pl -o jpg -f $image -p $payload`;
    $oPad->insert("end","Injecting Payload -> | $payload | \n");
    $oPad->insert("end","Injecting Image   -> | $image   | \n");
    $oPad->insert("end",@ret) if(scalar(@ret)>0);
    return;
}

sub bmpjection(){
    $oPad->insert("end","\n");
    my @ret = `sudo perl r6.pl -o bmp -f $image -p $payload`;
    $oPad->insert("end","Injecting Payload -> | $payload | \n");
    $oPad->insert("end","Injecting Image   -> | $image   | \n");
    $oPad->insert("end",@ret) if(scalar(@ret)>0);
    return;
}

sub anime() {
    $oPad->insert("end","\n");
    Tk::grid($lab,'-','-',-sticky => 'nsew');
    $mw->gridRowconfigure(0,-weight => 1);
}
