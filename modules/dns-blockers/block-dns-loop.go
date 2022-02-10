// dns blocker that writes to the hosts file and attempts a TCP connection between the host
// if dial failed tool = = sucess if else then tool has failed blocking attempt

package main

//#include<stdio.h>
//#include<stdlib.h>
//#include<string.h>
//#include<time.h>
//#include <unistd.h>
//#include <errno.h>
//#include <sys/utsname.h>
//#include <sys/utsname.h>
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
//    printf("\e[0;35m|CPU is                  |=> ");
//    brandString(1);
//    brandString(2);
//    brandString(3);
//    printf("\n");
//}
//int main_system()
//{
//    struct utsname buf1;
//    errno =0;
//    if(uname(&buf1)!=0)
//    {
//        perror("Error => Uname returned 0");
//        exit(EXIT_FAILURE);
//    }
//    printf("\e[0;35m_________________________________________________\n");
//    printf("\e[0;35m|System Name             |=> %s\n", buf1.sysname);
//    printf("\e[0;35m|Node/System Name        |=> %s\n", buf1.nodename);
//    printf("\e[0;35m|System Current Version  |=> %s\n", buf1.version);
//    printf("\e[0;35m|Release Version         |=> %s\n", buf1.release);
//    printf("\e[0;35m|Machine ARCH            |=> %s\n", buf1.machine);
//}
import "C"
import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"
)

func system_info_beggin() {
	C.main_system()
	C.getCpuID()
	fmt.Print("\n\n")
}

var (
	chex         = "\x1b[H\x1b[2J\x1b[3J"
	BLK          = "\033[0;30m"
	RED          = "\033[0;31m"
	GRN          = "\033[0;32m"
	YEL          = "\033[0;33m"
	BLU          = "\033[0;34m"
	MAG          = "\033[0;35m"
	CYN          = "\033[0;36m"
	WHT          = "\033[0;37m"
	BBLK         = "\033[1;30m"
	BRED         = "\033[1;31m"
	BGRN         = "\033[1;32m"
	BYEL         = "\033[1;33m"
	BBLU         = "\033[1;34m"
	BMAG         = "\033[1;35m"
	BCYN         = "\033[1;36m"
	BWHT         = "\033[1;37m"
	UBLK         = "\033[4;30m"
	URED         = "\033[4;31m"
	UGRN         = "\033[4;32m"
	UYEL         = "\033[4;33m"
	UBLU         = "\033[4;34m"
	UMAG         = "\033[4;35m"
	UCYN         = "\033[4;36m"
	UWHT         = "\033[4;37m"
	BLKB         = "\033[40m"
	REDB         = "\033[41m"
	GRNB         = "\033[42m"
	YELB         = "\033[43m"
	BLUB         = "\033[44m"
	MAGB         = "\033[45m"
	CYNB         = "\033[46m"
	WHTB         = "\033[47m"
	BLKHB        = "\033[0;100m"
	REDHB        = "\033[0;101m"
	GRNHB        = "\033[0;102m"
	YELHB        = "\033[0;103m"
	BLUHB        = "\033[0;104m"
	MAGHB        = "\033[0;105m"
	CYNHB        = "\033[0;106m"
	WHTHB        = "\033[0;107m"
	HBLK         = "\033[0;90m"
	HRED         = "\033[0;91m"
	HGRN         = "\033[0;92m"
	HYEL         = "\033[0;93m"
	HBLU         = "\033[0;94m"
	HMAG         = "\033[0;95m"
	HCYN         = "\033[0;96m"
	HWHT         = "\033[0;97m"
	BHBLK        = "\033[1;90m"
	BHRED        = "\033[1;91m"
	BHGRN        = "\033[1;92m"
	BHYEL        = "\033[1;93m"
	BHBLU        = "\033[1;94m"
	BHMAG        = "\033[1;95m"
	BHCYN        = "\033[1;96m"
	BHWHT        = "\033[1;97m"
	flagfilepath = flag.String("config", "", " config file for hosts blocking configuration")
)

const (
	lo = "127.0.0.1"
)

func checker() {
	flag.Parse()
	host_config := *flagfilepath
	if runtime.GOOS == "windows" {
		content, err := ioutil.ReadFile("windows.txt")
		ce(err)
		filepath := content
		block_hosts(host_config, string(filepath))
	}
	if runtime.GOOS == "linux" {
		filepath := "/etc/hosts"
		block_hosts(host_config, string(filepath))
	} else {
		filepath := "/etc/hosts"
		block_hosts(host_config, string(filepath))
	}
}

func timenow() {
	t := time.Now()
	fmt.Println(WHT, "Script started blocking At:")
	fmt.Println("_______________________________________")
	fmt.Println("|Current Year        |", t.Year())
	fmt.Println("|Current Month       |", t.Month())
	fmt.Println("|Current Day         | ", t.Day())
	fmt.Println("|Current Hour        |", t.Hour())
	fmt.Println("|Current Minute      |", t.Minute())
	fmt.Println("|Current Second      |", t.Second())
	fmt.Println("|Current Nanosecond  |", t.Nanosecond())
	fmt.Println("|____________________|__________________")
}

func ce(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func test_if_online() bool {
	_, err := http.Get("https://www.google.com")
	if err == nil {
		fmt.Println()
		return true
	} else {
		fmt.Println(RED, "Device may have been disconnected from the network")
		return false
	}
}

func block_hosts(host_config_path, filepath string) {
	file := host_config_path
	blacklist, err := os.Open(file)
	ce(err)
	//
	//
	defer blacklist.Close()
	//
	scanner := bufio.NewScanner(blacklist)
	scanner.Split(bufio.ScanWords)
	count := 0
	fmt.Println(BLU, "\n\n\n  ______________HOSTS IN FILE__________________")
	for scanner.Scan() {
		count += 1
		fmt.Println(MAG, "|", count, "|", scanner.Text())
		fp := filepath
		pathmain, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		ce(err)
		defer pathmain.Close()
		c, err := fmt.Fprintln(pathmain, lo, scanner.Text())
		ce(err)
		if c > 1 {
		} else {
			fmt.Println("")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func dial_attemptin(host_config_path string) {
	file := host_config_path
	blacklist, err := os.Open(file)
	ce(err)
	//
	//
	defer blacklist.Close()
	//
	scanner := bufio.NewScanner(blacklist)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		dialmethod := scanner.Text() + ":80"
		conn, err := net.Dial("tcp", dialmethod)
		if err != nil {
			fmt.Println(BLU, "\n[ * ] DATA: HOST -> ", dialmethod, "BLOCKED")
		} else {
			fmt.Println(RED, "\n[ ? ] Connection was established? was host down or blocked?")
			fmt.Println(RED, "\n[ * ] STAT: Trying TCP dialup again -> ", conn)
			dialmethod := scanner.Text() + ":80"
			conn, err := net.Dial("tcp", dialmethod)
			if err != nil {
				fmt.Println(BLU, "\n[ * ] DATA: HOST -> ", dialmethod, "BLOCKED")
			} else {
				fmt.Println(RED, "\n[ ? ] Connection was established? was host down or blocked?")
				fmt.Println(RED, "\n[ * ] STAT: Trying TCP dialup LAST TIME -> ", conn)
				dialmethod := scanner.Text() + ":80"
				conn, err := net.Dial("tcp", dialmethod)
				if err != nil {
					fmt.Println(BLU, "\n[ * ] DATA: HOST -> ", dialmethod, "BLOCKED")
				} else {
					fmt.Println(RED, "\n[ ? ] Connection was established? was host down or blocked? -> ", conn)
				}
			}
		}
	}
}

func main() {
	flag.Parse()
	system_info_beggin()
	test_if_online()
	timenow()
	checker()
	dial_attemptin(*flagfilepath)
}
