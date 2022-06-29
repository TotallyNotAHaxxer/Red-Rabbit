package YML_Config_Parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type SSH_settings struct {
	Username       string `yaml:"Username"`
	Password       string `yaml:"Password"`
	PrivateKeyFile string `yaml:"Private_Key_File"`
	Port           string `yaml:"Port"`
	Host           string `yaml:"Host"`
}

func trim(data string) string {
	return strings.Trim(fmt.Sprint(data), "{}")
}

func Parse_File_ssh(filename string) {
	config, x := ioutil.ReadFile(filename)
	if x != nil {
		fmt.Println("<RR6> File I/O: Could not read configuration file for SSH config....FATAL", x)
		os.Exit(1)
	} else {
		data := make(map[string]SSH_settings)
		x2 := yaml.Unmarshal(config, data)
		if x2 != nil {
			fmt.Println("<RR6> YAML I/O: Could not read the config.yaml file....FATAL -> ", x2)
			os.Exit(1)
		} else {
			for _, d := range data {
				host := d.Host
				pass := d.Password
				port := d.Port
				user := d.Username
				file := d.PrivateKeyFile
				fmt.Println("_______SSH Server Configuration______")
				fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Set Host -> ", trim(host))
				fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Set Pass -> ", trim(pass))
				fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Set Port -> ", trim(port))
				fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Set User -> ", trim(user))
				fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Set key  -> ", trim(file))
			}
		}
	}
}

func Parse_return(filename string) (string, string, string, string, string) {
	config, x := ioutil.ReadFile(filename)
	if x != nil {
		fmt.Println("<RR6> File I/O: Could not read configuration file for SSH config....FATAL", x)
		os.Exit(1)
	} else {
		data := make(map[string]SSH_settings)
		x2 := yaml.Unmarshal(config, data)
		if x2 != nil {
			fmt.Println("<RR6> YAML I/O: Could not read the config.yaml file....FATAL -> ", x2)
			os.Exit(1)
		} else {
			for _, d := range data {
				host := d.Host
				pass := d.Password
				port := d.Port
				user := d.Username
				file := d.PrivateKeyFile
				return host, pass, port, user, file
			}
		}
	}
	return "", "", "", "", ""
}
