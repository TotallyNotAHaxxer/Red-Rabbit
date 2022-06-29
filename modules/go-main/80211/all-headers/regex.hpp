#include <regex>
#include <bits/stdc++.h>

bool String_IP(string str) {
    std::regex ipv4("(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])");
    std::regex ipv6("((([0-9a-fA-F]){1,4})\\:){7}([0-9a-fA-F]){1,4}");
    if (regex_match(str, ipv4) || regex_match(str, ipv6))
        return true;
    else
        return false;
}

