package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"

	"github.com/ArkAngeL43/port-scanning/port"
)

func ce(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	chex = "\x1b[H\x1b[2J\x1b[3J"
	BLU  = "\033[0;94m"
	BLK  = "\033[0;30m"
	RED  = "\033[0;31m"
	GRN  = "\033[0;32m"
	YEL  = "\033[0;33m"
	MAG  = "\033[0;35m"
	CYN  = "\033[0;36m"
	WHT  = "\033[0;37m"
	//Regular bold
	BBLK = "\033[1;30m"
	BRED = "\033[1;31m"
	BGRN = "\033[1;32m"
	BYEL = "\033[1;33m"
	BBLU = "\033[1;34m"
	BMAG = "\033[1;35m"
	BCYN = "\033[1;36m"
	BWHT = "\033[1;37m"
	//Regular underline
	UBLK = "\033[4;30m"
	URED = "\033[4;31m"
	UGRN = "\033[4;32m"
	UYEL = "\033[4;33m"
	UBLU = "\033[4;34m"
	UMAG = "\033[4;35m"
	UCYN = "\033[4;36m"
	UWHT = "\033[4;37m"
	//Regular back
	BLKB = "\033[40m"
	REDB = "\033[41m"
	GRNB = "\033[42m"
	YELB = "\033[43m"
	BLUB = "\033[44m"
	MAGB = "\033[45m"
	CYNB = "\033[46m"
	WHTB = "\033[47m"
	//High intensty back
	BLKHB = "\033[0;100m"
	REDHB = "\033[0;101m"
	GRNHB = "\033[0;102m"
	YELHB = "\033[0;103m"
	BLUHB = "\033[0;104m"
	MAGHB = "\033[0;105m"
	CYNHB = "\033[0;106m"
	WHTHB = "\033[0;107m"
	//High intensty text
	HBLK = "\033[0;90m"
	HRED = "\033[0;91m"
	HGRN = "\033[0;92m"
	HYEL = "\033[0;93m"
	HBLU = "\033[0;94m"
	HMAG = "\033[0;95m"
	HCYN = "\033[0;96m"
	HWHT = "\033[0;97m"
	//Bold high intensity text
	BHBLK = "\033[1;90m"
	BHRED = "\033[1;91m"
	BHGRN = "\033[1;92m"
	BHYEL = "\033[1;93m"
	BHBLU = "\033[1;94m"
	BHMAG = "\033[1;95m"
	BHCYN = "\033[1;96m"
	BHWHT = "\033[1;97m"
)

func server_writer(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "[  ] 404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

		http.ServeFile(w, r, "html/main.htm")
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		name := r.FormValue("command")
		fmt.Println("[ + ] DEBUG: SET TARGET TO -> ", name)
		if name == "help" {
			content, err := ioutil.ReadFile("html/help.txt")
			checkErr(err)
			fmt.Fprintf(w, string(content))
		}
		str1 := name
		re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
		fmt.Println(re.MatchString(str1))
		if re.MatchString(str1) != true {
			fmt.Println(RED, "\n\n[ * ] DEBUG: HOST NOT SET, POSSIBLE OTHER COMMAND EXECUTED")
		} else {
			submatchall := re.FindAllString(str1, -1)
			for _, element := range submatchall {
				fmt.Println(element)
				IPAddr, err := net.ResolveIPAddr("ip", name)
				if err != nil {
					fmt.Println("Error in resolving IP")
					os.Exit(1)
				}

				addr := net.ParseIP(IPAddr.String())

				if addr == nil {
					fmt.Println("Invalid address")
					os.Exit(1)
				}
				mask := addr.DefaultMask()
				network := addr.Mask(mask)
				fmt.Fprintf(w, "Address : %s \nNetwork : %s \n", addr.String(), network.String())
				port.GetOpenPorts(name, port.PortRange{Start: 1, End: 8090})
			}
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

// other experiments

func validateDomainName(domain string) bool {
	RegExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)

	return RegExp.MatchString(domain)
}

func domain_checker() {

	domName := "www.golang.org"
	if !validateDomainName(domName) {
	} else {
		fmt.Println("[ * ] DATA: Domain input detected => ", domName)
	}

}

func main() {
	fmt.Println(WHT, "[  +  ] DEBUG: [ DATA ] => LOG: Server started")
	fmt.Println(WHT, "[  *  ] DEBUG: http://localhost:8080 ")
	http.HandleFunc("/", server_writer)
	fmt.Printf(GRN, "[ * ] DEBUG: Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
