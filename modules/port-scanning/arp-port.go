package main

// putting it all into one main file as adding onto easy make and call

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func checkErr(err error) {
	//var err error
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func handelreturncon(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nDetected Interupt.....")
			os.Exit(1)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func handeler(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nDetected Interupt.....")
			os.Exit(1)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func netcheck() bool {
	_, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("[!] Interface may be offline")
		os.Exit(1)
	} else {
		return false
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////////////

// gather of interfaces

func localaddr() {
	ifaces, err := net.Interfaces()
	checkErr(err)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		checkErr(err)
		for _, a := range addrs {
			log.Printf("\033[35m[\033[34mINTERFACE\033[35m] ->  %v %v\n", i.Name, a)
		}
	}
}

////////////

func scan(iface *net.Interface) error {
	go handelreturncon(make(chan os.Signal, 1))
	var addr *net.IPNet
	if addrs, err := iface.Addrs(); err != nil {
		return err
	} else {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok {
				if ip4 := ipnet.IP.To4(); ip4 != nil {
					addr = &net.IPNet{
						IP:   ip4,
						Mask: ipnet.Mask[len(ipnet.Mask)-4:],
					}
					break
				}
			}
		}
	}
	if addr == nil {
		return errors.New("\033[34m[\033[35m?\033[34m] Unstable Network on card") // check networks interface
	} else if addr.IP[0] == 127 {
		return errors.New("\033[34m[\033[35m*\033[34m] Skipping LO...") //skip local host
	} else if addr.Mask[0] != 0xff || addr.Mask[1] != 0xff {
		return errors.New("\033[34m[\033[35m?\033[34m] Mask is to large....")
	}
	log.Printf("\033[34m[\033[35m*\033[34m] USING CURRENT NET RANGE -> %v FOR INTERFACES -> %v", addr, iface.Name)
	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	checkErr(err)
	defer handle.Close()

	stop := make(chan struct{})
	go ARPR(handle, iface, stop)
	defer close(stop)
	for {

		if err := ARPW(handle, iface, addr); err != nil {
			log.Printf("[DATA]->[ERROR] An error has occured during the following write of packets-> %v: %v", iface.Name, err)
			return err
		}
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()
		time.Sleep(2 * time.Second)
		s.Stop()
	}
}

func ARPR(handle *pcap.Handle, iface *net.Interface, stop chan struct{}) {
	go handelreturncon(make(chan os.Signal, 1))
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		var packet gopacket.Packet
		select {
		case <-stop:
			return
		case packet = <-in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp := arpLayer.(*layers.ARP)
			if arp.Operation != layers.ARPReply || bytes.Equal([]byte(iface.HardwareAddr), arp.SourceHwAddress) {

				continue
			}

			log.Printf("\033[34m[\033[35m+\033[34m] Address [%v] Has MAC ~> [%v]", net.IP(arp.SourceProtAddress), net.HardwareAddr(arp.SourceHwAddress))
		}
	}
}

func ARPW(handle *pcap.Handle, iface *net.Interface, addr *net.IPNet) error {
	go handelreturncon(make(chan os.Signal, 1))
	eth := layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(iface.HardwareAddr),
		SourceProtAddress: []byte(addr.IP),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
	}

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	for _, ip := range ips(addr) {
		arp.DstProtAddress = []byte(ip)
		gopacket.SerializeLayers(buf, opts, &eth, &arp)
		if err := handle.WritePacketData(buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

func ips(n *net.IPNet) (out []net.IP) {
	num := binary.BigEndian.Uint32([]byte(n.IP))
	mask := binary.BigEndian.Uint32([]byte(n.Mask))
	network := num & mask
	broadcast := network | ^mask
	for network++; network < broadcast; network++ {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], network)
		out = append(out, net.IP(buf[:]))
	}
	return
}

func main() {
	fmt.Println("\033[34m[\033[35m+\033[34m] Module ARP Relay Loaded....")
	fmt.Println("\033[34m[\033[35m+\033[34m] Gathering Interfaces.....")
	localaddr()
	go handelreturncon(make(chan os.Signal, 1))
	var err error
	netcheck()
	ifaces, err := net.Interfaces()
	checkErr(err)
	var wg sync.WaitGroup
	for _, iface := range ifaces {
		wg.Add(1)
		go func(iface net.Interface) {
			defer wg.Done()
			if err := scan(&iface); err != nil {
				log.Printf("interface %v: %v", iface.Name, err)
			}
		}(iface)
		go handelreturncon(make(chan os.Signal, 1))
	}
	go handelreturncon(make(chan os.Signal, 1))
	wg.Wait()
}
