package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/vmihailenco/msgpack.v2"
)

type Metasploit struct {
	host  string
	user  string
	pass  string
	token string
}

type login_Request struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

type login_Response struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type logout_Request struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	LogoutToken string
}

type logout_Response struct {
	Result string `msgpack:"result"`
}

type sessionList_Request struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type Session_List_Response struct {
	ID          uint32 `msgpack:",omitempty"`
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload"`
	Description string `msgpack:"desc"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"session_host"`
	SessionPort int    `msgpack:"session_port"`
	Username    string `msgpack:"username"`
	UUID        string `msgpack:"uuid"`
	ExploitUUID string `msgpack:"exploit_uuid"`
}

// added banner
func banner(file string) {
	content, err := ioutil.ReadFile(file)
	checker_msf_err(err, "error when opening file", 1)
	fmt.Println("\x1b[H\x1b[2J\x1b[3J", string(content))
}

// called err checking function
func checker_msf_err(err error, msg string, exit_code int) bool {
	if err != nil {
		log.Fatal(msg, err)
		os.Exit(exit_code)
		return true
	}
	return false
}

func New(host, user, pass string) (*Metasploit, error) {
	msf := &Metasploit{host: host, user: user, pass: pass}

	if err := msf.Login(); err != nil {
		return nil, err
	}

	return msf, nil
}

func (msf *Metasploit) send(req interface{}, res interface{}) error {
	buf := new(bytes.Buffer)
	msgpack.NewEncoder(buf).Encode(req)
	// fixed POST REQ
	dest := fmt.Sprintf("http://%s:55552/api/", msf.host)
	r, err := http.Post(dest, "binary/message-pack", buf)
	checker_msf_err(err, "[ ERROR ] [ FATAL ] When making POST request to MSF API URL -> ", 1)
	defer r.Body.Close()

	if err := msgpack.NewDecoder(r.Body).Decode(&res); err != nil {
		return err
	}
	return nil
}

func (msf *Metasploit) Login() error {
	ctx := &login_Request{Method: "auth.login", Username: msf.user, Password: msf.pass}
	var res login_Response
	if err := msf.send(ctx, &res); err != nil {
		return err
	}
	msf.token = res.Token
	return nil
}

func (msf *Metasploit) Logout() error {
	ctx := &logout_Request{
		Method:      "auth.logout",
		Token:       msf.token,
		LogoutToken: msf.token,
	}
	var res logout_Response
	if err := msf.send(ctx, &res); err != nil {
		return err
	}
	msf.token = ""
	return nil
}

func (msf *Metasploit) SessionList() (map[uint32]Session_List_Response, error) {
	req := &sessionList_Request{Method: "session.list", Token: msf.token}
	res := make(map[uint32]Session_List_Response)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}

	for id, session := range res {
		session.ID = id
		res[id] = session
	}
	return res, nil
}

func main() {
	banner("msf.txt")
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := "msf"

	if host == "" || pass == "" {
		log.Fatalln("Missing required environment variable MSFHOST or MSFPASS")
	}

	msf, err := New(host, user, pass)
	checker_msf_err(err, "[ ERROR ] [ FATAL ] When CALLING NEW function ", 1)

	defer msf.Logout()

	sessions, err := msf.SessionList()

	checker_msf_err(err, "[ ERROR ] [ FATAL ] When calling sessions list ", 1)
	fmt.Println("Current MSF Sessions: ")
	for _, msfs := range sessions {
		fmt.Printf("\t%5d  %s\n", msfs.ID, msfs.Info)
	}
}
