/*

Author     => ArkAngeL43
Language   => Golang
Type APP   => CLI
Color      => Standard, GRN, BLU, DEEPGRN ASCII
name       => Ether_Rabbit

Tool DESC: This is a simple ethernet and tcp dump/traffic logger as a command line application which gives you the option to filter your capture

the filter types are tcp, ip, eth, application

filter tcp will log, sniff, and output only TCP packets and responses
filter ip will log, sniff, and output only IP packets and responses
filter eth will log, snff, and output only ETH packets and responses
filter application will output only type application packets/responses and payloads

interface -> Defualt for linux is Eth0

Support -> Linux OS's ONLY, windows and MAC have not been tested and windows is not supported yet

built off of Black hat GO's very small, and lightweight concurrent pcap lib
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bndr/gotabulate"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// base data
var (
	// time layout and log layout
	filter = flag.String("BPF", "DEFUALT: LISTEN: ALL ", "set the BPF filter on a port")
	//Defualt Capture ALL
	application_type_filters = flag.String("filter", "LISTEN ALL", " set a application filter ex listen for ether packets or HTTP packets")
	tru_fals_dump_pcap       = flag.Bool("dump", false, "activate pcap dumping")
	dump_pcap                = flag.String("pcapfile", "", "set a pcap file to dump")
	now                      = time.Now()
	formatDate               = now.Format("15:04:05")
	interface_use_flag       = flag.String("interface", defualt_interface, "")
	url_test                 = "https://www.google.com"
	defualt_interface        = "eth0"
	err                      error
	snapshot_main            int32         = 1024 // set as type 32
	timeout_shot_cap         time.Duration = 40 * time.Second
	handeler_pcap            *pcap.Handle
	clear_hex                = "\x1b[H\x1b[2J\x1b[3J"
	RETURN_COLOR_FOR         = "\033[0;39m"
	BLK                      = "\033[0;30m"
	RED                      = "\033[0;31m"
	GRN                      = "\033[0;32m"
	YEL                      = "\033[0;33m"
	BLU                      = "\033[0;34m"
	MAG                      = "\033[0;35m"
	CYN                      = "\033[0;36m"
	WHT                      = "\033[0;37m"
	BBLK                     = "\033[1;30m"
	BRED                     = "\033[1;31m"
	BGRN                     = "\033[1;32m"
	BYEL                     = "\033[1;33m"
	BBLU                     = "\033[1;34m"
	BMAG                     = "\033[1;35m"
	BCYN                     = "\033[1;36m"
	BWHT                     = "\033[1;37m"
	UBLK                     = "\033[4;30m"
	URED                     = "\033[4;31m"
	UGRN                     = "\033[4;32m"
	UYEL                     = "\033[4;33m"
	UBLU                     = "\033[4;34m"
	UMAG                     = "\033[4;35m"
	UCYN                     = "\033[4;36m"
	UWHT                     = "\033[4;37m"
	BLKB                     = "\033[40m"
	REDB                     = "\033[41m"
	GRNB                     = "\033[42m"
	YELB                     = "\033[43m"
	BLUB                     = "\033[44m"
	MAGB                     = "\033[45m"
	CYNB                     = "\033[46m"
	WHTB                     = "\033[47m"
	BLKHB                    = "\033[0;100m"
	REDHB                    = "\033[0;101m"
	GRNHB                    = "\033[0;102m"
	YELHB                    = "\033[0;103m"
	BLUHB                    = "\033[0;104m"
	MAGHB                    = "\033[0;105m"
	CYNHB                    = "\033[0;106m"
	WHTHB                    = "\033[0;107m"
	HBLK                     = "\033[0;90m"
	HRED                     = "\033[0;91m"
	HGRN                     = "\033[0;92m"
	HYEL                     = "\033[0;93m"
	HBLU                     = "\033[0;94m"
	HMAG                     = "\033[0;95m"
	HCYN                     = "\033[0;96m"
	HWHT                     = "\033[0;97m"
	BHBLK                    = "\033[1;90m"
	BHRED                    = "\033[1;91m"
	BHGRN                    = "\033[1;92m"
	BHYEL                    = "\033[1;93m"
	BHBLU                    = "\033[1;94m"
	BHMAG                    = "\033[1;95m"
	BHCYN                    = "\033[1;96m"
	BHWHT                    = "\033[1;97m"
	//
	ETHER_       = layers.LayerTypeEthernet
	TCP_         = layers.LayerTypeTCP
	IP_          = layers.LayerTypeIPv4
	IP_6         = layers.LayerTypeIPv6
	DHCMP_       = layers.LayerTypeDHCPv4
	DHCMP_6      = layers.LayerTypeDHCPv6
	ARP_         = layers.LayerTypeARP
	DNS_         = layers.LayerTypeDNS
	ICMP_6_NA_   = layers.LayerTypeICMPv6NeighborAdvertisement
	ICMP_6_SOL_  = layers.LayerTypeICMPv6NeighborSolicitation
	ICP_4_       = layers.LayerTypeICMPv4
	ICMP_6_      = layers.LayerTypeICMPv6
	ICMP_ECH_    = layers.LayerTypeICMPv6Echo
	ICMP_router_ = layers.LayerTypeICMPv6RouterAdvertisement
	ICMP_REDI_   = layers.LayerTypeICMPv6Redirect
	UDP_         = layers.LayerTypeUDP
	// DOT 11
	DOT11_             = layers.LayerTypeDot11
	DOT11_BEC_         = layers.LayerTypeDot11MgmtBeacon
	DOT_11_DEIS_       = layers.LayerTypeDot11MgmtDeauthentication
	DOT_11_DISASS_     = layers.LayerTypeDot11MgmtDisassociation
	DOT_11_PROBE_      = layers.Dot11InformationElementIDSSID
	DOT_11_PROBE_REQ_  = layers.LayerTypeDot11MgmtProbeReq
	DOT_11_PROBE_RESP_ = layers.LayerTypeDot11MgmtProbeResp
	DOT_11_WEP_        = layers.LayerTypeDot11WEP
	DOT_11_RAD_TAP_    = layers.LayerTypeRadioTap
)

func banner(filename string) {
	copntent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(clear_hex)
		fmt.Println(WHT, string(copntent))
		fmt.Println(BLU, "^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	}
}

func check_err_cl(err error) bool {
	if err != nil {
		log.Fatal("COULD NOT CAPTURE ON DEVICE OR NETWORK => CLIENT NET DOWN => ", err)
		return true
	}
	return false
}

func std_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func cap_proc_IEEE_ICMP_PROCMOD1(pkt_icmp gopacket.Packet) {
	ICMP_V6_PROC := pkt_icmp.Layer(ICMP_REDI_)
	ICMP_6_SOL_m := pkt_icmp.Layer(ICMP_6_SOL_)
	ICMP_4 := pkt_icmp.Layer(ICP_4_)
	// echo packet
	ECHO_ := pkt_icmp.Layer(ICMP_ECH_)
	if ICMP_V6_PROC != nil {
		fmt.Println(RED, "\n[ETHER_RABBIT] DATA:    ", BLU, formatDate, YEL, " ICMP VERSION 6 PACKET SNIFFED......")
		fmt.Println(BLU, "\n___________________________________________________")
		fmt.Println(MAG, "|", ICMP_V6_PROC.LayerPayload())
		fmt.Println(MAG, "|", ICMP_V6_PROC.LayerContents())
	}
	if ICMP_6_SOL_m != nil {
		fmt.Println(RED, "\n[ETHER_RABBIT] DATA:    ", BLU, formatDate, YEL, " ICMP VERSION 6 Neighbor SOLICITATION PACKET SNIFFED......")
		fmt.Println(BLU, "\n___________________________________________________")
		fmt.Println(MAG, "| CONTENTS => ", ICMP_6_SOL_m.LayerContents())
		fmt.Println(MAG, "| PAYLOAD  => ", ICMP_6_SOL_m.LayerPayload())
		fmt.Println(MAG, "| TYPE     => ", ICMP_6_SOL_m.LayerType())
	}
	if ICMP_4 != nil {
		fmt.Println(RED, "\n[ETHER_RABBIT] DATA:    ", BLU, formatDate, YEL, " ICMP VERSION 4 PACKET SNIFFED......")
		fmt.Println(BLU, "\n___________________________________________________")
		fmt.Println(MAG, "| CONTENTS => ", ICMP_4.LayerContents())
		fmt.Println(MAG, "| PAYLOAD  => ", ICMP_4.LayerPayload())
		fmt.Println(MAG, "| TYPE     => ", ICMP_4.LayerType())
	}
	if ECHO_ != nil {
		fmt.Println(RED, "\n[ETHER_RABBIT] DATA:    ", BLU, formatDate, YEL, " ICMP ECHO PACKET SNIFFED......")
		fmt.Println(BLU, "\n___________________________________________________")
		fmt.Println(MAG, "| LAYER CONTENT => ", ECHO_.LayerContents())
		fmt.Println(MAG, "| PAYLOAD       => ", ECHO_.LayerPayload())
		fmt.Println(MAG, "| CONTENT       => ", ECHO_.LayerContents())
	}
}

func _IEEE_CAP_PROCESS_Dot11(capture_main gopacket.Packet) {
	fmt.Println(RED, "[ETHER_RABBIT] DATA: ", BLU, formatDate, " Working to process...Awaiting capture...")
	dot11 := capture_main.Layer(DOT11_)
	dot11info := capture_main.Layer(layers.LayerTypeDot11InformationElement)
	dot11_WEP := capture_main.Layer(DOT_11_WEP_)
	dot11_DEAUTH_ID := capture_main.Layer(DOT_11_DEIS_)
	if nil != dot11 {
		dot11, _ := dot11.(*layers.Dot11)
		fmt.Println(RED, "\n[ETHER_RABBIT] DATA:    ", BLU, formatDate, YEL, " DOT_11 Beacon packet picked up ")

		// most flags come back empty
		fmt.Println(BLU, "\n______________________________________________________")
		fmt.Println("|", MAG, "BSSID Address 3 -> ", dot11.Address3)
		fmt.Println("|", MAG, "BSSID Address 2 -> ", dot11.Address2)
		fmt.Println("|", MAG, "BSSID Address 1 -> ", dot11.Address1)
		fmt.Println("|", MAG, "BSSID Address 4 -> ", dot11.Address4)
		fmt.Println("|", MAG, "Packet Flages   -> ", dot11.Flags)
		fmt.Println("|", MAG, "Checksum        -> ", dot11.Checksum)
		packet_layers_all(capture_main)
		fmt.Println(BLU, "|------------------------------------------------------------")
		fmt.Println(YEL, "\n......................................................................................")
		fmt.Println(YEL, "PACKET DATA => ", dot11)
		fmt.Println(YEL, "\n......................................................................................")
	}
	// is probe request?
	if nil != dot11info {
		fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, YEL, " DOT_11 PROBE REQUEST PACKET picked up")
		dot11info, _ := dot11info.(*layers.Dot11InformationElement)
		if dot11info.ID == DOT_11_PROBE_ {
			fmt.Println(BLU, "______________________________________________________")
			fmt.Println(MAG, "|NET ID CAP      -> ", BLU, dot11info.ID)
			fmt.Println(MAG, "|OUI             -> ", BLU, dot11info.OUI)
			fmt.Println(MAG, "|SSID            -> ", BLU, dot11info.Info)
			fmt.Println(MAG, "|Base Layer      -> ", BLU, dot11info.BaseLayer)
			fmt.Println(MAG, "|PKT length      -> ", BLU, dot11info.Length)
			fmt.Println(MAG, "|------------------------------------------------------------")
			packet_layers_all(capture_main)
		}
	}
	if dot11_WEP != nil {
		fmt.Println(RED, "[ETHER_RABBIT] INFO: WARNING: ", BLU, formatDate, RED, " PICKED UP POSSIBLE WEP PACKET ")
		fmt.Println(dot11_WEP.LayerContents())
		fmt.Println(dot11_WEP.LayerPayload())
		fmt.Println(dot11_WEP.LayerType())
		fmt.Println("./././././././././././/././././././././/././/./././/././././/./././/")
		packet_layers_all(capture_main)
		fmt.Println(dot11_WEP)

	}
	if dot11_DEAUTH_ID != nil {
		fmt.Println(RED, "[ETHER_RABBIT] INFO: WARNING: ", BLU, formatDate, RED, " PICKED UP DEAUTHENTICATION PACKET!!  ")
		fmt.Println(BLU, "\n_______________________________________________________________")
		fmt.Println(RED, "| LAYER PAYLOAD  -> ", dot11_DEAUTH_ID.LayerPayload())
		fmt.Println(RED, "| LAYER TYPE     -> ", dot11_DEAUTH_ID.LayerType())
		fmt.Println(RED, "| LAYER CONTENT  -> ", dot11_DEAUTH_ID.LayerContents())
		packet_layers_all(capture_main)
		fmt.Println("./././././././././././/././././././././/././/./././/././././/./././/")
		fmt.Println(dot11_DEAUTH_ID)
	}
}

func checker_pkt(pkt gopacket.Packet) {
	if pkt == nil {
		fmt.Println(RED, "[ETHER_RABBIT] INFO: ", BLU, formatDate, RED, " Have not picked up a packet yet....scanning..")
	}
}

func table4_data_STD(inter, path, filter, bpf string, is_root, root_required bool) {
	row_1 := []interface{}{inter, path, filter, bpf, is_root, root_required}
	t := gotabulate.Create([][]interface{}{row_1})
	t.SetHeaders([]string{"Interface", "FilePath", "Capture Filter", "Berkley Packet Filter", "Is user root", "Is root required"})
	t.SetEmptyString("None")
	t.SetAlign("right")
	fmt.Println(t.Render("grid"))
}

func dump_pcapNG_file(file string, hand *pcap.Handle) {
	hand, err = pcap.OpenOffline(file)
	if err != nil {
		log.Fatal(err)
	}
	defer hand.Close()
	pkt_src := gopacket.NewPacketSource(hand, hand.LinkType())
	for pkt_src := range pkt_src.Packets() {
		fmt.Println(pkt_src)
	}
}

func handeler_opener(interface_name string, snapshot int32, promiscuous bool) {
	// first test if the client is online
	flag.Parse()
	content_get, err_cltest := http.Get(url_test)
	check_err_cl(err_cltest)
	if content_get.StatusCode == 200 {
		fmt.Println(RED, "[ETHER_RABBIT] DATA: ", BLU, formatDate, WHT, " USER ONLINE, CONTINUING CAPTURE")
	}
	handeler_pcap, err = pcap.OpenLive(interface_name, snapshot, promiscuous, timeout_shot_cap)
	//
	// error func
	//main_buffer_filter := "tcp and port " + *filter
	//err = handeler_pcap.SetBPFFilter(main_buffer_filter)
	fmt.Println(RED, "[ETHER_RABBIT] DATA: ", BLU, formatDate, WHT, " LISTENING USING BPF STD FILTER -> ", *filter)

	// close the packet capture handeler
	defer handeler_pcap.Close()
	if *tru_fals_dump_pcap {
		dump_pcapNG_file(*dump_pcap, handeler_pcap)
		os.Exit(0)
	}
	packetSource := gopacket.NewPacketSource(handeler_pcap, handeler_pcap.LinkType())
	//checker_pkt(<-packetSource.Packets())
	// sig sev
	//[signal SIGSEGV: segmentation violation code=0x1 addr=0x38 pc=0x66ba12]
	for packet := range packetSource.Packets() {
		// application packet listener checker
		// ELNLI SIG SEVPAQ AT 0x812122
		if *application_type_filters == "ICMP" {
			fmt.Println(RED, "[ETHER_RABBIT] DATA: ", BLU, formatDate, WHT, " LISTENING FOR ICMP")
			cap_proc_IEEE_ICMP_PROCMOD1(packet)
			packet_layers_all(packet)
			fmt.Println("")
			fmt.Print(YEL, "########################################################################################################")
			fmt.Print("\n\n", packet, "\n\n")
			fmt.Print(YEL, "########################################################################################################")
		}
		if *application_type_filters == "tcp" {
			DETECT_TCP := packet.Layer(TCP_)
			if DETECT_TCP != nil {
				tcp, _ := DETECT_TCP.(*layers.TCP)
				fmt.Printf("\n\033[31m[ETHER_RABBIT] \033[35mP-CHAIN \033[39m FROM <%s>--TO-PORT--<%s> \n", tcp.SrcPort, tcp.DstPort)
				fmt.Println("Sequence-Num |=> ", tcp.Seq)
				fmt.Println("Destination  |=> ", tcp.DstPort)
				fmt.Println("Source       |=> ", tcp.SrcPort)
				fmt.Println("ACK          |=> ", tcp.Ack)
				fmt.Println("Data-Offset  |=> ", tcp.DataOffset)
				fmt.Println("Urg          |=> ", tcp.Urgent)
				fmt.Println("Checksum     |=> ", tcp.Checksum)
				fmt.Println("-------------|-----------------------------------------------------")
				packet_layers_all(packet)
				fmt.Println("")
				fmt.Print(YEL, "########################################################################################################")
				fmt.Print("\n\n", packet, "\n\n")
				fmt.Print(YEL, "########################################################################################################")

			}
		}
		if *application_type_filters == "eth" {
			ETH_HEADER := packet.Layer(ETHER_)
			if ETH_HEADER != nil {
				PACK_TYPE_ETHERconst, _ := ETH_HEADER.(*layers.Ethernet)
				fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, YEL, " ETHERNET PACKET PICKED UP ")
				fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, RETURN_COLOR_FOR, " SOURCE MAC ADDR => ", PACK_TYPE_ETHERconst.SrcMAC)
				fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, RETURN_COLOR_FOR, " DST    MAC ADDR => ", PACK_TYPE_ETHERconst.DstMAC)
				fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, RETURN_COLOR_FOR, " ETHER NET TYPE  => ", PACK_TYPE_ETHERconst.EthernetType)
				packet_layers_all(packet)
				fmt.Println("")
				fmt.Print(YEL, "########################################################################################################")
				fmt.Print("\n\n", packet, "\n\n")
				fmt.Print(YEL, "########################################################################################################")

			}
		}
		if *application_type_filters == "ip" {
			ipLayer := packet.Layer(IP_)
			if ipLayer != nil {
				fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, WHT, " INTERNET PROTOCAL VERSION 4 LAYER PICKED UP!")
				ip, _ := ipLayer.(*layers.IPv4)
				fmt.Printf("\033[31m[ETHER_RABBIT] \033[35mD-CHAIN \033[39m<%s>----<%s> \n", ip.SrcIP, ip.DstIP)
				fmt.Println(BLU)
				fmt.Println("IP-Protocal |=> ", ip.Protocol)
				fmt.Println("Length      |=> ", ip.Length)
				fmt.Println("IP-ID       |=> ", ip.Id)
				fmt.Println("IP-Flags    |=> ", ip.Flags)
				fmt.Println("Frag-Offset |=> ", ip.FragOffset)
				fmt.Println("TTL         |=> ", ip.TTL)
				fmt.Println("Checksum    |=> ", ip.Checksum)
				fmt.Println("IHL         |=> ", ip.IHL)
				fmt.Println("Version     |=> ", ip.Version)
				fmt.Println("------------|-----------------------------------------------------")
				packet_layers_all(packet)
				fmt.Println("")
				fmt.Print(YEL, "########################################################################################################")
				fmt.Print("\n\n", packet, "\n\n")
				fmt.Print(YEL, "########################################################################################################")
			}
		}
		if *application_type_filters == "dhcmp" {
			determin_DCHP_req := packet.Layer(DHCMP_) // not an if statement|| packet.Layer(layers.LayerTypeDHCPv6)
			if determin_DCHP_req != nil {
				fmt.Println(packet)
			}
		}
		if *application_type_filters == "application" {
			http_determin_layer := packet.ApplicationLayer()
			if http_determin_layer != nil {
				if strings.Contains(string(http_determin_layer.Payload()), "HTTPS") {
					fmt.Println(RED, "\033[31m[ETHER_RABBIT] ", BLU, formatDate, WHT, " DETECTED HTTP APPLICATION LAYER ")
					fmt.Println(http_determin_layer.Payload())
					packet_layers_all(packet)
					fmt.Print(YEL, "########################################################################################################")
					fmt.Print("\n\n", packet, "\n\n")
					fmt.Print(YEL, "########################################################################################################")

				}
			}
		}
		if *application_type_filters == "dot11" {
			_IEEE_CAP_PROCESS_Dot11(packet)

		}
		/* else {
			LOG_PACKET_INFO(packet)
			fmt.Println(RED, " WARN: FLag.EMPTY <NILL> ")
			fmt.Print(YEL, "########################################################################################################")
			fmt.Print(packet)
			fmt.Print(YEL, "########################################################################################################")
		}
		*/
		if *application_type_filters == "LISTEN ALL" {
			LOG_PACKET_INFO(packet)
			_IEEE_CAP_PROCESS_Dot11(packet)
			cap_proc_IEEE_ICMP_PROCMOD1(packet)
			determin_DCHP_req := packet.Layer(DHCMP_)
			if determin_DCHP_req != nil {
				fmt.Println(packet)
			}
		}
	}
}

func packet_layers_all(AF_PACK gopacket.Packet) {
	for _, pkt_layers := range AF_PACK.Layers() {
		fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, RETURN_COLOR_FOR, " PACKET LAYER => ", pkt_layers.LayerType())
	}
}

func LOG_PACKET_INFO(AF_PACK gopacket.Packet) {
	// TCP LAYERS
	// FLAGS: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS

	DETECT_TCP := AF_PACK.Layer(layers.LayerTypeTCP)
	if DETECT_TCP != nil {
		tcp, _ := DETECT_TCP.(*layers.TCP)
		fmt.Printf("\033[31m[ETHER_RABBIT] \033[35mP-CHAIN \033[39m FROM <%s>--TO-PORT--<%s> \n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence-Num |=> ", tcp.Seq)
		fmt.Println("Destination  |=> ", tcp.DstPort)
		fmt.Println("Source       |=> ", tcp.SrcPort)
		fmt.Println("ACK          |=> ", tcp.Ack)
		fmt.Println("Data-Offset  |=> ", tcp.DataOffset)
		fmt.Println("Urg          |=> ", tcp.Urgent)
		fmt.Println("Checksum     |=> ", tcp.Checksum)
		fmt.Println("-------------|-----------------------------------------------------")
		fmt.Println()
	}

	// ETHERNET LAYER

	ETH_HEADER := AF_PACK.Layer(layers.LayerTypeEthernet)
	if ETH_HEADER != nil {
		PACK_TYPE_ETHERconst, _ := ETH_HEADER.(*layers.Ethernet)
		fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, YEL, " ETHERNET PACKET PICKED UP ")
		fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, RETURN_COLOR_FOR, " SOURCE MAC ADDR => ", PACK_TYPE_ETHERconst.SrcMAC)
		fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, RETURN_COLOR_FOR, " DST    MAC ADDR => ", PACK_TYPE_ETHERconst.DstMAC)
		fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, RETURN_COLOR_FOR, " ETHER NET TYPE  => ", PACK_TYPE_ETHERconst.EthernetType)
	}

	ipLayer := AF_PACK.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println(RED, "[ETHER_RABBIT] DATA:    ", BLU, formatDate, WHT, " INTERNET PROTOCAL VERSION 4 LAYER PICKED UP!")
		ip, _ := ipLayer.(*layers.IPv4)
		fmt.Printf("\033[31m[ETHER_RABBIT] \033[35mD-CHAIN \033[39m<%s>----<%s> \n", ip.SrcIP, ip.DstIP)
		fmt.Println(BLU)
		fmt.Println("IP-Protocal |=> ", ip.Protocol)
		fmt.Println("Length      |=> ", ip.Length)
		fmt.Println("IP-ID       |=> ", ip.Id)
		fmt.Println("IP-Flags    |=> ", ip.Flags)
		fmt.Println("Frag-Offset |=> ", ip.FragOffset)
		fmt.Println("TTL         |=> ", ip.TTL)
		fmt.Println("Checksum    |=> ", ip.Checksum)
		fmt.Println("IHL         |=> ", ip.IHL)
		fmt.Println("Version     |=> ", ip.Version)
		fmt.Println("------------|-----------------------------------------------------")
	}

	// packet layer identification
	fmt.Println(RED, "\n[ETHER_RABBIT] PKT_INF: ", BLU, formatDate, RETURN_COLOR_FOR, "PACKET LAYERS")

	for _, PKT_LAYER := range AF_PACK.Layers() {
		fmt.Println(RED, "[ETHER_RABBIT] PKT_INF: ", BLU, formatDate, RETURN_COLOR_FOR, "LAYER -> ", PKT_LAYER.LayerType())
	}

	http_determin_layer := AF_PACK.ApplicationLayer()
	if http_determin_layer != nil {
		if strings.Contains(string(http_determin_layer.Payload()), "HTTPS") {
			fmt.Println(RED, "\033[31m[ETHER_RABBIT] ", BLU, formatDate, WHT, " DETECTED HTTP APPLICATION LAYER ")
			fmt.Println(http_determin_layer.Payload())
		}
	}
	if err := AF_PACK.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	}
}

func main() {
	flag.Parse()
	banner("banner.txt")
	table4_data_STD(*interface_use_flag, "modules/sniifed_ftp_go_other_bg", *application_type_filters, *filter, true, true)
	if *interface_use_flag == "" {
		handeler_opener(defualt_interface, snapshot_main, false)
	} else {
		handeler_opener(*interface_use_flag, snapshot_main, false)
	}
}
