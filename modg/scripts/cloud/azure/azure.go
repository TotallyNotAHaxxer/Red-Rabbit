package azure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	color "main/modg/colors"
)

// IP types

type IP_Preifx_check struct {
	ChangeNumber int    `json:"changeNumber"`
	Cloud        string `json:"cloud"`
	Values       []struct {
		Name       string `json:"name"`
		ID         string `json:"id"`
		Properties struct {
			ChangeNumber    int      `json:"changeNumber"`
			Region          string   `json:"region"`
			RegionID        int      `json:"regionId"`
			Platform        string   `json:"platform"`
			SystemService   string   `json:"systemService"`
			AddressPrefixes []string `json:"addressPrefixes"`
			NetworkFeatures []string `json:"networkFeatures"`
		} `json:"properties"`
	} `json:"values"`
}

func Check_ip(filename string, ip string) {
	azur_db, x := os.Open(filename)
	if x != nil {
		fmt.Println(color.REDHB, "<RR6> Azure cloud: Could not open AZURE struct file....")
	} else {
		defer azur_db.Close()
		v, _ := ioutil.ReadAll(azur_db)
		var prefixes IP_Preifx_check
		json.Unmarshal(v, &prefixes)
		for k := 0; k < len(prefixes.Values); k++ {
			for _, l := range prefixes.Values[k].Properties.AddressPrefixes {
				if strings.Compare(l, ip) == 0 {
					fmt.Println(color.UWHT, "\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Address | ", ip, " |  Has been matched to | ", l, " | ", color.RET_RED)
					fmt.Println(color.UWHT, "\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Address is an AZURE IP address", color.RET_RED)
				}
			}
		}
	}
}
