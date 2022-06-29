package SPOOFER

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func writePoison(arpPacket layers.ARP, etherPacket layers.Ethernet) {
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	gopacket.SerializeLayers(buf, opts, &etherPacket, &arpPacket)
	packetData := buf.Bytes()
	X = handle.WritePacketData(packetData[:42])
	if X != nil {
		fmt.Println("|-| Error: Got an error when attempting to write the packet data using the buffer -> ", X)
	}
}
