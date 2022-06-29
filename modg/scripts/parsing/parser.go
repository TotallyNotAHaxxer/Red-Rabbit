package host_parser

import (
	"bufio"
	"net/url"
	"os"
	"strings"

	v "main/modg/colors"
	str "main/modg/sub"
	ex "main/modg/warnings"
)

// host parser is a package to help the parsing, cutting, and appending of files, strings, integers, payloads, lists, hosts, dictionaries and more
// dev - ArkAngeL43
// Rdp = remove duplicate values

// parse a list or file with a mix of hosts and ports, filter ones with ports

/*
func call() {
	input := Scan_target_file("filenameexample")
	var result []string
	for _, elem := range input {
		item := GetHostWithoutPort(elem)
		if item != "" {
			result = append(result, item)
		}
	}
	for _, elem := range str.Rdv(result) {
		fmt.Println(elem)
	}
}


-------- premade function to take the args once the terminal sees this is active ------------

*/

func Scan_target_file(filename string) []string {
	var result []string
	f, err := os.Open(filename)
	ex.Warning_simple("<RR6> Parsing Module: Could not make filepath or read it, is it a directory? ", v.REDHB, err)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}
	return str.Rdv(result)
}

func GHWP(input string) string {
	u, err := url.Parse(input)
	ex.Warning_simple("<RR6> Parsing Module: Could not parse url, ", v.REDHB, err)
	if u.Scheme == "" {
		u.Scheme = "http"
	}
	if u.Host == "" {
		return ""
	}
	if len(strings.Split(u.Host, ":")) > 1 {
		if strings.Split(u.Host, ":")[1] == "80" || strings.Split(u.Host, ":")[1] == "443" {
			u.Host = strings.Split(u.Host, ":")[0]
		}
	}
	if u.RawQuery != "" {
		return u.Scheme + "://" + u.Host + u.Path + "?" + u.RawQuery
	}
	return u.Scheme + "://" + u.Host + u.Path
}
