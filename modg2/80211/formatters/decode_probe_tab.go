package FORMATION

import (
	"bytes"
	"fmt"
	m "main/modg2/80211/80211_oui"
	"strings"
	"time"

	v "main/modg2/80211/80211_constants"

	"github.com/google/gopacket"
	ll "github.com/google/gopacket/layers"
)

type Beacon struct {
	SSID           string
	BSSID          string
	MAC            string
	Freq           int8
	Chan           string
	Rate           string
	VendorSpecific []byte

	Time time.Time
}

func (b *Beacon) Deocde_Beacon(l *ll.Dot11MgmtProbeReq) {
	v.PB = l.LayerContents()
	for k := uint64(0); k < uint64(len(v.PB)); {
		id := ll.Dot11InformationElementID(v.PB[k])
		k++
		switch id {
		case ll.Dot11InformationElementIDSSID:
			e := uint64(v.PB[k])
			k++
			if e > 0 {
				b.SSID = string(v.PB[k : k+e])
				k += e
			}
			break
		case ll.Dot11InformationElementIDVendor:
			b.VendorSpecific = v.PB[k+1:]
			return
		default:
			e := uint64(v.PB[k])
			k += 1 + e
			break
		}
	}
}

func Pack(p gopacket.Packet) {

	typed := Beacon{Time: time.Now()}
	if lay := p.Layer(ll.LayerTypeDot11); lay != nil {
		q, _ := lay.(*ll.Dot11)
		typed.MAC = q.Address2.String()
		if j := p.Layer(ll.LayerTypeDot11MgmtProbeReq); j != nil {
			d, _ := j.(*ll.Dot11MgmtProbeReq)
			typed.Deocde_Beacon(d)
		}
		if la := p.Layer(ll.LayerTypeRadioTap); la != nil {
			r, _ := la.(*ll.RadioTap)
			typed.Freq = r.DBMAntennaSignal
			typed.Rate = r.Rate.String()
		}
	}
	addr4 := m.OUI(typed.MAC)
	if addr4 != nil {
		if len(addr4) <= 13 {
			n := 13
			addr := addr4
			s := Trim(addr[0], n)
			if len(s[0]) == 13 {
				if typed.MAC != "" && typed.Freq != 0x00 && typed.SSID != "" {
					fmt.Printf("│ %s│   %d   │  %s\t │   %s\t│  %v\n", typed.MAC, typed.Freq, typed.Rate, s[0], typed.SSID)
				}
			} else {
				a := 13 - len(s[0])
				chars := strings.Repeat("*", a)
				c := s[0] + chars
				if typed.MAC != "" && typed.Freq != 0x00 && typed.SSID != "" {
					fmt.Printf("│ %s│   %d   │  %s\t │   %s\t│  %v\n", typed.MAC, typed.Freq, typed.Rate, c, typed.SSID)
				}
			}

		}
	}

}

func Trim(m string, s int) []string {
	b := ""
	bb := []string{}
	r := bytes.Runes([]byte(m))
	k := len(r)
	for i, h := range r {
		b = b + string(h)
		if (i+1)%s == 0 {
			bb = append(bb, b)
		} else if (i + 1) == k {
			bb = append(bb, b)
		}
	}
	return bb
}
