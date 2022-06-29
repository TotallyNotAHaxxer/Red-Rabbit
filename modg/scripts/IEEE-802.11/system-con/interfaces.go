// grabs all IPA's and interfaces
package interfaces

import (
	"fmt"
	"net"

	v "main/modg/colors"
	ec "main/modg/warnings"
)

func HOME_Interfaces() {
	ifaces, err := net.Interfaces()
	ec.Warning_advanced("<RR6> Net Module: Could not grab interface names", v.REDHB, 1, false, false, true, err, 1, 233, "")
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		ec.Warning_advanced("<RR6> Net Module: Could not stat addresses for interfaces", v.REDHB, 1, false, false, true, err, 1, 233, "")
		for _, a := range addrs {
			fmt.Println(v.RED, "<RR6> Net Module: Found Device  | ", v.MAG, i.Name, v.RED, "\t | Addr:", a)
		}
	}
}

func HOME_Interfaces_return_list() []string {
	var array []string
	ifaces, err := net.Interfaces()
	ec.Warning_advanced("<RR6> Net Module: Could not grab interface names", v.REDHB, 1, false, false, true, err, 1, 233, "")
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		ec.Warning_advanced("<RR6> Net Module: Could not stat addresses for interfaces", v.REDHB, 1, false, false, true, err, 1, 233, "")
		for _, a := range addrs {
			fmt.Println(v.RED, "<RR6> Net Module: Found Device  | ", v.MAG, i.Name, v.RED, "\t | Addr:", a)
			array = append(array, i.Name)
		}
	}
	return array
}
