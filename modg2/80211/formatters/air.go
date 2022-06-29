package FORMATION

import (
	"fmt"
	m "main/modg2/80211/80211_oui"
	"strings"
	"time"

	"github.com/google/gopacket"
	ll "github.com/google/gopacket/layers"
)

// airodump-ng simple remake
/*

Outputs following data

Encryption
SSID
BSSID
Frequency
RSSI
*/

type Beacon_Data struct {
	BSSID      string
	SSID       string
	OUI        string
	ENCRYPTION string
	F_ENC      bool
	Frequency  string
	RSSI       int8
	date       time.Time
}

func Print_Packet_Data(
	BSSID, SSID,
	OUI, ENC,
	FREQ string,
	RSSI int8,
	date time.Time,
	F_ENC bool) {
	fmt.Printf(" │ \033[38;5;198m%v\033[0;37m │   \033[38;5;198m%v\033[0;37m   │  \033[38;5;198m%s\033[0;37m   │ \033[38;5;198m%s\033[0;37m │  \033[38;5;198m%s\033[0;37m │ \033[38;5;198m%s\033[0;37m  │\n",
		strings.ToUpper(BSSID),
		RSSI,
		FREQ,
		ENC,
		OUI,
		SSID)

}

func Caller(pkt gopacket.Packet) {
	r := Beacon_Data{date: time.Now()}
	if l := pkt.Layer(ll.LayerTypeDot11); l != nil {
		if la := pkt.Layer(ll.LayerTypeDot11InformationElement); la != nil {
			d11inf, _ := la.(*ll.Dot11InformationElement)
			if d11inf.ID == ll.Dot11InformationElementIDSSID {
				q, _ := l.(*ll.Dot11)
				r.BSSID = q.Address3.String()
				r.SSID = string(d11inf.Info)
				_, r.ENCRYPTION = Parse_ENC(pkt, q)
				if r.ENCRYPTION == "OPN" {
					r.ENCRYPTION += "*"
				}
				oui := m.OUI(r.BSSID)
				if lak := pkt.Layer(ll.LayerTypeRadioTap); lak != nil {
					rad, _ := lak.(*ll.RadioTap)
					if rad.DBMAntennaSignal != 0x00 {
						r.RSSI = rad.DBMAntennaSignal
					}
					if rad.ChannelFrequency.String() != "" {
						r.Frequency = rad.ChannelFrequency.String()
					}
				}
				if oui != nil {
					if len(oui) <= 13 {
						n := 13
						addr := oui
						s := Trim(addr[0], n)
						chars := "******************"
						if len(r.SSID) >= len(chars) {
							ads := Trim(r.SSID, len(chars))
							r.SSID = ads[0]
						} else {
							a := len(chars) - len(r.SSID)
							char := strings.Repeat("*", a)
							r.SSID = r.SSID + char
						}
						if len(s[0]) == 13 {
							r.OUI = s[0]
							if r.BSSID != "" {
								if r.ENCRYPTION != "" {
									if r.Frequency != "" {
										if r.SSID != "" {
											if r.OUI != "" {
												Print_Packet_Data(r.BSSID, r.SSID, r.OUI, r.ENCRYPTION, r.Frequency, r.RSSI, time.Now(), r.F_ENC)
											}
										}
									}
								}
							}
						} else {
							a := 13 - len(s[0])
							chars := strings.Repeat("*", a)
							c := s[0] + chars
							Print_Packet_Data(r.BSSID, r.SSID, c, r.ENCRYPTION, r.Frequency, r.RSSI, time.Now(), r.F_ENC)
						}
					}
				}
			}
		}
	}
}
