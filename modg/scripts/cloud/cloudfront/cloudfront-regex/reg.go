package cloudfront_regex

import (
	cloudfront_regex_constants "main/modg/scripts/cloud/cloudfront/cloudfront-constants"
	"regexp"
	"strings"
)

func Test(ip string) bool {
	ipa := strings.Trim(ip, " ")
	reg, _ := regexp.Compile(cloudfront_regex_constants.CIDR_Regex)
	if !reg.MatchString(ipa) {
		return false
	}
	return true
}
