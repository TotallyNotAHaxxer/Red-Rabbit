use strict;
use warnings;
use Tk;
use Tk::Clock;

my $mw = Tk::MainWindow->new();
my $datetime = localtime(); 





sub gui_warn() { 
    if ( $< != 0 ) {
    my $answer = $mw->messageBox(
		-icon => 'question',
        -type => 'yesno',
		-title => 'RR5 warning OUT-SUB-SHOWDIALOG-ERR',
		-message => "WARNING: DURING RUN OF RR5: RUBY: MAIN: =>  Red Rabbit Version 5 requires the user to be root, this is due to the 100+ commands in this script that mainly require root users\n\nTIME OF ERROR => $datetime",
        -height => '200',
        -width => '200',
        -background => 'blue'
	);
    $mw->MainLoop();
    }
}

gui_warn()