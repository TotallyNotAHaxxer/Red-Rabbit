package super_XML

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	structure "main/modules/go-main/xml/types"
)

// \033[37m
var NMAP structure.Auto_Gen_Nmaprun
var NMAP2 structure.Auto_Gen_Nmaprun_2

var list = []string{
	"host",
	"ports",
	"hosts",
	"service info",
	"port info",
	"Runstats",
	"Host Hint",
	"debug",
	"verbose",
	"scan info",
}

func Main(filename string, expression string) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &NMAP)
	if expression == "*" {
		for _, values := range list {
			Parser(filename, values, NMAP)

		}
	} else {
		Parser(filename, expression, NMAP)
	}
}

/*
func Parser2(filename, expression string, NMAP2 structure.Auto_Gen_Nmaprun_2) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &NMAP2)
	switch expression {
	case "os class":
		fmt.Printf(NMAP2.Taskprogress[])
	}
}
*/

func Parser(filename, choices string, NMAP structure.Auto_Gen_Nmaprun) {
	switch choices {
	case "host":
		fmt.Println("\033[37mAddresses scanned ~> ", NMAP.Host.Address.Addr)
	case "ports":
		for k := 0; k < len(NMAP.Host.Ports.Port); k++ {
			fmt.Printf("\033[37m< Port > | %v | < State >  | %s | < Service > | %s | \n", NMAP.Host.Ports.Extraports.Extrareasons.Ports[k], NMAP.Host.Ports.Extraports.State, NMAP.Host.Ports.Port[k].Service.Name)
		}
	case "service info":
		for k := 0; k < len(NMAP.Host.Ports.Port); k++ {
			fmt.Printf("\033[38;5;198m< Port > \033[38;5;21m| %v |\n", NMAP.Host.Ports.Extraports.Extrareasons.Ports[k])
			fmt.Printf("     \033[38;5;198m< State > \033[38;5;21m| %s | \n", NMAP.Host.Ports.Extraports.State)
			fmt.Printf("           \033[38;5;198m< Service information of port > \n")
			fmt.Printf("               \033[38;5;21mService | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Name)
			fmt.Printf("               \033[38;5;21mOS_type | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Ostype)
			fmt.Printf("               \033[38;5;21mProduct | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Product)
			fmt.Printf("               \033[38;5;21mVersion | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Version)
			fmt.Printf("               \033[38;5;21mCPE     | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Cpe)
			fmt.Printf("               \033[38;5;21mMethod  | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Method)
		}
	case "hostnames":
		for k := 0; k < len(NMAP.Host.Hostnames.Hostname); k++ {
			fmt.Printf("\033[38;5;198m< Host  > | %s | \n", NMAP.Host.Hostnames.Hostname[k].Name)
			fmt.Printf("\033[37m          |\n")
			fmt.Printf("\033[37m          |----<Host Type>: | %s | \n", NMAP.Host.Hostnames.Hostname[k].Type)
			fmt.Printf("\033[37m                            |    |\n")
			fmt.Printf("\033[37m          |----<Host Text>: | %s | \n", NMAP.Host.Hostnames.Hostname[k].Text)
			fmt.Printf("----------------------------------------------------------------------\n")
		}
	case "port info":
		for k := 0; k < len(NMAP.Host.Ports.Port); k++ {
			fmt.Printf("\033[38;5;198m< Port > | %v |\n", NMAP.Host.Ports.Extraports.Extrareasons.Ports[k])
			fmt.Printf("    \033[38;5;198m< Stat > | %s | \n", NMAP.Host.Ports.Extraports.State)
			fmt.Printf("           \033[38;5;198m< Service information of port > \n")
			fmt.Printf("               \033[38;5;21mService | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Name)
			fmt.Printf("               \033[38;5;21mOS_type | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Ostype)
			fmt.Printf("               \033[38;5;21mProduct | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Product)
			fmt.Printf("               \033[38;5;21mVersion | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Version)
			fmt.Printf("               \033[38;5;21mCPE     | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Cpe)
			fmt.Printf("               \033[38;5;21mMethod  | \t%s    | \n", NMAP.Host.Ports.Port[k].Service.Method)
			fmt.Printf("          \033[38;5;198m < Reason > \n")
			fmt.Printf("              \033[38;5;21mReason   | \t%s    | \n", NMAP.Host.Ports.Extraports.Extrareasons.Reason)
			fmt.Printf("              \033[38;5;21mProtocal | \t%s    | \n", NMAP.Host.Ports.Extraports.Extrareasons.Proto)
			fmt.Printf("              \033[38;5;21mCount    | \t%s    | \n", NMAP.Host.Ports.Extraports.Extrareasons.Count)
			fmt.Printf("              \033[38;5;21mText     | \t%s    | \n", NMAP.Host.Ports.Extraports.Extrareasons.Text)
			fmt.Printf("              \033[38;5;21mPorts    | \t%s    | \n", NMAP.Host.Ports.Extraports.Extrareasons.Ports)
		}
	case "Runstats":
		fmt.Print("\033[38;5;198m< Script Finished Area > \n")
		fmt.Printf("\033[37m   | \n")
		fmt.Printf("\033[37m   | \n")
		fmt.Printf("\033[37m   |>> \033[38;5;21mTime        | %s | \n", NMAP.Runstats.Finished.Time)
		fmt.Printf("\033[37m   |>> \033[38;5;21mTime String | %s | \n", NMAP.Runstats.Finished.Timestr)
		fmt.Printf("\033[37m   |>> \033[38;5;21mSummary     | %s | \n", NMAP.Runstats.Finished.Summary)
		fmt.Printf("\033[37m   |>> \033[38;5;21mElapsed Time| %s | \n", NMAP.Runstats.Finished.Elapsed)
		fmt.Printf("\033[37m   |>> \033[38;5;21mExit        | %s | \n", NMAP.Runstats.Finished.Exit)
	case "hosts":
		fmt.Printf("\033[38;5;198m< Run Statistics Hosts > \n")
		fmt.Printf("\033[37m   | \n")
		fmt.Printf("\033[37m   | \n")
		fmt.Printf("\033[37m   |>> \033[38;5;21mText        | %s | \n", NMAP.Runstats.Hosts.Text)
		fmt.Printf("\033[37m   |>> \033[38;5;21mDown        | %s | \n", NMAP.Runstats.Hosts.Down)
		fmt.Printf("\033[37m   |>> \033[38;5;21mUp          | %s | \n", NMAP.Runstats.Hosts.Up)
		fmt.Printf("\033[37m   |>> \033[38;5;21mTotal       | %s | \n", NMAP.Runstats.Hosts.Total)
	case "Host Hint":
		fmt.Printf("\033[38;5;198m< Host Hint Information > \n")
		fmt.Printf("\033[37m   |\n")
		fmt.Printf("\033[37m   |\n")
		fmt.Printf("\033[37m   |>> \033[38;5;21mStatus Text      |  %s | \n", NMAP.Hosthint.Text)
		fmt.Printf("\033[37m   |>> \033[38;5;21mStatus State     |  %s | \n", NMAP.Hosthint.Status.State)
		fmt.Printf("\033[37m   |>> \033[38;5;21mStatus Reason    | %s  |\n", NMAP.Hosthint.Status.Reason)
		fmt.Printf("\033[37m   |>> \033[38;5;21mStatus ReasonTTL | %s  |\n", NMAP.Hosthint.Status.ReasonTtl)
		fmt.Printf("\033[37m   |>> \033[38;5;21mAddress Text     | %s  |\n", NMAP.Hosthint.Address.Text)
		fmt.Printf("\033[37m   |>> \033[38;5;21mAddress Addr     | %s  |\n", NMAP.Hosthint.Address.Addr)
		fmt.Printf("\033[37m   |>> \033[38;5;21mAddress Type     | %s  |\n", NMAP.Hosthint.Address.Addrtype)
	case "debug":
		fmt.Printf("\033[38;5;198m< Debug Information > \n")
		fmt.Printf("\033[37m   |\n")
		fmt.Printf("\033[37m   |\n")
		fmt.Printf("\033[37m   |>> \033[38;5;21mLevel    |   %s   |\n", NMAP.Debugging.Level)
		fmt.Printf("\033[37m   |>> \033[38;5;21mText     |   %s   |\n", NMAP.Debugging.Text)
	case "verbose":
		fmt.Printf("\033[38;5;198m< Verbosity Information > \n")
		fmt.Printf("\033[37m   |\n")
		fmt.Printf("\033[37m   |\n")
		fmt.Printf("\033[37m   |>> \033[38;5;21mLevel    |   %s   |\n", NMAP.Verbose.Level)
		fmt.Printf("\033[37m   |>> \033[38;5;21mText     |   %s   |\n", NMAP.Verbose.Text)
	case "scan info":
		fmt.Printf("\033[38;5;198m< Scanning Information > \n")
		fmt.Printf("\033[37m   |\n")
		fmt.Printf("\033[37m   |\n")
		fmt.Printf("\033[37m   |>> \033[38;5;21mProtocal           |  %s   | \n", NMAP.Scaninfo.Protocol)
		fmt.Printf("\033[37m   |>> \033[38;5;21mType               |  %s   | \n", NMAP.Scaninfo.Type)
		fmt.Printf("\033[37m   |>> \033[38;5;21mNumber Of Services |  %s   | \n", NMAP.Scaninfo.Numservices)
		fmt.Printf("\033[37m   |>> \033[38;5;21mServices           |  %s   | \n", NMAP.Scaninfo.Services)
		fmt.Printf("\033[37m   |>> \033[38;5;21mText               |  %s   | \n", NMAP.Scaninfo.Text)
	case "scan time":
		fmt.Printf("\033[38;5;198m< Scanning Information > \n")
		fmt.Printf("\033[37m   |\n")
		fmt.Printf("\033[37m   |\n")
		fmt.Printf("\033[37m   |\n")
		fmt.Printf("\033[37m   |>> \t time Start  | %s | \n", NMAP.Host.Starttime)
		fmt.Printf("\033[37m   |>> \t Time end    | %s | \n", NMAP.Host.Endtime)
		fmt.Printf("\033[37m   |>> \t Time Finish | %s | \n", NMAP.Runstats.Finished.Time)
		fmt.Printf("\033[37m   |>> \t Time string | %s | \n", NMAP.Runstats.Finished.Timestr)
	default:
		fmt.Println("[-] Not a function or executable name")
	}
}

// port matching
func Find_and_Search(filename string, info string) string {
	f, x := os.Open(filename)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Close()
		b, _ := ioutil.ReadAll(f)
		var s structure.Registry
		xml.Unmarshal(b, &s)
		for i := 0; i < len(s.Record); i++ {
			num := s.Record[i].Number
			if strings.Compare(info, num) == 0 {
				fmt.Printf("\033[37m   |>> \033[38;5;21mPort number        |  \033[31m%s    \n", num)
				fmt.Printf("\033[37m   |>> \033[38;5;21mPort protocol      |  \033[31m%s    \n", s.Record[i].Protocol)
				fmt.Printf("\033[37m   |>> \033[38;5;21mPort service       |  \033[31m%s    \n", s.Record[i].Name)
				fmt.Printf("\033[37m   |>> \033[38;5;21mPort unauthorized? |  \033[31m%s    \n", s.Record[i].Unauthorized)
				fmt.Printf("\033[37m   |>> \033[38;5;21mPort Description   |  \033[31m%s    \n", s.Record[i].Description)
				fmt.Print("\n")
			}
		}
	}
	return ""

}
