package requests

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	c "main/modg/colors"
	client "main/modg/requests/vars"
	e "main/modg/warnings"

	"github.com/miekg/dns"
)

const (
	content_type_PUT = "Content-Type"
	application_PUT  = "application/json"
	CHARSET_PUT      = "charset=utf-8"
)

var (
	R           = ""
	HTTP_CLIENT = &http.Client{}
	Cc_chan     = make(chan error, 1)
)

type NIL struct{}

type SUB_RES struct {
	IPA  string
	HOST string
}

// defualt for all requests, this is a package to preform HEAD, PUT, DELETE, METHOD, POST, GET and other HTTP/HTTPS methods

// check errors
func Check_error(err error, msg, color string) {
	if err != nil {
		fmt.Println(color, msg, err)
	}
}

//make a GET
func Create_GET(target string) (string, int, error) {
	resp, err := http.Get(target)
	e.Warning_advanced("<RR6> Requests Module: Could not make a GET request to the target (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the body of the response? (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	a := string(rbody)
	return resp.Status, len(a), nil
}

//POST
func Create_POST(target string) (string, int, error) {
	pb, _ := json.Marshal("{data}")
	rb := bytes.NewBuffer(pb)
	response, err := http.Post(target, "application/json", rb)
	e.Warning_advanced("<RR6> Requests Module: Could not make a POST methodized request to the target? (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the body of the response? (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	a := string(body)
	return response.Status, len(a), nil
}

//POST with manual data
func Create_POST_data(target, data string) (string, int, string, error) {
	pb, _ := json.Marshal(data)
	rb := bytes.NewBuffer(pb)
	response, err := http.Post(target, "application/json", rb)
	e.Warning_advanced("<RR6> Requests Module: Could not make a POST methodized request to the target? (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the body of the response? (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	a := string(body)
	return response.Status, len(a), a, nil
}

//HEAD
func Create_HEAD(target string) (string, int, error) {
	response, ec := http.Head(target)
	e.Warning_advanced("<RR6> Requests Module: Could not make a HEAD methodized request to the target (>>>) ", c.REDHB, 1, false, false, true, ec, 1, 255, "")
	defer response.Body.Close()
	b, ea := ioutil.ReadAll(response.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the body of the response? (>>>) ", c.REDHB, 1, false, false, true, ea, 1, 255, "")
	sb := string(b)
	return response.Status, len(sb), nil
}

// Create GET and read body with the HTTP client and return body
func Create_GET_Body(url string) (string, uint, error) {
	client := &http.Client{}
	request, e := http.NewRequest("GET", url, nil)
	if e != nil {
		return "<RR6> Requests Module: Could not create a client GET request to the server -> ", 0x00, e
	} else {
		response, e := client.Do(request)
		if e != nil {
			return "<RR6> Requests Module: Could not make a GET request to the server -> ", 0x00, e
		} else {
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			return_B := string(body)
			return return_B, 0x00, nil
		}
	}
}

// scanning files for targets to test
func Scanners(filename string) []string {
	var result []string
	f, x := os.Open(filename)
	if x != nil {
		fmt.Println("<RR6> File I/O: Could not open file, find file, locate file, etc, got error -> ", x)
		os.Exit(0)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		domain := strings.ToLower(scanner.Text())
		result = append(result, domain)
	}
	return Remove_URL_vals(result)
}

// GET look for a code
func Create_Look_GET(target string, code int) {
	req, x := http.NewRequest("GET", target, nil)
	if x != nil {
		fmt.Println("<RR6> Requests Module: Could not prepare a GET request to the target URL")
	} else {
		response, e := client.HTTP_CLIENT.Do(req)
		if e != nil {
			fmt.Println(c.REDHB, "<RR6> Requests module: Got a major error -> Could not make a request to the server using the http client, got error -> ", e)
		} else {
			if response.StatusCode == code {
				fmt.Println("<RR6> Requests Module: Got the status code, request has been made -> ", response.StatusCode)
			} else {
				fmt.Println("<RR6> Requests Module: The status code did not match [ ", response.StatusCode, " ] With  [ ", response.StatusCode, " ] This might not be what you are looking for")
			}
		}
	}
}

//PUT
func Create_PUT(target string) (string, int, error) {
	client := &http.Client{}
	json, _ := json.Marshal("{data}")
	req, err := http.NewRequest(http.MethodPut, target, bytes.NewBuffer(json))
	e.Warning_advanced("<RR6> Requests Module: Could not make a new PUT methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	req.Header.Set(content_type_PUT, application_PUT)
	resp, err := client.Do(req)
	e.Warning_advanced("<RR6> Requests Module: Could not send the PUT methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	body, err := ioutil.ReadAll(resp.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the response body (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	b := string(body)
	return resp.Status, len(b), nil
}

// put with manual data
func Create_PUT_data(target, data string) (string, int, string, error) {
	client := &http.Client{}
	json, _ := json.Marshal(data)
	req, err := http.NewRequest(http.MethodPut, target, bytes.NewBuffer(json))
	e.Warning_advanced("<RR6> Requests Module: Could not make a new PUT methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	req.Header.Set(content_type_PUT, application_PUT)
	resp, err := client.Do(req)
	e.Warning_advanced("<RR6> Requests Module: Could not send the PUT methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	body, err := ioutil.ReadAll(resp.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the response body (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	b := string(body)
	return resp.Status, len(b), b, nil
}

//send the methodized request
func Request(target string, method string) (string, int, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, target, nil)
	e.Warning_advanced("<RR6> Requests Module: Could not make a new PUT methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	resp, err := client.Do(req)
	e.Warning_advanced("<RR6> Requests Module: Could not send the methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the response body (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	sb := string(body)
	return resp.Status, len(sb), nil
}

//send the methodized request with body response
func Request_body(target string, method string) (string, int, string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, target, nil)
	e.Warning_advanced("<RR6> Requests Module: Could not make a new PUT methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	resp, err := client.Do(req)
	e.Warning_advanced("<RR6> Requests Module: Could not send the methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the response body (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	sb := string(body)
	return resp.Status, len(sb), sb, nil
}

func Request_body_IO(target string, method string) (string, int, io.Reader, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, target, nil)
	e.Warning_advanced("<RR6> Requests Module: Could not make a new PUT methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	resp, err := client.Do(req)
	e.Warning_advanced("<RR6> Requests Module: Could not send the methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	body, err := ioutil.ReadAll(resp.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the response body (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	sb := string(body)
	return resp.Status, len(sb), resp.Body, nil
}

func GET_val(target, method, key string) (string, uint) {
	cl := &http.Client{}
	req, e := http.NewRequest(method, target, nil)
	if e != nil {
		fmt.Println("<RR6> Request module: Could not make a new methodized request (>>>) ", c.REDHB, e)
	}
	resp, e := cl.Do(req)
	if e != nil {
		fmt.Println("<RR6> Request module: Could not execute the new methodized request (>>>) ", c.REDHB, e)
	}
	defer resp.Body.Close()
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("<RR6> Request module: Could not read the response body (>>>) ", c.REDHB, e)
	}
	return resp.Header.Get("server"), 0x00
}

func LOOKUP_A_DNS_RECORD(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var ips []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return ips, err
	}
	if len(in.Answer) < 1 {
		return ips, errors.New("no answer")
	}
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
		}
	}
	return ips, nil
}

func Look_CNAME_DNS_RECORD(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var fqdns []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeCNAME)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return fqdns, err
	}
	if len(in.Answer) < 1 {
		return fqdns, errors.New("no answer")
	}
	for _, answer := range in.Answer {
		if c, ok := answer.(*dns.CNAME); ok {
			fqdns = append(fqdns, c.Target)
		}
	}
	return fqdns, nil
}

func Lookup_final_DIALDEF(fqdn, serverAddr string) []SUB_RES {
	var results []SUB_RES
	var cfqdn = fqdn
	for {
		s, x := Look_CNAME_DNS_RECORD(cfqdn, serverAddr)
		if x == nil && len(s) > 0 {
			cfqdn = s[0]
			continue
		}
		s, y := LOOKUP_A_DNS_RECORD(cfqdn, serverAddr)
		if y != nil {
			break
		}
		for _, ip := range s {
			results = append(results, SUB_RES{IPA: ip, HOST: fqdn})
		}
		break
	}
	return results
}

// read request data
func GET_RESP(url string) (string, int, error) {
	response, x := HTTP_CLIENT.Get(url)
	Ce(x, "<RR6> Requests module: Could not make a get request to the given URL.... -> ", "fmt", 1)
	defer response.Body.Close()
	body, x := ioutil.ReadAll(response.Body)
	Ce(x, "<RR6> Requests module: Could not read the response body properly........ -> ", "fmt", 1)
	return string(body), len(body), nil
}

func GET_URL_PROTO(i string) string {
	r := strings.Index(i, "://")
	if r >= 0 {
		return i[:r]
	} else {
		return ""
	}
}

func Remove_URL_vals(s []string) []string {
	l := []string{}
	k := make(map[string]bool)
	for _, e := range s {
		if _, v := k[e]; !v {
			k[e] = true
			l = append(l, e)
		}
	}
	return l
}

func Remove_Params(input string, payload string) string {
	y, x := url.Parse(input)
	Ce(x, "<RR6> Requests module: Could not parse the URL properly... -> ", "fmt", 1)
	v, x := url.QueryUnescape(y.RawQuery)
	Ce(x, "<RR6> Requests module: Could not parse the URL properly... -> ", "fmt", 1)
	c := strings.Split(v, "&")
	for _, a := range c {
		vals := strings.Split(a, "=")
		R += vals[0] + "=" + url.QueryEscape(payload) + "&"
	}
	return y.Scheme + "://" + y.Host + y.Path + "?" + R[:len(R)-1]
}

// request without errors
func Make_GET_NOERR(url string) {
	resp, x := http.Get(url)
	if x != nil {
		fmt.Println(c.WHT, "\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Was not able to make a request to the URL, URL has been blocked!")
	} else {
		if resp.StatusCode == 200 {
			fmt.Println(c.REDHB, "[!] Was able to make a complete request to the URL, meaning this host was not blocked. blocker has failed...")
		}
	}
}

func Ce(err error, msg string, typer string, exit_code int) bool {
	if err != nil {
		if typer == "fmt" {
			fmt.Println(err, msg, exit_code)
			os.Exit(exit_code)
		}
		if typer == "log" {
			log.Fatal(err, msg)
			os.Exit(exit_code)
		}
	} else {
		return true
	}
	return true
}

func Req_with_Response(uri string) (string, *http.Response, error) {
	var HTTP_ = http.Client{}
	request, X := http.NewRequest("GET", uri, nil)
	if X != nil {
		fmt.Printf("\033[38;5;55m|\033[38;5;88m-\033[38;5;55m| \033[38;5;88m Got a error when trying to make a request to the url, you might be offline, or connection may be private or unstable, please try again following this error -> %s ", X)
		return "", nil, X
	}
	response, X := HTTP_.Do(request)
	if X != nil {
		fmt.Printf("\033[38;5;55m|\033[38;5;88m-\033[38;5;55m| \033[38;5;88m Got a error when trying to make a request to the url, you might be offline, or connection may be private or unstable, please try again following this error -> %s ", X)
		return "", nil, X
	}
	defer response.Body.Close()
	b, X := ioutil.ReadAll(response.Body)
	if X != nil {
		fmt.Printf("\033[38;5;55m|\033[38;5;88m-\033[38;5;55m| \033[38;5;88m Got a error when trying to read the response body, maybe the body is corrupted, of wrong structure, of wrong type, or just messed up from a broken connection -> %s ", X)
		return "", nil, X
	}
	return string(b), response, X
}

func Isexisting(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Write(Filename_of_data, data string) {
	if !Isexisting(Filename_of_data) {
		d1 := []byte(data)
		err := os.WriteFile(Filename_of_data, d1, 0644)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("\033[31m<RR6> Logging: File and data have been saved...")
		}
	}
}
