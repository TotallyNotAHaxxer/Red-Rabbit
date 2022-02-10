#!/usr/bin/perl
#
# Author   => ArkAngeL43
# Language => Perl
# 
# PACKAGES:
#       Imager::QRCode
#       HTTP::Tiny
#       Term::ANSIColor
#       Image::ExifTools
#       Text::Table
#
#
#
# Program idea: the og idea of this script was to recrusively generate QR codes individually or in a list of URL's by generating the QR codes you are embedding each URL in the file
#
##                                                        =--_
#                                         .-""""""-.     |* _)
#                                        /          \   /  /
#                                       /            \_/  /
#           _                          /|                /
#       _-'"/\                        / |    ____    _.-"            _
#    _-'   (  '-_            _       (   \  |\  /\  ||           .-'".".
#_.-'       '.   `'-._   .-'"/'.      "   | |/ /  | |/        _-"   (   '-_
#             '.      _-"   (   '-_       \ | /   \ |     _.-'       )     "-._
#           _.'   _.-'       )     "-._    ||\\   |\\  '"'        .-'
#         '               .-'          `'  || \\  ||))
#   __  _  ___  _ ____________ _____  ___ _|\ _|\_|\\/ _______________  ___   _
#                       c  c  " c C ""C  " ""  "" ""
#                   c       C
#              C        C
#                   C
#    C     c
#
#
#Powered 
#      By
#        /$$$$$$$                     /$$
#        | $$__  $$                   | $$
#        | $$  \ $$ /$$$$$$   /$$$$$$ | $$
#        | $$$$$$$//$$__  $$ /$$__  $$| $$
#        | $$____/| $$$$$$$$| $$  \__/| $$
#        | $$     | $$_____/| $$      | $$
#        | $$     |  $$$$$$$| $$      | $$
#        |__/      \_______/|__/      |__/
#
#--------------------------------------------------------------------------------
#
#
#
#
# Execution time for 200+ url render and generation; 2.3s
# Execution time for single gen, .1s




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


my $Client = HTTP::Tiny->new();
my $file = "banner.txt";
my $extension = ".png";

my %opts = (
    o => 'QR.png',         
    e => 'M',                   
    s => '5',
    f => '',                   
);

getopt('o:e:s:f:', \%opts);

my $qr = Imager::QRCode->new(
    size =>  $opts{s},
    level => $opts{e},
);






# exif decleration 
my $exif = new Image::ExifTool;
my $info = $exif->ImageInfo($opts{o});
# checking dir
my $filename = $opts{o};
# filelist name 
my $list = $opts{f};

binmode STDOUT, ':encoding(utf8)';


sub banner() {
    print "\x1b[H\x1b[2J\x1b[3J";
    # banner starter
    open(F, '<', $file) or die $!;
    while (<F>) {
        print $_;
    }
    close(F);
}


sub list_qr_gen() {

}

sub filechecking() {
    # if opts{f} is a boolean based if statement while opts{f} is a string name, == is not astatement unless my defines opts as booltstring
    if ($opts{f}) {
        #say "[ INFO ] READ FILE SLEEPING, FOR 5 SECONDS WAITING FOR USER READ";
        #sleep(5);
        open(F, '<'. $list) or die $!;
        # while loop to open the file
        while (<F>) {
            # while file is true and is opened during and reading lined, parse with @url class
            my @urls = (
                $_, 
            );
            for my $url (@urls) {
                # while my urls is true parse each url and test it before being injected into the inital barcode and image
                print "\n\n";
                print color("red"),"#########################################\n", color("reset");
                print color("blue"),color("blink")," Testing URL ~~> $url\n", color("reset");
                print color("red"),"#########################################\n", color("reset");
                my $response = $Client->get($url);
                if($url, $response->{status} == '200'){
                say "\033[37m[ \033[34mINFO \033[37m] \033[32m URL turned with a 200 REQUEST during GET FRAME \n";      
                }
                elsif($url, $response->{status} == '307'){
                    say "\n[ INFO ] DATA: WARN: Temporary Redirection testing again.....";
                }
                else {
                    say "\n[ INFO ] DATA: WARN: im not sure about this response stat if its good or bad?";
                }
            }
            # once done parse the extension and . join the fileextension and the randfomly generated string
            my @set = ('0' ..'9', 'A' .. 'F');
            my $str = join '' => map $set[rand @set], 1 .. 8;
            my $finalgen = $str.$extension;
            # now write and parse the data with the qr code generation to generate it 
            my $main = $qr->plot("$_")->write( file => $finalgen );
            print "\n[ INFO ] DATA: WARN: NAME: WEBSITE => $_ IS EMBEDED IN IMAGE => $finalgen STAT: QR CODE GENERATED......";
            # then we parse all the data of every directory and URL filename into a ANSI table
            say "\033[37m[ \033[34mINFO \033[37m] \033[32m Chcking if $finalgen exists \n";      
            # now parse rthe EXIF 
            my $exif1 = new Image::ExifTool;
            my $infofinalgen = $exif->ImageInfo($finalgen);
            if (-e $finalgen) {
                say "\033[37m[ \033[34mINFO \033[37m] \033[32m File $finalgen exists \n";      
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
                foreach (keys %$infofinalgen) {
                    $tb->load( [1, $_,    $$infofinalgen{$_}] );
                    #print "\033[37m[ \033[34mEXIT DATA \033[37m] \033[32m $_ => $$info{$_}\n";
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
                print color("RED"),color("blink"),"[ INFO ] WARNING: FATAL: => COULD NOT GET EXIF OF $filename \n", color("reset");
                print color("RED"),color("blink"),"[ INFO ] WARNING: FATAL: => $filename SEEMS TO NOT EXIST\n", color("reset");
            }

        }
        close(F);
    }
}

# testing the URL's connection before using it or injecting it into the image 
sub TestVAL() {
    my @urls = (
        @ARGV,
    );

    for my $url (@urls) {
        print "\n\n";
        print color("red"),"#########################################\n", color("reset");
        print color("blue"),color("blink")," Testing URL ~~> $url\n", color("reset");
        print color("red"),"#########################################\n", color("reset");
        my $response = $Client->get($url);
        if($url, $response->{status} == '200'){
        say "\033[37m[ \033[34mINFO \033[37m] \033[32m URL turned with a 200 REQUEST during GET FRAME \n";      
        }
        elsif($url, $response->{status} == '307'){
            say "Temporary Redirection testing again.....";
        }
        else{
            say "im not sure about this";
        }
    }
}


#Get EXIF data from the image 
sub exifdata() {
    say "\033[37m[ \033[34mINFO \033[37m] \033[32m Chcking if $filename exists \n";      
    if (-e $filename) {
        say "\033[37m[ \033[34mINFO \033[37m] \033[32m File $filename exists \n";      
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
            #print "\033[37m[ \033[34mEXIT DATA \033[37m] \033[32m $_ => $$info{$_}\n";
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
        print color("RED"),color("blink"),"[ INFO ] WARNING: FATAL: => COULD NOT GET EXIF OF $filename \n", color("reset");
        print color("RED"),color("blink"),"[ INFO ] WARNING: FATAL: => $filename SEEMS TO NOT EXIST\n", color("reset");
    }
}

sub callqr() {
    my $main = $qr->plot("@ARGV")->write( file => $opts{o} )
}

sub file_exists_before_generation() {
    if (-e $filename) {
        print color("RED"),color("blink"),"\n\t[ INFO ] WARNING: FATAL: => $filename SEEMS TO ALREADY EXIST BEFORE GENERATION...\n", color("reset");
        print color("RED"),color("blink"),"\n\t[ INFO ] WARNING: FATAL: => GENERATING NEW AND RANDOM FILENAME TO PREVENT EXIT CODE 1...\n", color("reset");
        my @set = ('0' ..'9', 'A' .. 'F');
        my $str = join '' => map $set[rand @set], 1 .. 8;
        print color("YELLOW"),color("blink"),"\n\t[ INFO ] WARNING: NEW STRING NAME => $str$extension\n", color("reset");
        say "sleeping for 5 seconds....";
        sleep(5);
        my $str1 = $str.$extension;
        my $main = $qr->plot("@ARGV")->write( file => $str1 )
    }
}
 
sub main() {
    banner();
    TestVAL();
    filechecking();
    file_exists_before_generation();
    callqr();
    exifdata();
}


main();
