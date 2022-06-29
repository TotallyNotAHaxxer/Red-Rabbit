package SUPER_SYN_SCANNER

import (
	"fmt"
	"log"
	"time"

	netreturn "main/modg/scripts/IEEE-802.11/spawners-returners"
	netconst "main/modules/go-main/netc"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func capture(iface, target string) {
	handle, err := pcap.OpenLive(iface, netconst.Snap_Shot_Len, netconst.Promiscuous, netconst.Waiter)
	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(netconst.TCP_FILTER); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("Capturing packets")
	for packet := range source.Packets() {
		networkLayer := packet.NetworkLayer()
		if networkLayer == nil {
			continue
		}
		transportLayer := packet.TransportLayer()
		if transportLayer == nil {
			continue
		}

		srcHost := networkLayer.NetworkFlow().Src().String()
		srcPort := transportLayer.TransportFlow().Src().String()

		if srcHost != target {
			continue
		}
		netconst.TCP_SYN_SCAN_RES[srcPort]++
	}
}

func Runner(host string, l string) {
	a := netreturn.Locate()
	go capture(a, host)
	time.Sleep(netconst.PACKET_WAITE2 * time.Second)
	for port, confidence := range netconst.TCP_SYN_SCAN_RES {
		if confidence >= 1 {
			fmt.Printf("Port %s open (confidence: %d)\n", port, confidence)
		}
	}
}
