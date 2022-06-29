package F_Layers

import "github.com/google/gopacket"

func Layer(p gopacket.Packet) gopacket.LayerType {
	for _, l := range p.Layers() {
		return l.LayerType()
	}
	return 0x00
}
