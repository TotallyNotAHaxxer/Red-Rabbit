#include <iostream>
#include <sstream>
#include "stdlib.h"
#include "method.hpp"
#include "links.hpp"
#include "HTTP_Proc.hpp"
#include "/usr/local/include/pcapplusplus/SystemUtils.h"
#include "/usr/local/include/pcapplusplus/Packet.h"
#include "/usr/local/include/pcapplusplus/EthLayer.h"
#include "/usr/local/include/pcapplusplus/IPv4Layer.h"
#include "/usr/local/include/pcapplusplus/TcpLayer.h"
#include "/usr/local/include/pcapplusplus/HttpLayer.h"
#include "/usr/local/include/pcapplusplus/PcapFileDevice.h"
#include "/usr/local/include/pcapplusplus/RawPacket.h"




int Loader(pcpp::IFileReaderDevice* r, std::ostream* out, int pc, string output)
{
	pcpp::RawPacket pkt;
    int pcs = 0;
	while (r->getNextPacket(pkt) && pcs != pc) {

        pcpp::Packet parsedPacket(&pkt);
        for (pcpp::Layer* curLayer = parsedPacket.getFirstLayer(); curLayer != NULL; curLayer = curLayer->getNextLayer()) {

            if (getProtocolTypeAsString(curLayer->getProtocol()) == "TCP") {
            }
            if (getProtocolTypeAsString(curLayer->getProtocol()) == "Ethernet") {
            } 
            if (getProtocolTypeAsString(curLayer->getProtocol()) == "IPv4") {
            }
            if (getProtocolTypeAsString(curLayer->getProtocol()) == "HTTP") {
                Process_HTTP(pkt, output);
            }
	}
    	pcs++;
	}
	return pcs;
}
