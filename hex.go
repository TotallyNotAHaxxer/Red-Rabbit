package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	. "github.com/logrusorgru/aurora"
)

func listdir() {
	ex := "ls"

	cmd := exec.Command(ex)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(Cyan(string(stdout)))
}

func main() {
	fmt.Println(Cyan("hello"))
	listdir()
	var name string
	fmt.Println(Red("------------------------------------------------------------------------------"))
	fmt.Println("File to Dump | ")
	fmt.Scanf("%s", &name)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("[+] Opening File....", name)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("[+] Hex Dumping....")
	time.Sleep(1000 * time.Millisecond)
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 256)

	for {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(Red("[+] HEX FORMAT | "))
		fmt.Printf("%s", hex.Dump(buf))
		//fmt.Printf(Red("%s", hex.Dump(buf))
	}
}
