// file to get and downlaod the html file in golang

package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	v     = flag.String("targ", "", "")
	chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"
	rev   = "\033[0;39m"
	reb   = "\033[49m"
	blk   = "\033[0;30m"
	red   = "\033[0;31m"
	grn   = "\033[0;32m"
	yel   = "\033[0;33m"
	blu   = "\033[0;34m"
	mag   = "\033[0;35m"
	cyn   = "\033[0;36m"
	wht   = "\033[0;37m"
	blkb  = "\033[40m"
	redb  = "\033[41m"
	grnb  = "\033[42m"
	yelb  = "\033[43m"
	blub  = "\033[44m"
	magb  = "\033[45m"
	cynb  = "\033[46m"
	whtb  = "\033[47m"
)

func shortID(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

func writter_v2(url string) {
	flag.Parse()
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("INFO: FATAL: ERR: COULD NOT GET URL? ")
	}
	defer r.Body.Close()
	s := ".html"
	fmt.Println(wht+"[ "+redb+"GET"+wht+" ]"+rev+reb+blu+" INFO: GET request made to ", *v, " Generating character list")
	a := shortID(19) + s
	fmt.Println(wht+"[ "+redb+"GET"+wht+" ]"+rev+reb+blu+" INFO: GET DOWNLOAD: Filename and key generated -> ", a)
	f, err := os.Create(a)

	if err != nil {
		fmt.Println("INFO: FATAL: ERR: COULD NOT GET URL AND WRITE URL's HTML")
	}
	defer f.Close()
	_, err = f.ReadFrom(r.Body)
	if err != nil {
		fmt.Println("INFO: FATAL: ERR: COULD NOT READ FROM r.BODY!!!! ")
	} else {
		fmt.Println(wht + "[ " + redb + "GET" + wht + " ]" + rev + reb + blu + " INFO: GET DOWNLOAD SUCCESSFUL")
	}
}

func main() {
	flag.Parse()
	fmt.Println(wht + "[ " + redb + "GET" + wht + " ]" + rev + reb + blu + " DATA: Making GET request...")
	writter_v2(*v)
}
