package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"time"

	"github.com/bndr/gotabulate"
)

func ch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	flagHelp   = flag.Bool("h", false, `Print the help menu and exit`)
	flagList   = flag.String("l", "", `Use a list for SQL Testing`)
	flagTor    = flag.String("p", "", `Use tor proxies to connect to host`)
	flagTarget = flag.String("t", "", `target URL`)
)

var Proxy string = "socks5://127.0.0.1:9050"

func online() {
	response, err := http.Get("https://www.google.com")
	ch(err)
	if response.StatusCode != 200 {
		fmt.Println("\033[31m[-] You seem to not be online.....")
		fmt.Println("\033[31m[-] Exiting.....")
		os.Exit(0)
	} else {
		fmt.Println("[*] Online test passed.....")
	}
}

func res(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fastadmin() {
	fmt.Println("\033[31m[*] Starting Admin Panel finding.....")
	file, err := os.Open("payloads/admin.txt")
	res(err)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		u, err := url.Parse(*flagTarget)
		res(err)
		rel, err := u.Parse(reader.Text())
		res(err)
		fmt.Println("\033[37m-------------------------------------------------------------")
		resp, err := http.Get(rel.String())
		res(err)
		if resp.StatusCode != 200 {
			fmt.Println(rel, "\033[31mHas come back NEGATIVE")
		} else {
			fmt.Println("[+] Response given from server -> ", resp.StatusCode)
			fmt.Println(rel, "\033[32mHas come back POSITIVE")
		}
		go sighandel(make(chan os.Signal, 1))
	}
}

func sighandel(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nInteruption Recieved....")
			t := time.Now()
			fmt.Println("\n\n\t\033[31m[>] Script Ended At -> ", t)
			os.Exit(0)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func listed() {
	flag.Parse()
	f, err := os.Open(*flagList)
	if err != nil {
		fmt.Println("[-] Sorry could not parse the list -> ", *flagList)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		jector := []string{
			scan.Text(),
		}
		errors := []string{
			"SQL",
			"MySQL",
			"ORA-",
			"syntax", // better verticle
		}

		errRegexes := []*regexp.Regexp{}
		for _, e := range errors {
			re := regexp.MustCompile(fmt.Sprintf(".*%s.*", e))
			errRegexes = append(errRegexes, re)
		}

		for _, payload := range jector {

			client := new(http.Client)
			body := []byte(fmt.Sprintf("username=%s&password=p", payload))

			req, err := http.NewRequest(
				"POST",
				*flagTarget,
				bytes.NewReader(body),
			)

			if err != nil {
				log.Fatalf("\033[31m\t[!] Unable to generate request: %s\n", err)
			}

			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalf("[!] Unable to process response: %s\n", err)
			}

			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("[!] Unable to read response body: %s\n", err)
			}

			resp.Body.Close() // close response

			for idx, re := range errRegexes {
				if re.MatchString(string(body)) {
					stringerror := "Server is vulnerable"
					errormsg := "An error | detected vulnerability"
					//				fmt.Printf("[+] SQL Error found [Server->%s] for payload: %s\n", errors[idx], payload)
					row_1 := []interface{}{errors[idx], payload}
					row_2 := []interface{}{errormsg}
					t := gotabulate.Create([][]interface{}{row_1, row_2})
					t.SetHeaders([]string{*flagTarget, stringerror})
					t.SetAlign("center")
					fmt.Println("\033[37m", t.Render("grid"))
					break
				}
			}
		}
	}

}

func main() {
	listed()
	fastadmin()
}
