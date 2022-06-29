/*
Developer | ArkangeL43
Package   | vulns
Module    | vulnscans
File      | modg/scripts/vulnscans/vulns.go
Nest      | scripts/vulnscans

Does:
	Automates the net lookups for CNAME, MX, NS, TXT, etc
*/
package vulns

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	v "main/modg/colors"
	"main/modg/system-runscript"
	ec "main/modg/warnings"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"regexp"
)

func Sighandel(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			os.Exit(1)
		case os.Kill:
			os.Exit(2)
		}
	}
}

// SQL injection tester
func SQLIfinderMAIN(target, list string) {
	f, err := os.Open(list)
	if err != nil {
		fmt.Println("[-] Sorry could not parse the list -> ", list)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		jector := []string{scan.Text()}
		errors := []string{"SQL", "MySQL", "ORA-", "syntax"}
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
				target,
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
				fmt.Println(v.REDHB, "<RR6> Net Module: Could not read response body > ", err)
			}
			resp.Body.Close()
			for idx, re := range errRegexes {
				if re.MatchString(string(body)) {
					fmt.Printf(v.WHT+"\n[ \033[0;100m\033[36msqli-finder\033[0;109m "+v.WHT+"] FOUND VULN IN [%s] TO PAYLOAD [%s] IN URL [%s]", errors[idx], payload, target)
					system.Sep("\n")
					break
				} else {
					fmt.Println("<RR6> SQLI Stat: Could not inject payload -> ", scan.Text(), " Payload failed, testing next")
				}
			}
		}
	}
}

// admin panel vuln finder
func AdminFinder(u, filename string) {
	file, err := os.Open(filename)
	ec.Warning_advanced("<RR6> Web recon module: Could not open up payload file", v.REDHB, 1, false, false, true, err, 1, 233, "")
	reader := bufio.NewScanner(file)
	counter_raid := 0
	under_raid := 0
	for reader.Scan() {
		u, err := url.Parse(u)
		ec.Warning_advanced("<RR6> Web recon module: Could not parse target URL", v.REDHB, 1, false, false, true, err, 1, 233, "")

		rel, err := u.Parse(reader.Text())
		ec.Warning_advanced("<RR6> Web Recon Module: could not read response body", v.REDHB, 1, false, false, true, err, 1, 233, "")

		fmt.Println("\033[37m-------------------------------------------------------------")
		resp, err := http.Get(rel.String())
		ec.Warning_advanced("<RR6> Web Recon Module: Could not make a GET request to the URL", v.REDHB, 1, false, false, true, err, 1, 233, "")
		if resp.StatusCode != 200 {
			counter_raid += 1
			fmt.Println(v.RED, "<RR6> Web Module: Test -> [", v.REDHB, counter_raid, "] \033[49m\033[39m Has come back false")
		} else {
			under_raid += 1
			fmt.Println(v.RED, "<RR6> Web Module: Test -> [", v.REDHB, counter_raid, "] \033[49m\033[39m Has come back TRUE, this URL is a admin panel -> ", reader.Text())
		}
		go Sighandel(make(chan os.Signal, 1))
	}
}

// AJAX spid
