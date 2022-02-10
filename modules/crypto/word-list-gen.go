package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

var (
	key_len = flag.Int("keylen", 16, "basic key generation length")
)

const (
	keyList string = "abcdefghijklmnopqrstuvwxyzABCDEFHFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func sighandel(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nDetected Interupt.....")
			t := time.Now()
			fmt.Println("\n\n\t\033[31m[>] Script Ended At -> ", t)
			os.Exit(0)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func key_gen(length int) string {
	ll := len(keyList)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = keyList[int(b[i])%ll]
	}
	return string(b)
}

func main() {
	flag.Parse()
	pathmain, err := os.OpenFile("wordlist.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	check(err)
	defer pathmain.Close()
	for {
		go sighandel(make(chan os.Signal, 1))
		key := key_gen(*key_len)
		fmt.Println(key)
		c, err := fmt.Fprintln(pathmain, key)
		check(err)
		fmt.Println("Bytes written to wordlist -> ", c)

	}

}
