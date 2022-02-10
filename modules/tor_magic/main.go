package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/proxy"
)

var (
	flagHelp   = flag.Bool("h", false, `Print the help menu and exit`)
	flagTarget = flag.String("t", "", `target URL`)
)

func clear() {
	if runtime.GOOS == "windows" {
		fmt.Println("[-] I Will not be able to execute this")
	} else {
		out, err := exec.Command("clear").Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out[:])
		fmt.Println(output)
	}
}

//
//
//

func res(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func banner() {
	content, err := ioutil.ReadFile("banner.txt")
	res(err)
	fmt.Println("\033[35m", string(content))
}

//
//
func isflagtrue() {
	if *flagHelp {
		fmt.Println("-----------------------------")
		fmt.Println("Usage: go run main.go -u <url> ")
		flag.PrintDefaults()
		os.Exit(1)
	}

}

func IsOnline() bool {
	_, err := http.Get("https://www.google.com")
	if err == nil {
		return true
	}
	fmt.Println("[-] Interface has been disconnected from the network, please connect or set a connection ")
	return false
}

func processElement(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

var Proxy string = "socks5://127.0.0.1:9050"

func testproxy() {
	torProxyUrl, err := url.Parse(Proxy)

	if err != nil {
		fmt.Println("[-] Error when running proxy, is tor offline? or not being uses")
		os.Exit(0)
	}

	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 5}
	resp, err := client.Get("https://api.ipify.org?format=text")

	if err != nil {
		fmt.Println("[-] Error when attempting connection using socket -> ", Proxy)
		fmt.Println("[-] Attempted to grab or make a GET request to server => https://www.google.com")
		log.Fatal(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\033[32mPublic Internet Address came back with ~>  %s\n", ip)
}

func main() {
	isflagtrue()
	clear()
	banner()
	testproxy()
	flag.Parse()
	socket := "socks5://127.0.0.1:9050"
	torProxyUrl, err := url.Parse("socks5://127.0.0.1:9050") // port 9150 is for Tor Browser
	if err != nil {
		fmt.Println("[?] -??? Unable to parse URL:", err)
		os.Exit(-1)
	}

	torDialer, err := proxy.FromURL(torProxyUrl, proxy.Direct)
	if err != nil {
		fmt.Println("[*] -??? Unable to setup Tor proxy:", err)
		os.Exit(-1)
	}

	torTransport := &http.Transport{Dial: torDialer.Dial}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 5}

	response, err := client.Get(*flagTarget)
	if err != nil {
		fmt.Println("[-] Request failed -> Unable to complete GET request:", err)
		os.Exit(-1)
	}

	defer response.Body.Close()
	parsedURL, err := url.Parse(*flagTarget)
	res(err)
	fmt.Println("\033[31m[*] Sent GET Request using -> ", socket)
	fmt.Println("--------------------------Server Response---------------------------")
	fmt.Println("\033[34m[+] Response Status  -> ", response.StatusCode, http.StatusText(response.StatusCode))
	fmt.Println("\033[34m[+] Date Of Request  -> ", response.Header.Get("date"))
	fmt.Println("\033[34m[+] Content-Encoding -> ", response.Header.Get("content-encoding"))
	fmt.Println("\033[34m[+] Content-Type     -> ", response.Header.Get("content-type"))
	fmt.Println("\033[34m[+] Connected-Server -> ", response.Header.Get("server"))
	fmt.Println("\033[34m[+] X-Frame-Options  -> ", response.Header.Get("x-frame-options"))
	fmt.Println("--------------------------Server X-Requests-----------------------------")
	fmt.Println("\033[31m-------------------------- URL PARSED -------------- ")
	fmt.Println("\033[34mScheme        --->  " + parsedURL.Scheme)
	fmt.Println("\033[34mHostname      --->  " + parsedURL.Host)
	fmt.Println("\033[34mPath in URL   --->  " + parsedURL.Path)
	fmt.Println("\033[34mQuery Strings --->  " + parsedURL.RawQuery)
	fmt.Println("\033[34mFragments     --->  " + parsedURL.Fragment)
	fmt.Println("\033[31m-------------- URL QUERY VALS ----------------------- ")
	for k, v := range response.Header {
		fmt.Print("\033[31m[+] -> " + k)
		fmt.Print("\033[38m -> ")
		fmt.Println(v)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("\033[34ma Error has occured while loading HTTP response body ", err)
	}
	fmt.Println("\033[34m----------------------------- GATHERING HTTPS HREF LINKS ----------------------")
	document.Find("a").Each(processElement)
	fmt.Println("\033[34m----------------------------- GATHERING CODE NOTES ----------------------")
	time.Sleep(2 * time.Second)
	response, err = http.Get(*flagTarget)
	res(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	res(err)
	re := regexp.MustCompile("<!--(.|/n)*?-->")
	comments := re.FindAllString(string(body), -1)
	if comments == nil {
		fmt.Println("[-] Domain has not had any Code notes to parse")
	} else {
		for _, comment := range comments {
			timenow := time.Now()
			fmt.Println("[`] Completed at -> ", timenow)
			fmt.Println(comment)
		}
	}
	w, err := os.Create("index.html")
	res(err)
	defer w.Close()
	_, err = w.ReadFrom(response.Body)
	if err != nil {
		fmt.Println("[CUSTOM]-[FATAL]-> Error occured when getting eaither response or writting to a file")
		log.Fatal(err)
	}
}
