package Engine_Shodan

type T struct {
	Matches []Hostname `json:"matches"`
}

type Host_GEO_Info struct {
	City         string  `json:"city"`
	RegionCode   string  `json:"region_code"`
	AreaCode     int     `json:"area_code"`
	Longitude    float32 `json:"longitude"`
	CountryCode3 string  `json:"country_code3"`
	CountryName  string  `json:"country_name"`
	PostalCode   string  `json:"postal_code"`
	DMACode      int     `json:"dma_code"`
	CountryCode  string  `json:"country_code"`
	Latitude     float32 `json:"latitude"`
}

type Hostname struct {
	OS        string        `json:"os"`
	Timestamp string        `json:"timestamp"`
	ISP       string        `json:"isp"`
	ASN       string        `json:"asn"`
	Hostnames []string      `json:"hostnames"`
	Location  Host_GEO_Info `json:"location"`
	IP        int64         `json:"ip"`
	Domains   []string      `json:"domains"`
	Org       string        `json:"org"`
	Data      string        `json:"data"`
	Port      int           `json:"port"`
	IPString  string        `json:"ip_str"`
}

var Results T
