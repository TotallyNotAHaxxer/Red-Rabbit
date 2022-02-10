/*

Author   => ArkAngeL43
Github   => https://github.com
Package  => Main
App Type => CLI
Language => Go
*/
package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	now         = time.Now()
	formatDate  = now.Format("15:04:05")
	target_list = flag.String("l", "", " Set a list of hosts")
	clear_hex   = "\x1b[H\x1b[2J\x1b[3J"
	BLK         = "\033[0;30m"
	RED         = "\033[0;31m"
	GRN         = "\033[0;32m"
	YEL         = "\033[0;33m"
	BLU         = "\033[0;34m"
	MAG         = "\033[0;35m"
	CYN         = "\033[0;36m"
	WHT         = "\033[0;37m"
	BBLK        = "\033[1;30m"
	BRED        = "\033[1;31m"
	BGRN        = "\033[1;32m"
	BYEL        = "\033[1;33m"
	BBLU        = "\033[1;34m"
	BMAG        = "\033[1;35m"
	BCYN        = "\033[1;36m"
	BWHT        = "\033[1;37m"
	UBLK        = "\033[4;30m"
	URED        = "\033[4;31m"
	UGRN        = "\033[4;32m"
	UYEL        = "\033[4;33m"
	UBLU        = "\033[4;34m"
	UMAG        = "\033[4;35m"
	UCYN        = "\033[4;36m"
	UWHT        = "\033[4;37m"
	BLKB        = "\033[40m"
	REDB        = "\033[41m"
	GRNB        = "\033[42m"
	YELB        = "\033[43m"
	BLUB        = "\033[44m"
	MAGB        = "\033[45m"
	CYNB        = "\033[46m"
	WHTB        = "\033[47m"
	BLKHB       = "\033[0;100m"
	REDHB       = "\033[0;101m"
	GRNHB       = "\033[0;102m"
	YELHB       = "\033[0;103m"
	BLUHB       = "\033[0;104m"
	MAGHB       = "\033[0;105m"
	CYNHB       = "\033[0;106m"
	WHTHB       = "\033[0;107m"
	HBLK        = "\033[0;90m"
	HRED        = "\033[0;91m"
	HGRN        = "\033[0;92m"
	HYEL        = "\033[0;93m"
	HBLU        = "\033[0;94m"
	HMAG        = "\033[0;95m"
	HCYN        = "\033[0;96m"
	HWHT        = "\033[0;97m"
	BHBLK       = "\033[1;90m"
	BHRED       = "\033[1;91m"
	BHGRN       = "\033[1;92m"
	BHYEL       = "\033[1;93m"
	BHBLU       = "\033[1;94m"
	BHMAG       = "\033[1;95m"
	BHCYN       = "\033[1;96m"
	BHWHT       = "\033[1;97m"
	client      http.Client
	flagname    = flag.String("url", "", "URL")
	chars       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"
	robots_path = "/robots.txt"
)

type Load struct {
	perchecknt int64
	cur        int64
	total      int64
	rate       string
	graph      string
}

func (bar *Load) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "."
	}
	bar.perchecknt = bar.getPerchecknt()
	for i := 0; i < int(bar.perchecknt); i += 2 {
		bar.rate += bar.graph
	}
}

func (bar *Load) getPerchecknt() int64 {
	return int64((float32(bar.cur) / float32(bar.total)) * 100)
}

func (bar *Load) Play(cur int64) {
	bar.cur = cur
	last := bar.perchecknt
	bar.perchecknt = bar.getPerchecknt()
	if bar.perchecknt != last && bar.perchecknt%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\b\033[36m\r[%-50s]%3d%% Loading and awaiting for user data %8d/%d", bar.rate, bar.perchecknt, bar.cur, bar.total)
}

func (bar *Load) Finish() {
	fmt.Println()
}

func main_bar() {
	var load_bar Load
	load_bar.NewOption(0, 100)
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		load_bar.Play(int64(i))
	}
	load_bar.Finish()
}

/*
FUNCTION: Check

Returns: Nothing unless error does not = nil then panic

arguments:
		ERR VALUE, OS EXIT CODE | USAGE: check(err, 1)

*/
func check(err error, exit_code int) bool {
	if err != nil {
		log.Fatal(err)
		os.Exit(exit_code)
		return true
	}
	return false
}

/*
FUNCTION: BANNER

RETURNS: OUTPUT OF A FILE ARG

TAKES ARGUMENTS:
			File string | this is the filename call this function with banner("filename.txt") or whatever your file is

*/
func banner(file string) {
	fmt.Println(clear_hex)
	content, err := ioutil.ReadFile(file)
	check(err, 1)
	fmt.Println(BLU, string(content))
}

/*
Function:
	print_user_target_data

Returns: information of the user and the targets the user is targetting

Takes arguments:
		list 		   | the list of hosts
		path 		   | the path to find in the hosts
		request_method | the request method being used to send the http request
		request_path   | the request path that will be searched for
*/

func print_user_target_data(list, path, request_method string) {
	fmt.Println(RED, " List      Searching path     HTTP Request method       ")
	fmt.Println(RED, "---------------------------------------------------------")
	fmt.Printf("\033[31m [ %s ]           [ %s ]            [ %s ] ", list, path, request_method)
}

/*
Function:
	Killer

Returns:
	Nothing, listens for OS interuptions

Takes arguments:
		Go, Channel, and os signal
*/

func Killer(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nInteruption Recieved....")
			os.Exit(1)
		case os.Kill:
			fmt.Println("\n\n\tKILL recheckived")
			os.Exit(1)
		}
	}
}

/*
Function:
	get_req_main

Does:
	Making the inital get request for every host in the file name

Returns:
	Boolean based HTTP request answers to see if the filepath exists via request, ERR BASED TESTING

Takes Arguments:
	  	filename | filename of hosts

*/

func get_req_main(hosts_file, def_path string) {
	std_file_hosts, err := os.Open(hosts_file)
	check(err, 1)
	defer std_file_hosts.Close()
	scanner := bufio.NewScanner(std_file_hosts)
	for scanner.Scan() {
		pathmain := scanner.Text() + def_path
		resp, err := http.Get(pathmain)
		check(err, 1)
		//
		if resp.StatusCode == 200 || resp.StatusCode == 202 || resp.StatusCode == 206 {
			fmt.Println("\n", RED, "\033[31m[PATH_FINDER] ", BLU, formatDate, WHT, " PATH FOUND ~~> ", pathmain, GRN, " CAME BACK TRUE")
		} else {
			fmt.Println("\n", RED, "\033[31m[PATH_FINDER] ", BLU, formatDate, " TEST FOR URL -> ", pathmain, RED, " TURNED FALSE ")
		}
		robot(scanner.Text(), robots_path)
	}

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

func robot(uri, path string) {
	flag.Parse()
	parsed_uri := uri + path
	resp, err := http.Get(parsed_uri)
	check(err, 1)
	if resp.StatusCode == 100 || resp.StatusCode == 200 || resp.StatusCode == 201 || resp.StatusCode == 202 {
		body, err := ioutil.ReadAll(resp.Body)
		check(err, 1)
		resp.Body.Close()
		resp, err = http.Head(parsed_uri)
		check(err, 1)
		resp.Body.Close()
		form := url.Values{}
		form.Add(parsed_uri, "name1")
		resp, err = http.Post(
			parsed_uri,
			"application/x-www-form-urlencoded",
			strings.NewReader(form.Encode()),
		)
		check(err, 1)

		resp.Body.Close()
		req, err := http.NewRequest("DELETE", parsed_uri, nil)
		check(err, 1)
		resp, err = client.Do(req)
		check(err, 1)
		resp.Body.Close()
		req, err = http.NewRequest("PUT", parsed_uri, strings.NewReader(form.Encode()))
		check(err, 1)
		resp, err = client.Do(req)
		check(err, 1)
		resp.Body.Close()
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
		check(err3, 1)
		fmt.Println("\n", RED, "\033[31m[PATH_LOGGER] ", BLU, formatDate, " File path of new generated file => ", path)
	}
	if isexisting("robots.txt") {
		fmt.Println("\n", RED, "\033[31m[PATH_LOGGER] ", BLU, formatDate, " FILE robots.txt must have already been in use, downloading a new one ")
		s := ".txt"
		a := shortID(4) + s
		f, err := os.Create(a)
		check(err, 1)
		defer f.Close()
		d1 := []byte(body1)
		err1 := os.WriteFile(a, d1, 0644)
		check(err1, 1)
		path, err3 := os.Getwd()
		check(err3, 1)
		fmt.Println("\n", RED, "\033[31m[PATH_LOGGER] ", BLU, formatDate, " File path of new generated file => ", path)
	}

}

func main() {
	flag.Parse()
	banner("conf/banner.txt")
	//main_bar()
	print_user_target_data(*target_list, ".env", "GET")
	fmt.Print("\n\n\n")
	get_req_main(*target_list, ".env")
}
