package Super_types

// API key config
// ALL TYPES WERE PRE GENERATED FROM EXAMPLE FILES SUCH AS THE BURP TARGET FILE, JSON TO GO WEBSITE AND SOURCE https://mholt.github.io/json-to-go/
// TYPE NAMES WERE CHANGED DUE TO CONFLICTING TYPES

type Super_config struct {
	Api_key string `yaml:"Key"`
	Knoxss  string `yaml:" XSS_KEY"`
}

type Super_Port_config struct {
	Hostname   string `yaml:"Hostname"`
	Port_Start string `yaml:"Port_Start"`
	Port_End   string `yaml:"Port_End"`
}

// domain data

type Domain_data struct {
	Enabled  bool   `json:"enabled"`
	File     string `json:"file"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
}

// burp file configuation

type Burp_conf struct {
	Burp_target Burp_target `json:"target"`
}

type Burp_target struct {
	Burp_Scope Burp_Scope `json:"scope"`
}

type Burp_Scope struct {
	Advanced_mode_burp bool          `json:"advanced_mode"`
	Exclude_data       []Domain_data `json:"exclude"`
	Include_data       []Domain_data `json:"include"`
}

//
