#include <iostream>
#include <cstring>
#include "color.hpp"
using namespace std;

string Download_File_types_to_Look_For[21] {
    "pdf",
    "gif",
    "jpg",
    "jpeg",
    "webp",
    "bmp",
    "mp4",
    "docx",
    "mp3",
    "ico",
    "tiff",
    "svg",
    "apng",
    "avif",
    "pjp",
    "pjpeg",
    "jfif",
    "cur",
    "tif",
    "heic",
    "hevc"};

int Get_File_Extension_and_match(string url, string prog, string arg, string arg2, string uri) {
    std::string fn = url;
    std::string fe = uri;
    for (int i = 0; i < 10; i++) {
        if(fn.substr(fn.find_last_of(".") + 1) == Download_File_types_to_Look_For[i]) {
            std::cout << RED << "<RR6>" << HIGH_BLUE << " Download    |" << YEL << Download_File_types_to_Look_For[i] << "" << std::endl;
            string a = fe.substr(fe.find_last_of("/"));
            std::cout << RED << "<RR6>" << HIGH_BLUE << " Directory   |" << HIGH_PINK << a << "\n" << std::endl; 
            cout << "\n";
            // prog = ./goul
            // arg  = output directory
            // arg2 = url
            string res = prog + arg2 + " " + arg;
            const char *cmd = res.c_str();
            system(cmd);
        }
    }
    return 0;
}