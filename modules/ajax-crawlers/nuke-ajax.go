package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
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

	"github.com/ArkAngeL43/port-scanning/port"
	"github.com/PuerkitoBio/goquery"
	"github.com/atotto/clipboard"
	"github.com/bndr/gotabulate"
	"github.com/steelx/extractlinks"
	"golang.org/x/net/html"
)

var (
	flagTarget   = flag.String("target", "", "")
	flagDownload = flag.Bool("filedfel", false, "") // download html for every link
	flagDomain   = flag.String("domain", "", "")
	flagbase     = flag.String("base", "", "")
	clear_hex    = "\x1b[H\x1b[2J\x1b[3J"
	BLK          = "\033[0;30m"
	RED          = "\033[0;31m"
	GRN          = "\033[0;32m"
	YEL          = "\033[0;33m"
	BLU          = "\033[0;34m"
	MAG          = "\033[0;35m"
	CYN          = "\033[0;36m"
	WHT          = "\033[0;37m"
	BBLK         = "\033[1;30m"
	BRED         = "\033[1;31m"
	BGRN         = "\033[1;32m"
	BYEL         = "\033[1;33m"
	BBLU         = "\033[1;34m"
	BMAG         = "\033[1;35m"
	BCYN         = "\033[1;36m"
	BWHT         = "\033[1;37m"
	UBLK         = "\033[4;30m"
	URED         = "\033[4;31m"
	UGRN         = "\033[4;32m"
	UYEL         = "\033[4;33m"
	UBLU         = "\033[4;34m"
	UMAG         = "\033[4;35m"
	UCYN         = "\033[4;36m"
	UWHT         = "\033[4;37m"
	BLKB         = "\033[40m"
	REDB         = "\033[41m"
	GRNB         = "\033[42m"
	YELB         = "\033[43m"
	BLUB         = "\033[44m"
	MAGB         = "\033[45m"
	CYNB         = "\033[46m"
	WHTB         = "\033[47m"
	BLKHB        = "\033[0;100m"
	REDHB        = "\033[0;101m"
	GRNHB        = "\033[0;102m"
	YELHB        = "\033[0;103m"
	BLUHB        = "\033[0;104m"
	MAGHB        = "\033[0;105m"
	CYNHB        = "\033[0;106m"
	WHTHB        = "\033[0;107m"
	HBLK         = "\033[0;90m"
	HRED         = "\033[0;91m"
	HGRN         = "\033[0;92m"
	HYEL         = "\033[0;93m"
	HBLU         = "\033[0;94m"
	HMAG         = "\033[0;95m"
	HCYN         = "\033[0;96m"
	HWHT         = "\033[0;97m"
	BHBLK        = "\033[1;90m"
	BHRED        = "\033[1;91m"
	BHGRN        = "\033[1;92m"
	BHYEL        = "\033[1;93m"
	BHBLU        = "\033[1;94m"
	BHMAG        = "\033[1;95m"
	BHCYN        = "\033[1;96m"
	BHWHT        = "\033[1;97m"
	fp           string
	filet        string
	size         int64
)

//structureing and handeling functions

func banner() {
	fmt.Println(clear_hex)
	content, err := ioutil.ReadFile("txt/banner.txt")
	checkErr(err)
	fmt.Println(RED, string(content))
}

type UrlTitle struct {
	idx   int
	url   string
	title string
}

var (
	wg        sync.WaitGroup
	urlQueue  = make(chan string)
	config    = &tls.Config{InsecureSkipVerify: true}
	transport = &http.Transport{
		TLSClientConfig: config,
	}
	hasCrawled = make(map[string]bool)
	netClient  *http.Client
)

func main_download(url string) {
	flag.Parse()
	if *flagDownload {

		resp, err := http.Get(url)
		if err != nil {
			print(err)
			return
		}
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("title").Each(func(i int, s *goquery.Selection) {
			fmt.Printf("Title of the page: %s\n", s.Text())
		})
	} else {
		fmt.Println(WHT, "\t[", BLU, "INFO", WHT, "] \033[31m -> SKIPPED HTML DOWNLOAD, FLAG NOT PARSED")

	}

}

func admin(urltarget string) {
	file, err := os.Open("payloads/payloads-admin.txt")
	checkErr(err)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		u, err := url.Parse(urltarget)
		checkErr(err)
		rel, err := u.Parse(reader.Text())
		checkErr(err)
		resp, err := http.Get(rel.String())
		checkErr(err)
		if resp.StatusCode != 200 {
			fmt.Println(WHT, "\t[", BLU, "INFO", WHT, "] \033[31m -> ", rel, RED, " IS NOT VULNERABLE")
			filepath := "output/tests_pass_list.txt"
			pathmain, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			checkErr(err)
			defer pathmain.Close()
			c, err := fmt.Fprintln(pathmain, "NUKE-AJAX: TESTS_INJECTIONS_ADMIN_PANEL [ INFO ]", rel, " IS NOT VULNERABLE => STATUS = 200")
			if err != nil {
				fmt.Println(c)
				os.Exit(1)
			}
		} else {
			fmt.Println(WHT, "\t[", BLU, "INFO", WHT, "] \033[32m -> ", rel, GRN, " IS VULNERABLE")
			filepath := "output/tests_pass_list.txt"
			pathmain, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			checkErr(err)
			defer pathmain.Close()
			c, err := fmt.Fprintln(pathmain, "NUKE-AJAX: TESTS_INJECTIONS_ADMIN_PANEL [ INFO ]", rel, " IS VULNERABLE => STATUS = 200")
			if err != nil {
				fmt.Print(c)
				os.Exit(1)
			}
		}
		go sighandel(make(chan os.Signal, 1))
	}
}

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

func online() bool {
	_, err := http.Get("https://www.google.com")
	if err == nil {
		fmt.Println(CYN, "\033[32m[+] Connection Good....")
		return true
	}
	fmt.Println(CYN, "[-] Interface has been disconnected from the network, please connect or set a connection ")
	os.Exit(1)
	return false
}

func main_sql(url string) {
	injections := []string{
		"baseline",
		")",
		"(",
		"\"",
		"'",
	}
	errors := []string{
		"SQL",
		"MySQL",
		"ORA-",
		"syntax",
	}

	errRegexes := []*regexp.Regexp{}
	for _, e := range errors {
		re := regexp.MustCompile(fmt.Sprintf(".*%s.*", e))
		errRegexes = append(errRegexes, re)
	}

	for _, payload := range injections {
		client := new(http.Client)
		body := []byte(fmt.Sprintf("username=%s&password=p", payload))

		res, err := http.NewRequest(
			"POST",
			url,
			bytes.NewReader(body),
		)

		if err != nil {
			log.Fatalf("\033[31m\t[X] Unable to Create request -> %s\n", err)
		}

		res.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(res)
		if err != nil {
			log.Fatalf("\033[31m[X] Unable to process response: %s\n", err)
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("\033[31m[X] Unable to read response body: %s\n", err)
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
				filepath := "output/tests_pass_list.txt"
				pathmain, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				checkErr(err)
				defer pathmain.Close()
				c, err := fmt.Fprintln(pathmain, "NUKE-AJAX: TESTS_INJECTIONS_SQLI [ INFO ] IS VULNERABLE TO => ", payload, " GIVIN ERROR => ", errors[idx])
				if err != nil {
					fmt.Print(c)
					os.Exit(1)
				}
				break
			} else {
				filepath := "output/tests_pass_list.txt"
				pathmain, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				checkErr(err)
				defer pathmain.Close()
				c, err := fmt.Fprintln(pathmain, "NUKE-AJAX: TESTS_INJECTIONS_SQLI [ INFO ] URL => ", url, "IS NOT VULNERABLE")
				if err != nil {
					fmt.Print(c)
					os.Exit(1)
				}
				fmt.Println(WHT, "\t[", BLU, "INFO", WHT, "] \033[32mServer HAS PASSED ALL INJECTIONS, NOT VULNERABLE")
			}
		}
	}
}

//// get each URLS title
/// as of now this code is in running and testing
/// will not work as a main function
func isValidUri(uri string) bool {
	_, err := url.ParseRequestURI(uri)

	return err == nil
}

func toUrlList(input string) []string {
	list := strings.Split(strings.TrimSpace(input), "\n")
	urls := make([]string, 0)

	for _, url := range list {
		if isValidUri(url) {
			urls = append(urls, url)
			file, fileErr := os.Create("output/urls.txt")
			if fileErr != nil {
				fmt.Println("[!] Could not Create a File.......")
				fmt.Println(fileErr)
			}
			fmt.Fprintf(file, "%v\n", url)
		}
	}

	return urls
}

func fetchUrlTitles(urls []string) []*UrlTitle {
	ch := make(chan *UrlTitle, len(urls))
	for idx, url := range urls {
		go func(idx int, url string) {
			doc, err := goquery.NewDocument(url)

			if err != nil {
				ch <- &UrlTitle{idx, url, ""}
			} else {
				ch <- &UrlTitle{idx, url, doc.Find("title").Text()}
			}
		}(idx, url)
	}
	urlsWithTitles := make([]*UrlTitle, len(urls))
	for range urls {
		urlWithTitle := <-ch
		urlsWithTitles[urlWithTitle.idx] = urlWithTitle
	}
	return urlsWithTitles
}

func toMarkdownList(urlsWithTitles []*UrlTitle) string {
	markdown := ""
	for _, urlWithTitle := range urlsWithTitles {
		markdown += fmt.Sprintf("- [%s](%s)\n", urlWithTitle.title, urlWithTitle.url)
	}
	return strings.TrimSpace(markdown)
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

func processElement(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

func grabparse() {
	hardurl := "placeholder" // figure out parsing with the command line arguments
	uro := hardurl
	parsedURL, err := url.Parse(uro)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-------------------------- URL PARSED -------------- ")
	fmt.Println("Scheme        --->  " + parsedURL.Scheme)
	fmt.Println("Hostname      --->  " + parsedURL.Host)
	fmt.Println("Path in URL   --->  " + parsedURL.Path)
	fmt.Println("Query Strings --->  " + parsedURL.RawQuery)
	fmt.Println("Fragments     --->  " + parsedURL.Fragment)
}

func write(UO string) {
	req, err := http.Get(UO)
	checkErr(err)
	defer req.Body.Close()
	w, err := os.Create("index.html")
	checkErr(err)
	defer w.Close()
	_, err = w.ReadFrom(req.Body)
	if err != nil {
		fmt.Println("[CUSTOM]-[FATAL]-> Error occured when getting eaither response or writting to a file")
		log.Fatal(err)
	}
}

/////////////////////////////////////////////////////////////////////////////////

func create_file(file, file2, file3 string) {
	f, err := os.Create(file)
	if err != nil {
	}
	f.Close()
	file2c, err := os.Create(file2)
	if err != nil {
	}
	defer file2c.Close()
	file3c, err3 := os.Create(file3)
	checkErr(err3)
	defer file3c.Close()

}

func main() {
	online()
	create_file("output/out.txt", "output/ip-out.txt", "output/tests_pass_list.txt")
	flag.Parse()

	args := *flagTarget   // scrape parsing
	secarg := *flagDomain //whois IP
	baseUrl := *flagbase
	addr, err := net.LookupIP(secarg)
	checkErr(err)
	resp, err := http.Get(baseUrl)
	t := time.Now()
	fmt.Println("\033[34m[>] Script Started At -> ", t)
	// error argument handeling
	if err != nil {
		fmt.Println(RED, "[-] Couldnt Get the hostname? ")
	} else {
		fmt.Println("\033[32m[*]Server IPA -> ", addr)
	}
	if len(args) == 0 {
		fmt.Println(RED, "[-] Url seems to be missing? try https://www.google.com")
		os.Exit(1)
	}

	// argument URO parsing
	if err != nil {
		log.Fatal(err)
	}

	input, _ := clipboard.ReadAll()

	urls := toUrlList(input)

	if len(urls) == 0 {
		fmt.Println("\033[31m[*] Skipping....No URLs found in Copy")
	}
	// cliboard finding titles
	urlsWithTitles := fetchUrlTitles(urls)
	markdown := toMarkdownList(urlsWithTitles)
	fmt.Println(markdown)
	clipboard.WriteAll(markdown)
	//
	fmt.Println("[*] Crawling URL >> ", baseUrl)
	uro := *flagTarget
	parsedURL, err := url.Parse(uro)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(CYN, "\t─────────────────────────Server Response─────────────────────────────")
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mResponse Status  -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mDate Of Request  -> ", resp.Header.Get("date"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mContent-Encoding -> ", resp.Header.Get("content-encoding"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mContent-Type     -> ", resp.Header.Get("content-type"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mConnected-Server -> ", resp.Header.Get("server"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mX-Frame-Options  -> ", resp.Header.Get("x-frame-options"))
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mScheme        --->  " + parsedURL.Scheme)
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mHostname      --->  " + parsedURL.Host)
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mPath in URL   --->  " + parsedURL.Path)
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mQuery Strings --->  " + parsedURL.RawQuery)
	fmt.Println("\033[34m[\033[35m*\033[34m] \033[35mFragments     --->  " + parsedURL.Fragment)
	for k, v := range resp.Header {
		fmt.Print(WHT, "\t[ ", BLU, " INFO ", WHT, " ] \033[32m "+k)
		fmt.Print(RED, "\t -> ")
		fmt.Println(v)
	}
	admin(*flagTarget)
	//grab content
	webPage := baseUrl
	data, err := getHtmlPage(webPage)

	if err != nil {
		log.Fatal(err)
	}

	parse(data)
	go func() {
		urlQueue <- baseUrl
	}()

	for href := range urlQueue {
		if !hasCrawled[href] {
			crawlLink(href, 0)
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

func requests(url string) {
	resp, err := http.Get(url)
	checkErr(err)
	for k, v := range resp.Header {
		fmt.Print(WHT, "\t[ ", BLU, " INFO", WHT, "  ] \033[32m "+k)
		fmt.Print(RED, "\t -> ")
		fmt.Println(v)
	}
}

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

func shortID(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

func writter_v2(url string) {
	flag.Parse()
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("INFO: FATAL: ERR: COULD NOT GET URL? ")
	}
	defer r.Body.Close()
	s := ".html"
	a := shortID(19) + s
	f, err := os.Create(a)

	if err != nil {
		fmt.Println("INFO: FATAL: ERR: COULD NOT GET URL AND WRITE URL's HTML")
	}
	defer f.Close()
	_, err = f.ReadFrom(r.Body)
	if err != nil {
		fmt.Println("INFO: FATAL: ERR: COULD NOT READ FROM r.BODY!!!! ")
	}
}

func crawlLink(baseHref string, counter int) {
	go sighandel(make(chan os.Signal, 1))
	// declaring name
	for counter < 1 {
		counter += 1
	}
	hasCrawled[baseHref] = true
	filepath := "output/out.txt"
	////////////////////──────────────────
	fmt.Print(WHT, "\t[ ", BLU, "INFO", WHT, " ] \033[32m FOUND URL => ", baseHref)
	go sighandel(make(chan os.Signal, 1))
	main_sql(baseHref)
	requests(baseHref)
	main_download(baseHref)
	//

	pathmain, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	checkErr(err)
	defer pathmain.Close()
	c, err := fmt.Fprintln(pathmain, "NUKE-AJAX: URL [ ", counter, " ]", baseHref)
	if err != nil {
		os.Exit(1)
	} else {
		fmt.Println(c)
	}
	u, err := url.Parse(baseHref)
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(u.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
	fmt.Print(WHT, "\t[ ", BLU, "INFO", WHT, " ] \033[32m FOUND DOMAIN NAME => ", domain)
	port.GetOpenPorts(domain, port.PortRange{Start: 1, End: 8090})
	addr, err := net.LookupIP(domain) //domain IP for each

	if err != nil {
		fmt.Println(RED, "[-] Couldnt Get the hostname? is there even one? ")
	} else {
		fip := "output/ip-out.txt"
		fmt.Print(WHT, "\t[ ", BLU, "INFO", WHT, " ] \033[32m FOUND DOMAIN IP => ", addr)
		pathmain, err := os.OpenFile(fip, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("file IP-OUT.txt DOES NOT EXIST, CREATING FILE")
			f, err := os.Create("output/ip-out.txt")
			if err != nil {
				fmt.Println(" error : ", err)
			}
			f.Close()
		}
		defer pathmain.Close()
		c, err := fmt.Fprintln(pathmain, "NUKE-AJAX: URL-IP [ ", addr, " ]")
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(c)
		}
	}
	resp, err := http.Get(baseHref)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(WHT, "\t[ ", BLU, "INFO", WHT, " ] \033[32m FOUND SERVER => ", resp.Header.Get("server"), "\n")
		fmt.Print(WHT, "\t[ ", BLU, "INFO", WHT, " ] \033[32m RESPO STATUS => ", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	resp, err = netClient.Get(baseHref)
	checkErr(err)
	defer resp.Body.Close()

	links, err := extractlinks.All(resp.Body)
	checkErr(err)
	writter_v2(baseHref)
	for _, l := range links {
		if l.Href == "" {
			continue
		}
		Url := fixedURL(l.Href, baseHref)
		if baseHref != Url {
		}
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
}
