#include <stdio.h>
#include <iostream>
#include <string>
#include <string_view>
#include "system-data.hpp"




using std::cout; 
using std::endl;
using std::string; 


std::string Defualt_interface_OS_Base() {
    #if __linux__
        return "wlan0";
    #elif __APPLE__
        return "en0";
    #elif __linux__
        return "wlan0";
    #elif TARGET_OS_MAC
        return "en0";
    #elif __ANDROID__
        return "wlan0";
    #elif __unix__
        return "wlan0";
    #elif _POSIX_VERSION
        return "wlan0";
    #elif __sun
        return "net0";
    #elif __hpux
        return "lan0";
    #endif
        return "Unknown interface";
}

std::string Command_parser(int System_type, string interface_name, string interface_switch_name) {
    switch(System_type) {
        case 1: // ip link start
            std::string parse_step_four  = System_command_interface_set_up + interface_switch_name + System_command_interface_set_up2;
            std::string parse_step_one   = System_command_interface + interface_name + System_command_interface_2;
            std::string parse_step_two   = System_command_interface_set_name + interface_name + System_command_interface_set_name2 + interface_switch_name;
            std::string parse_step_three = System_command_interface_set_mode + interface_switch_name + System_command_interface_set_mode2;
            std::cout << "[*] Using interface name for commands  -> " << interface_name        << std::endl;
            std::cout << "[*] Using interface name to change     -> " << interface_switch_name << std::endl;
            std::cout << "[*] Using command for interface change -> " << parse_step_one        << std::endl;
            std::cout << "[*] Using command for interface name   -> " << parse_step_two        << std::endl;
            std::cout << "[*] Using command for interface mode   -> " << parse_step_three      << std::endl;
            std::cout << "[*] Using command for interface up     -> " << parse_step_four       << std::endl;
            const char *cmd1 = parse_step_one.c_str();
            const char *cmd2 = parse_step_two.c_str();
            const char *cmd3 = parse_step_three.c_str();
            const char *cmd4 = parse_step_four.c_str();
            system(cmd1);
            system(cmd2);
            system(cmd3);
            system(cmd4);    
        }
    return "";
}

std::string Command_parser_down(int System_type, string interface_name, string interface_switch_name) {
        switch(System_type) {
        case 1: // ip link stop
            std::string parse_step_one_    = System_command_reverse_interface   + interface_name        + System_command_reverse_interface2;
            std::string parse_step_two_    = System_command_reverse_interface_n + interface_name        + System_command_reverse_interface_n2 + interface_switch_name;
            std::string parse_step_three_  = System_command_reverse_interface_m + interface_switch_name + System_command_reverse_interface_m2;
            std::string parse_step_four_   = System_command_interface_set_up    + interface_switch_name + System_command_interface_set_up_2;
            std::cout                      << "[*] Using interface name for commands  -> "              << interface_name        << std::endl;
            std::cout                      << "[*] Using interface name to change     -> "              << interface_switch_name << std::endl;
            std::cout                      << "[*] Using command for interface change -> "              << parse_step_one_        << std::endl;
            std::cout                      << "[*] Using command for interface name   -> "              << parse_step_two_        << std::endl;
            std::cout                      << "[*] Using command for interface mode   -> "              << parse_step_three_      << std::endl;
            std::cout                      << "[*] Using command for interface up     -> "              << parse_step_four_       << std::endl;
            const char *cmd0 = parse_step_one_.c_str();
            const char *cmd9 = parse_step_two_.c_str();
            const char *cmd8 = parse_step_three_.c_str();
            const char *cmd7 = parse_step_four_.c_str();
            system(cmd0);
            system(cmd9);
            system(cmd8);
            system(cmd7);    
    }
    return "";
}

