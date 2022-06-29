package FORMATION

import (
	"fmt"
	ct "main/modg2/80211/80211_types"

	lc "main/modg2/80211/80211_color"

	m "main/modg2/80211/80211_oui"

	l "github.com/google/gopacket/layers"
)

type Check struct {
	Set  *bool
	Uset *bool
}

func Parse_Dot11_normal(t *l.Dot11, c ct.Color) {
	box := c.Beginning_Color + c.Beginning_ASCII + c.Middle_Color + c.Middle_ASCII + c.End_Color + c.End_ASCII
	fmt.Println("_____________________________________________")
	fmt.Println(lc.CYN, ">>> New Packet Data")
	addr1 := m.OUI(t.Address1.String())
	addr2 := m.OUI(t.Address1.String())
	addr3 := m.OUI(t.Address1.String())
	addr4 := m.OUI(t.Address1.String())
	fmt.Println(box, lc.HIGH_BLUE, " Address 1     ", lc.WHT, ":: ", lc.HIGH_PINK, t.Address1.String())
	fmt.Println(box, lc.HIGH_BLUE, " Address 1 OUI ", lc.WHT, ":: ", lc.HIGH_PINK, addr1)
	fmt.Println(box, lc.HIGH_BLUE, " Address 2     ", lc.WHT, ":: ", lc.HIGH_PINK, t.Address2.String())
	fmt.Println(box, lc.HIGH_BLUE, " Address 2 OUI ", lc.WHT, ":: ", lc.HIGH_PINK, addr2)
	fmt.Println(box, lc.HIGH_BLUE, " Address 3     ", lc.WHT, ":: ", lc.HIGH_PINK, t.Address3.String())
	fmt.Println(box, lc.HIGH_BLUE, " Address 3 OUI ", lc.WHT, ":: ", lc.HIGH_PINK, addr3)
	fmt.Println(box, lc.HIGH_BLUE, " Address 4     ", lc.WHT, ":: ", lc.HIGH_PINK, t.Address4.String())
	fmt.Println(box, lc.HIGH_BLUE, " Address 4 OUI ", lc.WHT, ":: ", lc.HIGH_PINK, addr4)
	fmt.Println(box, lc.HIGH_BLUE, " Flags Power   ", lc.WHT, ":: ", lc.HIGH_PINK, t.Flags.PowerManagement())
	fmt.Println(box, lc.HIGH_BLUE, " Flags FromDS  ", lc.WHT, ":: ", lc.HIGH_PINK, t.Flags.FromDS())
	fmt.Println(box, lc.HIGH_BLUE, " Flags WEP     ", lc.WHT, ":: ", lc.HIGH_PINK, t.Flags.WEP())
	fmt.Println(box, lc.HIGH_BLUE, " Flags MF      ", lc.WHT, ":: ", lc.HIGH_PINK, t.Flags.MD())
	fmt.Println(box, lc.HIGH_BLUE, " Flags MD      ", lc.WHT, ":: ", lc.HIGH_PINK, t.Flags.MF())
	fmt.Println(box, lc.HIGH_BLUE, " Flags FromDS  ", lc.WHT, ":: ", lc.HIGH_PINK, t.Flags.String())
	fmt.Println(box, lc.HIGH_BLUE, " Checksum      ", lc.WHT, ":: ", lc.HIGH_PINK, t.Checksum)
	fmt.Println(box, lc.HIGH_BLUE, " Duration ID   ", lc.WHT, ":: ", lc.HIGH_PINK, t.DurationID)
	fmt.Println(box, lc.HIGH_BLUE, " Fragment Num  ", lc.WHT, ":: ", lc.HIGH_PINK, t.FragmentNumber)
	fmt.Println(box, lc.HIGH_BLUE, " HT Control    ", lc.WHT, ":: ", lc.HIGH_PINK, t.HTControl)
	fmt.Println(box, lc.HIGH_BLUE, " Protocal      ", lc.WHT, ":: ", lc.HIGH_PINK, t.Proto)
	fmt.Println(box, lc.HIGH_BLUE, "---------------", lc.WHT, "::-----------------------")

}
