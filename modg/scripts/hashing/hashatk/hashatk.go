/*
Developer | ArkangeL43
Package   | hashatk
Module    | hashatk
File      | modg/scripts/hashing/hashatk.go
Nest      | scripts/hashing

Does:
	MD5, SHA-1, SHA-2, UNIX hash, and SHA256 hash brute forcing

*/
package hashatk

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"log"
	v "main/modg/colors"
	ec "main/modg/warnings"

	"os"
)

/*


SHA256 hash attacks

*/

func brute_SHA256(wordlist, sha_hash string) {
	f, err := os.Open(wordlist)
	ec.Warning_simple("<RR6> Cryptography Module: Could not open wordlist > ", v.REDHB, err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password := scanner.Text()
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		if hash == sha_hash {
			fmt.Println(v.RED, "<RR6> Cryptography Module: Found HASH-> ["+v.REDHB+sha_hash+" \033[49m\033[31m] PASS-> ["+v.REDHB+password+"\033[49m\033[31m]")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

// list brute forcing SHA256
func Brute_SHA256_main(hashlist, wordlist string) {
	hash_file, err := os.Open(hashlist)
	ec.Warning_advanced("<RR6> Brute forcing Module: Could not open up hash file", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer hash_file.Close()
	hashes := bufio.NewScanner(hash_file)
	for hashes.Scan() {
		brute_SHA256(wordlist, hashes.Text())
	}
}

// single hash brute forcing [SHA256]
func Brute_SHA256_main_Single(wordlist, hash string) {
	brute_SHA256(wordlist, hash)
}

/*




MD5 brute forcing



*/

// MD5 hash attacks
func Brute_MD5_main(wordlist, md5_hash string) {
	f, err := os.Open(wordlist)
	ec.Warning_simple("<RR6> Cryptography Module: Could not open wordlist > ", v.REDHB, err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password := scanner.Text()
		hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if hash == md5_hash {
			fmt.Println(v.RED, "<RR6> Cryptography Module: FOUND PASSWORD HASH-> ["+v.REDHB+md5_hash+" \033[49m\033[31m] PASS-> ["+v.REDHB+password+"\033[49m\033[31m]")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

// brute force with MD5 a list of hashes
func Brute_MD5_wordlist(wordlist, hashlist string) {
	hash_file, err := os.Open(hashlist)
	ec.Warning_advanced("<RR6> Brute forcing Module: Could not open up hash file", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer hash_file.Close()
	hashes := bufio.NewScanner(hash_file)
	//
	for hashes.Scan() {
		Brute_MD5_main(wordlist, hashes.Text())
	}
}

// brute force with a single hash
func Brute_MD5_Single(wordlist, hash string) {
	Brute_MD5_main(wordlist, hash)
}

/*


SHA1 brute forcing



*/

func Brute_SHA1_main(list, sha1_hash string) {
	f, err := os.Open(list)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password := scanner.Text()
		hash := fmt.Sprintf("%x", sha1.Sum([]byte(password)))
		if hash == sha1_hash {
			fmt.Println(v.RED, "<RR6> Cyrptography: found password for hash -> ", sha1_hash, " | ", scanner.Text(), " | ")
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func Brute_SHA1_single(wordlist, hash string) {
	Brute_SHA1_main(wordlist, hash)
}

func Brute_SHA1_wordlist(wordlist, hash_list string) {
	f, err := os.Open(hash_list)
	ec.Warning_simple("<RR6> Crypto Module: Could not load hash list", v.REDHB, err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		Brute_SHA1_main(wordlist, scanner.Text())
	}
}

/*



SHA512 hash attacks



*/

func Brute_SHA512_main(list, sha512_hash string) {
	f, err := os.Open(list)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password := scanner.Text()
		hash := fmt.Sprintf("%x", sha256.Sum224([]byte(password)))
		if hash == sha512_hash {
			fmt.Println(v.RED, "<RR6> Cyrptography: found password for hash [SHA512] -> ", sha512_hash, " | ", scanner.Text(), " | ")
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func Brute_SHA512_single(wordlist, hash string) {
	Brute_SHA1_main(wordlist, hash)
}

func Brute_SHA512_wordlist(wordlist, hash_list string) {
	f, err := os.Open(hash_list)
	ec.Warning_simple("<RR6> Crypto Module: Could not load hash list", v.REDHB, err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		Brute_SHA1_main(wordlist, scanner.Text())
	}
}
