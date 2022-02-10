// robots.txt file/url parser/downloader
// pure go code
// Author: ArkAngeL43 | https://github.com/ArkAngeL43
// in workings with the Red-Rabbit project this will parse the
// robots.txt filepath to a url then if it returns with a status that is or close
// to OK or HTTPOK/STATOK
// package main decleration here
package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// flag variable and color strings
// check is a cross platform hex to clear the terminal, rather than
// using go's map, and system its better to clear or print this
var (
	BACKBLUE = "\033[0;44m"
	BLK      = "\033[0;30m"
	RED      = "\033[0;31m"
	GRN      = "\033[0;32m"
	YEL      = "\033[0;33m"
	BLU      = "\033[0;34m"
	MAG      = "\033[0;35m"
	CYN      = "\033[0;36m"
	WHT      = "\033[0;37m"
	BBLK     = "\033[1;30m"
	BRED     = "\033[1;31m"
	BGRN     = "\033[1;32m"
	BYEL     = "\033[1;33m"
	BBLU     = "\033[1;34m"
	BMAG     = "\033[1;35m"
	BCYN     = "\033[1;36m"
	BWHT     = "\033[1;37m"
	UBLK     = "\033[4;30m"
	URED     = "\033[4;31m"
	UGRN     = "\033[4;32m"
	UYEL     = "\033[4;33m"
	UBLU     = "\033[4;34m"
	UMAG     = "\033[4;35m"
	UCYN     = "\033[4;36m"
	UWHT     = "\033[4;37m"
	BLKB     = "\033[40m"
	REDB     = "\033[41m"
	GRNB     = "\033[42m"
	YELB     = "\033[43m"
	BLUB     = "\033[44m"
	MAGB     = "\033[45m"
	CYNB     = "\033[46m"
	WHTB     = "\033[47m"
	BLKHB    = "\033[0;100m"
	REDHB    = "\033[0;101m"
	GRNHB    = "\033[0;102m"
	YELHB    = "\033[0;103m"
	BLUHB    = "\033[0;104m"
	MAGHB    = "\033[0;105m"
	CYNHB    = "\033[0;106m"
	WHTHB    = "\033[0;107m"
	HBLK     = "\033[0;90m"
	HRED     = "\033[0;91m"
	HGRN     = "\033[0;92m"
	HYEL     = "\033[0;93m"
	HBLU     = "\033[0;94m"
	HMAG     = "\033[0;95m"
	HCYN     = "\033[0;96m"
	HWHT     = "\033[0;97m"
	BHBLK    = "\033[1;90m"
	BHRED    = "\033[1;91m"
	BHGRN    = "\033[1;92m"
	BHYEL    = "\033[1;93m"
	BHBLU    = "\033[1;94m"
	BHMAG    = "\033[1;95m"
	BHCYN    = "\033[1;96m"
	BHWHT    = "\033[1;97m"
	client   http.Client
	flagname = flag.String("url", "", "URL")
	chars    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"
)

// errror handeling function, where
// msg = warning/errror message to parse with err type
// typer is a type of statement that will be used to output the
// err data, such as println, log, fatal etc
// bool and exit code to return exit status
func ce(err error, msg string, typer string, exit_code int) bool {
	if err != nil {
		if typer == "fmt" {
			fmt.Println(err, msg, exit_code)
			os.Exit(exit_code)
		}
		if typer == "log" {
			log.Fatal(err, msg)
			os.Exit(exit_code)
		}
	} else {
		return true
	}
	return true
}

// file checking to make sure the file exists, if not make one, if so generate
// a new tage and add a new name, then warn user of secondary file generation
// and the name of the file

func isexisting(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// id generation function if file robots.txt exists
func shortID(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

//https://www.google.com/robots.txt
// make function return boolean based values if the url returns with a status code outside of or statement range
func robot(uri, path string) {
	flag.Parse()
	parsed_uri := uri + path
	resp, err := http.Get(parsed_uri)
	ce(err, "ERROR DURING GET REQUEST LINE 32 ", "panic", 1)
	if resp.StatusCode == 100 || resp.StatusCode == 200 || resp.StatusCode == 201 || resp.StatusCode == 202 {
		body, err := ioutil.ReadAll(resp.Body)
		ce(err, "ERROR: the following exception occured when reading the HTTP response body", "log", 1)
		fmt.Println(WHT, "[", BHRED, "********************************* ROBOTS.TXT INFO *********************************", WHT, "]")
		// string response body of robots.txt
		fmt.Println(string(body))
		fmt.Println(WHT, "[", BHRED, "********************************* END OF ROBOTS.TXT *********************************", WHT, "]")
		resp.Body.Close()
		fmt.Println(WHT, "[", BHRED, "INFO: RESPONSE BODY CLOSED", WHT, "]")
		resp, err = http.Head(parsed_uri)
		ce(err, "ERROR: the following exception occured when reading the HTTP HEAD", "log", 1)
		resp.Body.Close()
		fmt.Println(WHT, "[", BHRED, "INFO: RESPONSE BODY CLOSED", WHT, "]")
		form := url.Values{}
		form.Add(parsed_uri, "name1")
		resp, err = http.Post(
			parsed_uri,
			"application/x-www-form-urlencoded",
			strings.NewReader(form.Encode()),
		)
		ce(err, "ERROR: the following exception occured when POSTING the HTTP data", "log", 1)
		resp.Body.Close()
		fmt.Println(WHT, "[", BHRED, "INFO: RESPONSE BODY CLOSED", WHT, "]")
		req, err := http.NewRequest("DELETE", parsed_uri, nil)
		ce(err, "ERROR: the following exception occured when making the new request method DELETE", "log", 1)
		resp, err = client.Do(req)
		ce(err, "ERROR: the following exception occured when SENDING the HTTP DELETE method request", "log", 1)
		fmt.Println(" [", REDB, "DELETE request was made to\033[1;39m", BLU, "]", uri)
		resp.Body.Close()
		fmt.Println(WHT, "[", BHRED, "INFO: RESPONSE BODY CLOSED", WHT, "]")
		req, err = http.NewRequest("PUT", parsed_uri, strings.NewReader(form.Encode()))
		ce(err, "ERROR: the following exception occured when making the new request method PUT", "log", 1)
		resp, err = client.Do(req)
		ce(err, "ERROR: the following exception occured when executing the HTTP response body", "log", 1)
		fmt.Println(" [", REDB, "PUT request was made \033[1;39m", BLU, "     ]", uri)
		resp.Body.Close()
		fmt.Println(WHT, "[", BHRED, "INFO: RESPONSE BODY CLOSED", WHT, "]")
		counter := 0
		for k := range resp.Status {
			counter++
			fmt.Println(k)
		}
		fmt.Println("[ ", REDB, " TOTAL STATUS CODES => \033[1;39m", BLU, "   ]", counter)
		write(string(body))
	} else {
		fmt.Println("ERR: sorry but the http response code i got was not in range => ", resp.StatusCode)
	}
}

func write(body1 string) {
	if !isexisting("robots.txt") {
		d1 := []byte(body1)
		err := os.WriteFile("robots.txt", d1, 0644)
		if err != nil {
			log.Fatal(err)
		}
		path, err3 := os.Getwd()
		ce(err3, " ERROR WHEN GETTING WD ( WORKING DIR ) ", "panic", 1)
		fmt.Println("FILE PATH OF GENERATED FILE => ", path)
	}
	if isexisting("robots.txt") {
		fmt.Println(WHT, "[", BHRED, "INFO: FILE robots.txt ALREADY EXISTS GENERATING NEW ONE", WHT, "]")
		s := ".txt"
		a := shortID(4) + s
		f, err := os.Create(a)
		ce(err, "ERROR CREATING FILE", "panic", 1)
		defer f.Close()
		d1 := []byte(body1)
		err1 := os.WriteFile(a, d1, 0644)
		ce(err1, " ERROR WHEN WRITING TO FILE ", "panic", 1)
		path, err3 := os.Getwd()
		ce(err3, " ERROR WHEN GETTING WD ( WORKING DIR ) ", "panic", 1)
		fmt.Println("FILE PATH OF GENERATED FILE => ", path)
	}

}

func main() {
	flag.Parse()
	path1 := "/robots.txt"
	robot(*flagname, path1)
}
