/*

package main

program does: makes a few requests to a few websites to get outbound address information, grabs public IP, private IP, Net addresses, Interfaces
and general net names

author => literally the same person who developed RR3, RR2, RR1, RR4, RR5
fucking ArkAngeL
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
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
	BBLK                 = "\033[1;30m"
	BRED                 = "\033[1;31m"
	BGRN                 = "\033[1;32m"
	BYEL                 = "\033[1;33m"
	BBLU                 = "\033[1;34m"
	BMAG                 = "\033[1;35m"
	BCYN                 = "\033[1;36m"
	BWHT                 = "\033[1;37m"
	UBLK                 = "\033[4;30m"
	URED                 = "\033[4;31m"
	UGRN                 = "\033[4;32m"
	UYEL                 = "\033[4;33m"
	UBLU                 = "\033[4;34m"
	UMAG                 = "\033[4;35m"
	UCYN                 = "\033[4;36m"
	UWHT                 = "\033[4;37m"
	BLKB                 = "\033[40m"
	REDB                 = "\033[41m"
	GRNB                 = "\033[42m"
	YELB                 = "\033[43m"
	BLUB                 = "\033[44m"
	MAGB                 = "\033[45m"
	CYNB                 = "\033[46m"
	WHTB                 = "\033[47m"
	BLKHB                = "\033[0;100m"
	REDHB                = "\033[0;101m"
	GRNHB                = "\033[0;102m"
	YELHB                = "\033[0;103m"
	BLUHB                = "\033[0;104m"
	MAGHB                = "\033[0;105m"
	CYNHB                = "\033[0;106m"
	WHTHB                = "\033[0;107m"
	HBLK                 = "\033[0;90m"
	HRED                 = "\033[0;91m"
	HGRN                 = "\033[0;92m"
	HYEL                 = "\033[0;93m"
	HBLU                 = "\033[0;94m"
	HMAG                 = "\033[0;95m"
	HCYN                 = "\033[0;96m"
	HWHT                 = "\033[0;97m"
	BHBLK                = "\033[1;90m"
	BHRED                = "\033[1;91m"
	BHGRN                = "\033[1;92m"
	BHYEL                = "\033[1;93m"
	BHBLU                = "\033[1;94m"
	BHMAG                = "\033[1;95m"
	BHCYN                = "\033[1;96m"
	BHWHT                = "\033[1;97m"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func grab_text_response_API(url string) {
	response, err := http.Get(url)
	check(err)
	defer response.Body.Close()
	ip, err := ioutil.ReadAll(response.Body)
	check(err)
	fmt.Printf("\033[32m\t\tPublic Internet Address came back with ~>  %s\n", ip)
}

func localaddr() {
	ifaces, err := net.Interfaces()
	check(err)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		check(err)
		for _, a := range addrs {
			fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] PICKED UP INTERFACE > ", WHT, "[ ", BLU, i.Name, WHT, " ] IPA -> [ ", a, " ]", returnfore, "\n")
			//log.Printf("\033[35m[\033[34mINTERFACE\033[35m] ->  %v %v\n", i.Name, a)
		}
	}
}

func checktor(webname, proxyname string) {
	// make a standard get request to check tor IP through APIFYIP to verify IPA
	//
	// tor settings
	torProxyUrl, err1 := url.Parse(proxyname)
	// err
	check(err1)
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 20}
	responsetor, err := client.Get(webname)
	check(err)
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] CONNECTED: HTTP Status  > ", WHT, "[ ", BLU, responsetor.StatusCode, WHT, " ]", returnfore, "\n")
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] CONNECTED: TOR  SOCKET  > ", WHT, "[ ", BLU, proxyname, WHT, " ]", returnfore, "\n")

}

func main() {
	grab_text_response_API(uli_ip_get)
	localaddr()
	checktor(base_testurl, proxy)

}
