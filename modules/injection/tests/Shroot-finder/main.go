package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/briandowns/spinner"
)

var (
	flagFast    = flag.String("f", "", `Eliminate time waiting process to prevent False positives -> this could return false results`)
	flagWrite   = flag.String("w", "", `Write HTML data to a file | to enable this use -w true`)
	flagSQL     = flag.String("s", "", `SQL vulnerability parsing to enable this type true or | go run main.go -s true `)
	flagAdmin   = flag.String("a", "", `admin panel finding to enable this type true or | go run main.go -a true`)
	flagHelp    = flag.Bool("h", false, `Print usage instructions`)
	flagVersion = flag.Bool("v", false, `Print version`)
	flagTarget  = flag.String("t", "", `target URL`)
)

const (
	v = "1.0 BETA"
)

// same code before
func fastadmin() {
	fmt.Println("\033[31m[*] Starting Admin Panel finding.....")
	file, err := os.Open("txt/payloads.txt")
	res(err)
	reader := bufio.NewScanner(file)
	counter_raid := 0
	under_raid := 0
	for reader.Scan() {
		u, err := url.Parse(*flagTarget)
		res(err)
		rel, err := u.Parse(reader.Text())
		res(err)
		fmt.Println("\033[37m-------------------------------------------------------------")
		resp, err := http.Get(rel.String())
		res(err)
		if resp.StatusCode != 200 {
			counter_raid += 1
			fmt.Println(rel, "\033[35m[ TRY: ", counter_raid, "]\033[31mHas come back NEGATIVE")

			fmt.Println(rel, "\033[31mHas come back NEGATIVE")
		} else {
			under_raid += 1
			log.Println("Response given from server -> ", resp.StatusCode)
			fmt.Println(rel, "[ TRY: ", under_raid, "]\033[32mHas come back POSITIVE")
		}
		go sighandel(make(chan os.Signal, 1))
	}
}

func fastSQL() {
	fmt.Println(" OPENING SQL PAYLOAD FILE ")
	file, err := os.Open("txt/SQLpayloads.txt") // storing payloads in a txt file
	res(err)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		u, err := url.Parse(*flagTarget)
		if err != nil {
			log.Fatal(err)
		}
		rel, err := u.Parse(reader.Text())
		if err != nil {
			fmt.Println("[ - ] FATAL: ERROR PARSING, continuing script")
			continue
		}
		fmt.Println("\033[37m-------------------------------------------------------------")
		resp, err := http.Get(rel.String())
		if err != nil {
			fmt.Println("[ - ] error ")
			continue
		}
		if resp.StatusCode != 200 {
			fmt.Println(rel, "\033[31mHas come back NEGATIVE")
		} else {
			fmt.Println(rel, "\033[32mHas come back POSITIVE")
			fmt.Println("[+] Response given from server -> ", resp.StatusCode)
		}
		go sighandel(make(chan os.Signal, 1))
	}
}

func write() {
	req, err := http.Get(*flagTarget)
	res(err)
	defer req.Body.Close()
	w, err := os.Create("index.html")
	res(err)
	defer w.Close()
	_, err = w.ReadFrom(req.Body)
	if err != nil {
		fmt.Println("[CUSTOM]-[FATAL]-> Error occured when getting eaither response or writting to a file")
		log.Fatal(err)
	}
}

func res(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
func isline() {
	test, err := http.Get("https://www.google.com")
	res(err)
	if test.StatusCode != 200 {
		fmt.Println("[-] Server did not respond with a response code that was OK")
		fmt.Println("[-] Check your connection and try again")
		os.Exit(1)
	} else {
		fmt.Println("")
	}
}
func limit() { // function to limit the amount of requests being sent, attempt to prevent false positives
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	time.Sleep(1 * time.Second)
	s.Stop()
}
func flags() {
	flag.Parse()
	if *flagFast == "true-all" {
		fmt.Println(" starting SQL....")
		fastadmin()
		write()
		fastSQL()

	}
	if *flagFast == "true-sql" {
		fastSQL()
		write()
	}
	if *flagFast == "true-admin" {
		fastadmin()
	}
	if *flagWrite == "true" {
		fmt.Println("\033[34m[+] Writing main URL Source code to index.html")
		write()
	} else {
		fmt.Println("")
	}
	if *flagVersion {
		fmt.Println("\033[38;5;56mCurrent Version | \033[38;5;56m<\033[38;5;199m", v, "\033[38;5;56m>")
		os.Exit(0)
	}
	if *flagHelp {
		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("[*] Value usage")
		fmt.Println("go run main.go |-s|-a|-w true   -t <target> ")
		fmt.Println("--------------------------- Basic value usages-------------")
		fmt.Println("-> go run main.go -s true -a true -w true -t www.example.com | this will get server information, enable admin and SQL testing, then take the HTML and save it ")
		fmt.Println("-> Enable a host -> -a|-s|-w -> true -> -s true same with the others ")
		fmt.Println("U\033[36msage :> go run main.go -t https://www.parrot-pentest.com/")
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *flagTarget == "" {
		fmt.Println("\033[36m[-] Missing Target please input a URL -> |EX| https://www.google.com ")
		fmt.Println("\033[36m[+] Basic usage -> go run main.go -a|-s=TRUE -t <url>")
		fmt.Println("\033[36m[+] Advanced example -> go run main.go -a true -s true -t https://www.example.com")
		os.Exit(0)
	}
	if *flagAdmin == "true" {
		admin()
	} else {
		fmt.Println("") // leave empty
	}
	if *flagSQL == "true" {
		SQL()
	} else {
		fmt.Println("")
	}

}
func banner() {
	banner, err := ioutil.ReadFile("txt/banner.txt")
	res(err)
	fmt.Println("\033[37m", string(banner))
}
func admin() {
	file, err := os.Open("txt/payloads.txt")
	res(err)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		u, err := url.Parse(*flagTarget)
		res(err)
		rel, err := u.Parse(reader.Text())
		res(err)
		fmt.Println("\033[37m-------------------------------------------------------------")
		limit()
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
func SQL() {
	file, err := os.Open("txt/SQLpayloads.txt") // storing payloads in a txt file
	res(err)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		u, err := url.Parse(*flagTarget)
		res(err)
		rel, err := u.Parse(reader.Text())
		res(err)
		fmt.Println("\033[37m-------------------------------------------------------------")
		limit()
		resp, err := http.Get(rel.String())
		res(err)
		if resp.StatusCode != 200 {
			fmt.Println(rel, "\033[31mHas come back NEGATIVE")
		} else {
			fmt.Println(rel, "\033[32mHas come back POSITIVE")
			fmt.Println("[+] Response given from server -> ", resp.StatusCode)
		}
		go sighandel(make(chan os.Signal, 1))
	}

}
func main() {
	flag.Parse()
	fmt.Println("\x1b[H\x1b[2J\x1b[3J") // cross platform hexicode format for clear screen
	banner()
	isline()
	flags()

}
