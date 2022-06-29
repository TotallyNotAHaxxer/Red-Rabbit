package SUPER_SSH_functions

import (
	"fmt"
	"io/ioutil"
	"log"

	color "main/modg/colors"
	e "main/modg/service/ssh/ssh-errors"

	"golang.org/x/crypto/ssh"
)

func Load_file(file_location string) (string, ssh.Signer) {
	c, x := ioutil.ReadFile(file_location)
	e.Normal(x, "<RR6> OS/File : Could not read the filename -> ", color.REDHB)
	k, xe := ssh.ParsePrivateKey(c)
	e.Normal(xe, "<RR6> Net SSH : Could not parse the key from the file -> ", color.REDHB)
	return "", k
}

// runs

func Run(session *ssh.Session, command string) {
	x := session.Run(command)
	if x != nil {
		log.Fatal(x)
	} else {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Command executed....")
	}
}

// prints session information
func Print_session(session *ssh.Client) {
	fmt.Println("<RR6> Net SSH : Information Session ID     | ", session.SessionID(), " | ")
	fmt.Println("<RR6> Net SSH : Information Client version | ", session.ClientVersion(), " | ")
	fmt.Println("<RR6> Net SSH : Information Remote address | ", session.RemoteAddr(), " | ")
	fmt.Println("<RR6> Net SSH : Information Server Version | ", session.ServerVersion(), " | ")
	fmt.Println("<RR6> Net SSH : Information Session User   | ", session.User(), " | ")

}
