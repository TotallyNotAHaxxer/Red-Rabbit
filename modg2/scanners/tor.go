package scanners

import (
	"fmt"
	"io"
	"io/ioutil"
	c "main/modg/colors"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

var Proxy string = "socks5://127.0.0.1:9050"
var U string = "https://api.ipify.org?format=text"

func Test() {
	t, x := url.Parse(Proxy)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> Tor -> Sockets -> Reader -> Parser: Could not use the tor sockets address given which is < socks5://127.0.0.1:9050, got an error when trying to parse the string -> ", x)
	} else {
		tt := &http.Transport{Proxy: http.ProxyURL(t)}
		client := &http.Client{Transport: tt, Timeout: time.Second * 5}
		resp, x := client.Get("https://api.ipify.org?format=text")
		if x != nil {
			fmt.Println(c.REDHB, "<RR6> Tor -> Sockets -> Reader -> Requests: Could not make a request using the tor transport client with url < socks5://127.0.0.1:9050 > got a weird error -> ", x)
		} else {
			defer resp.Body.Close()
			ip, x := ioutil.ReadAll(resp.Body)
			if x != nil {
				fmt.Println(c.REDHB, "<RR6> Tor -> Requests -> I/O -> Reader: Could not read the response body for a reason, using the I/O Module -> ", x)
			} else {
				fmt.Println(c.RET_RED)
				fmt.Printf("|+| Tor sock online ->  [ %s ] \n", Proxy)
				fmt.Printf("|+| Possible Tor IP ->  [ %s ] \n", ip)
				fmt.Printf("|+| OUT Request     ->  [ %s ] \n", U)
			}

		}
	}
}

func Download_File_TOR(uri, directory string) {
	var p string
	t, x := url.Parse(Proxy)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> Tor -> Sockets -> Reader -> Parser: Could not use the tor sockets address given which is < socks5://127.0.0.1:9050, got an error when trying to parse the string -> ", x)
	} else {
		endpath, x := url.Parse(uri)
		if x != nil {
			fmt.Println(c.REDHB, "<RR6> File -> I/O -> Requests -> URL: Was unable to parse the url's endpoint and endpath, this is required so that the user can have a valid output file of the downloaded content, try again? -> ", x)
		} else {
			p = directory + path.Base(endpath.Path)
		}

		tt := &http.Transport{
			Proxy: http.ProxyURL(t),
		}
		client := &http.Client{
			Transport: tt,
			Timeout:   time.Second * 10,
		}
		resp, x := client.Get(uri)
		if x != nil {
			fmt.Println(c.REDHB, "<RR6> Tor -> Sockets -> Reader -> Requests: Could not make a request using the tor transport client with url < socks5://127.0.0.1:9050 > got a weird error -> ", x)
		} else {
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				fmt.Println(c.BLK, "<RR6> HTTP -> Sockets -> Requests: Could not make a decent request to this server enough to get a 200 anyway -> ", resp.StatusCode)
			} else {
				file, x := os.Create(p)
				if x != nil {
					fmt.Println(c.REDHB, "<RR6> OFFLINE -> Module -> File I/O: Got error when trying to create this new file from the URL to copy data from the response, this could be do to a conflicting filepath or maybe the filename / path did not exist -> ", x)
				} else {
					defer file.Close()
					_, x := io.Copy(file, resp.Body)
					if x != nil {
						fmt.Println(c.REDHB, "<RR6> OFFLINE -> Module -> File I/O: Got error when trying to copy data to the given filepath or created filename, this filepath could be malformed, wrong, corrupted, or not saved correctly -> ", x)
					} else {
						fmt.Println("[*] File saved to -> ", directory)
					}
				}
			}
		}
	}
}
