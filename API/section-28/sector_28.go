package RR6_OFFICAL_SECTOR_28_AUTO_GENERATED_CONFIGURATION_MODULE_BASED_MODULE_PARSER_MODG28_SECTION_13_0X0013_FILEBYTYPE_234

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var r RR6_SECTOR_API_GRABBER

type RR6_SECTOR_API_GRABBER struct {
	APIKeys struct {
		WhoisKey string `json:"WHOIS_KEY"`
		WhoisUsr string `json:"WHOIS_USR"`
		WhoisPas string `json:"WHOIS_PAS"`
		ApilaKey string `json:"APILA_KEY"`
	} `json:"API_Keys"`
}

func Ret_api(key string) (string, string, string, string) {
	f, x := os.Open("API/configuration/api_keys.json")
	if x != nil {
		fmt.Println("<RR6> API Section, Sector 28: Could not grab api keys, error when loading, opening, finding, or running file, got error -> ", x)
	} else {
		defer f.Close()
		d, x := ioutil.ReadAll(f)
		if x != nil {
			fmt.Println("<RR6> I/O -> JSON -> SECTOR 28 -> FILE -> UTIL: Got error when trying to read all of the file, -> ", x)
		} else {

			x := json.Unmarshal(d, &r)
			if x != nil {
				fmt.Println("<RR6> JSON -> SECTOR28 -> PROCESSOR: Got error when trying to unmarshal / parse json structure with the read body of the file -> ", x)
			} else {
				return r.APIKeys.ApilaKey, r.APIKeys.WhoisKey, r.APIKeys.WhoisUsr, r.APIKeys.WhoisPas
			}
		}
	}
	return "no keys in file ---- (IGNORE VALUE)", "for some fucking reason ---- (IGNORE VALUE)", "(IGNORE VALUE)", "(IGNORE VALUE)"

}
