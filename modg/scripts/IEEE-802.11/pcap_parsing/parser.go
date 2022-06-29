package PPARSER

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	v "main/modg/colors"
	IEEE_CONSTANTS "main/modg/constants"
	opc "main/modg/copt"
	"main/modg/system-runscript"
	ec "main/modg/warnings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/pflag"
)

var (
	rr6f          opc.RR6_options
	err           error
	handeler_pcap *pcap.Handle
	flags         = pflag.FlagSet{SortFlags: false}
)

func Run_CPP(pcap_file, outputdir, exe string) {
	prg := "./" + exe
	c := exec.Command(prg, pcap_file, outputdir)
	fmt.Println("running -> ", c)
	s, x := c.Output()
	if x != nil {
		fmt.Println("<RR6> Got error when running executeable -> ", x)
	} else {
		fmt.Print(string(s))
	}
}

// pcap file parser

// dumping PCAP parser, short [BYTE]
func Pcap_parser_OFFLINE_byte(filenametakein string, payload1 string) {
	var handle *pcap.Handle
	handle, err = pcap.OpenOffline(filenametakein)
	ec.Warning_advanced("<RR6> IEEE-802.11 Packet Parsing Module: Could not parse, open, or log file -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	counter := 0
	for packet := range packetSource.Packets() {
		counter++

		if packet.ApplicationLayer() != nil {
			payload := packet.ApplicationLayer().LayerContents()
			if bytes.Contains(payload, []byte(payload1)) {
				src := packet.NetworkLayer().NetworkFlow().Src()
				dst := packet.NetworkLayer().NetworkFlow().Dst()
				fmt.Println(v.RED, "------------------- TOP HEADER DATA ---------------------------")
				fmt.Println(v.MAG, "| Source      Address |> \033[49m", v.BLKHB, src, "\033[49m")
				fmt.Println(v.MAG, "| Destination Address |> \033[49m", v.BLKHB, dst, "\033[49m")
				fmt.Println(v.MAG, "| Detected on packet# |> \033[49m", v.BLKHB, counter, "\033[49m")
				fmt.Println(v.RED, "------------------- BOTTOM HEADER INF --------------------------")
				fmt.Println(v.MAG, "|", v.BLKHB, " Found Search Payload     |> \033[49m", v.BLUHB, string(payload1), "\033[49m")
				fmt.Println(v.MAG, "|", v.BLKHB, " Found Information Element|> \033[49m", v.RED, string(payload), "\033[49m")
			}
		}
	}
}

// dumping PCAP parser, short
func Pcap_parser_OFFLINE(filenametakein string, payload1 string) {
	var handle *pcap.Handle
	handle, err = pcap.OpenOffline(filenametakein)
	ec.Warning_advanced("<RR6> IEEE-802.11 Packet Parsing Module: Could not parse, open, or log file -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	counter := 0
	for packet := range packetSource.Packets() {
		counter++

		if packet.ApplicationLayer() != nil {
			payload := packet.ApplicationLayer().LayerContents()
			if strings.Contains(string(payload), string(payload1)) {
				src := packet.NetworkLayer().NetworkFlow().Src()
				dst := packet.NetworkLayer().NetworkFlow().Dst()
				fmt.Println(v.RED, "------------------- TOP HEADER DATA ---------------------------")
				fmt.Println(v.MAG, "| Source      Address |> \033[49m", v.BLKHB, src, "\033[49m")
				fmt.Println(v.MAG, "| Destination Address |> \033[49m", v.BLKHB, dst, "\033[49m")
				fmt.Println(v.MAG, "| Detected on packet# |> \033[49m", v.BLKHB, counter, "\033[49m")
				fmt.Println(v.RED, "------------------- BOTTOM HEADER INF --------------------------")
				fmt.Println(v.MAG, "|", v.BLKHB, " Found Search Payload     |> \033[49m", v.BLUHB, string(payload1), "\033[49m")
				fmt.Println(v.MAG, "|", v.BLKHB, " Found Information Element|> \033[49m", v.RED, string(payload), "\033[49m")
			}
		}
	}
}

// simple pcap parsing
func Parser(filename string) {
	handeler_pcap, err = pcap.OpenOffline(filename)
	ec.Warning_advanced("<RR6> IEEE-802.11 Packet Parsing Module: Could not open pcap listener offline, something went wrong ->", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer handeler_pcap.Close()
	packetSource := gopacket.NewPacketSource(handeler_pcap, handeler_pcap.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}

// FTP sniffing OFFLINE, parse through a PCAP file

func Ftp_sniff_OFFLINE_PCAP(filenametakein string, cside bool) {
	if cside {
		handeler_pcap, err = pcap.OpenOffline(rr6f.Filepath_general)
	}
	handeler_pcap, err = pcap.OpenOffline(filenametakein)
	if err != nil {
		log.Fatal(err)
	}
	defer handeler_pcap.Close()

	packetSource := gopacket.NewPacketSource(handeler_pcap, handeler_pcap.LinkType())
	for packet := range packetSource.Packets() {
		if packet.ApplicationLayer() != nil {
			payload := packet.ApplicationLayer().Payload()
			dst := packet.NetworkLayer().NetworkFlow().Dst()
			if bytes.Contains(payload, []byte("USER")) {
				fmt.Println(v.RED, "[Red Rabbit] [FTP_Authentication] Found FTP Username \t| ", v.BLU, system.FormatDate, "\t | ")
				fmt.Print(v.MAG, "---------------------------------------------------------------------------------------\n")
				fmt.Print(v.MAG, "| Destination -> ", v.BLU, dst, "\n")
				fmt.Print(v.MAG, "| Payload     -> ", v.BLU, string(payload), "\n")
				fmt.Print(v.MAG, "---------------------------------------------------------------------------------------\n")
				fmt.Println("\n\n", packet)
			} else if bytes.Contains(payload, []byte("PASS")) {
				fmt.Println(v.RED, "[Red Rabbit] [FTP_Authentication] Found FTP Password \t| ", v.BLU, system.FormatDate, "\t | ")
				fmt.Print(v.MAG, "---------------------------------------------------------------------------------------\n")
				fmt.Print(v.MAG, "| Destination -> ", v.BLU, dst, "\n")
				fmt.Print(v.MAG, "| Payload     -> ", v.BLU, string(payload), "\n")
				fmt.Print(v.MAG, "---------------------------------------------------------------------------------------\n")
				fmt.Println("\n\n", packet)
			}
		}
	}
}

// OSPF authentication

func OSPF_OFFLINE_Parsing(filename string) {
	handle, err := pcap.OpenOffline(filename)
	if err != nil {
		fmt.Println("<RR6> Got error -> ", err)
	}
	defer handle.Close()
	counter := 0
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ospf := packet.Layer(layers.LayerTypeOSPF)
		if nil != ospf {
			ospf, _ := ospf.(*layers.OSPFv2)
			switch {
			case ospf.AuType == 1:
				counter++
				fmt.Printf("|---------------------------------|\n")
				fmt.Printf("|        NEW BOX SESSION #%v\t  |\n", counter)
				fmt.Printf("|                                 |\n")
				fmt.Printf("|>> OSPF Authentication           |-> %v\n", ospf.Authentication)
				fmt.Printf("|>> OSPF Area ID                  |-> %v\n", ospf.AreaID)
				fmt.Printf("|>> OSPF Authentication Type      |-> %v\n", ospf.AuType)
				fmt.Printf("|>> OSPF Checksum                 |-> %v\n", ospf.Checksum)
				fmt.Printf("|>> OSPF Packet Length            |-> %v\n", ospf.PacketLength)
				fmt.Printf("|>> OSPF Router ID                |-> %v\n", ospf.RouterID)
				fmt.Printf("|>> OSPF Version                  |-> %v\n", ospf.Version)
				fmt.Printf("|>> OSPF Payload                  |-> %v\n", ospf.Payload)
			case ospf.AuType == 2:
				fmt.Printf("|---------------------------------|\n")
				fmt.Printf("|        NEW BOX SESSION #%v\t  |\n", counter)
				fmt.Printf("|                                 |\n")
				fmt.Printf("|>> OSPF Authentication           |-> %v\n", ospf.Authentication)
				fmt.Printf("|>> OSPF Area ID                  |-> %v\n", ospf.AreaID)
				fmt.Printf("|>> OSPF Authentication Type      |-> %v\n", ospf.AuType)
				fmt.Printf("|>> OSPF Checksum                 |-> %v\n", ospf.Checksum)
				fmt.Printf("|>> OSPF Packet Length            |-> %v\n", ospf.PacketLength)
				fmt.Printf("|>> OSPF Router ID                |-> %v\n", ospf.RouterID)
				fmt.Printf("|>> OSPF Version                  |-> %v\n", ospf.Version)
				fmt.Printf("|>> OSPF Payload                  |-> %s\n", ospf.Payload)

			}
		}
	}
}

func Open_parse_BSSID(pcapfile string) {
	handeler_pcap, err = pcap.OpenOffline(pcapfile)
	ec.Warning_advanced("<RR6> Network Module: Could not parse packet listener and make or open an offline handeler", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer handeler_pcap.Close()
	packetSource := gopacket.NewPacketSource(handeler_pcap, handeler_pcap.LinkType())
	for packet := range packetSource.Packets() {
		dot11 := packet.Layer(layers.LayerTypeDot11)
		if nil != dot11 {
			dot11, _ := dot11.(*layers.Dot11)
			if dot11.Address3 != nil {
				fmt.Println(v.BLKHB, "BSSID | ", dot11.Address3, " | Flags -> ", dot11.Flags, "\033[49m\033[31m")
			}
		}
		dot11info := packet.Layer(layers.LayerTypeDot11InformationElement)
		if nil != dot11info {
			dot11info, _ := dot11info.(*layers.Dot11InformationElement)
			if dot11info.ID == layers.Dot11InformationElementIDSSID {
				fmt.Printf(" \033[0;100mSSID | %q\n", dot11info.Info)
				fmt.Print("\033[49m")
			}
		}
		fmt.Printf("\n")
	}
}

func Find_And_Download(packet gopacket.Packet) {
	app := packet.ApplicationLayer()
	//payload := packet.ApplicationLayer().LayerContents()
	MLD_Report_message := packet.Layer(IEEE_CONSTANTS.MLDv2MulticastListenerReport)
	MLD_Report_Query := packet.Layer(IEEE_CONSTANTS.Multicase_listener)
	IP6_ := packet.Layer(IEEE_CONSTANTS.IP_6)
	IP6_Hop := packet.Layer(IEEE_CONSTANTS.IPv6HopByHop)
	IP6_router_AD := packet.Layer(IEEE_CONSTANTS.IPv6_Router_advertisement)
	ip6, _ := IP6_.(*layers.IPv6)
	ip6_hop, _ := IP6_Hop.(*layers.IPv6HopByHop)
	IP6_router_Ad, _ := IP6_router_AD.(*layers.ICMPv6RouterAdvertisement)
	MulticastListener, _ := MLD_Report_message.(*layers.MLDv2MulticastListenerReportMessage)
	MulticastListener_query, _ := MLD_Report_Query.(*layers.MLDv2MulticastListenerQueryMessage)
	if app != nil {
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
		if strings.Contains(string(app.Payload()), "HTTPS") || strings.Contains(string(app.Payload()), "HTTP") {
			fmt.Println((string(app.Payload())))

		}
	}
}

func Download_images_found_in_pcap(pcapfile string) {
	handle, x := pcap.OpenOffline(pcapfile)
	if ec.Warning_simple("<RR6> PCAP -> NET -> Gopcap: Could not open a offline handle got error -> ", v.REDHB, x) {
		fmt.Println(x)
	} else {
		defer handle.Close()
		counter := 0
		pkt := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range pkt.Packets() {
			counter++
			Find_And_Download(packet)
		}
	}
}
