package OUI

import (
	"net"

	"github.com/google/gopacket/macs"
)

func OUI(m string) []string {
	var alist []string
	if mac, x := net.ParseMAC(m); x == nil {
		prefix := [3]byte{
			mac[0],
			mac[1],
			mac[2],
		}
		manufacturer, e := macs.ValidMACPrefixMap[prefix]
		if e {
			alist = append(alist, manufacturer)
		}
		return alist
	}
	return []string{"Unknown"}
}
