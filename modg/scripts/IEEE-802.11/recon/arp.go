/*

Package and module dedicated to sending out ARP requests

and responding with the clients MAC, IP, OUI, and range

this is a remodified version of the offical ARP script written by the gopacket developers
i thought i should edit this alot and seperate the code to where its readable, and designed
a bit more better since it uses modules and will rewrite the packets from other modules
*/

package arp

import (
	"fmt"
	"log"
	"net"

	ARP_CONSTANTS "main/modg/scripts/IEEE-802.11/IEEE-802.11-c"
	ARP_RUN "main/modg/scripts/IEEE-802.11/recon/functions/run"
)

func Return_Values_andCall() {
	s, e := net.Interfaces()
	if e != nil {
		log.Fatal(e)
	} else {
		for _, r := range s {
			ARP_CONSTANTS.Wait_group.Add(1)
			go func(inter net.Interface) {
				defer ARP_CONSTANTS.Wait_group.Done()
				if x := ARP_RUN.Run(&inter); x != nil {
					fmt.Printf("Interface %v | %v \n", inter.Name, x)
				}
			}(r)
		}
	}
	ARP_CONSTANTS.Wait_group.Wait()
}
