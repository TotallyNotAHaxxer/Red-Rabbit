/*

Developer => ArkAngeL43
Code      => Go


Description:
		password generation using go
*/

package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
)

var (
	BLU  = "\033[0;94m"
	BLK  = "\033[0;30m"
	RED  = "\033[0;31m"
	GRN  = "\033[0;32m"
	YEL  = "\033[0;33m"
	MAG  = "\033[0;35m"
	CYN  = "\033[0;36m"
	WHT  = "\033[0;37m"
	BBLK = "\033[1;30m"
	BRED = "\033[1;31m"
	BGRN = "\033[1;32m"
	BYEL = "\033[1;33m"
	BBLU = "\033[1;34m"
	BMAG = "\033[1;35m"
	BCYN = "\033[1;36m"
	BWHT = "\033[1;37m"
	fk0  = flag.String("file", "", "")
	fk1l = flag.Int("ple", 16, "")
	a1   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"
)

func ID(len1a4 int) string {
	llB3 := len(a1)
	b8l := make([]byte, len1a4)
	rand.Read(b8l)
	for i := 0; i < len1a4; i++ {
		b8l[i] = a1[int(b8l[i])%llB3]
	}
	return string(b8l)
}

func sighandel(c chan os.Signal) {
	flag.Parse()
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("getting information for file -> ", *fk0)
			fk0inf, err := os.Stat(*fk0)
			cce_hk3(err)
			fk0si := fk0inf.Size()
			fk0id := fk0inf.IsDir()
			mTime := fk0inf.ModTime()
			fmt.Println("File size => ", fk0si)
			fmt.Println("Is Dir    => ", fk0id)
			fmt.Println("Last Mod  => ", mTime)

			os.Exit(0)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func cce_hk3(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func file_checker(fk0 string) bool {
	// check file existance
	info, err := os.Stat(fk0)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func write_saver_a3(inb3, fk0 string) {
	go killc91(make(chan os.Signal, 1))
	file_checker(fk0)
	var lentoa1tob1 = len([]rune(inb3))
	fmt.Println(BLU, fk0, WHT, "   Written     ", BLU, inb3, WHT, "     FILE LOADED ", BLU, " \t   STRING\t", WHT, lentoa1tob1)
}

func killc91(c chan os.Signal) {
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

func main() {
	flag.Parse()
	fmt.Println(" Wordlist    Write STAT     Password     Load STAT        Type STAT      Pass Length")
	fmt.Println("-------------------------------------------------------------------------------------------")
	pathmain, err := os.OpenFile(*fk0, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		// if the path is not there, then create the file
		f, err := os.Create(*fk0)
		cce_hk3(err)
		defer f.Close()
	}
	if *fk1l >= 100 {
		output, err := exec.Command("perl", "notice").Output()
		if err == nil {
			log.Println(output)
		}
		os.Exit(1)
	}
	for {
		// handeler
		a := ID(*fk1l)
		write_saver_a3(a, "wordlist.txt")
		fmt.Fprintln(pathmain, a)
	}
}
