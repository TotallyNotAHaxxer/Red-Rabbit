#include <iostream>
#include "stdlib.h"
#include "/usr/local/include/pcapplusplus/SystemUtils.h"
#include "/usr/local/include/pcapplusplus/Packet.h"
#include "/usr/local/include/pcapplusplus/EthLayer.h"
#include "/usr/local/include/pcapplusplus/IPv4Layer.h"
#include "/usr/local/include/pcapplusplus/TcpLayer.h"
#include "/usr/local/include/pcapplusplus/HttpLayer.h"
#include "/usr/local/include/pcapplusplus/PcapFileDevice.h"

std::string getProtocolTypeAsString(pcpp::ProtocolType protocolType)
{
	switch (protocolType)
	{
	case pcpp::Ethernet:
		return "Ethernet";
	case pcpp::IPv4:
		return "IPv4";
	case pcpp::TCP:
		return "TCP";
	case pcpp::HTTPRequest:
	case pcpp::HTTPResponse:
		return "HTTP";
	default:
		return "Unknown";
	}
}

std::string printTcpFlags(pcpp::TcpLayer* tcpLayer)
{
	std::string result = "";
	if (tcpLayer->getTcpHeader()->synFlag == 1)
		result += "SYN ";
	if (tcpLayer->getTcpHeader()->ackFlag == 1)
		result += "ACK ";
	if (tcpLayer->getTcpHeader()->pshFlag == 1)
		result += "PSH ";
	if (tcpLayer->getTcpHeader()->cwrFlag == 1)
		result += "CWR ";
	if (tcpLayer->getTcpHeader()->urgFlag == 1)
		result += "URG ";
	if (tcpLayer->getTcpHeader()->eceFlag == 1)
		result += "ECE ";
	if (tcpLayer->getTcpHeader()->rstFlag == 1)
		result += "RST ";
	if (tcpLayer->getTcpHeader()->finFlag == 1)
		result += "FIN ";

	return result;
}

std::string printTcpOptionType(pcpp::TcpOptionType optionType)
{
	switch (optionType)
	{
	case pcpp::PCPP_TCPOPT_NOP:
		return "NOP";
	case pcpp::PCPP_TCPOPT_TIMESTAMP:
		return "Timestamp";
	default:
		return "Other";
	}
}

std::string printHttpMethod(pcpp::HttpRequestLayer::HttpMethod httpMethod)
{
	switch (httpMethod)
	{
	case pcpp::HttpRequestLayer::HttpGET:
		return "GET";
	case pcpp::HttpRequestLayer::HttpPOST:
		return "POST";
	default:
		return "Other";
	}
}
