package SUPER_Dataparse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	super_color "main/modg/colors"
	super_requests "main/modg/requests"
	super_parser "main/modg/switch/scanners/url-scanner"
	super_data "main/modules/go-main/xml/types"
	"os"
)

func Call(filename, jsondb string) {
	var result []string
	input := super_requests.Scanners(filename)
	conf := Scan_file(jsondb)
	result = super_parser.Check_Check(input, conf)
	for _, elem := range result {
		if elem != "" {
			fmt.Println(super_color.HIGH_BLUE, "<RR6> Status: Found URL in json file -> ", elem)
		} else {
			continue
		}
	}
}

func Scan_file(filename string) super_data.Burp_Config {
	f, x := os.Open(filename)
	if x != nil {
		fmt.Println("<RR6> File I/O: Could not find, locate, open, or read filename, got error -> ", x)
	}
	defer f.Close()
	e, x := ioutil.ReadAll(f)
	if x != nil {
		fmt.Println("<RR6> File I/O: Could not read the file's data. got error -> ", x)
	}
	var conf super_data.Burp_Config
	json.Unmarshal(e, &conf)
	return conf
}
