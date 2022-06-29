package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

var parser string

func downloadFile(directory, URL string) error {
	endpath, x := url.Parse(URL)
	if x != nil {
		fmt.Println("<RR6> GOT ERROR WHEN GRABBING END OF URL - ", x)
	} else {
		parser = directory + path.Base(endpath.Path)
	}
	resp, x := http.Get(URL)
	if x != nil {
		fmt.Println("<RR6> Got error when trying to make a request to the URL -> ", x)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Println("\033[31m<RR6> Could not make a REQUEST -> ", resp.StatusCode)
		} else {
			f, x := os.Create(parser)
			if x != nil {
				fmt.Println("<RR6> Got error when trying to create filename -> ", x)
			} else {
				defer f.Close()
				_, x = io.Copy(f, resp.Body)
				if x != nil {
					fmt.Println("<RR6> Got error when trying to copy contents -> ", x)
				}
			}

		}
	}
	return nil
}

func main() {
	url := os.Args[1]
	dir := os.Args[2]
	err := downloadFile(dir, url)
	if err != nil {
		log.Fatal(err)
	}
}
