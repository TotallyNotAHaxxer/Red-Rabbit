package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var (
	proxy string = "socks5://127.0.0.1:9050"
	// unix AF standard socks5 portaddr
	AFUNIX_socket string = "socks5://127.0.0.1:9150"
	base_testurl         = "https://www.google.com"
	uli_ip_get           = "https://api.ipify.org?format=text"
	returnback           = "\033[0;49m"
	returnfore           = "\033[0;39m"
	t                    = time.Now()
	th                   = t.Hour()
	tm                   = t.Minute()
	ts                   = t.Second()
	BLK                  = "\033[0;30m"
	RED                  = "\033[0;31m"
	GRN                  = "\033[0;32m"
	YEL                  = "\033[0;33m"
	BLU                  = "\033[0;34m"
	MAG                  = "\033[0;35m"
	CYN                  = "\033[0;36m"
	WHT                  = "\033[0;37m"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ifaces, err := net.Interfaces()
	check(err)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		check(err)
		for _, a := range addrs {
			fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] PICKED UP INTERFACE > ", WHT, "[ ", BLU, i.Name, WHT, " ] IPA -> [ ", a, " ]", returnfore, "\n")
		}
	}
}
