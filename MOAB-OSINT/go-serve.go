package main

import (
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
	. "github.com/logrusorgru/aurora"
)

func urlidea() {
	var url string
	fmt.Scanf("%s", &url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		fmt.Println("Can not run this string on win32-64")
		os.Exit(1)
	} else {
		fmt.Println("[+] YAY STRING VARIABLES ARE WORKING!!! response code -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
		os.Exit(1)
	}
}
func IsOnline() bool {
	_, err := http.Get("https://www.google.com")
	if err == nil {
		return true
	}
	fmt.Println(Cyan("[-] Interface has been disconnected from the network, please connect or set a connection "))
	return false
}

func mndesknot() {
	if runtime.GOOS == "windows" {
		fmt.Println(Cyan("[-] Sorry, but t this time i can not run this command"))
	} else {
		out, err := exec.Command("notify-send", "Testing Server Conn and Node every 20-30 seconds").Output()
		if err != nil {
			log.Fatal(err)
		} else {
			output := string(out[:])
			fmt.Println(output)
		}
	}
}

func runrb() {
	if runtime.GOOS == "windows" {
		fmt.Println("[-] This can not run on windows")
	} else {
		out, err := exec.Command("job.sh").Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out[:])
		fmt.Println(output)
	}
}


func desk() {
	url := "https://google.com" //testing connection
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		fmt.Println("[-] Sorry will not be able to run this command")
	} else {
		if resp.StatusCode >= 200 {
			out, err := exec.Command("notify-send", "Server responded with code 200 Connection is stable  °˖✧◝(⁰▿⁰)◜✧˖° ✔️").Output()
			if err != nil {
				log.Fatal(err)
			}
			output := string(out[:])
			fmt.Println(output)
		} else {
			out, err := exec.Command("notify-send", "Server Responded with a code that is not within the indexed list or range").Output()
			if err != nil {
				log.Fatal(err)
			}
			output := string(out[:])
			fmt.Println(output)
		}
	}
}
func loggedtes() {
	if runtime.GOOS == "windows" {
		fmt.Println("This appends to a linux system only command, i will not be able to run it")
	} else {
		out, err := exec.Command("notify-send", "There was an error within the response").Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out[:])
		fmt.Println(output)
	}
}
func processElement(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}
func main() {
	uro := os.Args[1]
	IsOnline()
	//var uro string
	//fmt.Scanf("%s", &uro)
	resp, err := http.Get(uro)
	if err != nil {
		log.Fatal(err)
		fmt.Println("A error occured -> ", err)
	}
	if resp.StatusCode >= 400 {
		fmt.Println("[-] Server is not up or isnt a working direcotry now, try again later")
	} else {
		fmt.Println(Red("[+] Con Okay...."))
	}
	time.Sleep(10 * time.Second)
	resp, err = http.Get(uro)
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		fmt.Println("[-] Sorry will not be able to run this command")
	} else {
		if resp.StatusCode >= 200 {
			out, err := exec.Command("notify-send", "Server responded with code 200 Connection is stable  °˖✧◝(⁰▿⁰)◜✧˖° ✔️").Output()
			if err != nil {
				log.Fatal(err)
			}
			output := string(out[:])
			fmt.Println(output)
		} else {
			out, err := exec.Command("notify-send", "Server Responded with a code that is not within the indexed list or range").Output()
			if err != nil {
				log.Fatal(err)
			}
			output := string(out[:])
			fmt.Println(output)
		}
	}
	fmt.Println(Cyan("--------------------------Server Response---------------------------"))
	fmt.Println("[+] Response Status  -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Println("[+] Date Of Request  -> ", resp.Header.Get("date"))
	fmt.Println("[+] Content-Encoding -> ", resp.Header.Get("content-encoding"))
	fmt.Println("[+] Content-Type     -> ", resp.Header.Get("content-type"))
	fmt.Println("[+] Connected-Server -> ", resp.Header.Get("server"))
	fmt.Println("[+] X-Frame-Options  -> ", resp.Header.Get("x-frame-options"))
	fmt.Println(Cyan("--------------------------Server X-Requests-----------------------------"))
	for k, v := range resp.Header {
		fmt.Print(Cyan("[+] -> " + k))
		fmt.Print(Red(" -> "))
		fmt.Println(v)
	}
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
	fmt.Println("-------------- URL QUERY VALS ----------------------- ")
	time.Sleep(2 * time.Second)
	queryMap := parsedURL.Query()
	fmt.Println(queryMap)
	response, err := http.Get(uro)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("a Error has occured while loading HTTP response body ", err)
	}
	fmt.Println("[+] Scraping URLS......")
	document.Find("a").Each(processElement)
	fmt.Println(Red("----------------------------- GATHERING CODE NOTES ----------------------"))
	time.Sleep(2 * time.Second)
	response, err = http.Get(uro)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}
	re := regexp.MustCompile("<!--(.|/n)*?-->")
	comments := re.FindAllString(string(body), -1)
	if comments == nil {
		fmt.Println("No matches.")
	} else {
		for _, comment := range comments {
			timenow := time.Now()
			fmt.Println("[`] Completed at -> ", timenow)
			fmt.Println(comment)
		}
	}
}
