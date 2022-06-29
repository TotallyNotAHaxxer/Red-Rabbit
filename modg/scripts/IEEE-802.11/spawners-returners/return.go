package spawner_return

import (
	"fmt"
	"net"
)

func Locate() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("<RR6> Net Module: Could not make a request to find all network interfaces got error -> ", err)
	} else {
		for _, i := range ifaces {
			return i.Name
		}
	}
	return ""
}
