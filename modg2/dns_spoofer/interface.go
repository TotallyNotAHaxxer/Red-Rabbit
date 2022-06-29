package SPOOFER

import (
	"fmt"
	"net"
)

func interip(intf string) {
	itf, _ := net.InterfaceByName(intf)
	item, _ := itf.Addrs()
	var ip net.IP
	for _, addr := range item {
		switch v := addr.(type) {
		case *net.IPNet:
			if !v.IP.IsLoopback() {
				if v.IP.To4() != nil {
					ip = v.IP
				}
			}
		}
	}
	if ip != nil {
		fmt.Println("IP of interface " + intf + " => " + ip.String())
	} else {
		fmt.Println(intf, "has IP NIL>? => ", ip.String())
	}
}
