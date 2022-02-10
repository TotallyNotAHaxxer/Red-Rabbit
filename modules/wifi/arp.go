package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/ArkAngeL43/port-scanning/port"
	"github.com/briandowns/spinner"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	foarp      = flag.Bool("arp", false, "| Start modules for ARP relay |")
	foport     = flag.Bool("port", false, "| Load the mnodule for port scanning a host |")
	fohelp     = flag.Bool("help", false, "| further help for information")
	fohelpa    = flag.Bool("advanced-h", false, "| further ADVANCED help for information")
	foh        = flag.Bool("h", false, "| further help for information")
	foportl    = flag.Bool("portl", false, "| port scan with a list")
	foportlist = flag.String("list", "", "| port scan with a list")
	fowifi     = flag.Bool("wifi", false, "| scan and log BSSID's && RSSE's")
	foporthost = flag.String("host", "", "hostname after using flag -foport")
	//
	BLU  = "\033[0;94m"
	chex = "\x1b[H\x1b[2J\x1b[3J"
	BLK  = "\033[0;30m"
	RED  = "\033[0;31m"
	GRN  = "\033[0;32m"
	YEL  = "\033[0;33m"
	MAG  = "\033[0;35m"
	CYN  = "\033[0;36m"
	WHT  = "\033[0;37m"
	//Regular bold
	BBLK = "\033[1;30m"
	BRED = "\033[1;31m"
	BGRN = "\033[1;32m"
	BYEL = "\033[1;33m"
	BBLU = "\033[1;34m"
	BMAG = "\033[1;35m"
	BCYN = "\033[1;36m"
	BWHT = "\033[1;37m"
	//Regular underline
	UBLK = "\033[4;30m"
	URED = "\033[4;31m"
	UGRN = "\033[4;32m"
	UYEL = "\033[4;33m"
	UBLU = "\033[4;34m"
	UMAG = "\033[4;35m"
	UCYN = "\033[4;36m"
	UWHT = "\033[4;37m"
	//Regular back
	BLKB = "\033[40m"
	REDB = "\033[41m"
	GRNB = "\033[42m"
	YELB = "\033[43m"
	BLUB = "\033[44m"
	MAGB = "\033[45m"
	CYNB = "\033[46m"
	WHTB = "\033[47m"
	//High intensty back
	BLKHB = "\033[0;100m"
	REDHB = "\033[0;101m"
	GRNHB = "\033[0;102m"
	YELHB = "\033[0;103m"
	BLUHB = "\033[0;104m"
	MAGHB = "\033[0;105m"
	CYNHB = "\033[0;106m"
	WHTHB = "\033[0;107m"
	//High intensty text
	HBLK = "\033[0;90m"
	HRED = "\033[0;91m"
	HGRN = "\033[0;92m"
	HYEL = "\033[0;93m"
	HBLU = "\033[0;94m"
	HMAG = "\033[0;95m"
	HCYN = "\033[0;96m"
	HWHT = "\033[0;97m"
	//Bold high intensity text
	BHBLK = "\033[1;90m"
	BHRED = "\033[1;91m"
	BHGRN = "\033[1;92m"
	BHYEL = "\033[1;93m"
	BHBLU = "\033[1;94m"
	BHMAG = "\033[1;95m"
	BHCYN = "\033[1;96m"
	BHWHT = "\033[1;97m"
)

func ce(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func localaddr() {
	ifaces, err := net.Interfaces()
	ce(err)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		ce(err)
		for _, a := range addrs {
			log.Printf("\033[35m[\033[34mINTERFACE\033[35m] ->  %v %v\n", i.Name, a)
		}
	}
}

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
		return errors.New("\033[34m[\033[35m*\033[34m] Skipping LO") //skip local host
	} else if addr.Mask[0] != 0xff || addr.Mask[1] != 0xff {
		return errors.New("\033[34m[\033[35m?\033[34m] Mask is to large")
	}
	log.Printf("\033[34m[\033[35m*\033[34m] USING CURRENT NET RANGE -> %v FOR INTERFACES -> %v", addr, iface.Name)
	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	ce(err)
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

func port_scanner_list(list string) {
	listed, err := ioutil.ReadFile(list)
	ce(err)
	fmt.Println("____________Hosts in file_________________")
	fmt.Println("[ 1 ] File -> ", list)
	fmt.Println("[ 2 ] Are these hosts correct? \n\n")
	fmt.Println(string(listed))
	fmt.Print("n=No | Y=yes|")
	input := bufio.NewReader(os.Stdin)
	fmt.Print(BLU, "\n> ")
	for {
		go handelreturncon(make(chan os.Signal, 1))
		t, _ := input.ReadString('\n')
		t = strings.Replace(t, "\n", "", -1)
		if strings.Compare("yes", t) == 0 {
			fmt.Println(chex)
			if runtime.GOOS == "windows" {
				starter("txt\bn.txt")
			} else {
				starter("txt/bn.txt")
			}
			fill, err := os.Open(list)
			ce(err)
			defer fill.Close()
			watchdog := bufio.NewScanner(fill)
			for watchdog.Scan() {
				fmt.Println("\033[0;94m[", REDB, watchdog.Text(), "]\033[0;49m")
				port.GetOpenPorts(watchdog.Text(), port.PortRange{Start: 1, End: 8090})

			}
		}
		if strings.Compare("Y", t) == 0 {
			fmt.Println(chex)
			if runtime.GOOS == "windows" {
				starter("txt\bn.txt")
			} else {
				starter("txt/bn.txt")
			}
			fill, err := os.Open(list)
			ce(err)
			defer fill.Close()
			watchdog := bufio.NewScanner(fill)
			for watchdog.Scan() {
				fmt.Println("\033[0;94m[", REDB, watchdog.Text(), "]\033[0;49m")
				port.GetOpenPorts(watchdog.Text(), port.PortRange{Start: 1, End: 8090})

			}
		}
		if strings.Compare("YES", t) == 0 {
			fmt.Println(chex)
			if runtime.GOOS == "windows" {
				starter("txt\bn.txt")
			} else {
				starter("txt/bn.txt")
			}
			fill, err := os.Open(list)
			ce(err)
			defer fill.Close()
			watchdog := bufio.NewScanner(fill)
			for watchdog.Scan() {
				fmt.Println("\033[0;94m[", REDB, watchdog.Text(), "]\033[0;49m")
				port.GetOpenPorts(watchdog.Text(), port.PortRange{Start: 1, End: 8090})

			}
		}
		os.Exit(1)
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

			fmt.Println("\033[34m[\033[35m+\033[34m] IP-ADDR => ", net.IP(arp.SourceProtAddress), "Has MAC ADDR of => ", net.HardwareAddr(arp.SourceHwAddress))
			// possibly write it to a file idk at least try
			pathmain, err := os.OpenFile("in.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				if fileExists("in.txt") {
					fmt.Println("[ - ] Unexpected error? but why -> ", err)
					fmt.Println("[ * ] Creating file")
					f, err := os.Create("in.txt")
					ce(err)
					f.Close()
				} else {
					fmt.Println("[ - ] Unexpected error? but why -> ", err)
					fmt.Println("[ * ] Creating file")
					f, err := os.Create("in.txt")
					ce(err)
					f.Close()
					continue
				}
			}
			pathiplace, err := os.OpenFile("lace_ip.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			ce(err)
			fmt.Fprintln(pathiplace, net.IP(arp.SourceProtAddress))
			fmt.Fprintln(pathmain, "IPADDR => ", net.IP(arp.SourceProtAddress), "has MAC => ", net.HardwareAddr(arp.SourceHwAddress))
			if err != nil {
				log.Fatal(err)
			}
			b, err := ioutil.ReadFile("in.txt")
			if err != nil {
				if fileExists("in.txt") {
					fmt.Println("[ - ] Unexpected error? but why ")
					fmt.Println("[ * ] Creating file")
					f, err := os.Create("in.txt")
					ce(err)
					f.Close()
				} else {
					continue
				}
			}
			err = ioutil.WriteFile("out.txt", b, 0644)
			if err != nil {
				if fileExists("out.txt") {
					fmt.Println("[ - ] Unexpected error? but why ")
					fmt.Println("[ * ] Creating file")
					f, err := os.Create("in.txt")
					ce(err)
					f.Close()
				} else {
					continue
				}
			}

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

func arpmain() {
	localaddr()
	go handelreturncon(make(chan os.Signal, 1))
	var err error
	ifaces, err := net.Interfaces()
	ce(err)
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

func netip() {
	uli := "https://api.ipify.org?format=text"
	response, err := http.Get(uli)
	ce(err)
	defer response.Body.Close()
	ip, err := ioutil.ReadAll(response.Body)
	ce(err)
	fmt.Printf("\033[32m\t\t[PUBLIC]->[IPA-ADDR]      | %s\n", ip)
	addrs, err := net.InterfaceAddrs()
	ce(err)
	for i, addr := range addrs {
		fmt.Printf("\033[32m\t\t[PRIVATE]->[INTERFACE] #%d | %v\n", i, addr)
	}
	fmt.Println("[+] Getting Ranges......")
	ifaces, err := net.Interfaces()
	// handle function usage
	ce(err)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		ce(err)
		for _, addr := range addrs {
			//var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				fmt.Println("\t\t", ip)
			case *net.IPAddr:
				ip = v.IP
				fmt.Println("\t\t", ip)
			}
		}
	}
}

func starter(filename string) {
	content, err := ioutil.ReadFile(filename)
	ce(err)
	fmt.Println(RED, string(content))
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	} else {
		if resp.StatusCode >= 200 {
			if fileExists("in.txt") { // or or file exists out and in
				fmt.Println(BLU, "[", REDB, "Online\033[1;39m", BLU, "]", "\t\t\t[ ", REDB, "File in loaded\033[1;39m", BLU, "]")
			} else {
				f, err := os.Create("in.txt")
				ce(err)
				defer f.Close()
			}
			if fileExists("out.txt") {
				fmt.Println("\n\t\t\t\t[ ", REDB, "File out loaded\033[1;39m", BLU, "]")
			} else {
				f, err := os.Create("out.txt")
				ce(err)
				defer f.Close()
			}
		}
	}
}

func main() {
	flag.Parse()
	if *foarp {
		if !fileExists("lace_ip.txt") {
			os.Create("lace_ip.txt")
		}
		fmt.Println(chex)
		if runtime.GOOS == "windows" {
			starter("txt\bn.txt")
		} else {
			starter("txt/bn.txt")
		}
		fmt.Println("\n\n[ 1 ] STAT: Module ARP loaded ")
		arpmain()
		os.Exit(1)
	}
	if *foportl {
		files := *foportlist
		fmt.Println(chex)
		if runtime.GOOS == "windows" {
			starter("txt\bn.txt")
		} else {
			starter("txt/bn.txt")
		}
		port_scanner_list(files)
		os.Exit(0)
	}
}
