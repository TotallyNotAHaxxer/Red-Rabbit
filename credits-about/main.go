// at this point RR5 is written in ruby and powered by go XD
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/theckman/yacspin"
)

func main() {
	powered_by()
	port := ":8080"
	clear(clear_hex)
	banner()
	http.HandleFunc("/", process)
	is_online("https://www.google.com")
	fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mINITIATING SERVER")
	fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mSERVER URL https://localhost", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

const (
	v           = "1.0 BETA"
	DB_USER     = "user1"
	DB_PASSWORD = "fancy-bear-2021-fuckg"
	DB_NAME     = "testdb"
	host        = "localhost"
	port        = 5432
	user        = "user1"
	password    = "fancy-bear-2021-fuckg"
	dbname      = "testdb"
)

var (
	clear_hex = "\x1b[H\x1b[2J\x1b[3J"
	BLK       = "\033[0;30m"
	RED       = "\033[0;31m"
	GRN       = "\033[0;32m"
	YEL       = "\033[0;33m"
	BLU       = "\033[0;34m"
	MAG       = "\033[0;35m"
	CYN       = "\033[0;36m"
	WHT       = "\033[0;37m"
	BBLK      = "\033[1;30m"
	BRED      = "\033[1;31m"
	BGRN      = "\033[1;32m"
	BYEL      = "\033[1;33m"
	BBLU      = "\033[1;34m"
	BMAG      = "\033[1;35m"
	BCYN      = "\033[1;36m"
	BWHT      = "\033[1;37m"
	UBLK      = "\033[4;30m"
	URED      = "\033[4;31m"
	UGRN      = "\033[4;32m"
	UYEL      = "\033[4;33m"
	UBLU      = "\033[4;34m"
	UMAG      = "\033[4;35m"
	UCYN      = "\033[4;36m"
	UWHT      = "\033[4;37m"
	BLKB      = "\033[40m"
	REDB      = "\033[41m"
	GRNB      = "\033[42m"
	YELB      = "\033[43m"
	BLUB      = "\033[44m"
	MAGB      = "\033[45m"
	CYNB      = "\033[46m"
	WHTB      = "\033[47m"
	BLKHB     = "\033[0;100m"
	REDHB     = "\033[0;101m"
	GRNHB     = "\033[0;102m"
	YELHB     = "\033[0;103m"
	BLUHB     = "\033[0;104m"
	MAGHB     = "\033[0;105m"
	CYNHB     = "\033[0;106m"
	WHTHB     = "\033[0;107m"
	HBLK      = "\033[0;90m"
	HRED      = "\033[0;91m"
	HGRN      = "\033[0;92m"
	HYEL      = "\033[0;93m"
	HBLU      = "\033[0;94m"
	HMAG      = "\033[0;95m"
	HCYN      = "\033[0;96m"
	HWHT      = "\033[0;97m"
	BHBLK     = "\033[1;90m"
	BHRED     = "\033[1;91m"
	BHGRN     = "\033[1;92m"
	BHYEL     = "\033[1;93m"
	BHBLU     = "\033[1;94m"
	BHMAG     = "\033[1;95m"
	BHCYN     = "\033[1;96m"
	BHWHT     = "\033[1;97m"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func clear(clear_hex string) {
	fmt.Println(clear_hex)
}

func print() {
	fmt.Println("f")
}

func fileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func file_checker(filepath string) {
	if fileExists(filepath) {
		fmt.Println(BLU, "[", REDB, "Online\033[1;39m", BLU, "]", "\t\t\t[ ", REDB, "File "+filepath+" loaded\033[1;39m", BLU, "]")
	} else {
	}
}

func arr_files() {
	arr := [31]string{
		"log/logger.txt",
		"seperate/history.txt",
		"seperate/mentions.txt",
		"seperate/langs.txt",
		"seperate/ASCII-ART.txt",
		"seperate/comand1.txt",
		"seperate/install.txt",
		"seperate/script_help.txt",
		"seperate/ajax",
		"seperate/whois.txt",
		"seperate/sql.txt",
		"seperate/xss.txt",
		"seperate/port_all.txt",
		"seperate/dig.txt",
		"seperate/sshinject.txt",
		"seperate/interfaces.txt",
		"seperate/ftpmod.txt",
		"seperate/arp.txt",
		"seperate/port-lg.txt",
		"seperate/ftptcp=all.txt",
		"seperate/loopingdns.txt",
		"seperate/BSSIDs.txt",
		"seperate/fakeap.txt",
	}
	fmt.Println("----- Checking files -----")
	for i := 0; i < len(arr); i++ {
		file_checker(arr[i])
	}
	fmt.Print("\n\n\n")
}

func banner() {
	ascii :=
		`
	
______________                   _____            ________                                      
___  __ \__  /______________________(_)___  __    __  ___/______________   _____________________
__  /_/ /_  __ \  _ \  __ \_  __ \_  /__  |/_/    _____ \_  _ \_  ___/_ | / /  _ \_  ___/_  ___/
_  ____/_  / / /  __/ /_/ /  / / /  / __>  <      ____/ //  __/  /   __ |/ //  __/  /   _(__  ) 
/_/     /_/ /_/\___/\____//_/ /_//_/  /_/|_|______/____/ \___//_/    _____/ \___//_/    /____/  
                                           _/_____/                                             
            See command pheonix servers to learn about this service/server
------------------------------------------------------------------------------------------------
	`
	fmt.Println(ascii)
	arr_files()
}

func powered_by() {
	cfg := yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[59],
		Suffix:          "Loading => Red Rabbit V5 WEB-UI",
		SuffixAutoColon: true,
		Message:         "",
		StopCharacter:   "✓",
		StopColors:      []string{"fgGreen"},
	}
	spinner, err := yacspin.New(cfg)
	err = spinner.Start()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	spinner.Message("\nPowerd by GOlang->Pheonix_Databases")
	time.Sleep(2 * time.Second)
	err = spinner.Stop()
}

func is_online(website string) bool {
	get, err := http.Get(website)
	if err != nil {
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT ERROR  |=> ", get.StatusCode)
		return false
	} else {
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT PASSED |=> ", get.StatusCode)
		return true
	}
}

func out_msg(w http.ResponseWriter, filename string) {
	con, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, string(con))
	}
}

func logtxt(file string) {
	filepath := file
	pathmain, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	checkErr(err)
	defer pathmain.Close()
	c, err := fmt.Fprintln(pathmain, "[ INFO ] -> USER LOGGED IN AT => ", time.Now())
	if err != nil {
		fmt.Print(c)
		os.Exit(1)
	}
	fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mACCESS WRITTEN TO -> ", filepath)
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT  |=> ", http.StateActive)
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT  |=> ", http.StatusOK)
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		command := r.FormValue("command")
		apassword := r.FormValue("Acting Password")
		fmt.Fprintln(w, command, apassword)
		if command == "lo-get" {
			resp, err := http.Get("http://localhost:8080")
			checkErr(err)
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET DATA: |=> ", resp.Header.Get("Accept-Ranges"))
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET DATA: |=> ", resp.Header.Get("Content-Length"))
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET DATA: |=> ", resp.Header.Get("Content-Type"))
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET DATA: |=> ", resp.Header.Get("Last-Modified"))
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET DATA: |=> ", resp.Header.Get("Date"))
			fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Accept-Ranges"))
			fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Length"))
			fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Type"))
			fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Last-Modified"))
			fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Date"))
			logtxt("log/logger.txt")
		}
		if apassword == "RR5-admin" {
			logtxt("log/logger.txt")
			fmt.Print("\n")
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mAUTH: PASSWORD    |-> ", apassword, " = TRUE ")
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA: DB-PASSWORD |-> ", password)
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA: COMMAND-RUN |-> ", command)

		} else {
			fmt.Println(WHT, "\t\t[", BRED, "INFO", WHT, "] \033[34mAUTH: PASSWORD -> ", apassword, " ENTERED CAME BACK NEGATIVE, FAILED PASSWORD AUTH ")
			os.Exit(1)
		}
		if command == "chconn" {
			resp, err := http.Get("https://google.com")
			checkErr(err)
			if resp.StatusCode >= 100 {
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Accept-Ranges"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Length"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Type"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Last-Modified"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Date"))
				logtxt("log/logger.txt")
			} else {
				fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA=>GET  code not parsed")
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Accept-Ranges"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Length"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Type"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Last-Modified"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Date"))
				logtxt("log/logger.txt")
			}
		}
		if command == "help" {
			content, err := ioutil.ReadFile("log/help.txt")
			checkErr(err)
			fmt.Fprintln(w, string(content))
			logtxt("log/logger.txt")
			checkErr(err)
			outting, err2 := ioutil.ReadFile("log/help.txt")
			checkErr(err2)
			fmt.Fprintln(w, string(outting))
		}
		if command == "log" {
			content, err := ioutil.ReadFile("log/logger.txt")
			checkErr(err)
			fmt.Fprintln(w, string(content))
			logtxt("log/logger.txt")
			c := "Rscript"
			c1 := "main.r"
			c2 := "log/logger.txt"
			cf := exec.Command(c, c1, c2)
			stdout, err := cf.Output()
			checkErr(err)
			fmt.Fprintln(w, string(stdout))
		}
		// or or statement to keep down the case
		if command == "ch1 History" || command == "ch1 history" {
			out_msg(w, "seperate/history.txt")
		}
		if command == "ch1 Honorable mentions" || command == "honorable mentions" || command == "mentions" || command == "contribution" {
			out_msg(w, "seperate/mentions.txt")
		}
		if command == "ch1 Languages" || command == "languages" || command == "langs" || command == "langs used" {
			out_msg(w, "seperate/langs.txt")
		}
		if command == "ch1 art" || command == "art" || command == "banner" {
			out_msg(w, "seperate/ASCII-ART.txt")
		}
		if command == "ch1 commands" || command == "commands" || command == "command" {
			out_msg(w, "seperate/comand1.txt")
		}
		if command == "ch1 install" || command == "ch1 installs" || command == " installs" || command == "ch1 INSTALL" {
			out_msg(w, "seperate/install.txt")
		}
		if command == "ch2 Script_help" || command == "script_help" || command == "script help" || command == "script" {
			out_msg(w, "seperate/script_help.txt")
		}
		if command == "ch3 Ajax_spiders" || command == "Ajax_spiders" || command == "ch3 ajax" || command == "ch3 ajax_spiders" {
			out_msg(w, "seperate/ajax")
		}
		if command == "ch4 whois" || command == "whois" || command == "ch4 WHOIS" {
			out_msg(w, "seperate/whois.txt")
		}
		if command == "ch4 sql" || command == "ch4 sql-t" || command == "ch4 sqli-t" {
			out_msg(w, "seperate/sql.txt")
		}
		if command == "ch4 xss" || command == "ch4 xss-t" || command == "ch4 xss-i" || command == "ch4 xss" {
			out_msg(w, "seperate/xss.txt")
		}
		if command == "ch4 port-r" || command == "ch4 port-a" || command == "ch4 porta" || command == "ch4 portall" {
			out_msg(w, "seperate/port_all.txt")
		}
		if command == "ch4 dg" || command == "ch4 d-g" || command == "ch4 dg " {
			out_msg(w, "seperate/dig.txt")
		}
		if command == "ch5 SSH injection" || command == "ch5 ssh injection" || command == "ssh inject" {
			out_msg(w, "seperate/sshinject.txt")
		}
		if command == "ch5 FTP brute" || command == "ch5 brute forcing FTP" || command == "ch5 FTP brute" || command == "FTP brute" {
			out_msg(w, "seperate/ftpmod.txt")
		}
		if command == "ch5 interfaces" || command == "ch5 interface" || command == "ch5 inter" {
			out_msg(w, "seperate/interfaces.txt")
		}
		if command == "ch6 ARP" || command == "ch6 arp" || command == "ch6 Arp" {
			out_msg(w, "sepertae/arp.txt")
		}
		if command == "ch6 portlg" || command == "ch6 port-lg" || command == "ch6 port_listed" {
			out_msg(w, "seperate/port-lg.txt")
		}
		if command == "ch6 tcp/ftp=all" || command == "tcp/ftp=all" {
			out_msg(w, "seperate/ftptcp=all.txt")
		}
		if command == "ch7  DNS-Loop" || command == "ch7 dns-loop" {
			out_msg(w, "seperate/loopingdns.txt")
		}
		if command == "ch7 BSSID_scout" || command == "BSSID_scout" {
			out_msg(w, "seperate/BSSIDs.txt")
		}
		if command == "ch7 Fake_ap" || command == "ch7 Fake_AP" {
			out_msg(w, "seperate/fakeap.txt")
		}
		if command == "ch7 deauth" || command == "deauth" || command == "deauthentication" {
			out_msg(w, "seperate/deauth.txt")
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
