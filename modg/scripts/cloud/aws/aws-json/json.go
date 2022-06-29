/*

Package for parsing values and files from aws servers such as ip ranges, server names, bucket names, etc

*/

package awsj

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type AWS_Addresses struct {
	AWS_Addr []Prefixes `json:"prefixes"`
}

type Prefixes struct {
	Prefix       string `json:"ip_prefix"`
	Region       string `json:"region"`
	Service      string `json:"service"`
	Net_Border_g string `json:"network_border_group"`
}

func Read(file_json string) (*os.File, error) {
	filename, e := os.Open(file_json)
	if e != nil {
		log.Fatal("<RR6> AWS Package: Could not open AWS file / json structure, got error -> ", e)
	} else {
		defer filename.Close()
	}
	return filename, nil
}

func Output_test(ip string) {
	j, e := os.Open("json/amazonaws.com.json")
	if e != nil {
		log.Fatal(e)
	}
	defer j.Close()
	valueoftypebyte, _ := ioutil.ReadAll(j)
	var values AWS_Addresses
	json.Unmarshal(valueoftypebyte, &values)
	for i := 0; i < len(values.AWS_Addr); i++ {
		if strings.Compare(ip, values.AWS_Addr[i].Prefix) == 0 {
			fmt.Println(values.AWS_Addr[i].Prefix)
			fmt.Println("\033[32m<RR6> AWS Cloud module: Found matching address in range prefixes |")
			fmt.Println("--------------------------------------------------------------------------")
			fmt.Println("Range                |-> ", values.AWS_Addr[i].Prefix)
			fmt.Println("Range Region         |-> ", values.AWS_Addr[i].Region)
			fmt.Println("Range Service        |-> ", values.AWS_Addr[i].Service)
			fmt.Println("Range Network Border |-> ", values.AWS_Addr[i].Net_Border_g)
		}
	}
}
