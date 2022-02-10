require './net-module.rb'
require 'optiflag'

module Inputs extend OptiFlagSet
    flag "tport" # ectend TCP port with flags
    and_process!
end


tport = ARGV.flags.tport
if ARGV.flags.tport.empty?
    puts " TCP port disabledd "
end

def n(backslash)
    puts backslash
end

def main(tcp_port)
    rev       = "\033[0;39m"
    reb       = "\033[49m"
    blk       = "\033[0;30m"
	red       = "\033[0;31m"
	grn       = "\033[0;32m"
	yel       = "\033[0;33m"
	blu       = "\033[0;34m"
	mag       = "\033[0;35m"
	cyn       = "\033[0;36m"
	wht       = "\033[0;37m"
    blkb      = "\033[40m"
	redb      = "\033[41m"
	grnb      = "\033[42m"
	yelb      = "\033[43m"
	blub      = "\033[44m"
	magb      = "\033[45m"
	cynb      = "\033[46m"
	whtb      = "\033[47m"
    stats = TCPdump.new
    # packetfu connection and configuration
    iface = PacketFu::Utils.default_int
    my_ip = PacketFu::Utils.ifconfig(iface)[:ip_saddr]
    tcp_start_config = PacketFu::Capture.new(:iface => iface)
    tcp_start_config.bpf(:iface=> iface, :promisc => true, :filter => "ip and tcp port #{tcp_port}")
    #
    # start the dump
    tcp_start_config.start
    # for each packet format the data
    pack = 0
    tcp_start_config.stream.each do | packet |
        n("\n\n")
        pack += 1
        info = PacketFu::Packet.parse(packet)
        stats.process_connection(:source => info.ip_saddr, :destination => info.ip_daddr)
        puts stats.stats
        t = Time.now
        puts(wht, "[" + redb + t.strftime("%I:%M %p") + wht + "]" + wht + "[" + blub + " INFO " + wht + reb + "]"  + grn + " Packets captured => " + pack.to_s)
        puts(wht, "[" + redb + t.strftime("%I:%M %p") + wht + "]" + wht + "[" + blub + " INFO " + wht + reb + "]"  + blu, :source => info.ip_saddr.to_s )
    end
end


main("#{tport}")