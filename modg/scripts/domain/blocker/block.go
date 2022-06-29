// package just simply blocks a list of domains from connecting to you or you connecting to them
package host_blocker

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	colors "main/modg/colors"
	blocker_requests "main/modg/requests"
	blocker_constants "main/modg/scripts/domain/blocker/c"
)

// this will basically reset all data in the file, meaning everything even if you have extra hosts and what before there will be erased and reset to the following template
/*
127.0.0.1 localhost
127.0.0.1 NODE_NAME
::1        localhost ip6-localhost ip6-loopback
ff02::1    ip6-allnodes
ff02::2    ip6-allrouters


NO OTHER DATA WILL BE ADDED, this makes this functio dangerous to use....
*/
func Repair_File() {
	node_host, x := os.Hostname()
	if x != nil {
		panic(x)
	}
	e := os.Remove(blocker_constants.Linux_HOSTS_FILEPATH)
	if e != nil {
		log.Fatal(e)
	} else {
		fmt.Println(colors.WHT, "<RR6> OS MODULE: Was able to remove the filename, prepping to repair...")
	}
	_, x = os.Create(blocker_constants.Linux_HOSTS_FILEPATH)
	if x != nil {
		log.Fatal(colors.REDHB, "<RR6> OS MODULE: Was NOT able to remove the file -> ", x)
	} else {
		fmt.Println(colors.WHT, "<RR6> OS MODULE: Remade filepath, going to repair file....")
	}
	pathmain, x := os.OpenFile(blocker_constants.Linux_HOSTS_FILEPATH, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if x != nil {
		fmt.Println("<RR6> Advanced errors: Could not make a OS filepath opener, got error code (222 [Dev error]), on line 22 of main/modg/scripts/domain/blocker/block.go, in path nest, for file -> ", blocker_constants.Linux_HOSTS_FILEPATH, " Which got error -> ", x)
		os.Exit(0)
	}
	fmt.Fprintln(pathmain, "127.0.0.1 localhost")
	fmt.Println(colors.WHT, "<RR6> OS MODULE: 127.0.0.1 localhost definition repaired and remade....")
	fmt.Fprintln(pathmain, "127.0.0.1 "+node_host)
	fmt.Println(colors.WHT, "<RR6> OS MODULE: 127.0.0.1 node_host definition repaired and remade....")
	open_tmpl, x := os.Open(blocker_constants.TMPL_PATH)
	if x != nil {
		fmt.Println("<RR6> Advanced errors: Could not make a OS filepath opener, got error code (222 [Dev error]), on line 41 of main/modg/scripts/domain/blocker/block.go, in path nest, for file -> ", blocker_constants.Linux_HOSTS_FILEPATH, " Which got error -> ", x)
		os.Exit(0)
	}
	scanner := bufio.NewScanner(open_tmpl)
	for scanner.Scan() {
		fmt.Fprintln(pathmain, scanner.Text())
		fmt.Println(colors.WHT, "<RR6> OS MODULE: Wrote -> ", scanner.Text(), " and finished....")
	}
	open_file, x := os.Open(blocker_constants.Linux_HOSTS_FILEPATH)
	if x != nil {
		fmt.Println("<RR6> Advanced errors: Could not make a OS filepath opener, got error code (222 [Dev error]), on line 41 of main/modg/scripts/domain/blocker/block.go, in path nest, for file -> ", blocker_constants.Linux_HOSTS_FILEPATH, " Which got error -> ", x)
		os.Exit(0)
	}
	scan := bufio.NewScanner(open_file)
	fmt.Println("========================== REPAIRED FILE ==============")
	for scan.Scan() {
		fmt.Println(scan.Text())
	}

}

func Block_Single(hostname string) {
	pathmain, x := os.OpenFile(blocker_constants.Linux_HOSTS_FILEPATH, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if x != nil {
		fmt.Println("<RR6> Advanced errors: Could not make a OS filepath opener, got error code (222 [Dev error]), on line 22 of main/modg/scripts/domain/blocker/block.go, in path nest, for file -> ", blocker_constants.Linux_HOSTS_FILEPATH, " Which got error -> ", x)
		os.Exit(0)
	}
	fmt.Fprintln(pathmain, blocker_constants.Linux_HOSTS_FILEPATH_LO, hostname)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: (IGNORE) | Wrote host to filename, host should NOT be reachable")
	blocker_requests.Make_GET_NOERR(hostname)
}

func Parse_Settings(shost, filename string, fileyn bool) {
	if fileyn {
		if filename == "" {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: File of hosts detected as false....")
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Parsing single host.....")
			Block_Single(shost)
		}
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: File of hosts detected as true....")
		f, x := os.Open(filename)
		if x != nil {
			fmt.Println("<RR6> Simple errors: Could not open filename due to -> ", x)
			os.Exit(1)
		} else {
			defer f.Close()
			scanner := bufio.NewScanner(f)
			pathmain, x := os.OpenFile(blocker_constants.Linux_HOSTS_FILEPATH, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if x != nil {
				fmt.Println("<RR6> Advanced errors: Could not make a OS filepath opener, got error code (222 [Dev error]), on line 22 of main/modg/scripts/domain/blocker/block.go, in path nest, for file -> ", blocker_constants.Linux_HOSTS_FILEPATH, " Which got error -> ", x)
				os.Exit(0)
			}
			scanner.Split(bufio.ScanWords)
			counter := 0
			for scanner.Scan() {
				counter++
				fmt.Fprintln(pathmain, blocker_constants.Linux_HOSTS_FILEPATH_LO, scanner.Text())
				fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: (IGNORE) | Wrote host to filename, host should NOT be reachable")
			}
			dial_attemptin(filename)
		}
	}
	if !fileyn {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: File of hosts detected as false....")
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Parsing single host.....")
		Block_Single(shost)
	}
}

func dial_attemptin(host_config_path string) {
	file := host_config_path
	blacklist, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	//
	//
	defer blacklist.Close()
	//
	scanner := bufio.NewScanner(blacklist)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		dialmethod := scanner.Text() + ":80"
		conn, err := net.Dial("tcp", dialmethod)
		if err != nil {
			fmt.Println("\n[ * ] DATA: HOST -> ", dialmethod, "BLOCKED")
		} else {
			fmt.Println("\n[ ? ] Connection was established? was host down or blocked?")
			fmt.Println("\n[ * ] STAT: Trying TCP dialup again -> ", conn)
			dialmethod := scanner.Text() + ":80"
			conn, err := net.Dial("tcp", dialmethod)
			if err != nil {
				fmt.Println("\n[ * ] DATA: HOST -> ", dialmethod, "BLOCKED")
			} else {
				fmt.Println("\n[ ? ] Connection was established? was host down or blocked?")
				fmt.Println("\n[ * ] STAT: Trying TCP dialup LAST TIME -> ", conn)
				dialmethod := scanner.Text() + ":80"
				conn, err := net.Dial("tcp", dialmethod)
				if err != nil {
					fmt.Println("\n[ * ] DATA: HOST -> ", dialmethod, "BLOCKED")
				} else {
					fmt.Println("\n[ ? ] Connection was established? was host down or blocked? -> ", conn)
				}
			}
		}
	}
}
