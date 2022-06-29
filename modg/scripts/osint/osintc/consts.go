package osint_constants

import "net/http"

const (
	IP_url              = "https://api.protonmail.ch/vpn/logicals"
	EMAIL_url           = "https://api.protonmail.ch/pks/lookup?op=get&search="
	IP_Trace_url        = ""
	Regex__Email        = `([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+)`
	Regex__Ipaddr       = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	Verify_true_proton  = "info:1:1"
	Verify_false_proton = "info:1:0"
)

var (
	Client = &http.Client{}
)

var CloudFlareCIDR_IP6 = [7]string{
	"2400:cb00::/32",
	"2606:4700::/32",
	"2803:f800::/32",
	"2405:b500::/32",
	"2405:8100::/32",
	"2a06:98c0::/29",
	"2c0f:f248::/32",
}

var CloudFlareCIDR_IP4 = [14]string{
	"197.234.240.0/22",
	"198.41.128.0/17",
	"162.158.0.0/15",
	"104.16.0.0/12",
	"172.64.0.0/13",
	"131.0.72.0/22",
	"173.245.48.0/20",
	"103.21.244.0/22",
	"103.22.200.0/22",
	"103.31.4.0/22",
	"141.101.64.0/18",
	"108.162.192.0/18",
	"190.93.240.0/20",
	"188.114.96.0/20",
}
