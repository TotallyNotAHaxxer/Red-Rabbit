package Data_Table; 

use strict;
use warnings;
use feature 'say';
use utf8; # UTF8 formatting binmode is needed 
use Text::Table ();


binmode STDOUT, ':encoding(utf8)';

sub table{
    my $filename   = shift; # filename for output, this will be the data
    my $color      = shift; # color for output
    my $header     = shift; # table header name
    my @cols = qw//;
    push @cols,
        +{
        title => "$header",
        align => "center",
        };
    my $sep = \'│';
    
    my $major_sep = \'║';
    my $tb        = Text::Table->new( $sep, "Column number", $major_sep,
        ( map { +( ( ref($_) ? $_ : " $_ " ), $sep ) } @cols ) );
    
    my $num_cols = @cols;
    open(FH, '<', $filename) or die $!;
    while(<FH>) {
        $tb->load( [1, $_] );
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
    my $start_rule = $make_rule->(
        {
            left      => '┌',
            main_left => '╥',
            right     => '┐',
            middle    => '┬',
        }
    );
    my $mid_rule = $make_rule->(
        {
            left      => '├',
            main_left => '╫',
            right     => '┤',
            middle    => '┼',
        }
    );
    my $end_rule = $make_rule->(
        {
            left      => '└',
            main_left => '╨',
            right     => '┘',
            middle    => '┴',
        }
    );
    print $color, $start_rule, $tb->title,( map { $mid_rule, $_, } $tb->body() ), $end_rule;
}

