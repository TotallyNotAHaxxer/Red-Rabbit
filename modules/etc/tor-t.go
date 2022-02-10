/*

Simple tor socket testing

*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	proxy string = "socks5://127.0.0.1:9050"
)

// error handeler
func ce(err error, msg string) bool {
	if err != nil {
		log.Fatal(msg, err)
		return true
	}
	return false
}

func main() {
	torProxyUrl, err := url.Parse(proxy)
	ce(err, "couldnt parse proxy")
	tort := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: tort, Timeout: time.Second * 5}
	resp, err1 := client.Get("https://api.ipify.org?format=text")
	ce(err1, "could not have get client make tor request")
	if resp.StatusCode <= 200 {
		defer resp.Body.Close()
		ipa, err := ioutil.ReadAll(resp.Body)
		ce(err, "could not read body")
		fmt.Println("[Tor=Tester] HTTP stat > ", ipa)
		fmt.Println("Tor client tested and came back true -> online ")
	}
}
