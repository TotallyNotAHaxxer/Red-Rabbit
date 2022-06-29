#include <iostream>
#include "stdlib.h"
#include "/usr/local/include/pcapplusplus/PcapLiveDeviceList.h"
#include "/usr/local/include/pcapplusplus/SystemUtils.h"

std::string Dev_info_output(pcpp::PcapLiveDevice* dev) {
    std::cout 
        << "| Interface Name     -> " << dev->getName()            << std::endl
        << "| Interface MAC      -> " << dev->getMacAddress()      << std::endl
        << "| Interface MTU      -> " << dev->getMtu()             << std::endl
        << "| Interface Gateway  -> " << dev->getDefaultGateway()  << std::endl
        << "| Interface descript -> " << dev->getDesc()            << std::endl;
    return dev->getName();
}