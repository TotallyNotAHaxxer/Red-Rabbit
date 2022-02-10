use strict;
use warnings;
use Getopt::Std;

my %opts = (            
);

getopt('t:e:i:', \%opts);

my $target_   = $opts{t};
my $extension = $opts{e};
my $input     = $opts{i};

print "\n Target => $target_ \n Extension => $extension \n input => $input";