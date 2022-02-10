// this is a bit weird sorry lol tried my bets to find a decent module in golang

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/schollz/wifiscan"
)

func handelreturncon(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nDetected Interupt.....")
			os.Exit(1)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func ce(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	var chex = "\x1b[H\x1b[2J\x1b[3J"
	var BLU = "\033[0;94m"

	for {
		fmt.Println(" Sniffing every 5 seconds ")
		time.Sleep(1 * time.Second)
		go handelreturncon(make(chan os.Signal, 1))
		fmt.Println(chex)
		var wifiInterface string
		wifis, err := wifiscan.Scan(wifiInterface)
		if err != nil {
			log.Fatal(err)
		}
		if len(wifis) > 0 {
		} else {
			fmt.Println("no mac addresses found")
		}
		counter := 0
		for _, w := range wifis {
			counter += 1
			fmt.Println(BLU, counter, "| \t", w.SSID, w.RSSI)
			pathmain, err := os.OpenFile("in.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				if fileExists("in.txt") {
					fmt.Println("[ - ] Unexpected error? but why -> ", err)
					fmt.Println("[ * ] Creating file")
					f, err := os.Create("in.txt")
					ce(err)
					f.Close()
				} else {
					fmt.Println("[ - ] Unexpected error? but why -> ", err)
					fmt.Println("[ * ] Creating file")
					f, err := os.Create("in.txt")
					ce(err)
					f.Close()
					continue
				}
			}

			fmt.Fprintln(pathmain, w.SSID)
			if err != nil {
				log.Fatal(err)
			}
			b, err := ioutil.ReadFile("in.txt")
			if err != nil {
				if fileExists("in.txt") {
					fmt.Println("[ - ] Unexpected error? but why ")
					fmt.Println("[ * ] Creating file")
					f, err := os.Create("in.txt")
					ce(err)
					f.Close()
				} else {
					continue
				}
			}
			err = ioutil.WriteFile("out.txt", b, 0644)
			if err != nil {
				if fileExists("out.txt") {
					fmt.Println("[ - ] Unexpected error? but why ")
					fmt.Println("[ * ] Creating file")
					f, err := os.Create("in.txt")
					ce(err)
					f.Close()
				} else {
					continue
				}
			}
		}
	}
}

// monitor mode is NEEDED, find a function to take the net addr
// of a interface, and tune it to monitor mode on chan 1-100
// to automate channel surfing
