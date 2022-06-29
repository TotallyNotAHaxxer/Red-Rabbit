package Dump_Engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var torProxy string = "socks5://127.0.0.1:9050"
var c = &http.Client{}

func Check_CVE(cve string, tor bool) {
	if tor {
		proxy, err := url.Parse(torProxy)
		if err != nil {
			log.Fatal("Error parsing Tor proxy URL:", torProxy, ".", err)
		}
		Trans := &http.Transport{Proxy: http.ProxyURL(proxy)}
		c = &http.Client{Transport: Trans, Timeout: time.Second * 20}

	} else {
		c = &http.Client{
			Timeout: time.Second * 20,
		}
	}
	uri := "https://services.nvd.nist.gov/rest/json/cves/1.0?keyword=%s&resultsPerPage=10"
	URL := fmt.Sprintf(uri, cve)
	req, x := http.NewRequest("", URL, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_8_8 rv:5.0) Gecko/20170914 Firefox/35.0")
	req.Header.Set("method", "GET")
	if x != nil {
		fmt.Println("<RR6> OSINT Engine - Got error when making a new methodized request to the http client -> ", x)
	} else {
		resp, x := c.Do(req)
		if x != nil {
			fmt.Println("<RR6> OSINT Engine - Got error when actually attempting to make a request to the given URL, this may be due to shitty connection or your just fucking with the buffers -> ", x)
		} else {

			defer resp.Body.Close()
			b, x := ioutil.ReadAll(resp.Body)
			for i, k := range resp.Header {
				fmt.Println(i, k)
			}
			fmt.Println(resp.StatusCode)
			if x != nil {
				fmt.Println("<RR6> OSINT Engine - Got error when trying to read all of the response body to I/O -> ", x)
			} else {
				var Results AutoStruct
				x := json.Unmarshal(b, &Results)
				if x != nil {
					fmt.Println("<RR6> OSINT Engine - JSON - Marshal: Got error when trying to run the body with the json structure, this may be due to a corrupt tree or a currupt response body, please try again fixing this error -> ", x)
				} else {
					for i := 0; i < len(Results.Result.CVEItems); i++ {
						if Results.Result.CVEItems != nil {
							fmt.Println("\033[31m=============================================================================================================")
							fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| CVE Name    \t> \033[38;5;43m%s\n\n", Results.Result.CVEItems[i].Cve)
							fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Last Modded \t> \033[38;5;43m%s\n\n", Results.Result.CVEItems[i].LastModifiedDate)
							fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Published   \t> \033[38;5;43m%s\n\n", Results.Result.CVEItems[i].PublishedDate)
							fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Data type   \t> \033[38;5;43m%s\n\n", Results.Result.CVEItems[i].Cve.DataType)
							fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Data format \t> \033[38;5;43m%s\n\n", Results.Result.CVEItems[i].Cve.DataFormat)
							fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Description \t> \033[38;5;43m%s\n\n", Results.Result.CVEItems[i].Cve.Description)
							fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| CVE impact  \t> \033[38;5;43m%v\n\n", Results.Result.CVEItems[i].Impact)
							fmt.Println("\033[31m=============================================================================================================")
						}
					}
				}
			}
		}
	}
}
