package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	Server_pass_main              = "PQL-ADMIN-SERVER094"
	Database_pass_main            = "fancy-bear-2022-fuckg"
	version                       = "6.0"
	host                          = "localhost"
	port_DB                       = "5432"
	port_host                     = ":5501"
	second_port                   = ":5502"
	Server_UI_URL                 = "127.0.0.1:8080"
	method                        = "POST/GET"
	server_main_file              = ".go"
	powered_by                    = "Go"
	Processing                    = "r/perl5"
	banner_file                   = "banner.txt"
	defualt_route_test_connection = "https://www.google.com"
	_err_404                      = "html/error.html"
	_err_Passauth_failed_         = "html/auth-error.html"
	filepath                      = "/"
)

var (
	clear_hex = "\x1b[H\x1b[2J\x1b[3J"
	BLK       = "\033[0;30m"
	RED       = "\033[0;31m"
	GRN       = "\033[0;32m"
	YEL       = "\033[0;33m"
	BLU       = "\033[0;34m"
	MAG       = "\033[0;35m"
	CYN       = "\033[0;36m"
	WHT       = "\033[0;37m"
)

func file_server(http_writer_M1 http.ResponseWriter, http_request_reader *http.Request) {
	if http_request_reader.URL.Path != "/" {
		http.ServeFile(http_writer_M1, http_request_reader, _err_404)
	}
	switch http_request_reader.Method {
	case "GET":
		http.ServeFile(http_writer_M1, http_request_reader, "index.html")
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT  |=> ", http.StateActive)
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT  |=> ", http.StatusOK)
	}
}

func main() {
	http.HandleFunc(filepath, file_server)
	log.Fatal(http.ListenAndServe(port_host, nil))
}
