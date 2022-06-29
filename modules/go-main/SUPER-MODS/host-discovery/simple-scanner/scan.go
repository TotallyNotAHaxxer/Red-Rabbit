package SUPER_Scan

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	color "main/modg/colors"
	structure "main/modules/go-main/xml/types"
	"os"
	"strings"

	"net"
	"strconv"
	"time"
)

type Scan_range struct {
	Start int
	End   int
}

func Matcher(filename, data string) {
	f, x := os.Open(filename)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Close()
		b, _ := ioutil.ReadAll(f)
		var s structure.Registry
		xml.Unmarshal(b, &s)
		for i := 0; i < len(s.Record); i++ {
			num := s.Record[i].Number
			if strings.Compare(data, num) == 0 {
				fmt.Printf("\033[37m   |>> \033[38;5;21mPort number        |  \033[31m%s    \n", num)
				fmt.Printf("\033[37m   |>> \033[38;5;21mPort protocol      |  \033[31m%s    \n", s.Record[i].Protocol)
				fmt.Printf("\033[37m   |>> \033[38;5;21mPort service       |  \033[31m%s    \n", s.Record[i].Name)
				fmt.Printf("\033[37m   |>> \033[38;5;21mPort unauthorized? |  \033[31m%s    \n", s.Record[i].Unauthorized)
				fmt.Printf("\033[37m   |>> \033[38;5;21mPort Description   |  \033[31m%s    \n", s.Record[i].Description)
				fmt.Print("\n")
			}
		}
	}
}

func Scan(filename, location string, ports Scan_range) {
	ac := 0
	dc := make(chan bool)
	for i := ports.Start; i <= ports.End; i++ {
		go TO(filename, location, i, dc)
		ac++
	}
	for ac > 0 {
		<-dc
		ac--
	}
}

func TO(filename, ip string, port int, doneChannel chan bool) {
	_, x := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if x == nil {
		a := strconv.Itoa(port)
		if a != "" {
			fmt.Println(color.BHMAG, ": ", port, "\033[39m\033[49m", color.RET_RED)
			fmt.Println(color.BHMAG, ": ", "\033[39m\033[49m", color.RET_RED)
			fmt.Println(color.BHMAG, ": ", "\033[39m\033[49m", color.RET_RED)
			Matcher(filename, a)
		}
	}
	doneChannel <- true
}
