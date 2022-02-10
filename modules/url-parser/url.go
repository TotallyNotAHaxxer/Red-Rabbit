package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	ur = flag.String("url", "", "")
)

func processElement(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

func parse() {
	flag.Parse()
	URL := *ur
	parsedURL, err := url.Parse(URL)
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
	queryMap := parsedURL.Query()
	fmt.Println(queryMap)
}

func main() {
	flag.Parse()
	URL := *ur
	parsedURL, err := url.Parse(URL)
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
	queryMap := parsedURL.Query()
	fmt.Println(queryMap)
	response, err := http.Get(URL)
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
	fmt.Println("GATHERING CODE NOTES OF THE URL")
	time.Sleep(2 * time.Second)
	response, err = http.Get(URL)
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
			fmt.Println(comment)
		}
	}
}
