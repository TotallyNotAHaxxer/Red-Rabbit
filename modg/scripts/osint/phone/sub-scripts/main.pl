use warnings;
use strict;

my $phone_num = shift;
print $phone_num;
system("chmod +x ./modg/scripts/osint/phone/sub-scripts/sub.sh && ./modg/scripts/osint/phone/sub-scripts/sub.sh --num $phone_num");