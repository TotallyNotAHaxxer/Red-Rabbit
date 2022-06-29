#include <stdio.h>
#include <iostream>
#include <string>
#include <stdexcept>
#include "all-headers/system.hpp"
#include "all-headers/regex.hpp"
#include "all-headers/get_iface_byip.hpp"
using namespace std;

string x;  // representing x as an input value
string x1; // representing x1 as an input value 

int main(int argc, char* argv[]) {
    std::string chan         = argv[1];
    std::string iface        = argv[2]; 
    std::string change       = argv[3]; 
    std::string want_just_ip = argv[4];
    if (String_IP(iface)) {
        if (want_just_ip == "yes") {
            cout << "[*] Getting interface based on IP...." << std::endl;
            Get_iface(iface);
            return 0;
        }
        if (want_just_ip == "no") {
                iface = Get_iface(iface);
                cout << iface;
                if(chan == "down") {
                    Command_parser_down(1, iface, change);
                }
                if(chan == "up") {
                    Command_parser_down(1, iface, change);
                }
                return 0;
        }
    } else {
        if(chan == "down") {
            Command_parser_down(1, iface, change);
        }
        if(chan == "up") {
            Command_parser_down(1, iface, change);
        }
    }
}

