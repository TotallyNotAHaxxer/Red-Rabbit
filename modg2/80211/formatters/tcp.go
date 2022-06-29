package FORMATION

import (
	"fmt"
	ct "main/modg2/80211/80211_types"

	l "github.com/google/gopacket/layers"
)

// Output ICMP

func Parse_ICMP_test(t *l.ICMPv6, c ct.Color) {
	box := c.Beginning_Color + c.Beginning_ASCII + c.Middle_Color + c.Middle_ASCII + c.End_Color + c.End_ASCII
	fmt.Println(box, " New Data Segment -> ICMPv6 packer")
	fmt.Println("------------------------------------")
	fmt.Print(box, " ")
	fmt.Printf("Base layer |%v\n", t.BaseLayer)
	fmt.Print(box, " ")
	fmt.Printf("Payload    |%s\n", string(t.Payload))
}
