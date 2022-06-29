package SUPER_knoxss_API_SUPPORT

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	c "main/modg/colors"
	subdomains "main/modg/scripts/domain/subdom"
	structure "main/modules/go-main/SUPER-MODS/knoxss/types"
	"net/http"
	"strings"
	"time"
)

var (
	Client = &http.Client{
		Timeout: time.Second * 60,
	}
	X       error
	API_URL = "https://knoxss.me/api/v3"
)

func Make_Request(target, API string) (string, int, error) {
	postdata := subdomains.FILTER_VALUES("target=" + target)
	resp := bytes.NewBuffer([]byte(postdata))
	req, X := http.NewRequest("POST", API_URL, resp)
	if X != nil {
		fmt.Println(c.REDHB, "<RR6> Search: API: KNOXSS: -> Could not make a new POST methodized request to the given target -> ", X)
	} else {
		req.Header.Set("X-API-KEY", API)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		response, X := Client.Do(req)
		if X != nil {
			fmt.Println(c.REDHB, "<RR6> Requests->API->KNOXSS->RESPONSE => Could not make the new request, for some reason client failed -> ", X)
		} else {
			defer response.Body.Close()
			red, X := ioutil.ReadAll(response.Body)
			if X != nil {
				fmt.Println(c.REDHB, "<RR6> Requests->API->KNOXSS->READER => Could not read the response body of the clients request, got error -> ", X)
			} else {
				i := string(red)
				return i, response.StatusCode, nil
			}
		}
	}
	return "", 0, nil
}

func READER(results string) (structure.Body, error) {
	result := structure.Body{}
	if strings.Contains(results, "{") && strings.Contains(results, "XSS") {
		X = json.Unmarshal([]byte(results), &result)
	} else if strings.Contains(results, "Error Code: <b>HTTP 504</b>") {
		X = errors.New("<RR6> Requests->API => Connection seemed to have timed out, connection broken")
	} else if strings.Contains(results, "Incorrect API key") {
		fmt.Println("<RR6> Requests->API => The API key you used is incorrect or invalid, this may be due to the character limit, spelling, or the API key does not exist.")
	}
	return result, X
}
