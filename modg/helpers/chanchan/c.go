package chanchan

import (
	"fmt"
	"main/modg/requests"
	"strconv"

	lookups "main/modg/scripts/lookup"
)

func Worker(t chan requests.NIL, dns chan string, g chan []requests.SUB_RES, s string) {
	for n := range dns {
		r := requests.Lookup_final_DIALDEF(n, s)
		if len(r) > 0 {
			g <- r
		}
	}
	var NILRESP requests.NIL
	t <- NILRESP
}

func Worker_Lookup(ipa string, ac int, oschan chan bool) {
	fmt.Println("_____IP ADDR___________DOMAIN-LOCATION___________")
	for hostname := 0; hostname <= 255; hostname++ {
		fullIP := ipa + "." + strconv.Itoa(hostname)
		go lookups.Look_oc(fullIP, oschan)
		ac++
	}
	for ac > 0 {
		<-oschan
		ac--
	}
}
