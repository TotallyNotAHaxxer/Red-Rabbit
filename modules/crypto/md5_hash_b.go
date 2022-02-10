/*

DEVELOPER => ArkAngeL43
Github    => https://github.com/ArkAngeL43
Project   => In contribution to the red rabbit version 5 project
Type      => CLI
Functions => check_err, banner_hash, attack, main

Purpose => Brute force lists of MD5 hashes example (hashes.txt)

Inspiration => ch-11 Cryptography

Packages => bufio, crypto/md5, flag, fmt, io/ioutil, log, os

███    ███ ██████  ███████        ███    ███  █████  ██████  ██████  ███████ ███    ███
████  ████ ██   ██ ██              ████  ████ ██   ██ ██   ██ ██   ██ ██      ████  ████
██ ████ ██ ██   ██ ███████ █████ ██ ████ ██ ███████ ██   ██ ██   ██ █████   ██ ████ ██
██  ██  ██ ██   ██      ██         ██  ██  ██ ██   ██ ██   ██ ██   ██ ██      ██  ██  ██
██      ██ ██████  ███████         ██      ██ ██   ██ ██████  ██████  ███████ ██      ██


MD5 MADDEM

*/

package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	hash_list = flag.String("f", "hashes.txt", " set a file for hash cracking")
	pass_list = flag.String("w", "/usr/share/wordlists/rockyou.txt", "set a password file")
)

func check_err(err error, msg string, exit_code int) bool {
	if err != nil {
		fmt.Println(err, msg)
		os.Exit(exit_code)
		return true
	} else {
		return false
	}
}

func banner_hash(banner_file string) {
	content, err := ioutil.ReadFile(banner_file)
	check_err(err, " Could not open or read file ", 1)
	fmt.Println("\033[37m", string(content))
}

func attack(wordlist, md5_hash string) {
	f, err := os.Open(wordlist)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password := scanner.Text()
		hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if hash == md5_hash {
			fmt.Printf("\033[35m[MD5-MADDEM] STAT_CHECK: FOUND PASSWORD FOR HASH \033[32m[%s] \033[37m-> %s\n", md5_hash, password)
		}
		/*
			hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sha256hash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
			}
		*/
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func data_(filename string) {
	content, err := ioutil.ReadFile(filename)
	check_err(err, " Could not open or read file ", 1)
	fmt.Println("\t\t\tHASHES LOADED IN FILE ", filename)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("\033[31m", string(content))
}

func main() {
	banner_hash("banner.txt")
	flag.Parse()
	hash_file, err := os.Open(*hash_list)
	check_err(err, "could not open hash list", 1)
	defer hash_file.Close()
	hashes := bufio.NewScanner(hash_file)
	//
	data_(*hash_list)
	for hashes.Scan() {
		attack(*pass_list, hashes.Text())
	}
}
