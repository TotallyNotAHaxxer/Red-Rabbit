package URL_Scanner

import (
	"net/url"
	"regexp"

	super_requests "main/modg/requests"
	super_data "main/modules/go-main/xml/types"
)

func Check_Check(input []string, conf super_data.Burp_Config) []string {
	var result []string
	for _, item := range input {
		if super_requests.GET_URL_PROTO(item) == "" {
			continue
		}
		u, err := url.Parse(item)
		if err != nil {
			continue
		}
		var excludedItem = false
		for _, excluded := range conf.Target_Scope.Scope.Exclude {
			rHost, _ := regexp.Compile(excluded.Host)
			rFile, _ := regexp.Compile(excluded.File)
			if rHost.MatchString(u.Host) && rFile.MatchString(u.Path) {
				excludedItem = true
				break
			}
		}
		if excludedItem {
			continue
		}
		for _, included := range conf.Target_Scope.Scope.Include {
			rHost, _ := regexp.Compile(included.Host)
			rFile, _ := regexp.Compile(included.File)
			if rHost.MatchString(u.Host) && rFile.MatchString(u.Path) {
				result = append(result, item)
				break
			}
		}
	}
	return result
}
