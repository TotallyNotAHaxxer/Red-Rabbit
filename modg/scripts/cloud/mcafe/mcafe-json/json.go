package __mcafe__

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	constants "main/modg/scripts/cloud/mcafe/mcafe-constants"
	"os"
	"regexp"
	"strings"
)

type Mcafe_Addresses struct {
	Mcafe_Addresses []Prefixes `json:"prefixes"`
}

type Prefixes struct {
	Date    string `json:"date"`
	Inbound bool   `json:"inbound"`
	Prefix  string `json:"ip"`
}

func Test_addr(ip string) bool {
	ipa := strings.Trim(ip, " ")
	re, _ := regexp.Compile(constants.CIDR_Regex)
	return re.MatchString(ipa)
}

func Output_test(ip string) {
	j, e := os.Open("json/mcafe.json")
	if e != nil {
		log.Fatal(e)
	}
	defer j.Close()
	valueoftypebyte, _ := ioutil.ReadAll(j)
	var Mcafe_results Mcafe_Addresses
	json.Unmarshal(valueoftypebyte, &Mcafe_results)
	for i := 0; i < len(Mcafe_results.Mcafe_Addresses); i++ {
		if strings.Compare(ip, Mcafe_results.Mcafe_Addresses[i].Prefix) == 1 {
			fmt.Printf("\033[4;37m(Cloud Lookup)\033[4;39m \033[39m<-> \033[4;34m(Mcafe)\033[4;39m\033[39m IP |%s| Is NOT apart of the MCAFE cloud IP list.. \n", ip)
		}
	}
}

func Dev(ip, results string) bool {
	switch strings.Compare(results, ip) {
	case 1:
		return true
	default:
		return false
	}
}
