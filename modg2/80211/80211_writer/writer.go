package E_Writer

import (
	cc "main/modg2/80211/80211_color"
	cq "main/modg2/80211/80211_constants"
	format "main/modg2/80211/80211_formatter"
	ct "main/modg2/80211/80211_types"

	"github.com/google/gopacket"
)

var (
	packet_name       gopacket.LayerType
	packet_layer_name string
)

var arg string

var config = ct.Color{
	Beginning_Color: cc.RED,
	Beginning_ASCII: "|",
	Middle_Color:    cc.HIGH_BLUE,
	Middle_ASCII:    "+",
	End_Color:       cc.RED,
	End_ASCII:       "|",
}

func Process(t string, pkt gopacket.Packet) {
	switch t {

	case "IPV6":
		packet_name = cq.ICMPV6_
		packet_layer_name = "TCP"
	case "dot11":
		packet_name = cq.DOT11_plain
		packet_layer_name = "DOT11"
	case "Probe_request":
		packet_name = cq.DOT11_MGMT_Probe_Request
		packet_layer_name = "Probe Request"
	case "table":
		packet_name = cq.DOT11_MGMT_Probe_Request
		packet_layer_name = "Probe Request table"
		arg = "table"
	case "*":
		packet_name = cq.DOT11_plain
		arg = "func"
	case "air":
		packet_name = cq.DOT11_plain
		arg = "air"
	}
	format.Output_Results(pkt, packet_name, packet_layer_name, config, "(Some Protocal)", arg)

}
