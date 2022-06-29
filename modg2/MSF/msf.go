package MSF

import (
	"bytes"
	"fmt"
	"net/http"

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

var (
	x   error
	res login_Response
)

func Session_New(hostname, username, password string) (*Metasploit, error) {
	m := &Metasploit{
		host: hostname,
		user: username,
		pass: password,
	}
	if x = m.Session_Login(); x != nil {
		return nil, x
	}
	return m, nil
}

func (metasploit_structure *Metasploit) Session_Login() error {
	ctx := &login_Request{
		Method:   "auth.login",
		Username: metasploit_structure.user,
		Password: metasploit_structure.pass,
	}
	if x := metasploit_structure.Session_Send(ctx, &res); x != nil {
		return nil
	}
	metasploit_structure.token = res.Token
	return nil
}

func (Metasploit_Structure *Metasploit) Session_Send(request interface{}, response interface{}) error {
	buffer := new(bytes.Buffer)
	msgpack.NewEncoder(buffer).Encode(request)
	dest := fmt.Sprintf("http://%s:55552/api/", Metasploit_Structure.host)
	r, x := http.Post(dest, "binary/message-pack", buffer)
	if x != nil {
		return x
	} else {
		defer r.Body.Close()
		if x := msgpack.NewDecoder(r.Body).Decode(&response); x != nil {
			return x
		}
		return nil
	}

}

func (Metasploit_Structure *Metasploit) Session_Logout() error {
	ctx := &logout_Request{
		Method:      "auth.logout",
		Token:       Metasploit_Structure.token,
		LogoutToken: Metasploit_Structure.token,
	}
	var result logout_Response
	if x = Metasploit_Structure.Session_Send(ctx, &result); x != nil {
		return x
	}
	Metasploit_Structure.token = ""
	return nil
}

func (Metasploit_Structure *Metasploit) Session_List() (map[uint32]Session_List_Response, error) {
	request := &sessionList_Request{
		Method: "session.list",
		Token:  Metasploit_Structure.token,
	}
	response := make(map[uint32]Session_List_Response)
	if x = Metasploit_Structure.Session_Send(request, &response); x != nil {
		return nil, x
	}
	for I, K := range response {
		K.ID = I
		response[I] = K
	}
	return response, nil
}
