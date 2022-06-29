package writer

import (
	"net"

	ARP_CONSTANTS "main/modg/scripts/IEEE-802.11/IEEE-802.11-c"
	ARP_INTERNET_UTILS "main/modg/scripts/IEEE-802.11/recon/functions"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func WRITER(handle *pcap.Handle, iface *net.Interface, addr *net.IPNet) error {
	NULL1 := ARP_CONSTANTS.NULLMARK_1
	NULL2 := ARP_CONSTANTS.NULLMARK_2
	NULL3 := ARP_CONSTANTS.NULLMARK_3
	NULL4 := ARP_CONSTANTS.NULLMARK_4
	NULL5 := ARP_CONSTANTS.NULLMARK_5
	NULL6 := ARP_CONSTANTS.NULLMARK_6
	buffer := gopacket.NewSerializeBuffer()
	eth := layers.Ethernet{
		SrcMAC: iface.HardwareAddr,
		DstMAC: net.HardwareAddr{
			byte(NULL1), byte(NULL2),
			byte(NULL3), byte(NULL4),
			byte(NULL5), byte(NULL6)},
		EthernetType: ARP_CONSTANTS.ETH_ETHERNET_TYPE,
	}
	arp := layers.ARP{
		AddrType:          ARP_CONSTANTS.ARP_ADDRESS_PROTOCAL_TYPE,
		Protocol:          ARP_CONSTANTS.ARP_PACKET_PROTOCAL,
		HwAddressSize:     ARP_CONSTANTS.ARP_HW_ADDR_SIZE,
		ProtAddressSize:   ARP_CONSTANTS.ARP_PR_ADDR_SIZE,
		Operation:         ARP_CONSTANTS.ARP_OPERATION,
		SourceHwAddress:   []byte(iface.HardwareAddr),
		SourceProtAddress: []byte(addr.IP),
		DstHwAddress: []byte{
			ARP_CONSTANTS.NILLBYTE_0,
			ARP_CONSTANTS.NILLBYTE_0,
			ARP_CONSTANTS.NILLBYTE_0,
			ARP_CONSTANTS.NILLBYTE_0,
			ARP_CONSTANTS.NILLBYTE_0,
			ARP_CONSTANTS.NILLBYTE_0},
	}
	Packet_OPTIONS := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	for _, ip := range ARP_INTERNET_UTILS.Internet_Protocal_Addresses_To_Return(addr) {
		arp.DstProtAddress = []byte(ip)
		gopacket.SerializeLayers(buffer, Packet_OPTIONS, &eth, &arp)
		if err := handle.WritePacketData(buffer.Bytes()); err != nil {
			return err
		}
	}
	return nil
}
