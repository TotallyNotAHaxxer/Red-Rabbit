package Sub_domain

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	request_chan "main/modg/helpers/chanchan"
	requests "main/modg/requests"
)

var (
	results []requests.SUB_RES
)

func Run(wordlist, domain, server string) {
	s := make(chan string, 1000)
	g := make(chan []requests.SUB_RES)
	a := make(chan requests.NIL)
	h, x := os.Open(wordlist)
	if x != nil {
		fmt.Println("<RR6> File I/O SUPER: Could not read wordlist, open wordlist, read from wordlist, find wordlist, etc, got error -> ", x)
	} else {
		defer h.Close()
		scanner := bufio.NewScanner(h)
		for K := 0; K < 1000; K++ {
			go request_chan.Worker(a, s, g, server)
		}
		go func() {
			for value := range g {
				results = append(results, value...)
			}
			var e requests.NIL
			a <- e
		}()
		for scanner.Scan() {
			s <- fmt.Sprintf("%s.%s", scanner.Text(), domain)
		}
		close(s)
		for i := 0; i < 1000; i++ {
			<-a
		}
		close(g)
		<-a
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
		for _, r := range results {
			fmt.Fprintf(w, "\033[31m %s\t          | %s\n", r.HOST, r.IPA)
		}
		w.Flush()
	}
}

func FILTER_VALUES(input string) string {
	return strings.ReplaceAll(input, "&", "%26")
}
