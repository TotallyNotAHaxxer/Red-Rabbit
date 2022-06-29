use strict;
use Tk; 
use Tk::ROText; 

my $CIDR = shift;

sub main_gui_caller {
    my $netrange = shift;
    my $mw = MainWindow->new(
    -title=>"Host Scanner (GUI)",
    -background=>"black",
    -foreground=>"light blue"
    );

    $mw->geometry("600x270+100+300");

    $mw->Button(
        -text=>"Send ICMP",
        -command=>sub{ hosts_ICMP($netrange); }
    )->grid(
        -row=>0,
        -column=>1,
        -sticky=>"w"
    );

    $mw->Button(
        -text=>"Send TCP",
        -command=>sub{ hosts_TCP($netrange); }
    )->grid(
        -row=>0,
        -column=>2,
        -sticky=>"w"
    );

    $mw->Button(
        -text=>"Send UDP",
        -command=>sub{ hosts_UDP($netrange); }
    )->grid(
        -row=>0,
        -column=>3,
        -sticky=>"w"
    );


    $mw->Button(
        -text=>"Send SYN",
        -command=>sub{ hosts_SYN($netrange); }
    )->grid(
        -row=>0,
        -column=>4,
        -sticky=>"w"
    );


    my $oPad = $mw->Scrolled(
        'ROText',
        -scrollbars=>"e",
        -width=>"65",
        -height=>"10",
        -foreground=>"light blue",
        -background=>"black"
    )->grid(
        -row=>1,
        -columnspan=>5,
        -pady=>5,
        -column=>0
    );

    $mw->Button(
        -text=>"Exit",
        -command=>sub{ exit; }
    )->grid(
        -row=>3,
        -column=>5,
        -sticky=>"e"
    );

    MainLoop(); 
    sub hosts_ICMP {
        my $name = shift;
        $oPad->insert("end","Scanning...\n");
        my @hosts = `sudo perl r6.pl -o discover_icmp -t $name`;
        $oPad->insert("end",@hosts);
        return;
    }

    sub hosts_TCP {
        my $name = shift;
        $oPad->insert("end","Scanning...\n");
        my @hosts = `sudo perl r6.pl -o discover_tcp -t $name`;
        $oPad->insert("end",@hosts);
        return;
    }

    sub hosts_SYN {
        my $name = shift;
        $oPad->insert("end","Scanning...\n");
        my @hosts = `sudo perl r6.pl -o discover_syn -t $name`;
        $oPad->insert("end",@hosts);
        return;
    }

    sub hosts_UDP {
        my $name = shift;
        $oPad->insert("end","Scanning...\n");
        my @hosts = `sudo go run arp.go $name`;
        $oPad->insert("end",@hosts);
        return;
    }
}

main_gui_caller($CIDR);
