package FORMATION

import (
	"bytes"
	c "main/modg2/80211/80211_constants"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	ll "github.com/google/gopacket/layers"
)

func Parse_ENC(p gopacket.Packet, d *ll.Dot11) (bool, string) {
	switch d.Flags {
	case ll.Dot11FlagsWEP:
		c.F = true
		c.ENCRYPTION = "WEP"
	}
	for _, l := range p.Layers() {
		if l.LayerType() == ll.LayerTypeDot11InformationElement {
			i, k := l.(*ll.Dot11InformationElement)
			if k {
				c.F = true
				if i.ID == layers.Dot11InformationElementIDRSNInfo {
					c.ENCRYPTION = "WPA2"
				} else {
					if c.ENCRYPTION == "" && i.ID == ll.Dot11InformationElementIDVendor && i.Length >= 8 && bytes.Equal(i.OUI, c.WPAS) && bytes.HasPrefix(i.Info, []byte{1, 0}) {
						c.ENCRYPTION = "WPA"
					}
				}

			}
		}
	}
	if c.ENCRYPTION == "" && c.F {
		c.ENCRYPTION = "OPN"
	}
	return c.F, c.ENCRYPTION
}
