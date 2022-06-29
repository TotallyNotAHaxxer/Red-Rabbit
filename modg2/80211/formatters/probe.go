package FORMATION

import (
	"fmt"
	lc "main/modg2/80211/80211_color"
	ct "main/modg2/80211/80211_types"

	l "github.com/google/gopacket/layers"
)

func Parse_Probe_Req(t *l.Dot11MgmtProbeReq, c ct.Color) {
	fmt.Println("_____________________________________________")
	fmt.Println(lc.CYN, ">>> New Packet Data")
	fmt.Println(t.Payload)
}
