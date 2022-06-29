package IEEE80211CONSTANTS

import (
	"net"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// PKT         | Defines the packet name for capture
// Wait_group  | Defines the wait group async function
// Waiter      | Defines the time to sleep upon sending out packets
// Pcap_handle | Defines the handler for the packet capture
// Sniffc      | Defines the network interface to use for capture
// IP

var (
	PKT                    gopacket.Packet
	Wait_group             sync.WaitGroup
	Pcap_handle            *pcap.Handle
	Sniffc                 *net.Interface
	Buffer                 = gopacket.NewSerializeBuffer
	ETH_SOURCE_MAC_ADDRESS net.HardwareAddr
	Addresses              *net.IPNet
	WAITER                 = make(chan struct{})
)

const (
	Waiter                    = 15
	NULLMARK_1                = 0xff
	NULLMARK_2                = 0xff
	NULLMARK_3                = 0xff
	NULLMARK_4                = 0xff
	NULLMARK_5                = 0xff
	NULLMARK_6                = 0xff
	NILLMARK_0                = 0x00
	NILLMARK_1                = 0x00
	NILLMARK_2                = 0x00
	NILLMARK_3                = 0x00
	NILLMARK_4                = 0x00
	NILLMARK_5                = 0x00
	NILLBYTE_0                = 0
	LOADDR                    = 127
	ARP_HW_ADDR_SIZE          = 6
	ARP_PR_ADDR_SIZE          = 4
	ARP_OPERATION             = layers.ARPRequest
	ARP_PACKET_PROTOCAL       = layers.EthernetTypeIPv4
	ARP_ADDRESS_PROTOCAL_TYPE = layers.LinkTypeEthernet
	ETH_ETHERNET_TYPE         = layers.EthernetTypeARP
)
