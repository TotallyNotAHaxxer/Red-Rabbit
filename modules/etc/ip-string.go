// Verify an IP via Regex IPv4 string only
package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	test_addr = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	BBLU      = "\033[1;34m"
	fl_ip     = flag.String("ip", "", "")
)

func validate_addr(addr string) {
	ipa := strings.Trim(addr, " ")
	re, _ := regexp.Compile(test_addr)
	if re.MatchString(ipa) {
		fmt.Println(BBLU, "\n\b[+] IP Address has been validated -> ", addr)
	} else {
		fmt.Println("\033[31m\n\b[!] Address seems to not be a real IP, verified by regex str -> ", test_addr)
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	validate_addr(*fl_ip)
}
