package sockets

import (
	"fmt"
	"net"
	"net/http"
	"time"

	xo "main/modg/colors"
)

// check if a port is useable

func Check_Activity(port string) string {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", port), timeout)
	if err != nil {
		fmt.Println("Got error when trying to get port -> ", err)
	}
	if conn != nil {
		defer conn.Close()
		return port
	}
	return "no port is active?"
}

// run custom header

func Header_based(url, header, header_content, method string, timeout time.Duration) {
	c := &http.Client{Timeout: time.Second * timeout}
	r, x := http.NewRequest(method, url, nil)
	if x != nil {
		fmt.Println(xo.REDHB, "<RR6> Custom Requests Module: Could not customize request or make a new http response and request method got error -> ", x)
	} else {
		r.Header.Add("Host", "http://evil.com")
		response, x := c.Do(r)
		if x != nil {
			fmt.Println(xo.REDHB, "<RR6> Custom Requests Module: Could not customize request or make a firmed request to the URL or server using custom rule, got error -> ", x)
		} else {
			for k, v := range r.Header {
				fmt.Printf("%s %s\n", k, v)
			}
			defer response.Body.Close()
		}
	}
}
