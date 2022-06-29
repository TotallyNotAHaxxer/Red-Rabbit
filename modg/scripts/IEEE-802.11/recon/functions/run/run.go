package run

import (
	"errors"
	"fmt"
	"net"
	"time"

	ARP_COLORS "main/modg/colors"
	ARP_CONSTANTS "main/modg/scripts/IEEE-802.11/IEEE-802.11-c"
	ARP_ERRORS "main/modg/scripts/IEEE-802.11/recon/functions/errors"
	ARP_READERS "main/modg/scripts/IEEE-802.11/recon/functions/reader"
	ARP_WRITERS "main/modg/scripts/IEEE-802.11/recon/functions/writer"

	"github.com/google/gopacket/pcap"
)

func Run(iface *net.Interface) error {
	var r *net.IPNet
	if d, x := iface.Addrs(); x != nil {
		return x
	} else {
		for _, l := range d {
			if inet, k := l.(*net.IPNet); k {
				if i4 := inet.IP.To4(); i4 != nil {
					r = &net.IPNet{IP: i4, Mask: inet.Mask[len(inet.Mask)-4:]}
					break
				}
			}
		}
	}
	if r == nil {
		return errors.New("<RR6> Network: Skipping network card, value nil > empty")
	} else if r.IP[0] == 127 {
		return errors.New("<RR6> Network: Skipping network card, Localhost detected")
	} else if r.Mask[0] != 0xff || r.Mask[1] != 0xff {
		return errors.New("<RR6> Network: Skipping network card, mask is above legal limit")
	}
	p, x := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	ARP_ERRORS.Return_err(x, "<RR6> Network Module: Could not open up the network capture handle got error -> ", ARP_COLORS.REDB)
	defer p.Close()
	go ARP_READERS.Output_Packet_Information(p, iface, ARP_CONSTANTS.WAITER)
	defer close(ARP_CONSTANTS.WAITER)
	for {
		if x := ARP_WRITERS.WRITER(p, iface, r); x != nil {
			fmt.Println(ARP_COLORS.REDHB, "<RR6> Writer: Could not write out the packets to the specified interface | ", iface.Name, " | Got error ", x)
			return x
		}
		time.Sleep(ARP_CONSTANTS.Waiter * time.Second)
	}
}
