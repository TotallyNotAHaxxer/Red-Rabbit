/*
Developer | ArkAngeL43
Package   | warn
Module    | warn.go
File      | modg/warnings
Nest      | warnings

Does:
	Provides error module

Could be modified alot more, sticking to what i have right now

*/

package warnings

import (
	"bufio"
	"fmt"
	"log"
	v "main/modg/colors"
	"os"
)

func Warning_simple(msg, color string, err error) bool {
	if err != nil {
		fmt.Println(color, msg, err, "\033[39m\033[49m")
		return true
	}
	return false
}

func Che(err error, msg string, exit int) {
	if err != nil {
		fmt.Println(v.RED, "[!] Error: Fatal: ", msg, err, "\033[39m\033[49m")
	}
}

/*
Requrires:
	A message
	A color
	A message format : 1 for println, 2 for print, 3 for log println
	true or false value for if set a file
	true or false value if tests are needed to test paramaters
	true or false value for if its a message or file
	err
	exit code
	dev code
	filename, leave blank if value for set file is false

	normal message:
		ec.Warning_advanced("message", colorcall, 1, false, false, true, err, 1, 233, "")

	file message;
		ec.Warning_advanced("message", colorcall, 2, true, false, false, err, 1, 233, "path/to/error/file/file.txt")



*/
// add return error function

func Warning_advanced(msg, color string, msg_format int, file, extra, rr6 bool, err error, exit_code int, dev_code int, filename string) bool {
	if err != nil {
		if rr6 {
			if msg_format == 1 {
				fmt.Println(color, "<RR6> ERRORS MODULE: Got dev code < ", dev_code, "> | -> ", msg, err, "\033[39m\033[49m")
			}
			if msg_format == 2 {
				fmt.Print(color, "<RR6> ERRORS MODULE: Got dev code < ", dev_code, "> | -> ", msg, err, "\033[39m\033[49m")
			}
			if msg_format == 3 {
				log.Println(color, "<RR6> ERRORS MODULE: Got dev code < ", dev_code, "> | -> ", msg, err, "\033[39m\033[49m")
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
					fmt.Print(color, "| -> ", msg, err)
				}
				if msg_format == 3 {
					log.Println(color, "| -> ", msg, err)
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

// Ce(err, colorname, "message", code)
func Ce(err error, color, msg string, exit_code int) {
	if err != nil {
		fmt.Println(color, msg, err, "\033[39m\033[49m")
		os.Exit(0)
	}
}

//func Warning_a1dvanced(msg, color string, msg_format int, file, extra, rr6 bool, err error, exit_code int, dev_code int, filename string) bool {if err != nil {if rr6 {if msg_format == 1 {fmt.Println(color, "<RR6> ERRORS MODULE: Got dev code < ", dev_code, "> | -> ", msg)};if msg_format == 2 {fmt.Print(color, "<RR6> ERRORS MODULE: Got dev code < ", dev_code, "> | -> ", msg)}if msg_format == 3 {log.Println(color, "<RR6> ERRORS MODULE: Got dev code < ", dev_code, "> | -> ", msg)} else {if file {if msg_format == 1 {if extra {if file {fmt.Println(true)} else {fmt.Println(false)}}if msg_format == 2 {fmt.Print(color, "| -> ", msg)}fmt.Println("-------------------------------------- ERROR LOG ")fmt.Println(color, msg)fmt.Println("====Opening error log file...")content, err := os.Open(filename)if err != nil {fmt.Println(color, "<RR6> LOG MODULE: Could not open the warnings filename, note this output message should NOT appear, THIS IS NOT A WARNING THIS IS A FATAL ERROR, FILE might not exist, or have good premissions or enough premissions to open, read, and output, please make sure you are selecting the correct file or running as a sudo user")os.Exit(exit_code)} else {scanner := bufio.NewScanner(content)for scanner.Scan() {fmt.Println(color, content)}}}}return true}return false}
