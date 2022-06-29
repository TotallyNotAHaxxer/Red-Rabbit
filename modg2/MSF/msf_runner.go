package MSF

import (
	"fmt"
	"os"
)

func Run() {
	hostname := os.Getenv("MSFHOST")
	password := os.Getenv("MSFPASS")
	username := "msf"
	if hostname == "" || password == "" {
		fmt.Println("<RR6> Metasploit Module -> Got error when trying to get enviroment variables MSFHOST and MSFPASS for hostname and password which will be needed to extract things such as session lists...")
	}
	m, x := Session_New(hostname, username, password)
	if x != nil {
		fmt.Println("<RR6> Metasploit Module -> NEW: Got error when trying to use the enviroment variables to start a new module session -> ", x)
	} else {
		defer m.Session_Logout()
		sessions, x := m.Session_List()
		if x != nil {
			fmt.Println("<RR6> Metasploit Module -> SESSION: Got error when trying to get the sessions list, for some reason the enviroment variables might not have been set properly -> ", x)
		} else {
			fmt.Println("__Number____ID_____Info_______")
			for i, l := range sessions {
				fmt.Printf("%v | %5d | %s", i, l.ID, l.Info)
			}
		}

	}
}
