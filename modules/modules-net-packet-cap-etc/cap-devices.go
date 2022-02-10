/*

Author => ArkAngeL43

Program does -< gather listenable devices for TCP capture

packages: go tabulate, go packet
*/
package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	dev_pcaps, err := pcap.FindAllDevs()
	//
	if err != nil {
		log.Fatal(err, "RR5 -> STAT -> Go packet PCAP could not find all listenable devices")
	} else {
		for _, devices := range dev_pcaps {
			fmt.Println("DEVICE NAME => ", devices)
			for _, address := range devices.Addresses {
				fmt.Println(address.IP, address.Broadaddr)
			}
		}
	}
}
