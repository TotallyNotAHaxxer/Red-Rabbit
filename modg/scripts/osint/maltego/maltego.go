package Maltego_Utils

import (
	"fmt"
	"io/ioutil"
	c "main/modg/colors"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Get_Links_FORM_(url string) {
	var responses []string
	response, x := http.Get(url)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> Requests Module: Got error when making a request to the URL -> ", x)
	} else {
		defer response.Body.Close()
		if response.StatusCode == 200 {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Data: Status code returned 200, stat good...")
			document, x := goquery.NewDocumentFromReader(response.Body)
			if x != nil {
				fmt.Println(c.REDHB, "<RR6> Reader module: Got error when attempting to read the response body in a new document -> ", x)
			} else {
				f := func(i int, s *goquery.Selection) bool {
					link, _ := s.Attr("href")
					return strings.HasPrefix(link, "https")
				}
				document.Find("body a").FilterFunction(f).Each(func(_ int, tag *goquery.Selection) {
					link, _ := tag.Attr("href")
					responses = append(responses, link)
				})
			}
		} else {
			fmt.Println("<RR6> Requests Module: Got a invalid status code from the http response, can not continue unless code is 200...")
		}
	}
	fmt.Println("<MaltegoMessage>")
	fmt.Println("<MaltegoTransformResponseMessage>")
	fmt.Println("  <Entities>")
	for _, i := range responses {
		switch {
		case strings.HasPrefix(i, "http"):
			fmt.Println("<Entity Type=\"maltego.Domain\">")
			fmt.Println("<Value>" + i + "</Value>")
		case strings.HasPrefix(i, "https"):
			fmt.Println("<Entity Type=\"maltego.Domain\">")
			fmt.Println("<Value>" + i + "</Value>")
		case strings.HasPrefix(i, "https"):
			fmt.Println("<Entity Type=\"maltego.Domain\">")
			fmt.Println("<Value>" + url + i + "</Value>")
		}

	}
	fmt.Println("  </Entities>")
	fmt.Println("</MaltegoTransformResponseMessage>")
	fmt.Println("</MaltegoMessage>")
}

func Get_Emails_FORM_(url string) {
	var email []string
	response, x := http.Get(url)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> Requests module: Got error when attempting to make a request to the target URL, error -> ", x)
	} else {
		defer response.Body.Close()
		outb, x := ioutil.ReadAll(response.Body)
		if x != nil {
			fmt.Println(c.REDHB, "<RR6> Requests I/O Module: Got error when attempting to respond or read the response body from the target url, got ERROR -> ", x)
		} else {
			regex, _ := regexp.Compile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,6}`)
			if regex.MatchString(string(outb)) {
				email = regex.FindAllString(string(outb), -1)
			}
			fmt.Println("<MaltegoMessage>")
			fmt.Println("<MaltegoTransformResponseMessage>")
			fmt.Println("  <Entities>")
			for _, i := range email {
				fmt.Println("    <Entity Type=\"maltego.EmailAddress\">")
				fmt.Println("      <Value>" + i + "</Value>")
				fmt.Println("    </Entity>")
			}
			fmt.Println("  </Entities>")
			fmt.Println("</MaltegoTransformResponseMessage>")
			fmt.Println("</MaltegoMessage>")
		}

	}
}
