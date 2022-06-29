package MODULES_AND_PATHNAMES_SIMPLE_TXT_DIG_LIV_NONEDIT_AUTO_TMPL_BASED_GENERATED_FILE

var filename string
var parser string

const (
	PATH                                               = "text/modules"
	BRUTE_FILENAME                                     = "/brute.txt"
	SEARCH_FILENAME                                    = "/search.txt"
	PING_FILENAME                                      = "/ping.txt"
	DUMP_FILENAME                                      = "/dump.txt"
	CHECK_FILENAME                                     = "/check.txt"
	ENCODE_FILENAME                                    = "/encode.txt"
	CRACK_FILENAME                                     = "/crack.txt"
	STEGONOGRAPHY_FILENAME                             = "/inject.txt"
	SNIFFERS_FILENAME                                  = "/sniff.txt"
	ETC_UTILS_TOOLS_FILENAME_MAIN_EXTRA_TOOLS          = "/etc.txt"
	XML_JSON_YAML_CONF_ETC_PARSER_EXTRA_UTILS_FILENAME = "/parse.txt"
	ENGINE_MODULE_FILENAME                             = "/engine.txt"
)

func Return_File(name string) string {
	switch name {
	case "brute":
		filename = BRUTE_FILENAME
	case "search":
		filename = SEARCH_FILENAME
	case "ping":
		filename = PING_FILENAME
	case "dump":
		filename = DUMP_FILENAME
	case "check":
		filename = CHECK_FILENAME
	case "encode":
		filename = ENCODE_FILENAME
	case "crack":
		filename = CRACK_FILENAME
	case "inject":
		filename = STEGONOGRAPHY_FILENAME
	case "sniff":
		filename = SNIFFERS_FILENAME
	case "utils":
		filename = ETC_UTILS_TOOLS_FILENAME_MAIN_EXTRA_TOOLS
	case "parse":
		filename = XML_JSON_YAML_CONF_ETC_PARSER_EXTRA_UTILS_FILENAME
	case "engine":
		filename = ENGINE_MODULE_FILENAME
	}
	parser = PATH + filename
	return parser
}
