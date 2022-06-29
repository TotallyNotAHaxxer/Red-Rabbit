package scanners

import (
	"bufio"
	"fmt"
	d "main/modg/colors"
	"net/http"
	"os"
)

// this one is just a scattered module for enviroment URL vulnerabilities

func Parse_Hosts(filename string) []string {
	var hosts []string
	a, x := os.Open(filename)
	if x != nil {
		fmt.Println("<RR6> File I/O -> Scanner: Could not open file got the following error -> ", x)
	} else {
		defer a.Close()
		scanner := bufio.NewScanner(a)
		for scanner.Scan() {
			hosts = append(hosts, scanner.Text())
		}
	}
	return hosts
}

func Run_Check(hostname string) {
	fmt.Println("[*] Checking host     |     ", hostname)
	f, x := http.Get(hostname)
	if x != nil {
		fmt.Println(d.REDHB, "<RR6> HTTP-Module -> Requests -> Scanner: Got error when trying to make a methodized request of HTTP method GET to the target URL -> ", x)
	} else {
		if f.StatusCode == 200 || f.StatusCode == 202 || f.StatusCode == 206 {
			fmt.Println("[*] Status for host above came back positive for .env -> ", hostname)
		}
	}
}

func Check_Env(hostfileq bool, filename, url string) {
	var data []string
	if hostfileq {
		data = Parse_Hosts(filename)
		for _, k := range data {
			parser := k + "/.env"
			Run_Check(parser)
		}
	} else {
		Run_Check(url)
	}
}
