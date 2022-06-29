package ServicesEtcUtils

import (
	"fmt"
	"io"
	"net"

	c "main/modg/colors"

	"golang.org/x/crypto/ssh"
)

type E struct {
	Hostname  string
	Host_port int
}

type Tun struct {
	Loc_end    *E
	Server_end *E
	Remote_end *E

	Conf *ssh.ClientConfig
}

func (e *E) String() string {
	return fmt.Sprintf("%s:%d", e.Hostname, e.Host_port)
}

func (t *Tun) Run() uint32 {
	fmt.Println("[- Running Tunnel")
	l, x := net.Listen("tcp", t.Loc_end.String())
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> SSH/Service based utilities module: Could not listen on local endpoint got error when initalizing structure -> ", x)
		return 0x00
	} else {
		defer l.Close()
		for {
			o, x := l.Accept()
			if x != nil {
				fmt.Println(c.REDHB, "<RR6> SSH/Service based utilities module: Could not accept the connection it seems as if the ssh configuration has recieved an error -> ", x)
				return 0x00
			} else {
				go t.Call(o)
			}
		}
	}
}

func (t *Tun) Call(pass_connection net.Conn) uint16 {
	sc, x := ssh.Dial("tcp", t.Server_end.String(), t.Conf)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> SSH/Service based utilities module: Could not ping the servers endpoint, got some error when running ssh dial call function with method tcp, config=config, data=Tun ", x)
		return 0x00
	} else {
		rc, x := sc.Dial("tcp", t.Remote_end.String())
		if x != nil {
			fmt.Println(c.REDHB, "<RR6> SSH/Service based utilities module: Could not ping the previous dial methods remote end address with config being of type Tun, method=tcp, func=dial error=", x)
			return 0x00
		} else {
			Copy_Connection := func(w, r net.Conn) {
				_, x := io.Copy(w, r)
				if x != nil {
					fmt.Println("<RR6> SSH/Service based utilities module: Could not create / copy current living connections to the I/O Got error=0x01 mean=error not nil, error.error = ", x)
				}
			}
			go Copy_Connection(pass_connection, rc)
			go Copy_Connection(rc, pass_connection)
		}
	}
	return 0x00
}

func Parse_Args_conf(
	lo_end_host,
	server_end_host,
	remote_end_host,
	SSH_USERNAME string,
	lo_end_port,
	server_end_port,
	remote_end_port int) {
	LO := &E{
		Hostname:  lo_end_host,
		Host_port: lo_end_port,
	}
	fmt.Println("< Local Host Endpoint Config > ")
	fmt.Println("| Local Host Endpoint Host :: ", lo_end_host)
	fmt.Println("| Local host Endpoint Port :: ", lo_end_port)
	SERVER := &E{
		Hostname:  server_end_host,
		Host_port: server_end_port,
	}
	fmt.Println("< Server host endpoint config > ")
	REMOTE := &E{
		Hostname:  remote_end_host,
		Host_port: remote_end_port,
	}
	fmt.Println("| Remote Server endpoint Host :: ", remote_end_host)
	fmt.Println("| Remote Server endpoint Port :: ", remote_end_port)
	CONF := &ssh.ClientConfig{
		User: SSH_USERNAME,
		Auth: []ssh.AuthMethod{
			Agent(),
		},
	}
	fmt.Println("< SSH Configuration for auth > ")
	fmt.Println("| SSH Server Username     :: ", CONF.User)
	fmt.Println("| SSH Server Auth Method  :: ", CONF.Auth)
	TUN := &Tun{
		Conf:       CONF,
		Loc_end:    LO,
		Server_end: SERVER,
		Remote_end: REMOTE,
	}
	TUN.Run()
}
