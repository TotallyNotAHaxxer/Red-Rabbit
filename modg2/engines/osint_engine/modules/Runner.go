package OSINT_Engine

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	r "main/modg/requests"
	maps "main/modules/go-main/SUPER-MAPS"
)

func Run_Engine() {
	if OPTIONS.Run_Twitter {
		maps.Param["q"] = "site:twitter.com " + OPTIONS.Search_Query
	} else if OPTIONS.Run_Facebook {
		maps.Param["q"] = "site:facebook.com " + OPTIONS.Search_Query
	} else if OPTIONS.Run_Linkedin {
		maps.Param["q"] = "site:linkedin.com " + OPTIONS.Search_Query
	} else {
		maps.Param["q"] = OPTIONS.Search_Query
	}
	maps.Param["num"] = strconv.Itoa(OPTIONS.Results_per_page)
	page := 1
	for {
		paramss := url.Values{}
		for k, v := range maps.Param {
			paramss.Add(k, v)
		}
		uri := URL + "?" + paramss.Encode()
		content, resp, err := r.Req_with_Response(uri)
		if err != nil {
			log.Printf("\033[38;5;55m|\033[38;5;88m-\033[38;5;55m| \033[38;5;88mGot connection error -> %s\n", err.Error())
			Connection_Attempt++
			if Connection_Attempt_max == Connection_Attempt {
				break
			}
			continue
		}
		watcher := maps.Response_Watcher[resp.StatusCode]
		if watcher == "Whoops triggered a CAPTCHA" {
			fmt.Printf("\033[38;5;55m|\033[38;5;178m*\033[38;5;55m| \033[38;5;178mWarning: \033[38;5;88m%s \n", watcher)
		}
		if watcher == "Redirect" {
			red := resp.Header["location"]
			fmt.Printf("\033[38;5;55m|\033[38;5;178m*\033[38;5;55m| \033[38;5;178mWarning: \033[38;5;88m%s \t Location [ %s ]\n", watcher, red)
			content, resp, err = r.Req_with_Response(uri)
		}
		if watcher != "OK" {
			fmt.Printf("\033[38;5;55m|\033[38;5;178m*\033[38;5;55m| \033[38;5;178mWarning: \033[38;5;88mGot error when making request {0x01} - %s\n", resp.Status)
			continue
		}
		Pages_Crawled += content
		page++
		if page-1 >= OPTIONS.Pages_to_Crawl {
			break
		}
		maps.Param["start"] = strconv.Itoa(SET_PAGE(page, OPTIONS.Results_per_page))
	}
	Spider()
	Results()
}

func sio(msg string) string {
	var data string
	fmt.Printf(msg)
	fmt.Scanf("%s", &data)
	return data
}

func iio(msg string) int {
	var data int
	fmt.Printf(msg)
	fmt.Scanf("%v", &data)
	return data
}

func Caller(Res_per_page, Pages_to_Crawl int, search_query, website string) {
	if Res_per_page == 0 {
		OPTIONS.Results_per_page = 1
	} else {
		OPTIONS.Results_per_page = Res_per_page
	}
	if Pages_to_Crawl == 0 {
		OPTIONS.Pages_to_Crawl = 1
	} else {
		OPTIONS.Pages_to_Crawl = Pages_to_Crawl
	}
	if search_query == "" {
		fmt.Println("<RR6> Parser: Could not parse search query, data missing")
	} else {
		OPTIONS.Search_Query = search_query
	}
	if OPTIONS.Results_per_page >= 1 && OPTIONS.Pages_to_Crawl >= 1 && len(OPTIONS.Search_Query) > 0 {
		watcher := maps.Watcher[website]
		if watcher == "Twitter" {
			OPTIONS.Run_Twitter = true
		}
		if watcher == "Facebook" {
			OPTIONS.Run_Facebook = true
		}
		if watcher == "Linkedin" {
			OPTIONS.Run_Linkedin = true
		}
		Run_Engine()
	} else {
		fmt.Println("<RR6> Command parser: Got error when parsing data and information to run the OSINT engine, please specify some data")
		fmt.Println("====================")
		fmt.Println("Results per page (def =1) -> ", OPTIONS.Results_per_page)
		fmt.Println("Pages to crawl   (def =1) -> ", OPTIONS.Pages_to_Crawl)
		fmt.Println("Search query              -> ", search_query)
		fmt.Println("website to check for      -> ", website)
	}
}
