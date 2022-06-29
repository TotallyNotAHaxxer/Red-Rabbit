package constants

import "github.com/google/gopacket/pcap"

const (
	Promiscuous     = false
	Snap_Shot_Len   = int32(320)
	Waiter          = pcap.BlockForever
	TCP_FILTER      = "tcp[13] == 0x11 or tcp[13] == 0x10 or tcp[13] == 0x18"
	PACKET_WAITER   = 5
	PACKET_WAITE2   = 2
	DIAL_TIMEOUT_MS = 1000
	DIAL_METHODIZED = "tcp"
	DEVICES_FOUND   = false
)

var (
	TCP_SYN_SCAN_RES = make(map[string]int)
)
