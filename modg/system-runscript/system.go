/*
Developer   | ArkAngeL43
Package     | system
Module      | system
Nest        | system-math
FP          | modg/system-math/system

Does:
	This module is designed to take care of most system things or math
	like routines and functions, such as: root checks, user checks, os checks,
	user checks, math, wrappers, or just in general anything to do with os, sys,
	calc, etc


*/

package system

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	v "main/modg/colors"
	ec "main/modg/warnings"

	"rsc.io/pdf"
)

var Regex_type_image = []*regexp.Regexp{regexp.MustCompile(`(?i)jpg`), regexp.MustCompile(`(?i).jpg`), regexp.MustCompile(`(?i)jpeg`), regexp.MustCompile(`(?i).jpeg`)}

const (
	Match_domain = `^((?!-)[A-Za-z0-9-]{1, 63}(?<!-)\\.)+[A-Za-z]{2, 6}$`
	Match_portip = `/(\\d{1,3}\.\\d{1,3}\.\\d{1,3}\.\\d{1,3}\:\\d{1,5})/`
	Match_port   = `^([1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])`
	Match_IP     = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	Match_Mac    = `([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})|([0-9a-fA-F]{4}\\.[0-9a-fA-F]{4}\\.[0-9a-fA-F]{4})$`
	Match_Host   = `^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$`
	Match_URL    = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	Match_Email  = `([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+)`
	Match_SHA1   = `/\b([a-f0-9]{40})\b/`
)

var (
	now        = time.Now()
	FormatDate = now.Format("15:04:05")
)

// SECTION REGEX

func ValSHA1(hash string) string {
	re := regexp.MustCompile(Match_SHA1)
	if !re.MatchString(hash) {
		return "<RR6> Regex Module: Hash was not verified, this is not a proper SHA1 hash"
	}
	return "<RR6> Regex Module: Hash verified [!] "
}

func ValIP(ip string) {
	re := regexp.MustCompile(Match_IP)
	if !re.MatchString(ip) {
		fmt.Println("<RR6> Regex Module: <", ip, "> Has been verified as a real address")
	}
}

func ValMac(mac string) {
	re := regexp.MustCompile(Match_Mac)
	if re.MatchString(mac) {
		fmt.Println("<RR6> Regex Module: <", mac, "> Has been verified as a real address")
	} else {
		fmt.Println("<RR6> Regex Module: < ", mac, " > did not match the regex prefix")
	}
}

func ValPort(port string) {
	re := regexp.MustCompile(Match_port)
	if re.MatchString(port) {
		fmt.Println("<RR6> Regex Module: < ", port, " > Has been verified as a real port")
	} else {
		fmt.Println("<RR6> Regex Module: < ", port, " > Has failed to match Match_port regex string -> ", Match_port)
	}
}

func ValEmail(email string) {
	re := regexp.MustCompile(Match_Email)
	if re.MatchString(email) {
		fmt.Println("<RR6> Regex Module: < ", email, " > Has been verified as a real Email")
	} else {
		fmt.Println("<RR6> Regex Module: < ", email, " > Has failed to match Match_port regex string -> ", Match_port)
	}
}

func ValPortandIP(portip string) {
	re := regexp.MustCompile(Match_portip)
	if !re.MatchString(portip) {
		fmt.Println("<RR6> Regex Module: < ", portip, " > Has failed to match Match_port regex string -> ", Match_port)
	} else {
		fmt.Println("<RR6> Regex Module: < ", portip, " > Has been verified as a real port and ip EX: int:int")
	}
}

func Time() {
	t := time.Now()
	fmt.Println("<RR6> Clock::OS Module> LOG: Time - ")
	fmt.Println("_______________________________________")
	fmt.Println("|Current Year        |", t.Year())
	fmt.Println("|Current Month       |", t.Month())
	fmt.Println("|Current Day         | ", t.Day())
	fmt.Println("|Current Hour        |", t.Hour())
	fmt.Println("|Current Minute      |", t.Minute())
	fmt.Println("|Current Second      |", t.Second())
	fmt.Println("|Current Nanosecond  |", t.Nanosecond())
	fmt.Println("|____________________|__________________")
	fmt.Println()
}

func Warning_advanced(msg, color string, msg_format int, file, extra, rr6 bool, err error, exit_code int, dev_code int, filename string) bool {
	if err != nil {
		if rr6 {
			if msg_format == 1 {
				fmt.Println(color, "<RR6> ERRORS MODULE: Got dev code < ", dev_code, "> | -> ", msg)
			}
			if msg_format == 2 {
				fmt.Print(color, "<RR6> ERRORS MODULE: Got dev code < ", dev_code, "> | -> ", msg)
			}
			if msg_format == 3 {
				log.Println(color, "<RR6> ERRORS MODULE: Got dev code < ", dev_code, "> | -> ", msg)
			}
		} else {
			if file {
				if msg_format == 1 {
					if extra {
						fmt.Println(color, "| -> ", msg, " : MESSAGE LOG, FROM DEV, ERROR, WARNING MODULE SETTINGS PUSHED BELOW")
						fmt.Println("\033[34m ?????:::::::::::: FORMAT< msg, color, msg_format, file, extra, rr6, exit_code, dev_code, filename (bool)")
						fmt.Println("Checking Settings.......")
						fmt.Println("[*] Checking file  | BOOL   |")
						if file {
							fmt.Println(true)
						} else {
							fmt.Println(false)
						}
						fmt.Println("[*] Checking extra | BOOL   |")
						if extra {
							fmt.Println(true)
						} else {
							fmt.Println(false)
						}
						fmt.Println("[*] Checking rr6   | BOOL   |")
						if rr6 {
							fmt.Println(true)
						} else {
							fmt.Println(false)
						}
						fmt.Println("[*] Checking msg   | STRING | ")
						if msg == "" {
							fmt.Println(nil)
						} else {
							fmt.Println("filled paramater")
						}
						fmt.Println("[*] Checking color | STRING |")
						if color == "" {
							fmt.Println(nil)
						} else {
							fmt.Println("filled paramater")
						}
						fmt.Println("paramaters tested, note: warning: user based error: Developer: You should not have filled the file parmater as true with a message format of 1, this does not work, defualting to print output")
						os.Exit(1)
					}
				}
				if msg_format == 2 {
					fmt.Print(color, "| -> ", msg)
				}
				if msg_format == 3 {
					log.Println(color, "| -> ", msg)
				}
				fmt.Println("-------------------------------------- ERROR LOG ")
				fmt.Println(color, msg)
				fmt.Println("====Opening error log file...")
				content, err := os.Open(filename)
				if err != nil {
					fmt.Println(color, "<RR6> LOG MODULE: Could not open the warnings filename, note this output message should NOT appear, THIS IS NOT A WARNING THIS IS A FATAL ERROR, FILE might not exist, or have good premissions or enough premissions to open, read, and output, please make sure you are selecting the correct file or running as a sudo user")
					os.Exit(exit_code)
				} else {
					scanner := bufio.NewScanner(content)
					for scanner.Scan() {
						fmt.Println(color, content)
					}
				}
			}
		}
		return true
	}
	return false
}

// rootc
func Root_check() bool {
	currentUser, err := user.Current()
	Warning_advanced("<RR6> USER Module: Could not get username or host of this machine -> ", "\033[0;101m", 1, false, false, true, err, 1, 233, "")
	return currentUser.Username == "root"
}

func Check_parser() {
	if !Root_check() {
		fmt.Println("\033[31m[RR6] Root: User is not root, please run this script as `sudo go run main.go`")
		os.Exit(0)
	}
}

// seperator
func Sep(seperator string) {
	fmt.Print(seperator)
}

// basic file info
func File_inf(filename string) {
	var (
		err error
		fi  os.FileInfo
	)
	fi, err = os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}

	fmt.Println("File name            |", fi.Name())
	fmt.Println("Size in bytes        |", fi.Size())
	fmt.Println("Permissions          |", fi.Mode())
	fmt.Println("Last modified        | ", fi.ModTime())
	fmt.Println("Is Directory         | ", fi.IsDir())
	fmt.Printf("System interface type |  %T\n", fi.Sys())
	fmt.Printf("System info           | %+v\n\n", fi.Sys())
}

// hex dumper
func Dumper(file string, buffer int) {
	hexed, err := os.Open(file)
	ec.Che(err, " Could not read or open this file -> ", 1)
	defer hexed.Close()
	reader := bufio.NewReader(hexed)
	buffer_ := make([]byte, buffer)
	fmt.Println("\n\t\033[38m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| HEX DUMPING IN 5 SECONDS")
	time.Sleep(5 * time.Second)
	for {
		_, err := reader.Read(buffer_)
		ec.Che(err, "Could not read or properly set buffer", 1)
		fmt.Printf("\033[31m%s", hex.Dump(buffer_))
	}
}

// PDF metadata
func PDF_META(filename string) {
	pdff, err := pdf.Open(filename)
	ec.Warning_advanced("<RR6> File Module: Could not open PDF, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	inf := pdff.Trailer().Key("Info").String()
	inf = strings.TrimLeft(inf, "<")
	inf = strings.TrimRight(inf, ">")
	info := strings.Split(inf, "/")
	for _, v := range info {
		println("Data -> | ", v)
	}
}

func Call_perl_s(image string) {
	prg := "perl"
	arg1 := "r6.pl"
	arg00 := "-o"
	arg01 := "opt1"
	arg5 := "-f"
	arg2 := image
	cmd := exec.Command(prg, arg1, arg00, arg01, arg5, arg2)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -<> ", 1)
	fmt.Print(string(stdout))

}

func Call_perl_lfiscan(url string) {
	prg := "perl"
	arg1 := "r6.pl"
	arg00 := "-o"
	arg01 := "lfi"
	arg5 := "-r"
	arg2 := url
	cmd := exec.Command(prg, arg1, arg00, arg01, arg5, arg2)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -<> ", 1)
	fmt.Print(string(stdout))

}

// JPG and JPEG image walking
func WalkFn(path string, f os.FileInfo, err error) error {
	for _, r := range Regex_type_image {
		if r.MatchString(path) {
			fmt.Println("\033[31m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| FOUND JPEG/JPG IMAGE ->  ", path, "\033[32m")
			Call_perl_s(path)
		}
	}
	return nil
}

func Walk_without(path string, file os.FileInfo, err error) error {
	for _, r := range Regex_type_image {
		if r.MatchString(path) {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| FOUND FILEPATH > ", path, "\033[31m")
		}
	}
	return nil
}

func Walker_caller(fiename string) {
	if err := filepath.Walk(fiename, WalkFn); err != nil {
		log.Panicln(err)
	}
}

// basic file stomper
func Stomper(filename string) {
	futureTime := time.Now().Add(50 * time.Hour).Add(15 * time.Minute)
	lastAccessTime := futureTime
	lastModifyTime := futureTime
	err := os.Chtimes(filename, lastAccessTime, lastModifyTime)
	ec.Warning_advanced("<RR6> File Module: Could not change file timestamp, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	fmt.Printf("timestamp change OK - %s", futureTime.String())
}

// basic file premission stomper
func Perm_stomper(mode, filename string) {
	fileModeValue, err := strconv.ParseUint(mode, 8, 32)
	ec.Warning_advanced("<RR6> File Module: Could not change premissions, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	fileMode := os.FileMode(fileModeValue)
	err = os.Chmod(filename, fileMode)
	ec.Warning_advanced("<RR6> File Module: Could not change premissions, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Premissions have been changed to <", mode, "> on file -> "+filename)
}

// telnet
func Perl_telnet(host, wordlist, username string) {
	p := "perl"
	p1 := "r6.pl"
	p00 := "-o"
	arg00 := "brute_telnet"
	p2 := "-f"
	arg := wordlist
	p3 := "-u"
	arg1 := username
	p4 := "-t"
	arg2 := host
	fmt.Println("Executing -> ", p, p1, p00, arg00, p2, arg, p3, arg1, p4, arg2)
	exe := exec.Command(p, p1, p2, arg, p3, arg1, p4, arg2)
	stdout, e := exe.Output()
	ec.Warning_simple("<RR6> Executable Module: Could not execute command to run the needed script got error -> ", v.REDHB, e)
	fmt.Print(string(stdout))
}

// us number tracer
func Run_number(phone_number_part1, phone_number_part2, phone_number_part3 string) {
	parser := "perl"
	parser2 := "modg/scripts/osint/phone/sub-scripts/main.pl"
	parser3 := phone_number_part1
	parser4 := phone_number_part2
	parser5 := phone_number_part3
	parser_main := parser3 + "-" + parser4 + "-" + parser5
	exe := exec.Command(parser, parser2, parser_main)
	stdout, e := exe.Output()
	ec.Warning_simple("<RR6> Executable Module: Could not execute command to run the needed script got error -> ", v.REDHB, e)
	fmt.Print(string(stdout))
}

// file checking to make sure the file exists, if not make one, if so generate
// a new tage and add a new name, then warn user of secondary file generation
// and the name of the file

func Isexisting(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
