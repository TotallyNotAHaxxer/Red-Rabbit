#include "/usr/local/include/pcapplusplus/PcapLiveDeviceList.h"
#include "/usr/local/include/pcapplusplus/SystemUtils.h"
#include "dsc.hpp"

string Get_iface(std::string IPA) {
    pcpp::PcapLiveDevice* device = pcpp::PcapLiveDeviceList::getInstance().getPcapLiveDeviceByIp(IPA);
    if (device == NULL) {
        std::cout << "[-] ERROR: Could not get device information based on the given IP address - ";
        return "error";
    } else {
        IPA = Dev_info_output(device);
    }
    return IPA;
}