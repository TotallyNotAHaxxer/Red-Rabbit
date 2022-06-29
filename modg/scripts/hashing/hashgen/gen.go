/*
Developer | ArkangeL43
Package   | gen
Module    | hashgen
File      | modg/scripts/hashing/hashgen/gen.go
Nest      | scripts/hashing

Does:
	MD5, SHA1, SHA224, SHA256, BASE64, BASE32, ROT13, HMAC, CSPRNG, and more cyrptographic generation, this is mainly a package that will be ded
	icated the the hashing functions and callers inside of RR6

Finished:
	SHA1
	SHA256
	SHA224
	MD5
	base64
	base32
	hashing files
	hashing larger files
	CSPRING
	HMAC
	ROT13

*/

// cg = Crypto gen
package CG

import (
	"bufio"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	y "main/modg/colors"
	cn "main/modg/constants"
	logging "main/modg/switch/loggers"
	xc "main/modg/warnings"
	"math"
	"math/big"
	"os"
	"strings"
)

// wordlist generation
func Key_gen(length int) string {
	const keyList string = "abcdefghijklmnopqrstuvwxyzABCDEFHFGHIJKLMNOPQRSTUVWXYZ1234567890"
	ll := len(keyList)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = keyList[int(b[i])%ll]
	}
	return string(b)
}

// SHA1 generation (works)
func Sha1_gen(msg string) (string, error) {
	alg := sha1.New()
	alg.Write([]byte(msg))
	s := alg.Sum(nil)
	sa := string(fmt.Sprintf("%x\n", s))
	return sa, nil
}

// SHA256 generation (works)
func Sha256_gen(msg string) (string, error) {
	a := sha256.New()
	a.Write([]byte(msg))
	s := a.Sum(nil)
	sa := string(fmt.Sprintf("%x\n", s))
	return sa, nil
}

//SHA-512 generation (works)
func Sha_512_gen(msg string) (string, error) {
	a := sha512.New()
	a.Write([]byte(msg))
	s := a.Sum(nil)
	sa := string(fmt.Sprintf("%x\n", s))
	return sa, nil
}

// SHA224 generation (works)
func Sha224_gen(msg string) (string, error) {
	a := sha512.New512_224()
	a.Write([]byte(msg))
	s := a.Sum(nil)
	sa := string(fmt.Sprintf("%x\n", s))
	return sa, nil

}

// MD5 generation (works)
func MD5_gen(msg string) (string, error) {
	a := md5.New()
	a.Write([]byte(msg))
	s := a.Sum(nil)
	sa := string(fmt.Sprintf("%x\n", s))
	return sa, nil
}

// base64 encoding (works)
func Encode_base64(msg string) (string, error) {
	enc := base64.StdEncoding.EncodeToString([]byte(msg))
	return enc, nil
}

//base32 encoding (works)
func Encode_base32(msg string) (string, error) {
	enc := base32.StdEncoding.EncodeToString([]byte(msg))
	return enc, nil
}

// hash file (working)
func HF(filename string) error {
	d, e := ioutil.ReadFile(filename)
	xc.Che(e, "<RR6> File Module: Could not open file -> ", 1)
	fmt.Printf("MD5    FILE HASH -  %x\n\n", md5.Sum(d))
	fmt.Printf("SHA1   FILE HASH - %x\n\n", sha1.Sum(d))
	fmt.Printf("SHA256 FILE HASH - %x\n\n", sha256.Sum256(d))
	fmt.Printf("SHA512 FILE HASH - %x\n\n", sha512.Sum512(d))
	return nil
}

// hash a large file (working)
func HFL(filename string) error {
	d, r := os.Open(filename)
	xc.Ce(r, y.REDHB, "<RR6> File module: Could not open filename, error has occured -> ", 1)
	defer d.Close()
	hs := md5.New()
	_, err := io.Copy(hs, d)
	xc.Che(err, "<RR6> Crypto module: Could not hash filename or grab checksum", 1)
	checksum := hs.Sum(nil)
	fmt.Printf("\nMD5 HASH CHECKSUM -> %x\n", checksum)
	return nil
}

// CSPRNG (working)
func CS() error {
	m := int64(math.MaxInt64)
	r, e := rand.Int(rand.Reader, big.NewInt(m))
	xc.Che(e, "<RR6> Crypto Module: Could not generate random integer -> ", 1)
	fmt.Println("<RR6> | Random Number Value - ", r)
	re := binary.Read(rand.Reader, binary.BigEndian, &cn.Num)
	xc.Che(re, "<RR6> Crypto Module: Could not make binary reader -> | ", 1)
	fmt.Println("<RR6> | Random uint32 Value - ", cn.Num)
	nb := 4
	rb := make([]byte, nb)
	rand.Read(rb)
	fmt.Println("<RR6> | Random Byte Value -", rb)
	return nil
}

// HMAC (working)
func gsalt() (string, error) {
	r := make([]byte, 32)
	_, e := rand.Read(r)
	xc.Warning_simple("<RR6> Crypto Module: Could not load or make byte, error -> ", y.REDHB, e)
	return base64.URLEncoding.EncodeToString(r), nil
}

func hp(text, salt, key string) (string, error) {
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, text+salt)
	v := h.Sum(nil)
	return hex.EncodeToString(v), nil
}

func Final_generate(password, key string) {
	salted_hash, e := gsalt()
	xc.Ce(e, y.REDHB, "<RR6> Crypto Module: Could not generate salt", 1)
	hap, e := hp(password, salted_hash, key)
	xc.Ce(e, y.REDHB, "<RR6> Crypto Module: Could not generate hash passwords", 1)
	fmt.Println("<RR6> Crypto Module | Salt - ", salted_hash)
	fmt.Println("<RR6> Crypto Module | Hash - ", hap)
	fmt.Println("<RR6> Crypto Module | Key  - ", key)
	fmt.Println("<RR6> Crypto Module | Pass - ", password)
}

// COC (Cipher of Ceaser) (encode and decode)
// for some reason this keeps outputting as a decode not encode
// FIXME:
func Enc(key int, pass string) string {
	en := strings.Map(func(r rune) rune {
		return caesar(r, -key)
	}, pass)
	return en
}

func Dec(key int, pass string) string {
	de := strings.Map(func(r rune) rune {
		return caesar(r, +key)
	}, pass)
	return de
}

func caesar(r rune, pushl int) rune {
	s := int(r) + pushl
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

//ROT13 (also known as rotating the input and letters 13 different places)
// (working)
func R3(r rune) rune {
	if r >= 'a' && r <= 'z' {
		if r <= 'm' {
			return r - 13
		} else {
			return r + 13
		}
	} else if r >= 'A' && r <= 'Z' {
		if r >= 'M' {
			return r - 13
		} else {
			return r + 13
		}
	}
	return r
}

func R3c(pass string) {
	m := strings.Map(R3, pass)
	fmt.Println("<RR6> Cryptography Module: Password      | ", pass)
	fmt.Println("<RR6> Cryptography Module: ROT13 Version | ", m)
}

func Logger(data, filename string) {
	logging.Logger(data, filename)
}

// file encoding types
func Listed_Generation(filename, output string, hashingt int) (string, error) {
	c, e := os.Open(filename)
	xc.Ce(e, y.REDHB, "<RR6> Cryptography Module: Could not open filename ", 1)
	scanner := bufio.NewScanner(c)
	var counter = 0
	for scanner.Scan() {
		counter++
		fmt.Print("At line < ", counter, " > | Hash => ")
		switch hashingt {
		case 1:
			a, e := MD5_gen(scanner.Text())
			xc.Warning_simple("<RR6> Cryptography Module: Could not read or make a new reader for MD5 hash generation    -> ", y.REDB, e)
			fmt.Println(string(a), " -> \033[37m", scanner.Text(), "\033[39m")
			Logger(a, output)
		case 2:
			a, e := Sha1_gen(scanner.Text())
			xc.Warning_simple("<RR6> Cryptography Module: Could not read or make a new reader for SHA1 hash generation   -> ", y.REDB, e)
			fmt.Println(string(a), " -> \033[37m", scanner.Text(), "\033[39m")
			Logger(a, output)
		case 3:
			a, e := Sha256_gen(scanner.Text())
			xc.Warning_simple("<RR6> Cryptography Module: Could not read or make a new reader for SHA256 hash generation -> ", y.REDB, e)
			fmt.Println(string(a), " -> \033[37m", scanner.Text(), "\033[39m")
			Logger(a, output)
		case 4:
			a, e := Encode_base32(scanner.Text())
			xc.Warning_simple("<RR6> Cryptography Module: Could not read or make a new reader for Base32 String encoding -> ", y.REDB, e)
			fmt.Println(string(a), " -> \033[37m", scanner.Text(), "\033[39m")
			Logger(a, output)
		case 5:
			a, e := Encode_base64(scanner.Text())
			xc.Warning_simple("<RR6> Cryptography Module: Could not read or make a new reader for Base32 String encoding -> ", y.REDB, e)
			fmt.Println(string(a), " -> \033[37m", scanner.Text(), "\033[39m")
			Logger(a, output)

		case 6:
			R3c(scanner.Text())
		case 7:
			a, e := Sha_512_gen(scanner.Text())
			xc.Warning_simple("<RR6> Cryptography Module: Could not read or make a new reader for SHA256 hash generation -> ", y.REDB, e)
			fmt.Println(string(a), " Password for hash -> \033[37m", scanner.Text(), "\033[39m")
			Logger(a, output)
		}
	}
	return "", nil
}

func Call_all(type_hash, key string, msg string) {
	if type_hash == "md5" {
		newkey, err := MD5_gen(msg)
		xc.Warning_simple("<RR6> Cryptography Module: Could not encode string into a MD5 hash", y.REDHB, err)
		fmt.Println("<RR6> Cryptography Module: hash -> ", newkey)
	}
	if type_hash == "sha1" {
		newkey, err := Sha1_gen(msg)
		xc.Warning_simple("<RR6> Cryptography Module: Could not encode string into a SHA1 hash", y.REDHB, err)
		fmt.Println("<RR6> Cryptography Module: hash -> ", newkey)
	}
	if type_hash == "sha256" {
		newkey, err := Sha256_gen(msg)
		xc.Warning_simple("<RR6> Cryptography Module: Could not encode string into a SHA256 hash", y.REDHB, err)
		fmt.Println("<RR6> Cryptography Module: hash -> ", newkey)
	}
	if type_hash == "sha224" {
		newkey, err := Sha224_gen(msg)
		xc.Warning_simple("<RR6> Cryptography Module: Could not encode string into a SHA224 hash", y.REDHB, err)
		fmt.Println("<RR6> Cryptography Module: hash -> ", newkey)
	}
	if type_hash == "base64" {
		newkey, err := Encode_base64(msg)
		xc.Warning_simple("<RR6> Cryptography Module: Could not encode string into a BASE64 string", y.REDHB, err)
		fmt.Println("<RR6> Cryptography Module: hash -> ", newkey)
	}
	if type_hash == "base32" {
		newkey, err := Encode_base32(msg)
		xc.Warning_simple("<RR6> Cryptography Module: Could not encode string into a BASE32 string", y.REDHB, err)
		fmt.Println("<RR6> Cryptography Module: hash -> ", newkey)
	}
	if type_hash == "hash small file md5" {
		err := HF(msg)
		xc.Warning_simple("<RR6> Cryptography Module: Could not encode file into a MD5 hash type", y.REDHB, err)
	}
	if type_hash == "hash large file md5" {
		err := HFL(msg)
		xc.Warning_simple("<RR6> Cryptography Module: Could not encode large file into a MD5 hash type", y.REDHB, err)
	}
	if type_hash == "sha512" {
		newkey, e := Sha_512_gen(msg)
		if e != nil {
			log.Fatal(e)
		} else {
			fmt.Println(string(newkey))
		}
	}
	if type_hash == "rot13" {
		R3c(msg)
	}
	if type_hash == "HMAC" {
		Final_generate(msg, key)
	}
}
