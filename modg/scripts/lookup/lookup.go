/*
Developer | ArkangeL43
Package   | lookup
Module    | lookup
File      | modg/scripts/lookup/lookup.go
Nest      | scripts/lookup

Does:
	Automates the net lookups for CNAME, MX, NS, TXT, etc
*/
package lookup

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	v "main/modg/colors"
	constants "main/modg/constants"
	helpers "main/modg/helpers"
	str "main/modg/sub"
	ec "main/modg/warnings"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	macs "main/modg/dependant-sub" // macs.go made by google, taken to prevent need to add another go get
	script_requests "main/modg/requests"
	script_constants "main/modg/scripts/scriptc"

	"github.com/PuerkitoBio/goquery"
	"github.com/atotto/clipboard"
	"github.com/steelx/extractlinks"
	"golang.org/x/net/html"
)

var t = time.Now().Add(5 * time.Second)
var k int

type URL_Redirection struct {
	Status_Code int
	URL         string
}

func Looktxt(domain string) {
	rec, _ := net.LookupTXT(domain)
	for _, text := range rec {
		fmt.Println(v.BLKHB, "< Record > |> \033[49m\033[31m", text)
	}
}

func LookHead(url string) {
	response, err := http.Get(url)
	ec.Warning_advanced("<RR6> Lookup Module: Could not make a GET request to the host or URL to grab response -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	for name, val := range response.Header {
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| %s \t \033[38;5;43m%s \n", name, val)
	}
}

func LookSRV(server_t, hostname string) {
	// form      : tcp
	// hostname  : google.com
	// servert   : xmpp-server
	fmt.Println("parsing -> ", server_t, " tcp ", hostname)
	cname, srvs, err := net.LookupSRV(server_t, "tcp", hostname)
	if err != nil {
		fmt.Println(v.REDHB, "<RR6> Got error when trying to lookup the CNAME and SRV for the record -> ", err)
	}
	fmt.Printf("\n\033[1;100m\033[31mcname> \033[49m\033[31m \t|\033[37m  %s  \033[49m\033[31m|\n\n", cname)
	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}
}

func LookCNAME(hostname string) (string, error) {
	if hostname == "" {
		fmt.Println(v.REDHB, "<RR6> Lookup Module: Could not load hostname, error: Hostname was not parsed, was there one?")
	}
	cnames, err := net.LookupCNAME(hostname)
	ec.Warning_advanced("<RR6> Lookup Module: Could not make a direct CNAME record lookup, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	hostname = strings.TrimSuffix(hostname, ".")
	cnames = strings.TrimSuffix(cnames, ".")
	if hostname == "" || cnames == "" {
		fmt.Println("<RR6> Lookup CNAME: Could not find CNAME for host, this domain may not exist")
	}
	return cnames, nil
}

func LookNS(hostname string) {
	nS, err := net.LookupNS(hostname)
	ec.Warning_advanced("<RR6> Lookup Module: Could not make a direct NAME SERVER lookup, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	for _, nameserver := range nS {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| \t \033[38;5;43m", nameserver.Host)
	}
}

func LookMX(domain string) {
	mxRecords, err := net.LookupMX(domain)
	ec.Warning_advanced("<RR6> Lookup Module: Could not make a direct MX lookup, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	fmt.Println("         Mx Record Host     \t|    Mx record pref")
	for _, mxRecord := range mxRecords {
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| \t %s \t| \033[38;5;43m\t %d\n", mxRecord.Host, mxRecord.Pref)
	}

}

func LookHSIP(hostname string) {
	ip, err := net.LookupHost(hostname)
	ec.Warning_advanced("<RR6> Lookup Module: Could not make a direct HOST lookup, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	for _, ipa := range ip {
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| \t IP - \033[38;5;43m%s\n", ipa)
	}
}

func LookIP(ipa string) {
	ip := net.ParseIP(ipa)
	hostnames, err := net.LookupAddr(ip.String())
	ec.Warning_advanced("<RR6> Lookup Module: Could not make a direct IP ADDRESS or HOST lookup, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")

	for _, hostnames := range hostnames {
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Hostnames \t\033[38;5;43m%s\n ", hostnames)
	}
}

func Lookup_banner_main(IPADDR string) {
	ac_threads := 0
	cancel_chan := make(chan bool)
	for port := 0; port <= 1024; port++ {
		go LookBANNER(port, IPADDR, cancel_chan)
		ac_threads++
	}
	for ac_threads > 0 {
		<-cancel_chan
		ac_threads--
	}
}

func Look_oc(hostname string, hwchan chan bool) {
	addr, x := net.LookupAddr(hostname)
	if x == nil {
		fmt.Printf("%s \t| %s\t|\n", hostname, strings.Join(addr, ", "))
	}
	hwchan <- true
}

func LookRedirect(url string) URL_Redirection {
	HTTP_HOST_RESULT := []URL_Redirection{{1, ""}}
	HTTP_HOST := url
	for k > 10 {
		if len(url) == 0 {
			break
		} else {
			if url[0] == '/' {
				HTTP_HOST = Host_info(HTTP_HOST_RESULT[len(HTTP_HOST_RESULT)-1].URL) + HTTP_HOST
			}
			response, x := http.Get(HTTP_HOST)
			if x != nil {
				log.Fatal(x)
			} else {
				if response.StatusCode == 200 {
					fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| < STATUS > | 200 | < Endpoint > | Good | < Redirection > | None | ")
				} else {
					HTTP_HOST = response.Header.Get("Location")
					direction := URL_Redirection{
						URL:         response.Request.URL.String(),
						Status_Code: response.StatusCode,
					}
					HTTP_HOST_RESULT = append(HTTP_HOST_RESULT, direction)
					k += 1
				}
			}
		}
	}
	return HTTP_HOST_RESULT[len(HTTP_HOST_RESULT)-1]
}

func LookBANNER(p int, h string, cancel_chan chan bool) {
	con, e := net.DialTimeout("tcp", h+":"+strconv.Itoa(p), time.Second*10)
	if e != nil {
		cancel_chan <- true
		return
	} else {
		buf := make([]byte, 4096)
		con.SetReadDeadline(t)
		br, ex := con.Read(buf)
		if ex != nil {
			cancel_chan <- true
			return
		} else {
			fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|  Port \033[38;5;43m %d \033[38;5;55m\t banner \033[38;5;43m%s \n", p, buf[0:br])
			cancel_chan <- true
		}
	}

}

func LookRobot(uri, path string) {
	var client http.Client
	parsed_uri := uri + path
	resp, err := http.Get(parsed_uri)
	ec.Warning_advanced("<RR6> Web Module: Could not make a GET request -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	if resp.StatusCode == 100 || resp.StatusCode == 200 || resp.StatusCode == 201 || resp.StatusCode == 202 {
		body, err := ioutil.ReadAll(resp.Body)
		ec.Warning_advanced("<RR6> Web Module: Could not read the HTTP/HTTPS response body -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		resp.Body.Close()
		fmt.Println(v.WHT, "[", v.BHRED, "INFO: RESPONSE BODY CLOSED", v.WHT, "]")
		resp, err = http.Head(parsed_uri)
		ec.Warning_advanced("<RR6> Web Module: Could not read the header values of the HTTP/HTTPS GET request ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		resp.Body.Close()
		fmt.Println(v.WHT, "[", v.BHRED, "INFO: RESPONSE BODY CLOSED", v.WHT, "]")
		form := url.Values{}
		form.Add(parsed_uri, "name1")
		resp, err = http.Post(parsed_uri, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
		ec.Warning_advanced("<RR6> Web Module: Could not make a POST request with the data -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		resp.Body.Close()
		req, err := http.NewRequest("DELETE", parsed_uri, nil)
		ec.Warning_advanced("<RR6> Web Module: Could not make a DELETE request to the HTTP/HTTPS URL", v.REDHB, 1, false, false, true, err, 1, 233, "")
		resp, err = client.Do(req)
		ec.Warning_advanced("<RR6> Web Module: the following exception occured when SENDING the HTTP DELETE method request -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		resp.Body.Close()
		req, err = http.NewRequest("PUT", parsed_uri, strings.NewReader(form.Encode()))
		ec.Warning_advanced("<RR6> Web Module: the following exception occured when making the new request method PUT", v.REDHB, 1, false, false, true, err, 1, 233, "")
		resp, err = client.Do(req)
		ec.Warning_advanced("<RR6> Web Module: the following exception occured when executing the HTTP response body", v.REDHB, 1, false, false, true, err, 1, 233, "")
		resp.Body.Close()
		helpers.Write(string(body))
	} else {
		fmt.Println("ERR: sorry but the http response code i got was not in range => ", resp.StatusCode)
	}
}

// end point path lookup

func Scan_target_file(filename string) []string {
	var result []string
	f, err := os.Open(filename)
	ec.Warning_simple("<RR6> File module: Could not make filepath or read it, is it a directory? ", v.REDHB, err)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		fmt.Println(v.BLKHB, "<RR6> Lookup Module: Looking up endpoints for host's -> ", sc.Text(), "\033[49m\033[39m\033[31m")
		result = append(result, domain)
	}
	return result
}

func Open_and_return(filename string) []string {
	var res []string
	y, x := os.Open(filename)
	if x != nil {
		log.Fatal(x)
	} else {
		scanner := bufio.NewScanner(y)
		for scanner.Scan() {
			res = append(res, scanner.Text())
		}
	}
	return res
}

func Run_Redirect(filename string) {

}

func RetrieveHeaders(data []string) {
	for k, domain := range data {
		script_constants.Limiter <- domain
		script_constants.WaitGroup.Add(1)
		go func(k int, web string) {
			defer script_constants.WaitGroup.Done()
			defer func() {
				<-script_constants.Limiter
			}()
			response, x := http.Get(web)
			if x != nil {
				log.Fatal(x)
			} else {
				for l, element := range response.Header {
					_, e := script_constants.Reuslt_map[l]
					if !e {
						script_constants.Reuslt_map[l] = script_requests.Remove_URL_vals(element)
					} else {
						var u = script_constants.Reuslt_map[l]
						u = append(u, element...)
						script_constants.Reuslt_map[l] = script_requests.Remove_URL_vals(element)
					}
				}
				response.Body.Close()
			}
			script_constants.Mut.Lock()
		}(k, domain)
	}
	script_constants.WaitGroup.Wait()
	for k, x := range script_constants.Reuslt_map {
		fmt.Printf("%s | %s \n", k, script_requests.Remove_URL_vals(x))
	}
}

func RetrieveContents(input []string) []string {
	var R []string
	var Mut = &sync.Mutex{}
	r, _ := regexp.Compile(`\"\/[a-zA-Z0-9_\/?=&]*\"`)
	limiter := make(chan string, 10)
	wg := sync.WaitGroup{}

	for i, domain := range input {
		limiter <- domain
		wg.Add(1)
		go func(i int, domain string) {
			defer wg.Done()
			defer func() { <-limiter }()
			resp, err := http.Get(domain)
			Mut.Lock()
			if err != nil {
				log.Fatal(err)
				ec.Warning_simple("<RR6> Mut Module: COuld not lock mutex function, or maybe even make a request? is ok? false: test line 156", v.REDHB, err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil && len(body) != 0 {
				sb := string(body)
				results := r.FindAllString(sb, -1)
				R = append(R, str.Rdv(results)...)
			}
			resp.Body.Close()
			Mut.Unlock()
		}(i, domain)
	}
	wg.Wait()
	return str.Rdv(R)
}

// mac location tracing
func Mac_trace(mac string) {

	regMac, err := regexp.Compile(`([([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`)

	ec.Warning_advanced("<RR6> REGEX Module: Could not compile regex string, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	if !regMac.MatchString(mac) {
		fmt.Fprintf(os.Stderr, "%s not found? please get MAC format 00:00:00:00:00:00 or 00-00-00-00-00-00\n", mac)
	}
	macAdress := mf(mac)
	url := fmt.Sprintf(`http://mobile.maps.yandex.net/cellid_location/?clid=1866854&lac=-1&cellid=-1&operatorid=null&countrycode=null&signalstrength=-1&wifinetworks=%s:-65&app=ymetro`, macAdress)
	resp, err := http.Get(url)
	ec.Warning_advanced("<RR6> Web Module: Could not make a GET request to the URL -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	ec.Warning_advanced("<RR6> Web Module: Could not read response body, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	data := constants.Loco{}
	err = xml.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| No data was found for this mac\n")
	} else {
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| \033[38;5;43mLAT[%s] LON[%s] \n", data.Location.Longitude, data.Location.Longitude)
	}
}

func mf(mac string) string {
	mac = strings.ToLower(mac)
	mac = strings.Replace(mac, ":", "", -1)
	mac = strings.Replace(mac, "-", "", -1)
	return mac
}

// oui lookup FILE OF MACS
func Oui_filename(filename string) string {
	content, err := os.Open(filename)
	ec.Warning_advanced("<RR6> Could not open filename", v.REDHB, 1, false, false, true, err, 1, 233, "")
	scanner := bufio.NewScanner(content)
	fmt.Println(v.BLKHB, "_______MAC ADDR____________________BRAND NAME_______")
	for scanner.Scan() {
		if mac, err := net.ParseMAC(scanner.Text()); err == nil {
			prefix := [3]byte{
				mac[0],
				mac[1],
				mac[2],
			}
			manufacturer, good := macs.ValidMACPrefixMap[prefix]
			if good {
				fmt.Println(v.BLKHB, " | ", scanner.Text(), " | -> | ", manufacturer, " | ")
			}
		}
	}
	return ""
}

// ajax spider for lookup
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func init() {
	constants.NetClient = &http.Client{
		Transport: constants.Transport,
	}
	go sighandel(make(chan os.Signal, 1))
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
			file, fileErr := os.Create("urls.txt")
			if fileErr != nil {
				fmt.Println("[!] Could not Create a File.......")
				fmt.Println(fileErr)
			}
			fmt.Fprintf(file, "%v\n", url)
		}
	}

	return urls
}

func fetchUrlTitles(urls []string) []*constants.UrlTitle {
	ch := make(chan *constants.UrlTitle, len(urls))
	for idx, url := range urls {
		go func(idx int, url string) {
			doc, err := goquery.NewDocument(url)
			if err != nil {
				ch <- &constants.UrlTitle{idx, url, ""}
			} else {
				ch <- &constants.UrlTitle{idx, url, doc.Find("title").Text()}
			}
		}(idx, url)
	}
	urlsWithTitles := make([]*constants.UrlTitle, len(urls))
	for range urls {
		urlWithTitle := <-ch
		urlsWithTitles[urlWithTitle.Idx] = urlWithTitle
	}
	return urlsWithTitles
}

func toMarkdownList(urlsWithTitles []*constants.UrlTitle) string {
	markdown := ""
	for _, urlWithTitle := range urlsWithTitles {
		markdown += fmt.Sprintf("- [%s](%s)\n", urlWithTitle.Title, urlWithTitle.Url)
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

/////////////////////////////////////////////////////////////////////////////////

func main_ajax(target, base, dom string) {
	addr, err := net.LookupIP(dom)
	checkErr(err)
	resp, err := http.Get(target)
	t := time.Now()
	fmt.Println("\033[34m[>] Script Started At -> ", t)
	if err != nil {
		fmt.Println(v.RED, "[-] Couldnt Get the hostname? ")
	} else {
		fmt.Println("\033[32m[*]Server IPA -> ", addr)
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
	fmt.Println("[*] Crawling URL >> ", base)
	uro := target
	parsedURL, err := url.Parse(uro)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v.CYN, "─────────────────────────Server Response─────────────────────────────")
	fmt.Println("<RR6> Net Lookup:\033[35mResponse Status  -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Println("<RR6> Net Lookup: \033[35mDate Of Request  -> ", resp.Header.Get("date"))
	fmt.Println("<RR6> Net Lookup: \033[35mContent-Encoding -> ", resp.Header.Get("content-encoding"))
	fmt.Println("<RR6> Net Lookup: \033[35mContent-Type     -> ", resp.Header.Get("content-type"))
	fmt.Println("<RR6> Net Lookup: \033[35mConnected-Server -> ", resp.Header.Get("server"))
	fmt.Println("<RR6> Net Lookup: \033[35mX-Frame-Options  -> ", resp.Header.Get("x-frame-options"))
	fmt.Println("<RR6> Net Lookup: \033[35mScheme        --->  " + parsedURL.Scheme)
	fmt.Println("<RR6> Net Lookup: \033[35mHostname      --->  " + parsedURL.Host)
	fmt.Println("<RR6> Net Lookup: \033[35mPath in URL   --->  " + parsedURL.Path)
	fmt.Println("<RR6> Net Lookup: \033[35mQuery Strings --->  " + parsedURL.RawQuery)
	fmt.Println("<RR6> Net Lookup: \033[35mFragments     --->  " + parsedURL.Fragment)
	for k, v := range resp.Header {
		fmt.Println(k)
		fmt.Println(v)
	}
	//grab content
	webPage := base
	data, err := getHtmlPage(webPage)

	if err != nil {
		log.Fatal(err)
	}

	parse(data)
	go func() {
		constants.UrlQueue <- base
	}()

	for href := range constants.UrlQueue {
		if !constants.HasCrawled[href] {
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
			fmt.Println("\033[39m\033[49m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Returning color....")
			os.Exit(0)
		}
	}
}

func crawlLink(baseHref string, counter int) {
	for counter < 1 {
		counter += 1
	}
	constants.HasCrawled[baseHref] = true
	fmt.Println(v.RED, "──────────────────────────────────────────────────────")
	fmt.Println(v.RED, "<RR6>  \033[35mURL Found -> ", baseHref)
	u, err := url.Parse(baseHref)
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(u.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
	fmt.Println(v.RED, "<RR6>  \033[35mDomain Name -> ", domain)
	addr, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println(v.RED, "[-] Couldnt Get the hostname? is there even one? ")
	}
	fmt.Println("server addr -> ", addr)
	resp, err := http.Get(baseHref)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(v.RED, "<RR6>  \033[35mConnected-Server -> ", resp.Header.Get("server"))
		fmt.Println(v.RED, "<RR6>  \033[35mResponse Status  -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	resp, err = constants.NetClient.Get(baseHref)
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
			constants.UrlQueue <- url
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

func Search_urls(url, domain, http string) {
	main_ajax(url, http, domain)
}

// single OUI tracing
func Single_oui(mac string) string {
	if mac, err := net.ParseMAC(mac); err == nil {
		prefix := [3]byte{mac[0], mac[1], mac[2]}
		manufacturer, fast := macs.ValidMACPrefixMap[prefix]
		if fast {
			fmt.Println(v.BLKHB, " | ", mac, " | -> | ", manufacturer, " | \033[39m\033[49m\033[31m")
		}
	}
	return ""
}

func Sub_main(ipa string) {
	ac_threads := 0
	dc := make(chan bool)
	for pi := 0; pi <= 255; pi++ {
		addr := ipa + "." + strconv.Itoa(pi)
		go host_miner(addr, dc)
		ac_threads++
	}
	fmt.Println(v.BLKHB, "\n<RR6> Concurency Module: Started threads - ", ac_threads, v.RET_RED)
	for ac_threads > 0 {
		<-dc
		ac_threads--
	}
}

// host lookup
func host_miner(p string, chann chan bool) {
	a, ex := net.LookupAddr(p)
	if ex != nil {
		chann <- true
	} else {
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| \t %s \t \033[38;5;43m%s \n", p, strings.Join(a, ","))
	}
}

func Host_info(Data string) string {
	u, x := url.Parse(Data)
	if x != nil {
		return "<RR6> (Skip): Got an error when parsing the url"
	}
	return u.Scheme + u.Host
}
