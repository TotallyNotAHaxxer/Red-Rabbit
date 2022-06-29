/*
| Protocal/Packet Type | Color code  | Color name                         |
| -------------------- | ----------  | ---------------------------------- |
|        IGMP          |  \033[4;32m | Underline green                    |
|        ICMP          |  \033[4;33m | Underline Yellow                   |
|        TCP/IP        |  \033[4;31m | Underline RED                      |
|        Ethernet      |  \033[4;37m | Underline white                    |
|        UDP           |  \033[4;35m | Underline Magenta                  |
|        DHCP          |  \033[4;34m | Underline blue                     |
|        MLD_V4        |  \033[4;36m | Underline cyan                     |
|        IPAv6         |  \033[4;30m | Underline black                    |
|        IP6 Hop By Hop|  \033[4;30m | Underline black w red foreground   |
|        Stop Underline|  \033[4349m]| underline none                     |
*/
package IEEE_Wire_Capture

import (
	"bytes"
	"fmt"
	v "main/modg/colors"
	IEEE_CONSTANTS "main/modg/constants"
	macs "main/modg/dependant-sub" // macs.go made by google, taken to prevent need to add another go get
	IEEE "main/modg/scripts/IEEE-802.11/system-con"
	ec "main/modg/warnings"

	"net"
	"os"
	"time"

	"main/modg/system-runscript"

	opc "main/modg/copt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/pflag"
)

var (
	rr6f  opc.RR6_options
	flags = pflag.FlagSet{SortFlags: false}
)

func Check_and_get_All_Values() {
	IEEE.HOME_Interfaces()
}

func HOME_Interfaces() {
	ifaces, err := net.Interfaces()
	ec.Warning_advanced("<RR6> Net Module: Could not grab interface names", v.REDHB, 1, false, false, true, err, 1, 233, "")
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		ec.Warning_advanced("<RR6> Net Module: Could not stat addresses for interfaces", v.REDHB, 1, false, false, true, err, 1, 233, "")
		for _, a := range addrs {
			fmt.Println("\033[38;5;21m<RR6> Net Module: Found Device  | \033[38;5;198m", i.Name, "\t | \033[38;5;21mAddr:", a)
		}
	}
}

// ftp sniffing ONLINE, NO PCAP FILE NEEDED FOR THIS SECTION
func FTP_Harvester_Layer_Application(packet gopacket.Packet) {
	app := packet.ApplicationLayer()
	if app != nil {
		payload := app.Payload()
		dst := packet.NetworkLayer().NetworkFlow().Dst()
		if bytes.Contains(payload, []byte("USER")) {
			fmt.Println(v.RED, "[RR6] [FTP_Authentication] Found FTP Username \t| ", v.BLU, system.FormatDate, "\t | ")
			fmt.Print(v.MAG, "---------------------------------------------------------------------------------------")
			fmt.Print(v.MAG, "| Destination -> ", dst)
			fmt.Print(v.MAG, "| Payload     -> ", string(payload))
			fmt.Print(v.MAG, "---------------------------------------------------------------------------------------")
			fmt.Println("\n\n", packet)
		} else if bytes.Contains(payload, []byte("PASS")) {
			fmt.Println(v.RED, "[RR6] [FTP_Authentication] Found FTP Password \t| ", v.BLU, system.FormatDate, "\t | ")
			fmt.Print(v.MAG, "---------------------------------------------------------------------------------------")
			fmt.Print(v.MAG, "| Destination -> ", dst)
			fmt.Print(v.MAG, "| Payload     -> ", string(payload))
			fmt.Print(v.MAG, "---------------------------------------------------------------------------------------")
			fmt.Println("\n\n", packet)
		}
	}
}

// harvester
func SMTP_Harvester(pkt gopacket.Packet) {
	layer := pkt.ApplicationLayer()
	if layer != nil {
		p := layer.Payload()
		dst := pkt.NetworkLayer().NetworkFlow().Dst()
		src := pkt.NetworkLayer().NetworkFlow().Src()
		End_t := pkt.NetworkLayer().NetworkFlow().EndpointType()
		End_t_dst := pkt.NetworkLayer().NetworkFlow().Dst().EndpointType()
		End_t_src := pkt.NetworkLayer().NetworkFlow().Src().EndpointType()
		if bytes.Contains(p, []byte("user")) {
			fmt.Println("\n------------------------------------------------------------------------------- ")
			fmt.Println("\033[38;5;198m| Packet type              |-> \033[38;5;21mApplication packet ")
			fmt.Println("\033[38;5;198m| Packet SRC               |-> \033[38;5;21m", src)
			fmt.Println("\033[38;5;198m| Packet DST               |-> \033[38;5;21m", dst)
			fmt.Println("\033[38;5;198m| Packet endpoint type     |-> \033[38;5;21m", End_t)
			fmt.Println("\033[38;5;198m| Packet SRC endpoint type |-> \033[38;5;21m", End_t_src)
			fmt.Println("\033[38;5;198m| Packet DST endpoint type |-> \033[38;5;21m", End_t_dst)
			fmt.Println("\033[38;5;198m| SMTP Username            |-> \033[38;5;21m", string(p))
		}
		if bytes.Contains(p, []byte("pass")) {
			fmt.Println("\n------------------------------------------------------------------------------- ")
			fmt.Println("\033[38;5;198m| Packet type              |-> \033[38;5;21mApplication packet ")
			fmt.Println("\033[38;5;198m| Packet SRC               |-> \033[38;5;21m", src)
			fmt.Println("\033[38;5;198m| Packet DST               |-> \033[38;5;21m", dst)
			fmt.Println("\033[38;5;198m| Packet endpoint type     |-> \033[38;5;21m", End_t)
			fmt.Println("\033[38;5;198m| Packet SRC endpoint type |-> \033[38;5;21m", End_t_src)
			fmt.Println("\033[38;5;198m| Packet DST endpoint type |-> \033[38;5;21m", End_t_dst)
			fmt.Println("\033[38;5;198m| SMTP Username            |-> \033[38;5;21m", string(p))
		}
	}
}

// ftp credential analyzation
func Ftp_starter_credential_Applayer(buffer int32, filter, sniffc string) {
	handler, err := pcap.OpenLive(sniffc, buffer, false, pcap.BlockForever)
	ec.Warning_advanced("<RR6> Net Module: Could not open a live pcap capture for interface, this might have been due to premission errors or that the device does not exist", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer handler.Close()
	if err := handler.SetBPFFilter(filter); err != nil {
		fmt.Println("<RR6> Errors Module: Could not set a BPF Filter ( Berkly Packet Filter ) got error -> ", err)
		os.Exit(0)
	}
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Started listener    |-> \033[4;30m(IGNORE VALUE)\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Set given BPF       |-> \033[4;30m", filter, "\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Set given buffer    |-> \033[4;30m", buffer, "\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Listen User payload |-> \033[4;30m117 115 101 114", filter, "\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Listen User payload |-> \033[4;30m112 97 115 115", filter, "\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Listening for pkt   |-> \033[4;30mApplication", "\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Listening for pkt_  |-> \033[4;30mApplication SMTP\033[0;39m\033[38;5;198m")
	source := gopacket.NewPacketSource(handler, handler.LinkType())
	for packet := range source.Packets() {
		FTP_Harvester_Layer_Application(packet)
	}
}

//smtp credential analyzation
func Smtp_Starter_crednetial_listener(buffer int32, filter, sniffc string) {
	handler, x := pcap.OpenLive(sniffc, buffer, false, pcap.BlockForever)
	ec.Warning_advanced("<RR6> Net Module: Could not open a live pcap capture for interface, this might have been due to premission errors or that the device does not exist", v.REDHB, 1, false, false, true, x, 1, 233, "")
	defer handler.Close()
	if x := handler.SetBPFFilter(filter); x != nil {
		fmt.Println(v.REDHB, "<RR6> Errors Module: Could not set BPF Filter -> ", x)
	}

	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Started listener    |-> \033[4;30m(IGNORE VALUE)\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Set given BPF       |-> \033[4;30m", filter, "\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Set given buffer    |-> \033[4;30m", buffer, "\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Listen User payload |-> \033[4;30m117 115 101 114", filter, "\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Listen User payload |-> \033[4;30m112 97 115 115", filter, "\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Listening for pkt   |-> \033[4;30mApplication", "\033[0;39m\033[38;5;198m")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting    : Listening for pkt_  |-> \033[4;30mApplication SMTP\033[0;39m\033[38;5;198m")
	src := gopacket.NewPacketSource(handler, handler.LinkType())
	for pkt := range src.Packets() {
		SMTP_Harvester(pkt)
	}
}

// parse packet layers
func parse_packet_layers(packet gopacket.Packet) {
	for _, l := range packet.Layers() {
		fmt.Println("\033[38;5;198m| Packet (Layer)          |-> \033[38;5;21m", l.LayerType())
	}
}

// parse all packet information, this will be the main parser
func Parse_All_packets(packet_type string, packet_to_parse gopacket.Packet, show_packet_full bool, counter int) {
	//app := packet_to_parse.ApplicationLayer()
	udplayer := packet_to_parse.Layer(IEEE_CONSTANTS.UDP_)
	ipLayer := packet_to_parse.Layer(IEEE_CONSTANTS.IP_)
	ether_layer := packet_to_parse.Layer(IEEE_CONSTANTS.ETHER_)
	dhcp_layer_v4 := packet_to_parse.Layer(IEEE_CONSTANTS.DHCMP_)
	dhcp_layer_v6 := packet_to_parse.Layer(IEEE_CONSTANTS.DHCMP_6)
	DETECT_TCP := packet_to_parse.Layer(IEEE_CONSTANTS.TCP_)
	ICMP_V6_PROC := packet_to_parse.Layer(IEEE_CONSTANTS.ICMP_REDI_)
	ICMP_6_SOL_m := packet_to_parse.Layer(IEEE_CONSTANTS.ICMP_6_SOL_)
	ICMP_4 := packet_to_parse.Layer(IEEE_CONSTANTS.ICP_4_)
	ECHO_ := packet_to_parse.Layer(IEEE_CONSTANTS.ICMP_ECH_)
	IGMP := packet_to_parse.Layer(IEEE_CONSTANTS.ICMP_IGMP)
	MLD_Report_message := packet_to_parse.Layer(IEEE_CONSTANTS.MLDv2MulticastListenerReport)
	MLD_Report_Query := packet_to_parse.Layer(IEEE_CONSTANTS.Multicase_listener)
	IP6_ := packet_to_parse.Layer(IEEE_CONSTANTS.IP_6)
	IP6_Hop := packet_to_parse.Layer(IEEE_CONSTANTS.IPv6HopByHop)
	IP6_router_AD := packet_to_parse.Layer(IEEE_CONSTANTS.IPv6_Router_advertisement)
	// functions
	ip6, _ := IP6_.(*layers.IPv6)
	ip6_hop, _ := IP6_Hop.(*layers.IPv6HopByHop)
	IP6_router_Ad, _ := IP6_router_AD.(*layers.ICMPv6RouterAdvertisement)
	MulticastListener, _ := MLD_Report_message.(*layers.MLDv2MulticastListenerReportMessage)
	MulticastListener_query, _ := MLD_Report_Query.(*layers.MLDv2MulticastListenerQueryMessage)
	dhcp, _ := dhcp_layer_v4.(*layers.DHCPv4)
	dhcp6, _ := dhcp_layer_v6.(*layers.DHCPv6)
	igmp, _ := IGMP.(*layers.IGMP)
	icmp6, _ := ICMP_V6_PROC.(*layers.ICMPv6)
	icmp6SOL, _ := ICMP_6_SOL_m.(*layers.ICMPv6NeighborSolicitation)
	icmp4, _ := ICMP_V6_PROC.(*layers.ICMPv4)
	icmpECHO, _ := ICMP_V6_PROC.(*layers.ICMPv6Echo)
	eth, _ := ether_layer.(*layers.Ethernet)
	ip, _ := ipLayer.(*layers.IPv4)
	udp_function, _ := udplayer.(*layers.UDP)
	switch packet_type {
	case "eth":
		ETH_Header := packet_to_parse.Layer(IEEE_CONSTANTS.ETHER_)
		if ETH_Header != nil {
			PACK_TYPE_ETHERconst, _ := ETH_Header.(*layers.Ethernet)
			fmt.Println("\n------------------------------------------------------------------------------- ")
			fmt.Println("\033[38;5;198m| Packet type             |-> \033[38;5;21mEthernet packet ")
			fmt.Println("\033[38;5;198m| Packet Source Mac       |-> \033[38;5;21m", PACK_TYPE_ETHERconst.SrcMAC)
			fmt.Println("\033[38;5;198m| Packet Destination Mac  |-> \033[38;5;21m", PACK_TYPE_ETHERconst.DstMAC)
			fmt.Println("\033[38;5;198m| Packet Ether Type       |-> \033[38;5;21m", PACK_TYPE_ETHERconst.EthernetType)
			fmt.Println("\033[38;5;198m| Packet Length           |-> \033[38;5;21m", PACK_TYPE_ETHERconst.Length)
			parse_packet_layers(packet_to_parse)
			if show_packet_full {
				fmt.Println(packet_to_parse)
			} else {
				fmt.Println("\033[38;5;198m| User Information        |-> \033[38;5;21m Packet output disabled not outputting packet")
			}
		}
	case "tcp":
		if DETECT_TCP != nil {
			tcp, _ := DETECT_TCP.(*layers.TCP)
			fmt.Println("\n-------------------------------------------------------------------------")
			fmt.Printf("\n\033[35mCHAIN \033[39m FROM <%s>--TO-PORT--<%s> \n", tcp.SrcPort, tcp.DstPort)
			fmt.Println("\033[38;5;198m| Packet type             |-> \033[38;5;21mTCP/IP packet ")
			fmt.Println("\033[38;5;198m| Sequence Number         |-> \033[38;5;21m", tcp.Seq)
			fmt.Println("\033[38;5;198m| ACK                     |-> \033[38;5;21m", tcp.Ack)
			fmt.Println("\033[38;5;198m| Data Offset             |-> \033[38;5;21m", tcp.DataOffset)
			fmt.Println("\033[38;5;198m| Checksum                |-> \033[38;5;21m", tcp.Checksum)
			fmt.Println("\033[38;5;198m| URG                     |-> \033[38;5;21m", tcp.URG)
			fmt.Println("\033[38;5;198m| TCP/IP Source Port      |-> \033[38;5;21m", tcp.SrcPort)
			fmt.Println("\033[38;5;198m| TCP/IP Destination Port |-> \033[38;5;21m", tcp.DstPort)
			parse_packet_layers(packet_to_parse)
			fmt.Println("---------------------------------------------------------------------------")
			if show_packet_full {
				fmt.Println(packet_to_parse)
			} else {
				fmt.Println("\033[38;5;198m| User Information        |-> \033[38;5;21m Packet output disabled not outputting packet")
			}
		}
	case "icmp":
		if ICMP_V6_PROC != nil {
			fmt.Println("\033[38;5;198m| Packet type             |-> \033[38;5;21m ", icmp6.LayerType())
			fmt.Println("\033[38;5;198m| Packet Base Layer       |-> \033[38;5;21m ", icmp6.BaseLayer)
			fmt.Println("\033[38;5;198m| Packet contents         |-> \033[38;5;21m ", icmp6.Contents)
			parse_packet_layers(packet_to_parse)
			fmt.Println("---------------------------------------------------------------------------")
			if show_packet_full {
				fmt.Println(packet_to_parse)
			} else {
				fmt.Println("\033[38;5;198m| User Information        |-> \033[38;5;21m Packet output disabled not outputting packet")
			}
		}
		if ICMP_6_SOL_m != nil {
			fmt.Println("\033[38;5;198m| Packet type                    |-> \033[38;5;21m ", icmp6SOL.LayerType())
			fmt.Println("\033[38;5;198m| Packet Target Address          |-> \033[38;5;21m ", icmp6SOL.TargetAddress)
			fmt.Println("\033[38;5;198m| Packet Target Address Mask     |-> \033[38;5;21m ", icmp6SOL.TargetAddress.DefaultMask())
			fmt.Println("\033[38;5;198m| Packet target Address unicast  |-> \033[38;5;21m ", icmp6SOL.TargetAddress.IsGlobalUnicast())
			fmt.Println("\033[38;5;198m| Packet contents                |-> \033[38;5;21m ", icmp6SOL.LayerContents())
			parse_packet_layers(packet_to_parse)
			fmt.Println("---------------------------------------------------------------------------")
			if show_packet_full {
				fmt.Println(packet_to_parse)
			} else {
				fmt.Println("\033[38;5;198m| User Information        |-> \033[38;5;21m Packet output disabled not outputting packet")
			}
		}
		if ICMP_4 != nil {
			fmt.Println("\033[38;5;198m| Packet type             |-> \033[38;5;21m ", ICMP_4.LayerType())
			fmt.Println("\033[38;5;198m| Packet Base Layer       |-> \033[38;5;21m ", icmp4.BaseLayer)
			fmt.Println("\033[38;5;198m| Packet Checksum         |-> \033[38;5;21m ", icmp4.Checksum)
			fmt.Println("\033[38;5;198m| Packet ID               |-> \033[38;5;21m ", icmp4.Id)
			fmt.Println("\033[38;5;198m| Packet Sequence         |-> \033[38;5;21m ", icmp4.Seq)
			fmt.Println("\033[38;5;198m| Packet Type Code        |-> \033[38;5;21m ", icmp4.TypeCode)
			fmt.Println("\033[38;5;198m| Packet contents         |-> \033[38;5;21m ", ICMP_4.LayerContents())
			parse_packet_layers(packet_to_parse)
			fmt.Println("---------------------------------------------------------------------------")
			if show_packet_full {
				fmt.Println(packet_to_parse)
			} else {
				fmt.Println("\033[38;5;198m| User Information        |-> \033[38;5;21m Packet output disabled not outputting packet")
			}
		}
		if ECHO_ != nil {
			fmt.Println("\033[38;5;198m| Packet type             |-> \033[38;5;21m ", ECHO_.LayerType())
			fmt.Println("\033[38;5;198m| Packet Base Layer       |-> \033[38;5;21m ", icmpECHO.BaseLayer)
			fmt.Println("\033[38;5;198m| Packet Identifier       |-> \033[38;5;21m ", icmpECHO.Identifier)
			fmt.Println("\033[38;5;198m| Packet Sequence Number  |-> \033[38;5;21m ", icmpECHO.SeqNumber)
			fmt.Println("\033[38;5;198m| Packet payloads         |-> \033[38;5;21m ", ECHO_.LayerPayload())
			fmt.Println("\033[38;5;198m| Packet contents         |-> \033[38;5;21m ", ECHO_.LayerContents())
			parse_packet_layers(packet_to_parse)
			fmt.Println("---------------------------------------------------------------------------")
			if show_packet_full {
				fmt.Println(packet_to_parse)
			} else {
				fmt.Println("\033[38;5;198m| User Information        |-> \033[38;5;21m Packet output disabled not outputting packet")
			}
		}
	case "ip":
		if ipLayer != nil {
			fmt.Printf("\n\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| New Packet \n\033[35mCHAIN         | \033[39m<%s>----<%s> |\n", ip.SrcIP, ip.DstIP)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) IP protocal      |-> \033[38;5;21m ", ip.Protocol)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) Length           |-> \033[38;5;21m ", ip.Length)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) IP-ID            |-> \033[38;5;21m ", ip.Id)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) Fragment Offset  |-> \033[38;5;21m ", ip.FragOffset)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) TTL              |-> \033[38;5;21m ", ip.TTL)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) Checksum         |-> \033[38;5;21m ", ip.Checksum)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) IHL              |-> \033[38;5;21m ", ip.IHL)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) Version          |-> \033[38;5;21m ", ip.Version)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) Source Address   |-> \033[38;5;21m ", ip.SrcIP)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) Destination Addr |-> \033[38;5;21m ", ip.DstIP)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) Flags / INF      |-> \033[38;5;21m ", ip.Flags)
			fmt.Println("\033[38;5;198m| Packet (\033[4;31mIP\033[0;39m\033[38;5;198m) Padding          |-> \033[38;5;21m ", ip.Padding)
			if ether_layer != nil {
				fmt.Println("\033[38;5;198m| Packet (\033[4;37mether\033[0;39m\033[38;5;198m) SRC Mac       |-> \033[38;5;21m ", eth.SrcMAC)
				fmt.Println("\033[38;5;198m| Packet (\033[4;37mether\033[0;39m\033[38;5;198m) DST Mac       |-> \033[38;5;21m ", eth.SrcMAC)
			}
			if udplayer != nil {
				fmt.Println("\033[38;5;198m| Packet (\033[4;35mUDP\033[0;39m\033[38;5;198m) DST Port        |-> \033[38;5;21m ", udp_function.DstPort)
				fmt.Println("\033[38;5;198m| Packet (\033[4;35mUDP\033[0;39m\033[38;5;198m) SRC Port        |-> \033[38;5;21m ", udp_function.SrcPort)
				fmt.Println("\033[38;5;198m| Packet (\033[4;35mUDP\033[0;39m\033[38;5;198m) Checksum        |-> \033[38;5;21m ", udp_function.Checksum)
				fmt.Printf("\033[35m|P-Chain (\033[1;31mPort Chain\033[0;39m)          |->  \033[39m<%s>----<%s>\n", udp_function.SrcPort, udp_function.DstPort)
			}
			if IGMP != nil {
				if igmp != nil {
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Base Layer     |-> \033[38;5;21m ", igmp.BaseLayer)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Checksum       |-> \033[38;5;21m ", igmp.Checksum)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Group Addr     |-> \033[38;5;21m ", igmp.GroupAddress)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Group Records  |-> \033[38;5;21m ", igmp.GroupRecords)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Interval Time  |-> \033[38;5;21m ", igmp.IntervalTime)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Max Response T |-> \033[38;5;21m ", igmp.MaxResponseTime)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Num group Rec  |-> \033[38;5;21m ", igmp.NumberOfGroupRecords)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Number of src's|-> \033[38;5;21m ", igmp.NumberOfSources)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Robust Value   |-> \033[38;5;21m ", igmp.RobustnessValue)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Source Addr    |-> \033[38;5;21m ", igmp.SourceAddresses)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Supress Router |-> \033[38;5;21m ", igmp.SupressRouterProcessing)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Version        |-> \033[38;5;21m ", igmp.Version)
					fmt.Println("\033[38;5;198m| Packet (\033[4;32mIGMP\033[0;39m\033[38;5;198m) Type           |-> \033[38;5;21m ", igmp.Type)
				}
			}
			parse_packet_layers(packet_to_parse)
			if show_packet_full {
				fmt.Println(packet_to_parse)
			} else {
				fmt.Println("\033[38;5;198m| User Information        |-> \033[38;5;21m Packet output disabled not outputting packet")
			}
			fmt.Println("------------------------------------------------------------------------------------------------")

		}
	case "dhcmp":
		if dhcp != nil {
			fmt.Println("____________________________________________NEW DHCP PACKET_____________________________________________________________")
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Packet Number         |-> \033[38;5;21m ", counter)
			fmt.Printf("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Base Layer            |-> \033[38;5;21m %s\n", dhcp.BaseLayer)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Client HW Address     |-> \033[38;5;21m ", dhcp.ClientHWAddr.String())
			mac := string(dhcp.ClientHWAddr)
			a, e := OUI_Mac(mac)
			if e != nil {
				fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Client OUI HW Addr    |-> \033[38;5;21m Could not trace mac OUI")
			}
			if a == "" {
				fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Client OUI HW Addr    |-> \033[38;5;21m  No manufacturer found")
			}
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Client OUI HW Addr    |-> \033[38;5;21m ", a)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Client IP Address     |-> \033[38;5;21m ", dhcp.ClientIP)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Your Client IPA       |-> \033[38;5;21m ", dhcp.YourClientIP)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Hardware Length       |-> \033[38;5;21m ", dhcp.HardwareLen)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Hardware Options      |-> \033[38;5;21m ", dhcp.HardwareOpts)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Hardware Type         |-> \033[38;5;21m ", dhcp.HardwareType)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Next Server IP        |-> \033[38;5;21m ", dhcp.NextServerIP)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Operation             |-> \033[38;5;21m ", dhcp.Operation)
			if len(dhcp.Options.String()) >= 3000 {
				fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Options               |-> \033[38;5;21m Length to long Silencing output.....")
			} else {
				fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Options               |-> \033[38;5;21m ", dhcp.Options.String())
			}
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Relay Agent IP        |-> \033[38;5;21m ", dhcp.RelayAgentIP)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Secs                  |-> \033[38;5;21m ", dhcp.Secs)
			fmt.Printf("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Server name           |-> \033[38;5;21m %s\n", dhcp.ServerName)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) XID                   |-> \033[38;5;21m ", dhcp.Xid)
			fmt.Printf("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Contents              |-> \033[38;5;21m %s\n", string(dhcp.Contents))
			fmt.Printf("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Files                 |-> \033[38;5;21m %s\n", dhcp.File)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Flags                 |-> \033[38;5;21m ", dhcp.Flags)
			parse_packet_layers(packet_to_parse)
			if show_packet_full {
				fmt.Println(packet_to_parse)
			} else {
				fmt.Println("\033[38;5;198m| User Information        |-> \033[38;5;21m Packet output disabled not outputting packet")
			}
		}
		if dhcp_layer_v6 != nil {
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Base Layer            |-> \033[38;5;21m ", dhcp6.BaseLayer)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Contents              |-> \033[38;5;21m ", dhcp6.Contents)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Hop Count             |-> \033[38;5;21m ", dhcp6.HopCount)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Link Address          |-> \033[38;5;21m ", dhcp6.LinkAddr)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Message Type          |-> \033[38;5;21m ", dhcp6.MsgType)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Transaction ID        |-> \033[38;5;21m ", dhcp6.TransactionID)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Peer Address          |-> \033[38;5;21m ", dhcp6.PeerAddr)
			fmt.Println("\033[38;5;198m| Packet (\033[4;34mDHCP\033[0;39m\033[38;5;198m) Payload               |-> \033[38;5;21m ", dhcp6.Payload)
			parse_packet_layers(packet_to_parse)
			if show_packet_full {
				fmt.Println(packet_to_parse)
			} else {
				fmt.Println("\033[38;5;198m| User Information        |-> \033[38;5;21m Packet output disabled not outputting packet")
			}
		}
	case "app":
		fmt.Print("\n")
		/*
			Defualt Packet layers that appear in a Application packet
			IPv6HopByHop
			IPv6
			ICMPv6RouterAdvertisement
			MLDv2MulticastListenerReport
			ICMPv6
			Ethernet
			UDP

			We will just pick apart every average application layer
			\033[4;36m
		*/
		if MLD_Report_message != nil {
			fmt.Println("----------------------------------------------- NEW PACKET (\033[4;36mMultiCast Listener Report\033[4;39m) ----------------------------------------------------------")
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Listener Report\033[0;39m\033[38;5;198m) Base layer            |-> \033[38;5;21m ", MulticastListener.BaseLayer)
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Listener Report\033[0;39m\033[38;5;198m) Payload               |-> \033[38;5;21m ", MulticastListener.Payload)
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Listener Report\033[0;39m\033[38;5;198m) Address Records       |-> \033[38;5;21m ", MulticastListener.MulticastAddressRecords)
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Listener Report\033[0;39m\033[38;5;198m) Number of ADDR Records|-> \033[38;5;21m ", MulticastListener.NumberOfMulticastAddressRecords)
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Listener Report\033[0;39m\033[38;5;198m) Contents              |-> \033[38;5;21m ", MulticastListener.Contents)
		}
		if MLD_Report_Query != nil {
			fmt.Println("----------------------------------------------- NEW PACKET (\033[4;36mMulticast Query Response\033[4;39m) ----------------------------------------------------------")
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Query Response\033[0;39m\033[38;5;198m) Supress Router proc   |-> \033[38;5;21m ", MulticastListener_query.SuppressRoutersideProcessing)
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Query Response\033[0;39m\033[38;5;198m) Source Address        |-> \033[38;5;21m ", MulticastListener_query.SourceAddresses)
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Query Response\033[0;39m\033[38;5;198m) Number Of Sources     |-> \033[38;5;21m ", MulticastListener_query.NumberOfSources)
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Query Response\033[0;39m\033[38;5;198m) Multi Cast Address    |-> \033[38;5;21m ", MulticastListener_query.MulticastAddress)
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Query Response\033[0;39m\033[38;5;198m) Maximum Response code |-> \033[38;5;21m ", MulticastListener_query.MaximumResponseCode)
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Query Response\033[0;39m\033[38;5;198m) Query Interval Code   |-> \033[38;5;21m ", MulticastListener_query.QueriersQueryIntervalCode)
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Query Response\033[0;39m\033[38;5;198m) Queries Robustnes     |-> \033[38;5;21m ", MulticastListener_query.QueriersRobustnessVariable)
			fmt.Println("\033[38;5;198m| Packet (\033[4;36mMulticast Query Response\033[0;39m\033[38;5;198m) Payload               |-> \033[38;5;21m ", MulticastListener_query.Payload)
		}
		if IP6_ != nil {
			fmt.Println("----------------------------------------------- NEW PACKET (\033[4;30mIPA Version 6\033[0;39m) ----------------------------------------------------------")
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP Version 6\033[0;39m\033[38;5;198m) Version         |-> \033[38;5;21m ", ip6.Version)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP Version 6\033[0;39m\033[38;5;198m) Traffic Class   |-> \033[38;5;21m ", ip6.TrafficClass)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP Version 6\033[0;39m\033[38;5;198m) Source IPA      |-> \033[38;5;21m ", ip6.SrcIP)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP Version 6\033[0;39m\033[38;5;198m) Destination IPA |-> \033[38;5;21m ", ip6.DstIP)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP Version 6\033[0;39m\033[38;5;198m) Next Header     |-> \033[38;5;21m ", ip6.NextHeader)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP Version 6\033[0;39m\033[38;5;198m) Length          |-> \033[38;5;21m ", ip6.Length)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP Version 6\033[0;39m\033[38;5;198m) Hop Limit       |-> \033[38;5;21m ", ip6.HopLimit)
			fmt.Printf("\033[38;5;198m| Packet (\033[4;30mIP Version 6\033[0;39m\033[38;5;198m) Hop by Hop      |-> \033[38;5;21m %v\n", ip6.HopByHop)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP Version 6\033[0;39m\033[38;5;198m) Flow Label      |-> \033[38;5;21m ", ip6.FlowLabel)
			fmt.Printf("\033[38;5;198m| Packet (\033[4;30mIP Version 6\033[0;39m\033[38;5;198m) Base Layer      |-> \033[38;5;21m %s\n", ip6.BaseLayer)
			fmt.Printf("\033[38;5;198m| Packet (\033[4;30mIP Version 6\033[0;39m\033[38;5;198m) Contents        |-> \033[38;5;21m %s\n", ip6.Contents)
		}
		if IP6_Hop != nil {
			fmt.Println("----------------------------------------------- NEW PACKET (\033[4;30mIP6 Hop By Hop\033[4;39m) ----------------------------------------------------------")
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP6 Hop By Hop\033[0;39m\033[38;5;198m) Contents        |-> \033[38;5;21m ", ip6_hop.Contents)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP6 Hop By Hop\033[0;39m\033[38;5;198m) Base layer      |-> \033[38;5;21m ", ip6_hop.BaseLayer)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP6 Hop By Hop\033[0;39m\033[38;5;198m) Full Length     |-> \033[38;5;21m ", ip6_hop.ActualLength)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP6 Hop By Hop\033[0;39m\033[38;5;198m) Header Lengths  |-> \033[38;5;21m ", ip6_hop.HeaderLength)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP6 Hop By Hop\033[0;39m\033[38;5;198m) Header Next     |-> \033[38;5;21m ", ip6_hop.NextHeader)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP6 Hop By Hop\033[0;39m\033[38;5;198m) Options         |-> \033[38;5;21m ", ip6_hop.Options)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30mIP6 Hop By Hop\033[0;39m\033[38;5;198m) Payload         |-> \033[38;5;21m ", ip6_hop.Payload)
		}
		if IP6_router_AD != nil {
			fmt.Println("----------------------------------------------- NEW PACKET (\033[4;30m\033[31mIP6 Router Advertisement\033[39m) ----------------------------------------------------------")
			fmt.Printf("\033[38;5;198m| Packet (\033[4;30m\033[31mIP6 Router Advertisement\033[0;39m\033[38;5;198m) Base Layer           |-> \033[38;5;21m %s\n", IP6_router_Ad.BaseLayer)
			fmt.Printf("\033[38;5;198m| Packet (\033[4;30m\033[31mIP6 Router Advertisement\033[0;39m\033[38;5;198m) Contents             |-> \033[38;5;21m %s\n", IP6_router_Ad.Contents)
			fmt.Printf("\033[38;5;198m| Packet (\033[4;30m\033[31mIP6 Router Advertisement\033[0;39m\033[38;5;198m) Payload              |-> \033[38;5;21m %s\n", IP6_router_Ad.Payload)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30m\033[31mIP6 Router Advertisement\033[0;39m\033[38;5;198m) Flags                |-> \033[38;5;21m ", IP6_router_Ad.Flags)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30m\033[31mIP6 Router Advertisement\033[0;39m\033[38;5;198m) Hop limit            |-> \033[38;5;21m ", IP6_router_Ad.HopLimit)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30m\033[31mIP6 Router Advertisement\033[0;39m\033[38;5;198m) Options              |-> \033[38;5;21m ", IP6_router_Ad.Options)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30m\033[31mIP6 Router Advertisement\033[0;39m\033[38;5;198m) Reach time           |-> \033[38;5;21m ", IP6_router_Ad.ReachableTime)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30m\033[31mIP6 Router Advertisement\033[0;39m\033[38;5;198m) Retransmission Timer |-> \033[38;5;21m ", IP6_router_Ad.RetransTimer)
			fmt.Println("\033[38;5;198m| Packet (\033[4;30m\033[31mIP6 Router Advertisement\033[0;39m\033[38;5;198m) Router Lifetime      |-> \033[38;5;21m ", IP6_router_Ad.RouterLifetime)
		}

	case "dot11":
	case "all":
	}
}

func OUI_Mac(mac string) (string, error) {
	if mac, err := net.ParseMAC(mac); err == nil {
		prefix := [3]byte{
			mac[0],
			mac[1],
			mac[2],
		}
		manufacturer, good := macs.ValidMACPrefixMap[prefix]
		if good {
			return manufacturer, nil
		}
	}
	return "", nil
}

// func pcap opener
func Live_Run(iface string, snapshot int32, promisc, filter bool, BPF string, timeout time.Duration, type_to_sniff string) (string, uint) {
	HOME_Interfaces()
	pcap_handler, e := pcap.OpenLive(iface, snapshot, promisc, timeout)
	if e != nil {
		fmt.Println("<RR6> Live Wire Sniffer: Could not make a live network capture, error came out as  < 0x01 > instead of < 0x00 > got error -> ", e)
	}
	defer pcap_handler.Close()
	if filter {
		e = pcap_handler.SetBPFFilter(BPF)
		if e != nil {
			fmt.Println("<RR6> Live Wire Sniffe: Was not able to set a BPF (Berkly Packet Filter) Got error <0x01> -> ", e)
		} else {
			fmt.Println("\033[38;5;21m<RR6> Net Module: Setting Data  |  \033[38;5;198mBPF Filter\033[38;5;198m    | \033[38;5;21mSetting: \033[31m", BPF)
		}
	}
	pkt := gopacket.NewPacketSource(pcap_handler, pcap_handler.LinkType())
	fmt.Println("sniffing type -> ", type_to_sniff)
	packets := 0
	for packet := range pkt.Packets() {
		packets++
		Parse_All_packets(type_to_sniff, packet, false, packets)

	}
	return "", 0x00
}
