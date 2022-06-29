package H_Formation

import (
	"fmt"
	t "main/modg2/80211/80211_types"

	cc "main/modg2/80211/80211_color"
	ct "main/modg2/80211/80211_types"
	sl "main/modg2/80211/formatters"

	"github.com/google/gopacket"
	l "github.com/google/gopacket/layers"
)

func Output_Results(pkt gopacket.Packet, pkt_layer gopacket.LayerType, pkt_layer_name string, c t.Color, n string, ea string) {
	TCP := pkt.Layer(pkt_layer)
	tcp, _ := TCP.(*l.ICMPv6)
	DOT11 := pkt.Layer(pkt_layer)
	dot11, _ := DOT11.(*l.Dot11)
	PROBE := pkt.Layer(pkt_layer)
	probe, _ := PROBE.(*l.Dot11MgmtProbeReq)
	RADIO := pkt.Layer(pkt_layer)
	radio, _ := RADIO.(*l.Dot11MgmtProbeReq)
	var config = ct.Color{
		Beginning_Color: cc.RED,
		Beginning_ASCII: "|",
		Middle_Color:    cc.HIGH_BLUE,
		Middle_ASCII:    "+",
		End_Color:       cc.RED,
		End_ASCII:       "|",
	}
	if tcp != nil {
		sl.Parse_ICMP_test(tcp, config)
	}
	if dot11 != nil {
		if ea == "air" {
			sl.Caller(pkt)
		}
		if ea == "func" {
			fmt.Println(pkt)
		}
		if ea != "air" && ea != "func" {
			sl.Parse_Dot11_normal(dot11, config)
		}
	}
	if probe != nil {
		if ea != "table" {
			sl.Parse(pkt, config)
		}
	}
	if radio != nil {
		if ea == "table" {
			sl.Pack(pkt)
		}
	}

}
