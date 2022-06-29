package conf

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	super_types "main/modules/go-main/SUPER-MODS/types"

	"gopkg.in/yaml.v2"
)

// literally the same thing different shit
func Open_and_return(filename string) (string, string, string, error) {
	yfile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]super_types.Super_Port_config)
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	for _, l := range data {
		PortEnd := strings.Trim(fmt.Sprint(l.Port_End), "{}")
		PortStart := strings.Trim(fmt.Sprint(l.Port_Start), "{}")
		Hostname := strings.Trim(fmt.Sprint(l.Hostname), "{}")
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting -> Scanning      | ", Hostname)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting -> Scanning from | ", PortStart)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting -> Scanning to   | ", PortEnd)
		return Hostname, PortStart, PortEnd, nil
	}
	return "", "", "", nil
}
