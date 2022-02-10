package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	filename = flag.String("f", "", "")
	bytes    = flag.Int("b", 256, "")
)

func dump(color, file string, b int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	if b >= 90000 || b >= 956 {
		fmt.Println(" [ - ] FATAl: CAUGHT: BYTE SIZE TO LARGE")
		os.Exit(1)
	} else {
		buf := make([]byte, b)
		for {
			_, err := reader.Read(buf)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			fmt.Println(hex.Dump(buf))
		}
	}
}

func main() {
	flag.Parse()
	dump("\033[31m", *filename, *bytes)
}
