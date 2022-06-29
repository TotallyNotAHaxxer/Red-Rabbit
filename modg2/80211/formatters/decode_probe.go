//apt install libpcap0.8 libpcap0.8-dev
package FORMATION

import (
	"fmt"
	"time"

	//modg2/80211
	lc "main/modg2/80211/80211_color"
	v "main/modg2/80211/80211_constants"
	m "main/modg2/80211/80211_oui"

	cc "main/modg2/80211/80211_types"

	"github.com/google/gopacket"
	ll "github.com/google/gopacket/layers"
)

type REQ struct {
	MAC                       string
	SSID                      string
	Radio_Tap_Frequency       string
	Radio_Tap_Flags           string
	Radio_Tap_Chan            bool
	Radio_Tap_Present_Rate    string
	Radio_Tap_DBAntennaSignal uint8
	Radio_Tap_DBAntennaNoise  uint8
	Radio_Tap_Lock_Quality    uint16
	Radio_Tap_Version         uint8
	OUI                       []byte
	Info                      []byte
	RSSI                      int8
	VendorSpecific            []byte
	CaptureDTS                time.Time
}

func (PROBE *REQ) Decoder(l *ll.Dot11MgmtProbeReq) {
	v.PB = l.LayerContents()
	for k := uint64(0); k < uint64(len(v.PB)); {
		id := ll.Dot11InformationElementID(v.PB[k])
		k++
		switch id {
		case ll.Dot11InformationElementIDSSID:
			e := uint64(v.PB[k])
			k++
			if e > 0 {
				PROBE.SSID = string(v.PB[k : k+e])
				k += e
			}
			break
		case ll.Dot11InformationElementIDVendor:
			PROBE.VendorSpecific = v.PB[k+1:]
			return

		default:
			e := uint64(v.PB[k])
			k += 1 + e
			break
		}
	}
}

func Parse(lp gopacket.Packet, c cc.Color) {
	r := REQ{
		CaptureDTS: time.Now(),
	}
	if k := lp.Layer(ll.LayerTypeDot11); k != nil {
		q, _ := k.(*ll.Dot11)
		r.MAC = q.Address2.String()
		if j := lp.Layer(ll.LayerTypeDot11MgmtProbeReq); j != nil {
			d, _ := j.(*ll.Dot11MgmtProbeReq)
			r.Decoder(d)
		}
		if o := lp.Layer(ll.LayerTypeRadioTap); o != nil {
			h, _ := o.(*ll.RadioTap)
			r.RSSI = h.DBMAntennaSignal
			r.Radio_Tap_Frequency = h.ChannelFrequency.String()
			r.Radio_Tap_Version = h.Version
			r.Radio_Tap_Lock_Quality = h.LockQuality
			r.Radio_Tap_DBAntennaNoise = h.DBAntennaNoise
			r.Radio_Tap_DBAntennaSignal = h.DBAntennaSignal
			r.Radio_Tap_Present_Rate = h.Rate.String()
			r.Radio_Tap_Chan = h.Present.Channel()
			r.Radio_Tap_Flags = h.Flags.String()
		}
		if oo := lp.Layer(ll.LayerTypeDot11MgmtProbeResp); oo != nil {
			q, _ := oo.(*ll.Dot11InformationElement)
			r.Info = q.Info
			r.OUI = q.OUI
		}
	}
	addr4 := m.OUI(r.MAC)
	box := c.Beginning_Color + c.Beginning_ASCII + c.Middle_Color + c.Middle_ASCII + c.End_Color + c.End_ASCII
	fmt.Println("_____________________________________________")
	fmt.Println(lc.CYN, ">>> New Packet Data")
	fmt.Println(box, lc.HIGH_BLUE, " SSID            ", lc.WHT, ":: ", lc.HIGH_PINK, r.SSID)
	fmt.Println(box, lc.HIGH_BLUE, " MAC             ", lc.WHT, ":: ", lc.HIGH_PINK, r.MAC)
	fmt.Println(box, lc.HIGH_BLUE, " OUI             ", lc.WHT, ":: ", lc.HIGH_PINK, r.OUI)
	fmt.Println(box, lc.HIGH_BLUE, " MAC->OUI        ", lc.WHT, ":: ", lc.HIGH_PINK, addr4)
	fmt.Println(box, lc.HIGH_BLUE, " Info            ", lc.WHT, ":: ", lc.HIGH_PINK, r.Info)
	fmt.Println(box, lc.HIGH_BLUE, " RSSI            ", lc.WHT, ":: ", lc.HIGH_PINK, r.RSSI)
	fmt.Println(box, lc.HIGH_BLUE, " DTS             ", lc.WHT, ":: ", lc.HIGH_PINK, r.CaptureDTS.String())
	fmt.Println(box, lc.HIGH_BLUE, " Radio Tap Freq  ", lc.WHT, ":: ", lc.HIGH_PINK, r.Radio_Tap_Frequency)
	fmt.Println(box, lc.HIGH_BLUE, " Radio Tap Flags ", lc.WHT, ":: ", lc.HIGH_PINK, r.Radio_Tap_Flags)
	fmt.Println(box, lc.HIGH_BLUE, " Present chan?   ", lc.WHT, ":: ", lc.HIGH_PINK, r.Radio_Tap_Chan)
	fmt.Println(box, lc.HIGH_BLUE, " DBAntennaNoise  ", lc.WHT, ":: ", lc.HIGH_PINK, r.Radio_Tap_DBAntennaNoise)
	fmt.Println(box, lc.HIGH_BLUE, " DBAntennaSignal ", lc.WHT, ":: ", lc.HIGH_PINK, r.Radio_Tap_DBAntennaSignal)
	fmt.Println(box, lc.HIGH_BLUE, " Lock Quality    ", lc.WHT, ":: ", lc.HIGH_PINK, r.Radio_Tap_Lock_Quality)
	fmt.Println(box, lc.HIGH_BLUE, " Version         ", lc.WHT, ":: ", lc.HIGH_PINK, r.Radio_Tap_Version)
	fmt.Println(box, lc.HIGH_BLUE, " Present Rate    ", lc.WHT, ":: ", lc.HIGH_PINK, r.Radio_Tap_Present_Rate)
	fmt.Println(box, lc.HIGH_BLUE, "-------------------"+lc.WHT+"::-----------------------")
}

/*

 */
