package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func finder(country string) {
	formattedurl := "http://www.insecam.org/en/bycountry/" + country
	resp, err := http.Get(formattedurl)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	last_page = r.findall("pagenavigator\\('\\?page=', (\\d+)', res.text)[0]")
	find_ip = r.findall("http://\\d+.\\d+.\\d+.\\d+:\\d+", responseData)

}
