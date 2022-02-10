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

 pro_email will validate proton EMAIL addresses using regex and the proton api



*/

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	vpn_u               = "https://api.protonmail.ch/vpn/logicals"
	vpn_a               = "https://api.protonmail.ch/pks/lookup?op=index&search=test@protonmail.com"
	email_url           = "https://api.protonmail.ch/pks/lookup?op=get&search="
	regexEmail          = `([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+)`
	email_main          = flag.String("email", "", "")
	test                = "jefer3401@gmail.com"
	BBLK                = "\033[1;30m"
	BRED                = "\033[1;31m"
	BGRN                = "\033[1;32m"
	BYEL                = "\033[1;33m"
	BBLU                = "\033[1;34m"
	BMAG                = "\033[1;35m"
	BCYN                = "\033[1;36m"
	BWHT                = "\033[1;37m"
	ch                  = "\x1b[H\x1b[2J\x1b[3J"
	verify_true_proton  = "info:1:1"
	verify_false_proton = "info:1:0"
	client              = &http.Client{}
)

func ce(err error, exit_code int) {
	if err != nil {
		log.Fatal(err)
		os.Exit(exit_code)
	}
}

func banner_main(file string) {
	content, err := ioutil.ReadFile(file)
	ce(err, 1)
	fmt.Println(ch, BCYN, string(content))
}

func validate_email(email string, regex string) {
	mail := strings.Trim(email, " ")
	val := regexp.MustCompile(regex)
	if val.MatchString(mail) {
		fmt.Println(BBLU, "\t\t\t[+] EMAIL -> ", email, " Turned up as a real email [ Regex Verify ] ")
	} else {
		os.Exit(1)
	}

}

func find_resp(file, str string) {
	flag.Parse()
	read, err := ioutil.ReadFile(file)
	ce(err, 1)
	if strings.Contains(string(read), str) {
		fmt.Println(BBLU, "\t\t\t[+] EMAIL is apart of the proton mail Service -> ", *email_main)
		fmt.Println(BBLU, "\t\t\t[+] Code for verification found in file was   -> ", str)
	} else {
		fmt.Println(BRED, "\t\t\t[!] EMAIL is NOT apart of the proton mail service -> ", str)
	}
}

func filee(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func write(path *os.File, data string) {
	time.Sleep(1 * time.Second)
	fmt.Fprintln(path, data)
}

func test_service(vpn_url, api_url string) bool {
	vpn, err := http.Get(vpn_url)
	ce(err, 1)
	if vpn.StatusCode >= 200 {
		fmt.Println(BBLU, "\t\t\t[+] VPN Service is UP -> ", vpn.StatusCode)
		return true
	}
	api, err_ap := http.Get(api_url)
	ce(err_ap, 1)
	if api.StatusCode >= 200 {
		fmt.Println(BBLU, "\t\t\t[+] API Service is UP -> ", api.StatusCode)
		return true
	}
	return false
}

func val_main_get(proton_url, email string) {
	email_str := "https://api.protonmail.ch/pks/lookup?op=index&search=" + email
	//em := proton_url + email
	req, err := http.NewRequest("GET", email_str, nil)
	fmt.Println(BBLU, "\t\t\t[+] Request URL    -> ", email_str)
	fmt.Println(BBLU, "\t\t\t[+] Request Method -> GET")
	ce(err, 1)
	req.Header.Set("Accept", "application/json")
	fmt.Println(BBLU, "\t\t\t[+] Request Header SET ")
	resp, err := client.Do(req)
	ce(err, 1)
	defer resp.Body.Close()
	fmt.Println(BBLU, "\t\t\t[+] Request Made -> GET")
	b, err := io.ReadAll(resp.Body)
	ce(err, 1)
	fmt.Println(BRED, "[+] RESPONSE BODY => ", string(b))
	fmt.Println(BBLU, "\t\t\t[+] Opening in file looking for -> ", verify_true_proton)
	pathmain, err := os.OpenFile("in.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		if !filee("in.txt") {
			fmt.Println(BRED, "\t\t\t[!] WARNING: File in.txt does not exist creating new one")
			cr, err := os.Create("in.txt")
			ce(err, 1)
			fmt.Println(BBLU, "\t\t\t[+] DATA: File created: -> ", cr)
			write(pathmain, string(b))
		}
	}
	write(pathmain, string(b))
	find_resp("in.txt", verify_true_proton)
}

func main() {
	flag.Parse()
	banner_main("banner.txt")
	validate_email(*email_main, regexEmail)
	test_service(vpn_u, vpn_a)
	val_main_get(email_url, *email_main)
}
