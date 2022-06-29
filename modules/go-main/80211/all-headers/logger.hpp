#include <iostream>
#include <ctime>
#include "color.hpp"
#include <iomanip>
using namespace std;



int Logger(string message, int method) {
    auto time = std::time(nullptr);
    switch(method) {
        case 1: // Red Blue, red is the logo < WOOFED > and blue is the time
            std::cout << RED << "\n<RR6> " << message << BLU << std::put_time(std::gmtime(&time), "%D"); 
            break;
        case 2:
            std::cout << RED << "\n<WOOFER> " << YEL << "WARNING -> " << message << std::put_time(std::gmtime(&time), "%D");
            break;
    }
    return 0;
}