package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"

	. "github.com/logrusorgru/aurora"
)

func clear() {
	ex := "clear"
	cmd := exec.Command(ex)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(Cyan(string(stdout)))
}

func banner() {
	clear()
	f, err := os.Open("banner.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(Cyan(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processElement(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

func parse() {
	URL := "https://www.amazon.com/AURSINC-Deauther-Wristband-Development-Wearable/dp/B08YX7FGZC/?_encoding=UTF8&pd_rd_w=GfMya&pf_rd_p=628b38b6-dbee-442c-9e9c-e8813ea9e367&pf_rd_r=XSAEDMNF8YA0CMAFXQX1&pd_rd_r=23e976a4-6fb9-4e02-8641-ba206fe373fb&pd_rd_wg=S3m1q&ref_=pd_gw_ci_mcx_mr_hp_d"
	parsedURL, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Cyan("-------------------------- URL PARSED -------------- "))
	fmt.Println(Cyan("Scheme        --->  " + parsedURL.Scheme))
	fmt.Println(Cyan("Hostname      --->  " + parsedURL.Host))
	fmt.Println(Cyan("Path in URL   --->  " + parsedURL.Path))
	fmt.Println(Cyan("Query Strings --->  " + parsedURL.RawQuery))
	fmt.Println(Cyan("Fragments     --->  " + parsedURL.Fragment))
	fmt.Println("-------------- URL QUERY VALS ----------------------- ")
	queryMap := parsedURL.Query()
	fmt.Println(queryMap)
}

func main() {
	banner()
	//parse()
	fmt.Println(Red("[+] Starting Go Diver...."))
	time.Sleep(1 * time.Second)
	var uro string
	fmt.Println("------------ Complex URl -------------- ")
	fmt.Scanf("%s", &uro)
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
			fmt.Println(comment)
		}
	}
}
