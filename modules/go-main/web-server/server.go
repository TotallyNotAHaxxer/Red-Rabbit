/*
Info about this package

this package should not import anything other than color modules, type modules, structure modules, authentication/encryption, or constant modules
everything like errors, loggers, formatters etc will be kept defualt


*/

package rr6_web_server

import (
	"fmt"
	"log"
	emsg "main/modules/go-main/web-server/errors"
	okmsg "main/modules/go-main/web-server/ok++"
	"net"
	"net/http"
)

const (
	//_____________________________________________________________________________________//
	Server_port_primary = 5501             // Primary port for the webserver and fileserver
	Server_port_backup  = 5502             // Backup port in case 5501 is not ready to be used
	Server_pass_primary = "RR6__ADMIN"     // Password for server authentication
	Server_host_primary = "localhost"      // Host for the webserver
	Server_Port_DBPrime = 5432             // Common postgreSQL port in the case we run a DB
	Server_WebU_primary = "127.0.0.1:5501" // This is the primary URL for the primary port
	Server_WebU_backup  = "127.0.0.1:5502" // If the primary port can not be used, this url will be set
	Server_port_backup2 = "127.0.0.1:8080" // standard HTTP
	//--------------------------------------------------------------------------------------------//
)

func T_port(portnum int) (bool, string, uint) {
	port := fmt.Sprint(rune(portnum))
	listener, e := net.Listen("tcp", ":"+port)
	if e != nil {
		fmt.Println(emsg.Server_TCP_Listener_FAIL, portnum)
		return false, port, 0x00
	} else {
		fmt.Println(okmsg.Server_Stat_Using_Port_For, " -> ", portnum)
		e := listener.Close()
		if e != nil {
			log.Fatal(okmsg.Server_Stat_Failed_To_Listen, e)
		}
		return true, port, 0x00
	}
}

func Call_port() string {
	a, b, e := T_port(Server_port_primary)
	if a {
		if e != 0x00 {
			fmt.Println(b)
			return b
		} else {
			fmt.Println(emsg.Internal_Server_Error, e)
		}
	} else {
		fmt.Println(okmsg.Server_Stat_Primary_Port_Good)
		a, b, e := T_port(Server_port_backup)
		if a {
			if e != 0x00 {
				fmt.Println(okmsg.Server_Stat_Backup_Port_Good)
				fmt.Println(b)
				return b
			} else {
				fmt.Println(emsg.Internal_Server_Error, e)
			}
		}
	}
	return b
}

func sample_handeler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, there\n")
}

func call_server(port int) {
	parser := ":" + fmt.Sprint(rune(port))
	http.HandleFunc("/", sample_handeler)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(parser, nil))
}
