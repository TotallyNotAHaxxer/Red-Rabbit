package C_Listener

import (
	xcol "main/modg2/80211/80211_color"
	data "main/modg2/80211/80211_constants"
	xerr "main/modg2/80211/80211_errors"
	ypar "main/modg2/80211/80211_writer"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Run(packet_type string) {
	data.Controller, data.X = pcap.OpenLive(data.Interface, data.Snaplen, data.Monitor, data.Timeout)
	xerr.Error(data.X, "<RR6> Controller: Failed to launch or create a new live capture on card, got error -> ", xcol.CYN, 255, 1, true)
	defer data.Controller.Close()
	s := gopacket.NewPacketSource(data.Controller, data.Controller.LinkType())
	counter := 0
	for p := range s.Packets() {
		counter++
		ypar.Process(packet_type, p)
	}
}
