// interactive deamon mode for RR5 port scanning
// packet capture, and some other things in grave detail
// as of January 3rd Red Rabbit V5 Interactive deamon mode is experimental
// and will not exceed everyones expectations or even come close to meeting them
//
// RR5 is a massive project so far and has been an amazing road, to improove it
// i figured why not make the interactive web mode, give people a little bit more
// of a prespective to how truly deep this project has come in the past 4-7 months
// and how far it will go in the further future
//
// the developers of this project are  => ArkAngeL43
// github                              => https://github.com/ArkAngeL43
// Discord Server for RR5              => https://discord.gg/JTDFMXuTkZ
// Server type                         => PostGreSQL
// Server ran by                       => Pheonix services
// encryption for pass                 => SHA256 reversed
// Languages used here                 => Inline C/Assembly, Go, Ruby, HTML/CSS/JS/Inline GO
// Type actions for server             => POST, GET
// Type actions NOW ALLOWED for server => DELETE, MOVE, SOAP, etc
// Type form result                    => POST
// gen key type                        => Rune
// String type for password            => String, int
// Allowed Post Data commands          => Help, SysINF, commands, command, exit, data, node name, net etc MORE INFO IN DOCUMENTATION
// Date of stat                        => Monday January 1st
// Type color string                   => ANSI Hexadecimal colors
// Type structures                     => SystemInformation
// Type server structure for CLI GET   => Tor
// Definitions                         => Main, inf, net, root, rootc, c, caller, etc
// Most used hex                       => Clear hex => \x1b[H\x1b[2J\x1b[3J
// Third Party packages used           => Prompt UI
// Filetype of the server              => /
// File launch index/main index        => index.html
// Seperators                          => \t, \n, \a
// Port of WEB Server                  => 8080/80/random if neaither work
// Port of SQL Server                  => 5432
//
// Aboout
//
// this server again is for the interactive deamon mode for RR5, the way this works
// is a bit awkward as it requires some user navigation throughout the entirety of the
// UI, the first panel you are brought to will be labeled `command type` this will differ
// weather the UI is targeting a URL, System, IP, localhost, or domain server etc
// once inputted the server will terminate and a new URL server will be launched
// a code file will be generated and finally spit out and ran by go, however this will
// be very very weird
// there will be more in the documentation of RR5
//
// if you want to know more about a designated WEB UI for RR5 offensive, then you will
// launch the web interface for documentation with the command `doc-dem` which will launch
// a local web UI on http://localhost:8080 where you will be directed to a similiar input box
// in the input box as a command input the following `DOC LOCAL CMD=RR5Local` this will show
// every possible command and detail on the RR5 web UI and how it works in multiple ways
// and the easier way to manuver, this backend has ALOT to do in general and even once done is still
// in BETA. it will remain that way until i decide to maybe make a version 6 of RR5
//
//
// main package decleration of main.go starts here
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/theckman/yacspin"
)

// server constants
const (
	port          = "8080"
	host          = "8080"
	local         = "localhost:"
	hosted_server = host + local
	version       = ".01 BETA"
	ch            = "\x1b[H\x1b[2J\x1b[3J" // hexidecimal value to clear a terminal screen
	statpass      = "RR5_UI"               // different from system UI
	docfilepath   = "RR5/credits-about/seperate/web-ui-RR5.txt"
)

// Variables and colors
var (
	BLK   = "\033[0;30m"
	RED   = "\033[0;31m"
	GRN   = "\033[0;32m"
	YEL   = "\033[0;33m"
	BLU   = "\033[0;34m"
	MAG   = "\033[0;35m"
	CYN   = "\033[0;36m"
	WHT   = "\033[0;37m"
	BBLK  = "\033[1;30m"
	BRED  = "\033[1;31m"
	BGRN  = "\033[1;32m"
	BYEL  = "\033[1;33m"
	BBLU  = "\033[1;34m"
	BMAG  = "\033[1;35m"
	BCYN  = "\033[1;36m"
	BWHT  = "\033[1;37m"
	UBLK  = "\033[4;30m"
	URED  = "\033[4;31m"
	UGRN  = "\033[4;32m"
	UYEL  = "\033[4;33m"
	UBLU  = "\033[4;34m"
	UMAG  = "\033[4;35m"
	UCYN  = "\033[4;36m"
	UWHT  = "\033[4;37m"
	BLKB  = "\033[40m"
	REDB  = "\033[41m"
	GRNB  = "\033[42m"
	YELB  = "\033[43m"
	BLUB  = "\033[44m"
	MAGB  = "\033[45m"
	CYNB  = "\033[46m"
	WHTB  = "\033[47m"
	BLKHB = "\033[0;100m"
	REDHB = "\033[0;101m"
	GRNHB = "\033[0;102m"
	YELHB = "\033[0;103m"
	BLUHB = "\033[0;104m"
	MAGHB = "\033[0;105m"
	CYNHB = "\033[0;106m"
	WHTHB = "\033[0;107m"
	HBLK  = "\033[0;90m"
	HRED  = "\033[0;91m"
	HGRN  = "\033[0;92m"
	HYEL  = "\033[0;93m"
	HBLU  = "\033[0;94m"
	HMAG  = "\033[0;95m"
	HCYN  = "\033[0;96m"
	HWHT  = "\033[0;97m"
	BHBLK = "\033[1;90m"
	BHRED = "\033[1;91m"
	BHGRN = "\033[1;92m"
	BHYEL = "\033[1;93m"
	BHBLU = "\033[1;94m"
	BHMAG = "\033[1;95m"
	BHCYN = "\033[1;96m"
	BHWHT = "\033[1;97m"
)

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
}

// func check errors, this will check errors with customizeable typers, errors, and messgaes
func che(err error, msg string, typer int) bool {
	// func to log err
	if typer == 1 {
		if err != nil {
			log.Fatal(err, msg)
			return true
		}
	}
	// func to panic error
	if typer == 2 {
		if err != nil {
			fmt.Print(msg)
			panic(err)
		}
	}
	// func to continue despite error
	if typer == 3 {
		fmt.Println(err, msg)
		fmt.Println("CONTINUING DESPITE ERROR TYPER <= 1 ")
	}
	return false
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static/main.html")
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT  |=> ", http.StateActive)
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT  |=> ", http.StatusOK)
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		command := r.FormValue("command")
		apassword := r.FormValue("Acting Password")
		if apassword == statpass {
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] AUTHENTICATION PASSED")
		} else {
			fmt.Println("WARN: PASSWORD INCORRECT")
			os.Exit(0)
		}
		//////////////////////////// SERVER PROCESSING WEATHER TARGET IS A URL, IP, or OTHER /////////////////////////////////////
		//
		// This function starts by checking via regex wether the inputted host is a IPV4, IPV6, MAC, or Domin name
		//
		// Depending on the type of string I.E Host, it will call functions with the hostname as string, ex whois.
		//
		// Once done the function will grab the output of the function and call the code generation and HTML result
		//
		// parser then leading the user to start a new server via sudo on port 80 with their host results organized into certian
		//
		// divs and positions in boxes, idea is to test the HTML before using it as a .TMPL file for template generatio
		//
		// If host turns out to be a BSSID ( as said and stated in documentation ) IT WILL NOT RELAY INFO OR GENERATE A FILE
		//
		// as of current version .01 BETA of the web UI there is no support for BSSID information and tagging
		//
		//
		// Function testing gets called here
		fmt.Fprintln(w, command, apassword)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
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

func clear(clear_hex string) {
	fmt.Println(clear_hex)
}

func main() {
	powered_by()
	clear(ch)
	go banner()
	go http.HandleFunc("/", process)
	is_online("https://www.google.com")
	fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mINITIATING SERVER")
	fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mSERVER URL ", hosted_server)
	fmt.Println(WHT, "\t\t[", GRN, "INFO", WHT, "] \033[34mSERVER PASSWORD => ", statpass)

	log.Fatal(http.ListenAndServe(port, nil))
}
