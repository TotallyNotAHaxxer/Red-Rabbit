package SSH_SUPER_session

import (
	"fmt"
	"log"
	ssh_key "main/modg/service/ssh/ssh-functions"
	"os"

	"golang.org/x/crypto/ssh"
)

func SSH_Shell(user, host, file, sys string, h, w int) {
	a, k := ssh_key.Load_file(file)
	fmt.Println("<RR6> SSH Session -> Using key | ", k)
	fmt.Println("<RR6> SSH Session -> Session   | Starting....")
	fmt.Println(a)
	ssh_conf := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(k),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	c, x := ssh.Dial("tcp", host, ssh_conf)
	if x != nil {
		fmt.Println("<RR6> SSH Session -> Could not dial server |", x)
	} else {
		ssh_shell, x := c.NewSession()
		ssh_shell.Stdout = os.Stdout
		ssh_shell.Stdin = os.Stdin
		ssh_shell.Stderr = os.Stderr
		if x != nil {
			fmt.Println("<RR6> SSH Session -> Could not make a new session -> ", x)
		} else {
			defer ssh_shell.Close()
			switch sys {
			case "windows":
				x = ssh_shell.RequestPty("vt100", h, w, ssh.TerminalModes{ssh.ECHO: 0})
			case "linux":
				x = ssh_shell.RequestPty("xterm", h, w, ssh.TerminalModes{ssh.ECHO: 0})
			}
			if x != nil {
				log.Fatal(x)
			} else {
				x = ssh_shell.Shell()
				if x != nil {
					fmt.Println("<RR6> SSH Session -> Could not make a new shell...")
				} else {
					fmt.Println("[+] Command executed...")
				}
				ssh_shell.Wait()
			}
		}
	}
}
