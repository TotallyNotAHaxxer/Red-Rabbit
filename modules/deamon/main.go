// interactive deamon mode for RR5 SYSTEM STATS MAINLY SEE DEAMON_RR5 DIR IN MODULES DIR FOR
// CODE AND GO MAIN TO GENERATE HTML BASED INPUTS
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
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"

	"github.com/sirupsen/logrus"
	"github.com/theckman/yacspin"
)

type ServerVals struct {
	Server1   string
	ServerRR5 string
}

var (
	//go:embed templates/*.tmpl
	rootFs embed.FS
)

var (
	err    error
	result string
)

// type set structure for system information
type SysInfo struct {
	Hostname string `bson:hostname`
	Platform string `bson:platform`
	CPU      string `bson:cpu`
	RAM      uint64 `bson:ram`
	Disk     uint64 `bson:disk`
}

// application values
type appValues struct {
	AppName  string
	YourName string
}

func ce(err error, msg string, msagetype string, exitcode int) bool {
	if msagetype == "log" {
		if err != nil {
			log.Fatal(err, msg)
			os.Exit(exitcode)
			return true
		}
	}
	if msagetype == "print" {
		if err != nil {
			fmt.Println(err, msg)
			os.Exit(exitcode)
		}
	}
	if msagetype == "panic" {
		if err != nil {
			panic(err)
		}
	}
	return false
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

const (
	port           = ":8080"
	hostname       = "http://localhost"
	host1          = hostname + port
	staticpassword = "RR5_admin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func clear(clear_hex string) {
	fmt.Println(clear_hex)
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

func execute(command1, command2, file string) {
	prg := "go"
	prg2 := "run"
	arg1 := file
	cmd := exec.Command(prg, prg2, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))
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
		if apassword == "RR5_admin" || apassword == "RR5_Admin" {
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] AUTHENTICATION PASSED")
		} else {
			fmt.Println("WARN: PASSWORD INCORRECT")
			os.Exit(0)
		}
		if command == "system" {
			fmt.Println(" Gathering data ")
			// code generation caller
			main_1()

		}
		fmt.Fprintln(w, command, apassword)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func process_system_html(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "RR5_host/index.html")
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT  |=> ", http.StateActive)
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT  |=> ", http.StatusOK)
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		command := r.FormValue("command")
		apassword := r.FormValue("Acting Password")
		if apassword == "RR5-Admin" {
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] AUTHENTICATION PASSED")
		} else {
			fmt.Println("WARN: PASSWORD INCORRECT")
		}
		if command == "system" {
			fmt.Println(" Gathering data ")
		}
		fmt.Fprintln(w, command, apassword)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	powered_by()
	clear(clear_hex)
	banner()
	http.HandleFunc("/", process)
	is_online("https://www.google.com")
	fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mINITIATING SERVER")
	fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mSERVER URL ", host1)
	fmt.Println(WHT, "\t\t[", GRN, "INFO", WHT, "] \033[34mSERVER PASSWORD => ", staticpassword)

	log.Fatal(http.ListenAndServe(port, nil))
}

//////////////////// code generation //////////////////////

// function main_1
// function do: Code generate system information and stats to load prev
// function returns: nothing
func main_1() {

	var (
		err       error
		fp        *os.File
		templates *template.Template
		subdirs   []string
	)

	values := SysInfo{}
	value2name := appValues{}
	info := new(SysInfo)
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024
	info.Hostname = hostStat.Hostname
	value2name.AppName = "RR5_host"
	// setting variables for code generatioon and static files
	values.Hostname = info.Hostname
	values.CPU = info.CPU
	values.Platform = info.Platform
	values.RAM = info.RAM
	values.Disk = info.Disk

	rootFsMapping := map[string]string{
		"index.html.tmpl": "static/index.html",
		"main.go.tmpl":    "main.go",
	}

	if err = os.Mkdir(value2name.AppName, 0755); err != nil {
		logrus.WithError(err).Errorf("error attempting to create application directory '%s'", values.Hostname)
	}

	if err = os.Chdir(value2name.AppName); err != nil {
		logrus.WithError(err).Errorf("error changing to new directory '%s'", values.Hostname)
	}

	subdirs = []string{
		"static",
	}

	for _, dirname := range subdirs {
		if err = os.MkdirAll(dirname, 0755); err != nil {
			logrus.WithError(err).Fatalf("unable to create subdirectory %s", dirname)
		}
	}

	if templates, err = template.ParseFS(rootFs, "templates/*.tmpl"); err != nil {
		logrus.WithError(err).Fatal("error parsing root templates files")
	}

	for templateName, outputPath := range rootFsMapping {
		if fp, err = os.Create(outputPath); err != nil {
			logrus.WithError(err).Fatalf("unable to create file %s for writing", outputPath)
		}

		defer fp.Close()

		if err = templates.ExecuteTemplate(fp, templateName, values); err != nil {
			logrus.WithError(err).Fatalf("unable to exeucte template %s", templateName)
		}
	}
	execute("go", "run", "RR5_host/main.go")
	fmt.Println("command executed")
	fmt.Println("Please CTRL+c and execute the following command")
	fmt.Println("cd RR5_host ; sudo go run main.go ")
}

func stringPrompt(label, defaultValue string) string {
	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}
	if result, err = prompt.Run(); err != nil {
		logrus.WithError(err).Fatalf("error asking for '%s'", label)
	}
	return result
}
