#
# Using perl to grab EXIF data and tabulate it into an image 
#
# Why perl? Why anything bud, why anything 
#
#
use strict;
use warnings;
use Getopt::Std;
use Imager::QRCode;
use feature 'say';
use HTTP::Tiny;
use Term::ANSIColor;
use Image::ExifTool;
use utf8;
use Text::Table ();

my %opts = (                            
    f => '',                   
);

getopt('f:', \%opts);
my $filep = $opts{f};
my $exif = new Image::ExifTool;
my $info = $exif->ImageInfo($filep);
my $bnn  = "banner.txt";
binmode STDOUT, ':encoding(utf8)';

sub banner() {
    print "\x1b[H\x1b[2J\x1b[3J";
    open(FA, "<", $bnn) or die $!;
    while (<FA>) {
        print $_;
    }
    close(FA);
} 

sub exifdata() {
    banner();
    say "\033[37m[ \033[34mINFO \033[37m] \033[32m Chcking if $filep exists \n";      
    if (-e $filep) {
        say "\033[37m[ \033[34mINFO \033[37m] \033[32m File $filep exists \n";      
        my @cols = qw/Data/;
        push @cols,
            +{
            title => "After DATA EXIF",
            align => "center",
            };
        my $sep = \'│';
        
        my $major_sep = \'║';
        my $tb        = Text::Table->new( $sep, " Data Number ", $major_sep,
            ( map { +( ( ref($_) ? $_ : " $_ " ), $sep ) } @cols ) );
        
        my $num_cols = @cols;
        # load table 
        foreach (keys %$info) {
            $tb->load( [1, $_,    $$info{$_}] );
        }


        my $make_rule = sub {
            my ($args) = @_;
        
            my $left      = $args->{left};
            my $right     = $args->{right};
            my $main_left = $args->{main_left};
            my $middle    = $args->{middle};
        
            return $tb->rule(
                sub {
                    my ( $index, $len ) = @_;
        
                    return ( '─' x $len );
                },
                sub {
                    my ( $index, $len ) = @_;
        
                    my $char = (
                        ( $index == 0 )             ? $left
                        : ( $index == 1 )             ? $main_left
                        : ( $index == $num_cols + 1 ) ? $right
                        :                               $middle
                    );
        
                    return $char x $len;
                },
            );
        };
        
        # ASCII TC LOAD FLOOR
        my $start_rule = $make_rule->(
            {
                left      => '┌',
                main_left => '╥',
                right     => '┐',
                middle    => '┬',
            }
        );
        
        # ASCII TC LOAD CENTER
        my $mid_rule = $make_rule->(
            {
                left      => '├',
                main_left => '╫',
                right     => '┤',
                middle    => '┼',
            }
        );
        
        # ASCII TC LOAD ROOF
        my $end_rule = $make_rule->(
            {
                left      => '└',
                main_left => '╨',
                right     => '┘',
                middle    => '┴',
            }
        );
        print color("RED"), $start_rule, $tb->title,( map { $mid_rule, $_, } $tb->body() ), $end_rule;
        print color("YELLOW"),color("blink"),"[ WARN ] GENERAL: EXIF DATA COLLECTED: QR CODE GENERATED\n", color("reset");
    } else {
        print color("RED"),color("blink"),"[ INFO ] WARNING: FATAL: => COULD NOT GET EXIF OF $filep \n", color("reset");
        print color("RED"),color("blink"),"[ INFO ] WARNING: FATAL: => $filep SEEMS TO NOT EXIST\n", color("reset");
    }
}

exifdata();