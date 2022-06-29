// package will run all SSH related attacks outside of brute forcing
package SUPER_SSH

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"

	ssh_constants "main/modg/service/ssh/ssh-constants"
	ssh_functions "main/modg/service/ssh/ssh-functions"

	"golang.org/x/crypto/ssh"
)

// checks the port on the host
func Check_Port(hostname string) (string, error) {
	for _, p := range ssh_constants.Ports_list {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Checking port -> ", p)
		to := time.Second
		connect, x := net.DialTimeout("tcp", net.JoinHostPort(hostname, p), to)
		if x != nil {
			return "<RR6> Net SSH: Got an error when checking ports activity -> ", x
		}
		if connect != nil {
			defer connect.Close()
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Host on port -> ", net.JoinHostPort(hostname, p), " is open")
			return "\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Port active true", nil
		}
	}
	return "[!] Port could not be found, this might be a fatal error -> ", nil
}

// checks connection to the host
func Make_Request(host string) (string, error) {
	conn, x := net.Dial("tcp", net.JoinHostPort(host, "22"))
	if x != nil {
		return "<RR6> Net SSH: Got an error when trying to dial the host -> ", x
	}
	if conn != nil {
		fmt.Println("<RR6> Net SSH: Information | Remote Host | > ", conn.RemoteAddr())
		fmt.Println("<RR6> Net SSH: Information | Remote Host | > Host is alive")
		fmt.Print("\n")
		return "<RR6> Net SSH: Status -> Was able to make a connection to the host\n", nil
	}
	return "<RR6> Net SSH: Was not able to make a dial to the target, got error -> ", x
}

// Test authentication
func Make_Auth(username, password, hostname string) (string, error) {
	for _, p := range ssh_constants.Ports_list {
		fmt.Println("[*] Trying port -> ", p)
		config := &ssh.ClientConfig{
			User:            username,
			Auth:            []ssh.AuthMethod{ssh.Password(password)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		_, x := ssh.Dial("tcp", net.JoinHostPort(hostname, p), config)
		if x != nil {
			return "<RR6> Net SSH: Could not make a network dial to the host -> ", x
		} else {
			return "<RR6> Net SSH: Was able to make a proper connection to the host, you have logged in sucessfully -> ", nil
		}
	}
	return "[!] What???? ", nil
}

// Test private key authentication
func Private_Auth(username, password, hostname, port, filename string) {
	get_key, _ := ssh_functions.Load_file(filename)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Using key -> ", get_key)
	ssh_key_conf := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	session, x := ssh.Dial("tcp", net.JoinHostPort(hostname, port), ssh_key_conf)
	if x != nil {
		fmt.Println("<RR6> Net SSH: Could not dial the server, got error -> ", x)
	} else {
		fmt.Println("<RR6> Net SSH: Was able to make a connection to the server")
		fmt.Println("<RR6> Net SSH: Using key -> ", get_key)
		ssh_functions.Print_session(session)
	}
}

// Execute
func Exec_SSH(command, password, username, port, host string) {
	ssh_key_conf := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	session, x := ssh.Dial("tcp", net.JoinHostPort(host, port), ssh_key_conf)
	if x != nil {
		fmt.Println("<RR6> Net SSH: Could not dial the server, got error -> ", x)
	} else {
		fmt.Println("<RR6> Net SSH: Was able to make a connection to the server")
		ssh_functions.Print_session(session)
	}
	session_exec, x2 := session.NewSession()
	if x2 != nil {
		fmt.Println("<RR6> Net SSH: Could not make a second executable SSH session, x -> ", x2)
	} else {
		fmt.Println("<RR6> Net SSH: Executing command.....")
		defer session_exec.Close()
		session_exec.Stdout = os.Stdout
		ssh_functions.Run(session_exec, command)
	}
}

// transfer
// didnt feel like importing a new SCP client for go, this might work a bit better
func SCP_ssh(file, username, host string) {
	parser := username + "@" + host
	a := "scp"
	b := file
	c := parser
	cmd := exec.Command(a, b, c)
	_, x := cmd.Output()
	if x != nil {
		fmt.Println("<RR6> Net SSH: < SCP >  => Was not able to trasnport file to client -> ", parser)
	}
}
