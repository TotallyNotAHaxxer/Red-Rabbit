package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/ArkAngeL43/port-scanning/port"
)

var (
	foportl    = flag.Bool("portl", false, "| port scan with a list")
	foportlist = flag.String("list", "", "| port scan with a list")
	foporthon  = flag.Bool("pscanh", false, "true or false for base single host scanning")
	foporthost = flag.String("host", "", "hostname after using flag -foport")
	//
	BLU  = "\033[0;94m"
	chex = "\x1b[H\x1b[2J\x1b[3J"
	BLK  = "\033[0;30m"
	RED  = "\033[0;31m"
	GRN  = "\033[0;32m"
	YEL  = "\033[0;33m"
	MAG  = "\033[0;35m"
	CYN  = "\033[0;36m"
	WHT  = "\033[0;37m"
	BBLK = "\033[1;30m"
	BRED = "\033[1;31m"
	BGRN = "\033[1;32m"
	BYEL = "\033[1;33m"
	BBLU = "\033[1;34m"
	BMAG = "\033[1;35m"
	BCYN = "\033[1;36m"
	BWHT = "\033[1;37m"
	//Regular underline
	UBLK = "\033[4;30m"
	URED = "\033[4;31m"
	UGRN = "\033[4;32m"
	UYEL = "\033[4;33m"
	UBLU = "\033[4;34m"
	UMAG = "\033[4;35m"
	UCYN = "\033[4;36m"
	UWHT = "\033[4;37m"
	//Regular back
	BLKB = "\033[40m"
	REDB = "\033[41m"
	GRNB = "\033[42m"
	YELB = "\033[43m"
	BLUB = "\033[44m"
	MAGB = "\033[45m"
	CYNB = "\033[46m"
	WHTB = "\033[47m"
)

func ce(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handelreturncon(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nDetected Interupt.....")
			os.Exit(1)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func port_scanner_list(list string) {
	listed, err := ioutil.ReadFile(list)
	ce(err)
	fmt.Println("____________Hosts in file_________________")
	fmt.Println("[ 1 ] File -> ", list)
	fmt.Println("[ 2 ] Are these hosts correct? \n", "")
	fmt.Println(string(listed))
	fmt.Print("n=No | Y=yes|")
	input := bufio.NewReader(os.Stdin)
	fmt.Print(BLU, "\n> ")
	for {
		go handelreturncon(make(chan os.Signal, 1))
		t, _ := input.ReadString('\n')
		t = strings.Replace(t, "\n", "", -1)
		if strings.Compare("yes", t) == 0 {
			fmt.Println(chex)
			fill, err := os.Open(list)
			ce(err)
			defer fill.Close()
			watchdog := bufio.NewScanner(fill)
			for watchdog.Scan() {
				fmt.Println("\033[0;94m[", REDB, watchdog.Text(), "]\033[0;49m")
				port.GetOpenPorts(watchdog.Text(), port.PortRange{Start: 1, End: 8090})

			}
		}
		if strings.Compare("Y", t) == 0 {
			fmt.Println(chex)
			fill, err := os.Open(list)
			ce(err)
			defer fill.Close()
			watchdog := bufio.NewScanner(fill)
			for watchdog.Scan() {
				fmt.Println("\033[0;94m[", REDB, watchdog.Text(), "]\033[0;49m")
				port.GetOpenPorts(watchdog.Text(), port.PortRange{Start: 1, End: 8090})

			}
		}
		if strings.Compare("YES", t) == 0 {
			fmt.Println(chex)
			fill, err := os.Open(list)
			ce(err)
			defer fill.Close()
			watchdog := bufio.NewScanner(fill)
			for watchdog.Scan() {
				fmt.Println("\033[0;94m[", REDB, watchdog.Text(), "]\033[0;49m")
				port.GetOpenPorts(watchdog.Text(), port.PortRange{Start: 1, End: 8090})

			}
		}
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	if *foportl {
		port_scanner_list(*foportlist)
	}
	if *foporthon {
		port.GetOpenPorts(*foporthost, port.PortRange{Start: 1, End: 8090})
	}
}
