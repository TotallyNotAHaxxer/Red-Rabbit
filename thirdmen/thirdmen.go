package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	. "github.com/logrusorgru/aurora"
)

func clear() {
	ex := "clear"
	cmd := exec.Command(ex)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(Red(string(stdout)))
}

func banner() {
	clear()
	f, err := os.Open("men.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(Red(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	banner()
}
