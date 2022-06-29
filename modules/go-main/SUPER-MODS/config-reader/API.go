/*

Apart of the super mods kit, the api or super_mods_api will be the package to read all config file API keys, parse them and find the data about the account

to check if this is still a valid key, shodan, knoxss are the only API keys currently in use. why is this file named SUPER-MODS, super mods means that this is

an essential mod to have in red rabbit in order for it to work, sure all modules inside of the modg file are needed but the suprt mod file will mainly be

logging, parsing, data, databases, web ui's, encryption, etc meaning this again is a VERY ESSENTIAL package to have tinside of red rabbit, without it

you could not execute API's, some SSH tunnlers, etc since main config is parsed by this file or other super modules ion the red rabbit filepath

*/

package super_mods_API

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	super_constants "main/modules/go-main/SUPER-MODS/consts"
	super_types "main/modules/go-main/SUPER-MODS/types"

	"gopkg.in/yaml.v2"
)

func Reader(config_file_path string) string {
	yfile, err := ioutil.ReadFile("config/api_key.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]super_types.Super_config)
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	for _, l := range data {
		fmt.Println(strings.Trim(fmt.Sprint(l), "{}"))
		a := Test_label_API((strings.Trim(fmt.Sprint(l), "{}")))
		fmt.Println(a)
	}
	return ""
}

func Return_KNOXSS(config_file_path string) string {
	yfile, err := ioutil.ReadFile("config/xssapi/api.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]super_types.Super_config)
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	for _, l := range data {
		a := strings.Trim(fmt.Sprint(l.Api_key), "{}")
		return a
	}
	return ""
}

func Test_label_API(key string) string {
	if len(key) == super_constants.MAX_CHARACTER_COUNT_SHODAN {
		super_constants.API_KEY = "Shodan API Key"
		fmt.Println("Key detected as Shodan (len) based:", super_constants.API_KEY)
	}
	if len(key) == super_constants.MAX_CHARACTER_COUNT_KNOXSS {
		super_constants.API_KEY = "KNOXSS API Key"
		fmt.Println("Key detected as KNOXSS (len) based: ", super_constants.API_KEY)
	}
	if len(key) != super_constants.MAX_CHARACTER_COUNT_KNOXSS || len(key) != super_constants.MAX_CHARACTER_COUNT_SHODAN {
		fmt.Println("[!!!!!] API KEY IS NOT KNOXSS OR SHODAN API KEYS!!!! ERROR...........")
		fmt.Print("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Please head over to config/api_key.yaml to fix this issue and input a valid API key")
		os.Exit(1)
	}
	return ""
}
