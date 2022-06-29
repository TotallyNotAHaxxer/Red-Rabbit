use strict;
use warnings;
use Digest::MD5 qw(md5_hex);

# define an empty array for push
my @hashes;

sub open_and_dump {
    my $file = shift;
    my $dictionary = shift;
    open(my $QLF, $file) or die $!;
    my @hf = $QLF;
    foreach my $line (@hf) {
        if ($line =~ m/'([0-9a-zA-Z]+)\'\, \'([0-9a-fA-F]{32})\'/g) {
            push(@hashes, "$1:$2");
            print "<RR6> SQL Dumper: Found hash $1:$2\n";
        }
    }
    print "<RR6> SQL Dumper: Found hashes -> ".($#hashes + 1).".\n";
    # crack hashes
    open(my $D, $dictionary) or die $!;
    foreach (@hashes) {
        my ($user, $hash) = split(":", $_);
        my $foundhash = 0;
        seek(D, 0, 0);
        while (chomp(my $line = <D>)) {
            if ($hash eq md5_hex($line)) {
                print "<RR6> SQL Dumper: Cracked hash for user $user : $line\n";
                $foundhash = 1;
                my $num_cracked++;
                last;
            }
        }
    }
    close(D);
    print "<RR6> Stat Module: Attempted to crack all hashes [ MD5 ] \n"
}
