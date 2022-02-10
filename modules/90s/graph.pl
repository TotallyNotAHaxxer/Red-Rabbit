
use strict;
use warnings;

# use POSIX qw( strftime );
use Tk;
use Tk::Clock;

my $m = MainWindow->new (-title => "World Clock");

sub popup() {
   my $answer = $m->Dialog(-title => 'NOTICE', 
    -text => 'would you like to continue'
    -default_button => 'yay', -buttons => [ 'yay', 'nay'], 
    -bitmap => 'question')->Show( );
    if ($answer eq 'yay') {
        print "hello";
    }
}

popup();
my %TZ;
while (<DATA>) {
    if (m/^(\S+)\s+(GMT\S*)\s+(.*)$/) {
	$TZ{$1}   = [ $1, $2, undef, $3 ];
	$TZ{$2} ||= [ $1, $2, undef, $3 ];
	}
    if (m/^(.*?)(?:\s+\*)?\s+(GMT\S*)$/) {
	$TZ{$2} ||= [ $2, $2, undef, "" ];
	$TZ{$2}[2] ||= $1;
	exists $TZ{$TZ{$2}[1]} and
	    $TZ{$TZ{$2}[1]}[2] ||= $1;
	}
    }

# my $tz = strftime ("%Z", localtime);
foreach my $cd (
	[ "UCT",		"UCT",			"#ff0040"	],
	[ "Local",		$ENV{TZ}||"",		"Red"		],
	[ "London",		"Europe/London",	"OrangeRed"	],
	[ "Amsterdam",		"Europe/Amsterdam",	"Orange"	],
	[ "Moscow",		"Europe/Moscow",	"Yellow"	],
	[ "Tokyo",		"Asia/Tokyo",		"YellowGreen"	],
	[ "Los Angeles",	"America/Los_Angeles",	"Green"		],
	[ "New York",		"America/New_York",	"Turquoise"	],
	[ "Darwin",		"Australia/Darwin",	"Blue"		],
	[ "Catham",		"GMT+13:45",		"Violet"	],
	) {
    my ($city, $tz, $color) = @$cd;
    if (exists $TZ{$tz}) {
	$tz = $TZ{$tz}[1];
	$city = $TZ{$tz}[2];
	}

    my $c = $m->Clock (-background => "Black");
    $c->config (
	anaScale	=> 200,
	secsColor	=> "Green",
	tickColor	=> "Blue",
	tickFreq	=> 1,
	timeFont	=> "{Liberation Mono} 12",
	timeColor	=> "lightBlue",
	timeFormat	=> "HH:MM:SS",
	dateFont	=> "{Liberation Mono} 12",
	dateColor	=> "Gold",

	dateFormat => $city,
	timeZone   => $tz,
	handColor  => $color,
	);
    $c->pack (-side => "left", -expand => 1, -fill => "both");
    }

MainLoop;