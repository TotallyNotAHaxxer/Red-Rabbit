package arp_reader

import (
	"bytes"
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/macs"
	"github.com/google/gopacket/pcap"

	ARP_CONSTANTS "main/modg/scripts/IEEE-802.11/IEEE-802.11-c"
)

func Output_Packet_Information(handle *pcap.Handle, iface *net.Interface, stop chan struct{}) {
	counter := 0
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		counter++
		select {
		case <-stop:
			return
		case ARP_CONSTANTS.PKT = <-in:
			arpLayer := ARP_CONSTANTS.PKT.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp := arpLayer.(*layers.ARP)
			if arp.Operation != layers.ARPReply || bytes.Equal([]byte(iface.HardwareAddr), arp.SourceHwAddress) {
				continue
			}
			ARP_SOURCE_ADDR := net.IP(arp.SourceProtAddress)
			ARP_SOURCE_HWADDR := net.HardwareAddr(arp.SourceHwAddress)
			var alist []string
			if mac, err := net.ParseMAC(ARP_SOURCE_HWADDR.String()); err == nil {
				prefix := [3]byte{
					mac[0],
					mac[1],
					mac[2],
				}
				manufacturer, e := macs.ValidMACPrefixMap[prefix]
				if e {
					alist = append(alist, manufacturer)
				}
				fmt.Printf("\033[4;32mIPA:  %s \t\033[4;34mMAC: %v \t\033[4;37mOUI: %v   \033[4;39m\n", ARP_SOURCE_ADDR, ARP_SOURCE_HWADDR, alist)
			}
		}
	}
}
