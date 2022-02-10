package main

import (
	"fmt"
	"net/http"
)

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/main":
		http.ServeFile(w, req, "server-conf-html/main.html")
	default:
		http.ServeFile(w, req, "server-conf-html/err.html")
	}
}

func main() {
	var r router
	fmt.Println("[+] Please navigate to http://localhost:8000/main in you're browser")
	http.ListenAndServe(":8000", &r)

}
