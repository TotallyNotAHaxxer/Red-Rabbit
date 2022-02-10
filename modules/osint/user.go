/*
Dev => ArkAngel43
Program => A more better, remodified, faster, and more equipped user scanner than user recon
Lang => Go
package => main

issues: None asides a url tracky sometimes the dial will timeout i/o


*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	user = flag.String("user", "", "")
	BLK  = "\033[0;30m"
	RED  = "\033[0;31m"
	GRN  = "\033[0;32m"
	YEL  = "\033[0;33m"
	BLU  = "\033[0;34m"
	MAG  = "\033[0;35m"
	CYN  = "\033[0;36m"
	WHT  = "\033[0;37m"
	BBLK = "\033[1;30m"
	BRED = "\033[1;31m"
	BGRN = "\033[1;32m"
	BYEL = "\033[1;33m"
	BBLU = "\033[1;34m"
	BMAG = "\033[1;35m"
	BCYN = "\033[1;36m"
	ch   = "\x1b[H\x1b[2J\x1b[3J"
)

func banner(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ch, BBLU, string(content))
}

func get(filename string) {
	flag.Parse()
	content, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		scanner := bufio.NewScanner(content)
		for scanner.Scan() {
			scan_parser := scanner.Text() + *user
			t, err := http.Get(scan_parser)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(BBLU, "\t\t\t[+] Testing URL -> ", scan_parser)
				if t.StatusCode == 200 {
					fmt.Println(BGRN, "\t\t\t[+] URL -> ", scan_parser, " Came back with -> ", t.StatusCode, " USER FOUND! ")
				} else {
					fmt.Println(BRED, "\t\t\t[!] URL -> ", scan_parser, " Came back      -< ", t.StatusCode, " USER NOT FOUND!!!")
				}
			}
		}
	}
}

func main() {
	banner("user_banner.txt")
	get("url-list.txt")
}
