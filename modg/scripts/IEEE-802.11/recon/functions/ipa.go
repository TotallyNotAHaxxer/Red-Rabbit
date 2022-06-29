package ipret

import (
	"encoding/binary"
	"net"
)

func Internet_Protocal_Addresses_To_Return(network *net.IPNet) (iptoout []net.IP) {
	num := binary.BigEndian.Uint32([]byte(network.IP))
	mask := binary.BigEndian.Uint32([]byte(network.Mask))
	network_Stream := num & mask
	broadcast := network_Stream | ^mask
	for network_Stream++; network_Stream < broadcast; network_Stream++ {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], network_Stream)
		iptoout = append(iptoout, net.IP(buf[:]))
	}
	return
}
