package Brute_Forcing_

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/alexmullins/zip"
)

type Settings struct {
	Filename    string
	Wordlist    string
	Concurrency int
}

var (
	Wait_Chan sync.WaitGroup
	DONE_CHAN = make(chan bool)

	Word  = make(chan string, 0)
	Found = make(chan string, 0)
)

func ZIP_CRACK(zip_file, wordlist string, current int) string {
	fmt.Println("Starting brute......")
	settings := Settings{
		Filename:    zip_file,
		Wordlist:    wordlist,
		Concurrency: current,
	}
	for i := 0; i < settings.Concurrency; i++ {
		Wait_Chan.Add(1)
		go Workers(settings, Word, Found, &Wait_Chan)
	}
	f, X := os.Open(settings.Wordlist)
	if X != nil {
		fmt.Println("<RR6> File I/O -> Brute forcing module -> Readers: Could not find, locate, load, or open the given file, please check your settings, maybe the directory path is invalid, or the file could be corrupted or missing please try again -> ", X)
	} else {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		go func() {
			for scanner.Scan() {
				data := scanner.Text()
				Word <- data
			}
			close(Word)

		}()
		go func() {
			Wait_Chan.Wait()
			DONE_CHAN <- true
		}()
		select {
		case f := <-Found:
			return f

		case <-DONE_CHAN:
			return "FAILED: JOB - Could not find password CODE 901"
		}
	}
	return "{No Data, odd error if error was not 0x00 but not 0x01 what was it?}"
}

func Workers(information Settings, Password <-chan string, password_found chan<- string, wait_group *sync.WaitGroup) {
	reader, X := zip.OpenReader(information.Filename)
	if X != nil {
		fmt.Println("<RR6> Brute forcing -> File -> I/O: Got error when trying to create a new open reader for the given zip file, this could be due to the fact that the file may be corrupted, of the wrong directory, of the wrong type, of the wrong permissions or corrupted, please try again and troubleshoot the following error -> ", X)
	} else {
		defer reader.Close()
		defer wait_group.Done()
		for k := range Password {
			for _, Q := range reader.File {
				Q.SetPassword(k)
				_, X := Q.Open()
				if X == nil {
					password_found <- k
				}
			}
		}

	}
}
