package script_constants

import (
	"crypto/tls"
	"net/http"
	"sync"
)

// package holds all basic constants used by every mut mod

var (
	Reuslt_map = make(map[string][]string)
	Mut        = &sync.Mutex{}
	WaitGroup  = sync.WaitGroup{}
	Limiter    = make(chan string, 10)
	Results    []string
)

var HTTP_Client_Transport_Configuration = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}

var HTTP_Client_Main = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}, Transport: HTTP_Client_Transport_Configuration,
}
