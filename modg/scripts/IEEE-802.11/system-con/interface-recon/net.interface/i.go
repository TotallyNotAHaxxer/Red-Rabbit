package i

import (
	"fmt"
	"net"
)

func Address_Net_Information(addr net.Addr) string {
	var scope string
	ipAddr, ipNet, err := net.ParseCIDR(addr.String())
	if err != nil {
		return "unknown"
	}
	switch {
	case ipAddr.IsLoopback():
		scope = "loopback"
	case ipAddr.IsGlobalUnicast():
		scope = "global unicast"
	case ipAddr.IsMulticast():
		scope = "global multicast"
	case ipAddr.IsLinkLocalUnicast():
		scope = "link local unicast"
	case ipAddr.IsLinkLocalMulticast():
		scope = "link local multicast"
	case ipAddr.IsInterfaceLocalMulticast():
		scope = "interface multicast"
	case ipAddr.IsUnspecified():
		scope = "unspecified"
	default:
		scope = "unknown"
	}

	return fmt.Sprintf("\n< NET TYPE > %s \n\t< NETWORK > \t\t\t\t\t|  %s   \n\t\t< NETWORK ADDRESS > \t\t\t|  %s   \n\t\t\t< NETWORK MASK > \t\t|  %v   \n\t\t\t\t< NETWORK SCOPE > \t|  %s  ", ipNet.Network(), ipNet.IP.String(), ipAddr.String(), ipAddr.DefaultMask(), scope)
}
