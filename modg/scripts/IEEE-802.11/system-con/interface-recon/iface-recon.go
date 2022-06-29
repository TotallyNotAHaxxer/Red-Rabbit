package interfacerecon

import (
	"fmt"
	"net"
	"os"
	"strings"

	constants "main/modg/scripts/IEEE-802.11/system-con/interface-recon/const"
	i "main/modg/scripts/IEEE-802.11/system-con/interface-recon/net.interface"
)

func Recon(iface string) {
	if iface != "*" {
		constants.Sniffc = iface
	}
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, iface := range ifaces {
		if constants.Sniffc == "" || constants.Sniffc == iface.Name {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}
			fmt.Printf("\n\033[32m********************************************\n* All information for interface [ %s ] *\n********************************************\n\033[39m< INTERFACE > | %s | < INTERFACE DATA > | %s |  < MTU > | %d | \n", iface.Name, iface.Name, strings.ToUpper(iface.Flags.String()), iface.MTU)
			if len(iface.HardwareAddr.String()) > 0 {
				fmt.Printf("\tETHERNET ADDRESS -> %s\n", iface.HardwareAddr.String())
			}
			if len(addrs) > 0 {
				for _, addr := range addrs {
					fmt.Printf("\t\033[0;97m%s\n", i.Address_Net_Information(addr))
				}
			}
		}
	}
}
