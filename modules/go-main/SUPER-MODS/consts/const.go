package Super_constants

import SUPER_TYPESTRUCT "main/modules/go-main/SUPER-MODS/types"

var (
	API_KEY                               string
	MAX_CHARACTER_COUNT_SHODAN            = 32
	MAX_CHARACTER_COUNT_KNOXSS            = 36
	STANDARD_BURP_CONFIGURATION_FILE_NAME = "targetscope.json"
	RETURN_DATA_BURP_CONFIG_FUNC          []string
	STANDARD_BURP_CONFIG_FILE_PREMISSIONS = 0644
	Domains                               []SUPER_TYPESTRUCT.Domain_data
	PORT_80_DEFUALT_PROTOCAL              = "http"
	PORT_443_DEFUALT_PROTOCAL             = "https"
	PROTCAL_PORT_80_DEFUALT_REGEX         = "^80$"
	PROTCAL_PORT_443_DEFUALT_REGEX        = "^443$"
	PROTOCAL_ALL_STANDARD_FILEPATH        = "^/.*"
	SUPER_KEY_STR_BOOL                    = make(map[string]bool)
	SUPER_DATA_STRING                     []string
)
