/*
Developer | ArkAngeL43
Package   | nettest
Module    | ntestgo
File      | modg/scripts/ntest
Nest      | scripts/ntest

Does:
	Provides network testing for tor, get requests, etc

*/
package nettest

import (
	"fmt"
	"io/ioutil"
	v "main/modg/colors"
	xe "main/modg/warnings"
	"net/http"
	"net/url"
	"time"
)

func Tor_test() {
	var proxy string = "socks5://127.0.0.1:9050"
	torProxyUrl, err := url.Parse(proxy)
	xe.Warning_advanced("<RR6> TOR Module: Could not parse the URL with the proxy, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	tort := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: tort, Timeout: time.Second * 5}
	resp, err1 := client.Get("https://api.ipify.org?format=text")
	xe.Warning_advanced("<RR6> TOR Module: Could not make a proper GET request to the target URL, something went wrong -> ", v.REDHB, 1, false, false, true, err1, 1, 233, "")
	if resp.StatusCode <= 200 {
		defer resp.Body.Close()
		ipa, err := ioutil.ReadAll(resp.Body)
		xe.Warning_advanced("<RR6> TOR Module: Could not read response body, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		fmt.Println("[Tor=Tester] HTTP stat > ", ipa)
		fmt.Println("Tor client tested and came back true -> online ")
	}
}

func Net_test() {
	resp, err := http.Get("http://www.google.com")
	xe.Warning_simple("<RR6> Net module: Could not make a get request to the URL", v.REDHB, err)
	if resp.Status == "200" {
		fmt.Println(v.BLKHB, "<RR6> Net Module: Was able to make a proper request, 200 status code")
	} else {
		fmt.Println(v.REDHB, "<RR6> Net Module: Was not able to make a proper request, repsonse was -> ", resp.StatusCode, "  | This may mean the network is not connected or the domain name does not exist")
	}
}
