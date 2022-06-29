#include <ctime>
#include <iostream>

using namespace std;

char* Time() {
    time_t d = time(0);
    char* dt = ctime(&d);
    return dt;
}