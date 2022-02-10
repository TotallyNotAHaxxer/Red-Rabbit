/*

Domain name guesser in go using basic wordlists to parse information

author => ArkAngeL43

issues, this code was built off of BHG LABS, a book that teahce sblack hat golang, while the subdoamin guesser was a great idea there are some organization that
was needed and left out explanation, for some reason if you make m dns.msg a public variable ( the variable visible to all classes ) then you get a smaller
output and not as in detailed response, during changing and maniuplation of the code to modern go programming language standards from 2016 standards and styles
i cam across another issue which included a NIL memory difference in changing the way the workers are called, even though this was the way to operate workers
in modern golang, there still were some issues


Libraries Third party:
			Gotabulate
			dns

what is the goal
the goal of this script is to guess subdomains and grab A and CNAME DNS record which will return the data in a table


*/
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"text/tabwriter"
	"time"

	"github.com/bndr/gotabulate"
	"github.com/miekg/dns"
)

func base_ascii() {
	content, err := ioutil.ReadFile("banner.txt")
	// err
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\033[31m", string(content))
	}
}

type empty struct{}

type result struct {
	IPADDR   string
	HOSTNAME string
}

var (
	results        []result
	flagdomain     = flag.String("domain", "", "Set the domain name")
	flaglist       = flag.String("list", "names.txt", "set wordlist for guessing")
	flagworker     = flag.Int("workers", 100, "Set workers")
	flagserveraddr = flag.String("serverIPAPort", "8.8.8.8:53", "DNS server to use")
)

func LOOKUP_A_DNS_RECORD(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var ips []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return ips, err
	}
	if len(in.Answer) < 1 {
		return ips, errors.New("no answer")
	}
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
		}
	}
	return ips, nil
}

func look_CNAME_DNS_RECORD(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var fqdns []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeCNAME)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return fqdns, err
	}
	if len(in.Answer) < 1 {
		return fqdns, errors.New("no answer")
	}
	for _, answer := range in.Answer {
		if c, ok := answer.(*dns.CNAME); ok {
			fqdns = append(fqdns, c.Target)
		}
	}
	return fqdns, nil
}

func lookup_final_DIALDEF(fqdn, serverAddr string) []result {
	var results []result
	var cfqdn = fqdn // Don't modify the original.
	for {
		cnames, err := look_CNAME_DNS_RECORD(cfqdn, serverAddr)
		if err == nil && len(cnames) > 0 {
			cfqdn = cnames[0]
			continue // process the CNAME records
		}
		ips, err := LOOKUP_A_DNS_RECORD(cfqdn, serverAddr)
		if err != nil {
			break
			// if there are no A records then break
		}
		for _, ip := range ips {
			results = append(results, result{IPADDR: ip, HOSTNAME: fqdn})
		}
		break // once done processing break the for loop and
	}
	return results // return the results
}

// during goroutines and using go's amazing concurrecy its nice to have a control
// or a os interupt listner, to listen for any hangups ex CTRL+c or OS>KILL, or KILL
// this will just give better handeling
func killer(c chan os.Signal) {
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

// checking if the file currently exists
func file_checker_ava1(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func worker(tracker chan empty, fqdns chan string, gather chan []result, serverAddr string) {
	for fqdn := range fqdns {
		results := lookup_final_DIALDEF(fqdn, serverAddr)
		if len(results) > 0 {
			gather <- results
		}
	}
	var e empty
	tracker <- e
}

// setting data

func data_table_targetINF(filename, domain_name, domain_server string, workers_in_use int) {
	// check if the file exists
	if file_checker_ava1(filename) {
		// do something with this later
	} else {
		fmt.Println("WARN: FATAL: DEFUALT WORDLIST USED AS ", filename, " COULD NOT BE FOUND OR PARSED")
	}
	row_1 := []interface{}{filename, domain_name, domain_name, workers_in_use}
	t := gotabulate.Create([][]interface{}{row_1})
	t.SetHeaders([]string{"filename", "domain name", "defualt domain server", "go workers in use"})
	t.SetEmptyString("None")
	t.SetAlign("center")
	fmt.Println("\033[34m", t.Render("grid"))

}

// ooutput the finalized and parsed data in a data table
func data_table_dump_all_main(filename, host, ipaddr string) {
	if file_checker_ava1(filename) {
		// filename is the file to dump all the data in, check first to see if it exists if so pass if else pass then create a new file names in.txt and output out.txt
	} else {
		fmt.Println("WARN: FATAL: ERROR, creating file since out file doesnt exist")
	}
}

func main() {
	flag.Parse()
	base_ascii()
	data_table_targetINF(*flaglist, *flagdomain, *flagserveraddr, *flagworker)
	fqdns := make(chan string, *flagworker)
	gather := make(chan []result)
	tracker := make(chan empty)

	fh, err := os.Open(*flaglist)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)

	for i := 0; i < *flagworker; i++ {
		go worker(tracker, fqdns, gather, *flagserveraddr)
	}
	go func() {
		for r := range gather {
			results = append(results, r...)
		}
		var e empty
		tracker <- e
	}()
	for scanner.Scan() {
		fqdns <- fmt.Sprintf("%s.%s", scanner.Text(), *flagdomain)
	}
	close(fqdns)
	for i := 0; i < *flagworker; i++ {
		<-tracker
	}
	close(gather)
	<-tracker
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
	// declare and opem pathmain before the for loop to prevwent NIL MEM INDIFFERENCE error
	pathmain, err_file := os.OpenFile("in.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err_file != nil {
		file_create, err := os.Create("in.txt")
		if err != nil {
			panic(err)
		} else {
			// close file use variable filec
			file_create.Close()
		}
	}
	for _, r := range results {
		// r.HOSTNAME is the host domain parsed
		// r.IPADDR is the hosts IP address
		// IPA = Internet Protocal Address
		// write this data to a file, then after for loop call from a function to output the data in a table
		fmt.Fprintf(w, "\033[31m HOST ~~~> %s\t IP ADDR OF HOST ~~~> %s\n", r.HOSTNAME, r.IPADDR)
		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		fmt.Println("\033[32m NOTE: output saved in path RR5/modules/domain/in.txt")
		fmt.Fprintln(pathmain, r.HOSTNAME, r.IPADDR)
		//
		//fmt.Fprintln(pathmain, "IPADDR => ", net.IP(arp.SourceProtAddress), "has MAC => ", net.HardwareAddr(arp.SourceHwAddress))

	}
	w.Flush()
}
