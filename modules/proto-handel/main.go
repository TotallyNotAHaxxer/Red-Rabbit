/*
Developer => ArkAngeL43
Program T => CLI
POG       => MAIN
PM        => MAIN
FP        => main.go


Does:
	apart of my personal information gathering tool kit, aka


  _ \         |                  _ _| __ __|
  __/ _| _ \   _|   _ \    \ ____| |     |
 _| _| \___/ \__| \___/ _| _|    ___|   _|

 Proton it

 Validate emails, IP addresses and more belonging to Proton mail services
 PMS

 Main will search and validate an IP address



*/

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	url           = "https://api.protonmail.ch/vpn/logicals"
	ipa           = flag.String("ip", "", "")
	defualt_entry = "5.8.16.147"
	defualt_exit  = "5.8.16.147"
	test_addr     = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	BBLK          = "\033[1;30m"
	BRED          = "\033[1;31m"
	BGRN          = "\033[1;32m"
	BYEL          = "\033[1;33m"
	BBLU          = "\033[1;34m"
	BMAG          = "\033[1;35m"
	BCYN          = "\033[1;36m"
	BWHT          = "\033[1;37m"
	ch            = "\x1b[H\x1b[2J\x1b[3J"
)

func err_c(err error, exit_code int) {
	if err != nil {
		fmt.Printf("\033[31m[!] Error: %s", err)
		os.Exit(exit_code)
	}
}

func find(file, str string) {
	read, err := ioutil.ReadFile(file)
	err_c(err, 1)
	if strings.Contains(string(read), str) {
		fmt.Println(BBLU, "\t\t\t[+] IP is apart of the proton mail VPN -> ", str)
	} else {
		fmt.Println(false)
	}
}

func validate_addr(addr string) {
	ipa := strings.Trim(addr, " ")
	re, _ := regexp.Compile(test_addr)
	if re.MatchString(ipa) {
		fmt.Println(BBLU, "\t\t\t[+] IP Address has been validated -> ", addr)
	} else {
		fmt.Println("\033[31m[!] Address seems to not be a real IP, verified by regex str -> ", test_addr)
		os.Exit(0)
	}
}

func request(ip, url string) {
	resp, err := http.Get(url)
	err_c(err, 1)
	defer resp.Body.Close()
	fmt.Println(BBLU, "\t\t\t[+] GET request made, saving body to file out.txt")
	if resp.StatusCode >= 200 {
		fmt.Println(BBLU, "\t\t\t[+] STAT -> ", resp.StatusCode)
		out, err_cr := os.Create("out.txt")
		err_c(err_cr, 1)
		defer out.Close()
		io.Copy(out, resp.Body)
		fmt.Println(BBLU, "\t\t\t[+] RESPONSE body copied to I/O")
		find("out.txt", ip)
	}

}

func banner(file string) {
	content, err := ioutil.ReadFile(file)
	err_c(err, 1)
	fmt.Println(ch, BCYN, string(content))
}

func main() {
	flag.Parse()
	banner("banner.txt")
	validate_addr(*ipa)
	request(*ipa, url)
}
