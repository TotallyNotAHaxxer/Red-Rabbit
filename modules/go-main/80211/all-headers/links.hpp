#include "/usr/local/include/pcapplusplus/SystemUtils.h"
#include "/usr/local/include/pcapplusplus/Packet.h"
#include "/usr/local/include/pcapplusplus/EthLayer.h"
#include "/usr/local/include/pcapplusplus/IPv4Layer.h"
#include "/usr/local/include/pcapplusplus/TcpLayer.h"
#include "/usr/local/include/pcapplusplus/HttpLayer.h"
#include "/usr/local/include/pcapplusplus/PcapFileDevice.h"
#include "/usr/local/include/pcapplusplus/RawPacket.h"

std::string LINK_LAYER(pcpp::LinkLayerType link)
{

	if (link == pcpp::LINKTYPE_ETHERNET)
		return "Ethernet";
	if (link == pcpp::LINKTYPE_IEEE802_5)
		return "IEEE 802.5 Token Ring";
	else if (link == pcpp::LINKTYPE_LINUX_SLL)
		return "Linux cooked capture";
	else if (link == pcpp::LINKTYPE_NULL)
		return "Null/Loopback";
	else if (link == pcpp::LINKTYPE_RAW || link == pcpp::LINKTYPE_DLT_RAW1 || link == pcpp::LINKTYPE_DLT_RAW2)
	{
		std::ostringstream stream;
		stream << "Raw IP (" << link << ")";
		return stream.str();
	}

	std::ostringstream stream;
	stream << (int)link;
	return stream.str();
}
