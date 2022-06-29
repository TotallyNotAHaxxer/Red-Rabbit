package SPOOFER

import (
	"net"

	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	X             error
	handle        *pcap.Handle
	ipAddr        net.IP
	macAddr       net.HardwareAddr
	target        string
	udpLayer      layers.UDP
	dnsLayer      layers.DNS
	ipLayer       layers.IPv4
	ethernetLayer layers.Ethernet
)
