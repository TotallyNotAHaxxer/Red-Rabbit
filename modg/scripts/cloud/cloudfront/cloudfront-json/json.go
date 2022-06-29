package cloudfront

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	cloudfront_constants "main/modg/scripts/cloud/cloudfront/cloudfront-constants"
	cloudfront_regex "main/modg/scripts/cloud/cloudfront/cloudfront-regex"
	"os"
	"strings"
)

type CloudFront_Addresses struct {
	Prefix_Global   string `json:"CLOUDFRONT_GLOBAL_IP_LIST"`
	Prefix_Regional string `json:"CLOUDFRONT_REGIONAL_EDGE_IP_LIST"`
}

func Compare(ip string) {
	if !cloudfront_regex.Test(ip) {
		fmt.Println("<RR6> Cloudfront Regex Module: Could not verify that this is a real IP Addr, must fit the following regex -> ", cloudfront_constants.IP_regex)
		os.Exit(0)
	}
	j, e := os.Open("json/list-cloudfront-ips.json")
	if e != nil {
		fmt.Println("<RR6> Cloudfront package: Could not open json file for parsing....")
	} else {
		defer j.Close()
		byteValue, _ := ioutil.ReadAll(j)
		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)
		sender := fmt.Sprintf("%s", result["CLOUDFRONT_GLOBAL_IP_LIST"])
		sender2 := fmt.Sprintf("%s", result["CLOUDFRONT_GLOBAL_IP_LIST"])
		if !Dev(sender, ip) {
			fmt.Printf("\033[4;31m(Cloud Check)\033[0;39m\033[38;5;198m \033[39m\033[49mIP < %s > is NOT apart of the cloudfront \033[39m(CLOUDFRONT_GLOBAL_IP_LIST)\033[38;5;198m \n", ip)
		} else {
			fmt.Printf("\033[4;34m(Cloud Check)\033[0;39m\033[38;5;198m | IP < %s > is apart of the cloudfront \033[39m(CLOUDFRONT_GLOBAL_IP_LIST)\033[38;5;198m\n", ip)
		}
		if !Dev(sender2, ip) {
			fmt.Printf("\033[4;31m(Cloud Check)\033[0;39m\033[38;5;198m \033[39m\033[49mIP < %s > is NOT apart of the cloudfront \033[39m(CLOUDFRONT_REGIONAL_EDGE_IP_LIST)\033[38;5;198m\n", ip)
		} else {
			fmt.Printf("\033[4;34m(Cloud Check)\033[0;39m\033[38;5;198m | IP < %s > is apart of the cloudfront  \033[39m(CLOUDFRONT_REGIONAL_EDGE_IP_LIST)\033[38;5;198m\n", ip)
		}
	}
}

func Dev(list string, ip string) bool {
	switch strings.Contains(list, ip) {
	case true:
		return true
	case false:
		return false
	}
	return false
}
