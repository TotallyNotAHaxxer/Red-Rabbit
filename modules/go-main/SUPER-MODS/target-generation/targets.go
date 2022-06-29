package super_targets

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	SUPER_CONSTANTS "main/modules/go-main/SUPER-MODS/consts"
	SUPER_TYPESTRUCT "main/modules/go-main/SUPER-MODS/types"
)

func Generate_Burp_Config(filename, outfile string) {
	c, x := os.Open(filename)
	if x != nil {
		fmt.Println("<RR6> File error: Could not open the os using the standard syscall, something went wrong -> ", x)
		os.Exit(0)
	} else {
		defer c.Close()
		scanner := bufio.NewScanner(c)
		for scanner.Scan() {
			SUPER_CONSTANTS.RETURN_DATA_BURP_CONFIG_FUNC = append(SUPER_CONSTANTS.RETURN_DATA_BURP_CONFIG_FUNC, scanner.Text())
		}
		Make(outfile, SUPER_CONSTANTS.RETURN_DATA_BURP_CONFIG_FUNC)
	}
}

func Make(filename string, data []string) {
	for _, element := range data {
		slices := "^" + strings.ReplaceAll(strings.ReplaceAll(element, ".", "\\."), "*", ".*") + "$"
		URL_IF_443 := SUPER_TYPESTRUCT.Domain_data{
			Enabled:  true,
			File:     SUPER_CONSTANTS.PROTOCAL_ALL_STANDARD_FILEPATH,
			Host:     slices,
			Port:     SUPER_CONSTANTS.PROTCAL_PORT_443_DEFUALT_REGEX,
			Protocol: SUPER_CONSTANTS.PORT_443_DEFUALT_PROTOCAL}
		URL_IF_80 := SUPER_TYPESTRUCT.Domain_data{
			Enabled:  true,
			File:     SUPER_CONSTANTS.PROTOCAL_ALL_STANDARD_FILEPATH,
			Host:     slices,
			Port:     SUPER_CONSTANTS.PROTCAL_PORT_80_DEFUALT_REGEX,
			Protocol: SUPER_CONSTANTS.PORT_80_DEFUALT_PROTOCAL}
		SUPER_CONSTANTS.Domains = append(SUPER_CONSTANTS.Domains, URL_IF_443)
		SUPER_CONSTANTS.Domains = append(SUPER_CONSTANTS.Domains, URL_IF_80)
	}
	var results = SUPER_TYPESTRUCT.Burp_conf{
		Burp_target: SUPER_TYPESTRUCT.Burp_target{
			Burp_Scope: SUPER_TYPESTRUCT.Burp_Scope{
				Advanced_mode_burp: true,
				Exclude_data:       []SUPER_TYPESTRUCT.Domain_data{},
				Include_data:       SUPER_CONSTANTS.Domains}}}
	write_indent, _ := json.MarshalIndent(results, "", "	")
	_ = ioutil.WriteFile(filename, write_indent, fs.FileMode(SUPER_CONSTANTS.STANDARD_BURP_CONFIG_FILE_PREMISSIONS))
}
