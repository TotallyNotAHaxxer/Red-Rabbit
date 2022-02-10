// SQLIFINDER is a remake of the og Sqlifinder made in python
// credits to the og developer see the based project here
// => https://github.com/americo/sqlifinder
//
//
// This rewrite was written and recreated by ArkAngeL43 with new features and general
// and base ideas of personally what i thought sqlifinder might have needed to make it
// a in general better script to use, not only with command line flags but with the idea
// of implimenting tor sockets, getting the database name, automating base server grabbing,
// parsing payload and URL lists, adding command line flags, and using GOlangs amazing concurrency
// to make both the crawler and SQL injection scanner unque.
//
//
//
// as a developer i try to make my weapons and self defense cyber weapons as weird
// and wacky as possible, sometimes i take peoples scripts, rewrite them in a new language
// with faster speeds and better concurency, i will be 100% honest with you i get general
// enjoyments out of rewritting peoples ideas but better and sometimes worse, lets admit it
// noone is the best at programming but its my way, my way of also learning new languages or
// putting my examples to the test, with my base knowlege of perl and regex in perl i was going
// to rewrite it in perl but i figured noone really uses perl anymore and the chances most people
// know how to run a perl file or install dependencies is really rare
//
// package decleration main
//
package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/steelx/extractlinks"
	"golang.org/x/net/html"
)

// variables
var (
	// flags
	flagTarget          = flag.String("target", "", "HTTP/HTTPS main target URL")
	flagDomain          = flag.String("domain", "", "WWW url")
	flagbase            = flag.String("base", "", "http url")
	flagfile            = flag.String("file", "main.txt", "")
	flagtargetlist      = flag.Bool("targetl", false, "list of urls")
	flagtargetlistfinal = flag.String("l", "", "list of urls")
	flagTOR             = flag.Bool("tor", false, "use tor sockets to connect to the host and crawl it [ INJECTION NOT INCLUDED ] ")

	// forground return color hex
	returnc          = "\033[0;39m"
	BLK              = "\033[0;30m"
	RED              = "\033[0;31m"
	GRN              = "\033[0;32m"
	YEL              = "\033[0;33m"
	BLU              = "\033[0;34m"
	MAG              = "\033[0;35m"
	CYN              = "\033[0;36m"
	WHT              = "\033[0;37m"
	BBLK             = "\033[1;30m"
	BRED             = "\033[1;31m"
	BGRN             = "\033[1;32m"
	BYEL             = "\033[1;33m"
	BBLU             = "\033[1;34m"
	BMAG             = "\033[1;35m"
	BCYN             = "\033[1;36m"
	BWHT             = "\033[1;37m"
	UBLK             = "\033[4;30m"
	URED             = "\033[4;31m"
	UGRN             = "\033[4;32m"
	UYEL             = "\033[4;33m"
	UBLU             = "\033[4;34m"
	UMAG             = "\033[4;35m"
	UCYN             = "\033[4;36m"
	UWHT             = "\033[4;37m"
	BLKB             = "\033[40m"
	REDB             = "\033[41m"
	GRNB             = "\033[42m"
	YELB             = "\033[43m"
	BLUB             = "\033[44m"
	MAGB             = "\033[45m"
	CYNB             = "\033[46m"
	WHTB             = "\033[47m"
	BLKHB            = "\033[0;100m"
	REDHB            = "\033[0;101m"
	GRNHB            = "\033[0;102m"
	YELHB            = "\033[0;103m"
	BLUHB            = "\033[0;104m"
	MAGHB            = "\033[0;105m"
	CYNHB            = "\033[0;106m"
	WHTHB            = "\033[0;107m"
	HBLK             = "\033[0;90m"
	HRED             = "\033[0;91m"
	HGRN             = "\033[0;92m"
	HYEL             = "\033[0;93m"
	HBLU             = "\033[0;94m"
	HMAG             = "\033[0;95m"
	HCYN             = "\033[0;96m"
	HWHT             = "\033[0;97m"
	BHBLK            = "\033[1;90m"
	BHRED            = "\033[1;91m"
	BHGRN            = "\033[1;92m"
	BHYEL            = "\033[1;93m"
	BHBLU            = "\033[1;94m"
	BHMAG            = "\033[1;95m"
	BHCYN            = "\033[1;96m"
	BHWHT            = "\033[1;97m"
	ch               = "\x1b[H\x1b[2J\x1b[3J"
	proxy     string = "socks5://127.0.0.1:9050"
	wg        sync.WaitGroup
	urlQueue  = make(chan string)
	config    = &tls.Config{InsecureSkipVerify: true}
	transport = &http.Transport{
		TLSClientConfig: config,
	}
	hasCrawled = make(map[string]bool)
	netClient  *http.Client
)

// error handeler
func ce(err error, msg string) bool {
	if err != nil {
		log.Fatal(msg, err)
		return true
	}
	return false
}

// first im going to start off with a connection tester, if you have known my
// scripts for a while then you know this is standard for me, tying this with tor
// socks as well

// tortf = tor true or false set this as a bool
func islineneat(website string, tortf bool, proxyname string, timer int) bool {
	if tortf {
		torProxyUrl, err := url.Parse(proxyname)
		ce(err, "couldnt parse proxy")
		tort := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
		client := &http.Client{Transport: tort, Timeout: time.Second * 5}
		resp, err1 := client.Get("https://api.ipify.org?format=text")
		ce(err1, "could not have get client make tor request")
		if resp.StatusCode <= 200 {
			defer resp.Body.Close()
			ipa, err := ioutil.ReadAll(resp.Body)
			ce(err, "could not read body")
			fmt.Println("[Tor=Tester] HTTP stat > ", ipa)
		}
		return true
	} else {
		response, err := http.Get(website)
		ce(err, "COULD NOT MAKE GET REQUEST USER OFFLINE")
		if response.StatusCode <= 200 {
			fmt.Println("[ \033[0;101msqli-finder\033[0;109m ] HTTP stat > ", response.StatusCode)
		} else {
			fmt.Println("[\033[0;106msqli-finder\033[0;109m] HTTP stat > ", response.StatusCode, "THIS MIGHT MEAN YOU ARE OFFLINE WARN:")
		}
	}
	return true
}

// second function uses hexes to clear the screen
func clear(hex string) {
	fmt.Println(hex)
}

// third initation of banner, use banner file banner.txt
func banner(color, file, returntype string) {
	content, err := ioutil.ReadFile(file)
	ce(err, "could not parse file ")
	fmt.Println(color, string(content), returntype)
}

// during goroutines and using go's amazing concurrecy its nice to have a control
// or a os interupt listner, to listen for any hangups ex CTRL+c or OS>KILL, or KILL
// this will just give better handeling
func killer(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nDetected Interupt.....")
			t := time.Now()
			fmt.Println("\n\n\t\033[31m[>] Script Ended At -> ", t)
			os.Exit(0)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

// sql finder is a weird but neat tool, further looking at the code it looks like its crawling the URL's
// then just testing them as well? so i decided im going to make a crawler, that crawls URl's to the domain
// and website, but then just save them to a file named crawlout.txt which will save the URL's for later
// when the cralwer stops it will exit then start testing each URL in that list, i might make it to where the
// sqlfinder can be alone and just test urls but first im going to make the very very basic SQL injection with customized
// payload lists then take that as a flag and finish that
//////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//
//
// CRALWER STARTS HERE
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func init() {
	netClient = &http.Client{
		Transport: transport,
	}
	go sighandel(make(chan os.Signal, 1))
}

func getHtmlPageTORONLY(webPage, torsocket string) (string, error) {
	torProxyUrl, err1 := url.Parse(torsocket)

	if err1 != nil {
		stat := "offline"
		fmt.Println("[-] Error when running proxy, is tor offline? or not being uses => TOR STAT => ", stat)
		os.Exit(0)
	} else {
		stat := "ONLINE TRUE"
		fmt.Printf(WHT+"[ \033[0;100mTOR STAT\033[0;109m"+WHT+"] Using [ %s ] AS SOCKET NAME | TOR:STATUS => [ |%s| ] ", torsocket, stat)
		sep("\n")

	}
	fmt.Println(WHT + " TOR STAT " + WHT + "] Waiting.....")
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Start()
	time.Sleep(6 * time.Second)
	s.Stop()
	testproxy()
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 5}
	responsetor, err := client.Get(webPage)
	if err != nil {
		return "", err
	}
	defer responsetor.Body.Close()
	body, err2 := ioutil.ReadAll(responsetor.Body)
	if err2 != nil {
		return "", err
	}
	return string(body), nil
}

/// get URL ID's

func getHtmlPage(webPage string) (string, error) {
	resp, err := http.Get(webPage)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func parse(text string) {
	tkn := html.NewTokenizer(strings.NewReader(text))
	var isTd bool
	var n int
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := tkn.Token()
			isTd = t.Data == "td"
		case tt == html.TextToken:
			t := tkn.Token()
			if isTd {

				fmt.Printf("%s ", t.Data)
				n++
			}
			if isTd && n%3 == 0 {
				fmt.Println()
			}
			isTd = false
		}
	}
}

//////////////////////////////////////// complex url shifting //////////////////

/////////////////////////////////////////////////////////////////////////////////

/// main

func sqlifinderURL() {
	flag.Parse()
	secarg := *flagDomain //whois IP
	baseUrl := *flagbase
	addr, err := net.LookupIP(secarg)
	checkErr(err)
	if err != nil {
		fmt.Println("[ \033[0;101mpath-finder\033[0;109m ] COULD NOT FIND HOST IPA ")
	} else {
		fmt.Println("[ \033[0;101mpath-finder\033[0;109m ] HTTP IPA > ", addr)
	}
	if err != nil {
		log.Fatal(err)
	}
	webPage := baseUrl
	if *flagTOR {
		// sleep method to pause timeout and keep from causing tor socket errors
		fmt.Println("Sleeping for tor, 6 seconds")
		s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
		s.Start()
		time.Sleep(6 * time.Second)
		s.Stop()
		testproxy()
		sep("\n")
		data, err := getHtmlPageTORONLY(webPage, proxy)
		if err != nil {
			log.Fatal(err)
		}

		parse(data)
		go func() {
			urlQueue <- baseUrl
		}()

		for href := range urlQueue {
			if !hasCrawled[href] {
				crawlLink(href)
			}
		}

	} else {
		data, err := getHtmlPage(webPage)
		//data, err := getHtmlPage(webPage)

		if err != nil {
			log.Fatal(err)
		}

		parse(data)
		go func() {
			urlQueue <- baseUrl
		}()

		for href := range urlQueue {
			if !hasCrawled[href] {
				crawlLink(href)
			}
		}
	}
}

func sighandel(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nDetected Interupt.....")
			t := time.Now()
			fmt.Println("\n\n\t\033[31m[>] Script Ended At -> ", t)
			os.Exit(0)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func crawlLink(baseHref string) {
	flag.Parse()
	// declaring name
	hasCrawled[baseHref] = true
	sep("\n")
	fmt.Println(WHT+"[ \033[0;101mpath-finder\033[0;109m "+WHT+"] URL DISCOVERED > ", baseHref)
	SQLIfinderMAIN(baseHref, *flagfile)
	u, err := url.Parse(baseHref)
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(u.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
	fmt.Println(WHT+"[ \033[0;101m\033[37mpath-finder\033[0;109m"+WHT+"] Domain Name > ", domain)
	addr, err := net.LookupIP(domain) //domain IP for each
	if err != nil {
		fmt.Println("[ \033[0;101mpath-finder\033[0;109m ] Could not get domain name> ")
	} else {
		fmt.Println(WHT+"[ \033[0;101m\033[37mpath-finder\033[0;109m"+WHT+"] Domain IPA > ", addr)
	}
	resp, err := http.Get(baseHref)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(WHT+"[ \033[0;101m\033[37mpath-finder\033[0;109m"+WHT+"] Domain SERVER > ", resp.Header.Get("server"))
	}
	// finally test the query if it is injectable or not with SQL error //
	resp, err = netClient.Get(baseHref)
	checkErr(err)
	defer resp.Body.Close()

	links, err := extractlinks.All(resp.Body)
	checkErr(err)

	for _, l := range links {
		if l.Href == "" {
			continue
		}
		Url := fixedURL(l.Href, baseHref)
		go func(url string) {
			urlQueue <- url
		}(Url)
	}
}

func fixedURL(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil || uri.Scheme == "mailto" || uri.Scheme == "tel" {
		return base
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
} //

func sep(sep string) {
	fmt.Print(sep)
}

//
//
//
//
//CRAWLER ENDS HERE
////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////

// function sqlfinderlist
// parses list of host
func SQLIfinderLIST(list, targetlist string) {
	flag.Parse()
	f, err := os.Open(list)
	if err != nil {
		fmt.Println("[-] Sorry could not parse the list -> ", list)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		f, err := os.Open(targetlist)
		if err != nil {
			fmt.Println("[-] Sorry could not parse the list -> ", list)
		}
		defer f.Close()
		scann1 := bufio.NewScanner(f)
		for scann1.Scan() {
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
					scann1.Text(),
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
						fmt.Printf(WHT+"\n[ \033[0;100m\033[36msqli-finder\033[0;109m "+WHT+"] FOUND VULN IN [%s] TO PAYLOAD [%s] IN URL [%s]", errors[idx], payload, scann1.Text())
						sep("\n")
						break
					}
				}
			}
		}
	}
}

// function SQLfinder
// tests SQL injection of the url
func SQLIfinderMAIN(target, list string) {
	flag.Parse()
	f, err := os.Open(list)
	if err != nil {
		fmt.Println("[-] Sorry could not parse the list -> ", list)
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
				log.Fatalf("[!] Unable to read response body: %s\n", err)
			}
			resp.Body.Close() // close response
			for idx, re := range errRegexes {
				if re.MatchString(string(body)) {
					fmt.Printf(WHT+"\n[ \033[0;100m\033[36msqli-finder\033[0;109m "+WHT+"] FOUND VULN IN [%s] TO PAYLOAD [%s] IN URL [%s]", errors[idx], payload, target)
					sep("\n")
					break
				}
			}
		}
	}
}

// before function name takes place make a tor socket and IP grabber to show changing of addr's

func testproxy() {
	fmt.Println("Sleeping for tor IP CHANGE , 3 seconds")
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Start()
	time.Sleep(3 * time.Second)
	s.Stop()
	torProxyUrl, err := url.Parse(proxy)

	if err != nil {
		fmt.Println("[-] Error when running proxy, is tor offline? or not being uses")
		os.Exit(0)
	}

	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 10}
	resp, err := client.Get("https://api.ipify.org?format=text")

	if err != nil {
		fmt.Println("[-] Error when attempting connection using socket -> ", proxy)
		fmt.Println("[-] Attempted to grab or make a GET request to server => https://www.google.com")
		log.Fatal(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf(WHT+"[ \033[0;100mTOR STAT\033[0;109m"+WHT+"] Using [ %s ] AS SOCKET NAME | TOR:CYCLE:ADDR => [ |%s| ] ", proxy, ip)
}

func main() {
	flag.Parse()
	clear(ch)
	banner(HCYN, "banner-sql.txt", returnc)
	islineneat("https://www.google.com", false, proxy, 1)
	go killer(make(chan os.Signal, 1))
	if *flagtargetlist {
		SQLIfinderLIST(*flagfile, *flagtargetlistfinal)
		os.Exit(0)
	}
	sqlifinderURL()

}

/*
Developer conclusion message.
this script is again a rewrite and remake of the original SQLIFINDER written in python3
https://github.com/americo/sqlifinder
i wanted to remake this to est my skill, turns out i did pretty decent, better
than I honestly ever thought I would even come close to doing XD. I tried as hard
as i could to finish and top it off with tor mining and searching but it just was not
working for shit XD, the injection part anyway as you can use option `tor` to start crawling
the url initally making the most traffic even outside the SQLI
now for common usages use `ch` which will pull out a massive fucking help menu and usages XD
*/
