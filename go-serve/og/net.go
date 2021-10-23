package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	. "github.com/logrusorgru/aurora"
)

func design() {
	f, err := os.Open("banner.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(Blue(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func IsOnline() bool {
	_, err := http.Get("https://www.google.com")
	if err == nil {
		return true
	}
	fmt.Println(Cyan("[-] Interface has been disconnected from the network, please connect or set a connection "))
	return false
}

func mndesknot() {
	if runtime.GOOS == "windows" {
		fmt.Println(Cyan("[-] Sorry, but t this time i can not run this command"))
	} else {
		out, err := exec.Command("notify-send", "Testing Server Conn and Node every 20-30 seconds").Output()
		if err != nil {
			log.Fatal(err)
		} else {
			output := string(out[:])
			fmt.Println(output)
		}
	}
}

func resplog() {
	url := "https://google.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		fmt.Println("[-] Sorry will not be able to run this command")
	} else {
		if resp.StatusCode >= 200 {
			out, err := exec.Command("notify-send", "Server responded with code 200 Connection is stable  °˖✧◝(⁰▿⁰)◜✧˖° ✔️").Output()
			if err != nil {
				log.Fatal(err)
			}
			output := string(out[:])
			fmt.Println(output)
		} else {
			out, err := exec.Command("notify-send", "Server Responded with a code that is not within the indexed list or range").Output()
			if err != nil {
				log.Fatal(err)
			}
			output := string(out[:])
			fmt.Println(output)
		}
	}
}

func logged() {
	if runtime.GOOS == "windows" {
		fmt.Println("This appends to a linux system only command, i will not be able to run it")
	} else {
		out, err := exec.Command("notify-send", "There was an error within the response").Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out[:])
		fmt.Println(output)
	}
}

func clear() {
	if runtime.GOOS == "windows" {
		fmt.Println(Red("[-] I Will not be able to execute this"))
	} else {
		out, err := exec.Command("clear").Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out[:])
		fmt.Println(output)
	}
	if runtime.GOOS == "windows" {
		os := "linux"
		fmt.Println("[-] Sorry, this command is system spacific to -> ", os, "Systems")
	} else {
		out, err := exec.Command("pwd").Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out[:])
		fmt.Println("[~] Working Directory ~> ", output)
	}
}

func get() {
	clear()
	url := "https://google.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	design()
	fmt.Println(Cyan("--------------------------Server Response---------------------------"))
	fmt.Println("[+] Response Status  -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Println("[+] Date Of Request  -> ", resp.Header.Get("date"))
	fmt.Println("[+] Content-Encoding -> ", resp.Header.Get("content-encoding"))
	fmt.Println("[+] Content-Type     -> ", resp.Header.Get("content-type"))
	fmt.Println("[+] Connected-Server -> ", resp.Header.Get("server"))
	fmt.Println("[+] X-Frame-Options  -> ", resp.Header.Get("x-frame-options"))
	fmt.Println(Cyan("--------------------------Server X-Requests-----------------------------"))
	for k, v := range resp.Header {
		fmt.Print(Cyan("[+] -> " + k))
		fmt.Print(Red(" -> "))
		fmt.Println(v)
	}
}

func urlidea() {
	var url string
	fmt.Scanf("%s", &url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		fmt.Println("Can not run this string on win32-64")
		os.Exit(1)
	} else {
		fmt.Println("[+] YAY STRING VARIABLES ARE WORKING!!! response code -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
		os.Exit(1)
	}
}

func main() {
	IsOnline()
	clear()
	time.Sleep(10 * time.Second)
	mndesknot()
	seconds := "20"
	time.Sleep(1 * time.Second)
	fmt.Println("[~] Testing Connection Every ", seconds, "Seconds")
	time.Sleep(1 * time.Second)
	for {
		time.Sleep(30 * time.Second)
		url := "https://google.com"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode >= 200 {
			fmt.Println("[+] Response Status Given -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println("[+] Response seems good")
			resplog()
			get()
		}
		if resp.StatusCode >= 300 && resp.StatusCode <= 400 {
			fmt.Println("[+] Response Status Given -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println("[~] Response may be laggy")
			logged()
			get()
		}
	}
}
