package API_MAKE_GET_PUT_POST_RUN_START_REQUESTSMODULE2_RR6_MODG78_FILE_239_REQUEST_FOR_API_TIMES_AND_RUN_USING_CONSTANTS_USING2_TYPES_USING2_MODGCOLOR_INCLUDING_API_KEYS_JSON_XML_TYPE

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	NETWORK_WEB_THREE_POINT_ZERO_UTILITIES_CONSTANTS "main/API/constants"
	NETWORK_WEB_THREE_POINT_ZERO_API_UTILITIES_TYPE_STRUCTURES "main/API/response-structure"
	colors "main/modg/colors"
	e "main/modg/errors/serr"
	PHONE_SERVICE_GEO_CODE "main/modg/scripts/osint/phone"
	PHONE_SERVICE_API_MAPS "main/modules/go-main/SUPER-MAPS"

	"github.com/PuerkitoBio/goquery"
)

var (
	KEY_FILE  = "configuration/api_keys.json"
	WHOIS_KEY string
	c         = http.Client{}
	struct1   = NETWORK_WEB_THREE_POINT_ZERO_API_UTILITIES_TYPE_STRUCTURES.WHOIS_XML_API_RESPONSE_STRUCT1_MODG3{}
	struct2   = NETWORK_WEB_THREE_POINT_ZERO_API_UTILITIES_TYPE_STRUCTURES.API_LAYER_PHONE{}
)

func Print_data(color1, color2, color3, title_color, title, content_color, content, end_content_color string) {
	box_drawer := color1 + "[" + color2 + "+" + color3 + "]" + title_color + " " + title + " " + content_color + " -> " + content + end_content_color
	fmt.Println(box_drawer)
}

func Print_Int(color1, color2, color3, title_color, title, content_color, end_content_color string, content int) {
	fmt.Println(color1, "["+color2+"+"+color3+"]", title_color, " ", title, " ", content_color+" -> ", content, end_content_color)
}

func Fetch_Tag(url string, tag_ string) string {
	resp, x := http.Get(url)
	e.See_errorbased(x, colors.REDHB, "<RR6> Requests Module: Could not make a request to the URL -> ", false)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("<RR6> Requests Module API: Could not make a request to the given URL, GOT NON 200 RESPONSE STAT ->  %d %s\n", resp.StatusCode, resp.Status)
	} else {
		doc, x := goquery.NewDocumentFromReader(resp.Body)
		e.See_errorbased(x, colors.REDHB, "<RR6> Requests Module: Could not get a new document reader got error -> ", false)
		tag_data := doc.Find(tag_).Text()
		return tag_data
	}
	return "nothing for this tag"
}

func Make_Request_WHOIS_XML_API_FREE_SERVICE(website, username, password string) {
	url := fmt.Sprintf(NETWORK_WEB_THREE_POINT_ZERO_UTILITIES_CONSTANTS.WHOIS_API_URL_WHOIS, website, username, password)
	fmt.Println("|-| Using -> ", url)
	f, x := http.Get(url)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Body.Close()
		b, _ := ioutil.ReadAll(f.Body)
		if err := json.Unmarshal(b, &struct1); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON -> \033[32m", err)
			os.Exit(1)
		} else {
			fmt.Println(colors.GRN, "\n---------------[ Contact and Administration ] ---------------------")
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Contact Email         : ", colors.HIGH_BLUE, struct1.WhoisRecord.ContactEmail, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Organization          : ", colors.HIGH_BLUE, struct1.WhoisRecord.AdministrativeContact.Organization, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Country Name          : ", colors.HIGH_BLUE, struct1.WhoisRecord.AdministrativeContact.Country, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Country Code          : ", colors.HIGH_BLUE, struct1.WhoisRecord.AdministrativeContact.CountryCode, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "State                 : ", colors.HIGH_BLUE, struct1.WhoisRecord.AdministrativeContact.State, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Contact Country       : ", colors.HIGH_BLUE, struct1.WhoisRecord.TechnicalContact.Country, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Contact Organiz       : ", colors.HIGH_BLUE, struct1.WhoisRecord.TechnicalContact.Organization, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Contact State         : ", colors.HIGH_BLUE, struct1.WhoisRecord.AdministrativeContact.State, colors.RET_RED)
			fmt.Println(colors.GRN, "\n---------------[ Dates, Times, and Registration ] ---------------------")
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Created Date          : ", colors.HIGH_BLUE, struct1.WhoisRecord.CreatedDate, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Created Date normal   : ", colors.HIGH_BLUE, struct1.WhoisRecord.CreatedDateNormalized, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Audit Created Date    : ", colors.HIGH_BLUE, struct1.WhoisRecord.Audit.CreatedDate, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Audit Updated Date    : ", colors.HIGH_BLUE, struct1.WhoisRecord.Audit.UpdatedDate, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Expires Date          : ", colors.HIGH_BLUE, struct1.WhoisRecord.ExpiresDate, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Expires Date normal   : ", colors.HIGH_BLUE, struct1.WhoisRecord.ExpiresDateNormalized, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Updated Date          : ", colors.HIGH_BLUE, struct1.WhoisRecord.UpdatedDate, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Date Create  : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.CreatedDateNormalized, colors.RET_RED)
			Print_Int(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Date Create  : ", colors.HIGH_BLUE, colors.RET_RED, struct1.WhoisRecord.EstimatedDomainAge)
			fmt.Println(colors.GRN, "\n---------------[ Whois Record | Registry Information ] ---------------------")
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Date Create  : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.DomainName, colors.RET_RED)
			box_drawer := colors.RED + "[" + "\033[90m" + "+" + colors.RED + "]" + colors.MAG + " " + "Domain NS Hostnames" + " " + colors.HIGH_BLUE + " -> "

			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Date Create  : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.DomainName, colors.RET_RED)
			fmt.Print(box_drawer, " ")
			fmt.Print(struct1.WhoisRecord.RegistryData.NameServers.Ips...)
			fmt.Print("\n")
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Data Status           : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.Status, colors.RET_RED)
			Print_Int(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Date Create           : ", colors.HIGH_BLUE, colors.RET_RED, struct1.WhoisRecord.RegistryData.ParseCode)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Data Header           : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.Header, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Data Footer           : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.Footer, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Data audit create     : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.Audit.CreatedDate, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Data audit update     : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.Audit.UpdatedDate, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Data registrar Name   : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.RegistrarName, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Data registrar IANAID : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.RegistrarIANAID, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Data registrar Create : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.CreatedDateNormalized, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Data registrar Update : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.UpdatedDateNormalized, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Data Expires Date     : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.ExpiresDateNormalized, colors.RET_RED)
			Print_data(colors.RED, "\033[90m", colors.RED, colors.MAG, "Registry Data WHOIS server     : ", colors.HIGH_BLUE, struct1.WhoisRecord.RegistryData.WhoisServer, colors.RET_RED)

		}

	}
}

func Make_Request_API_LAYER_PHONE_INFO(phone_number, api_key string, is_us, is_be bool, dial_code string) {
	url := fmt.Sprintf(NETWORK_WEB_THREE_POINT_ZERO_UTILITIES_CONSTANTS.API_LAYER_PHONE_CHECK, api_key, phone_number)
	c := http.Client{}
	fmt.Println("|-| Using -> ", url)
	f, x := http.NewRequest("GET", url, nil)
	if x != nil {
		log.Fatal(x)
	} else {
		f.Header.Add("access_key", api_key)
		r, x := c.Do(f)
		if x != nil {
			log.Fatal(x)
		} else {
			defer r.Body.Close()
			b, _ := ioutil.ReadAll(r.Body)
			for k, v := range r.Header {
				fmt.Printf("%s %s\n", k, v)
			}
			if err := json.Unmarshal(b, &struct2); err != nil { // Parse []byte to go struct pointer
				fmt.Println("Can not unmarshal JSON -> \033[32m", err)
				os.Exit(1)
			} else {
				Watch_Dog_5 := PHONE_SERVICE_API_MAPS.Countr_Code_To_FULL_Country_Code_AKA_ISO2_To_ISO3_To_Country_Capital[struct2.CountryCode]
				Watch_Dog_4 := PHONE_SERVICE_API_MAPS.Countr_Code_To_FULL_Country_Code_AKA_ISO2_To_ISO3[struct2.CountryCode]
				Watch_Dog_3 := PHONE_SERVICE_API_MAPS.Countr_Code_To_Country_Name[struct2.CountryCode]
				Watch_Dog_2 := PHONE_SERVICE_API_MAPS.International_Country_Code_To_Currency[struct2.CountryCode]
				Watch_Dog_1 := PHONE_SERVICE_API_MAPS.Countr_Codes[struct2.CountryCode]
				fmt.Println(colors.RED, "|\033[90m+\033[31m|", colors.MAG, "Phone information -> Valid                     : ", colors.HIGH_BLUE, struct2.Valid)
				fmt.Println(colors.RED, "|\033[90m+\033[31m|", colors.MAG, "Phone information -> Number                    : ", colors.HIGH_BLUE, struct2.Number)
				fmt.Println(colors.RED, "|\033[90m+\033[31m|", colors.MAG, "Phone information -> Local Format              : ", colors.HIGH_BLUE, struct2.LocalFormat)
				fmt.Println(colors.RED, "|\033[90m+\033[31m|", colors.MAG, "Phone information -> World Format              : ", colors.HIGH_BLUE, struct2.InternationalFormat)
				fmt.Println(colors.RED, "|\033[90m+\033[31m|", colors.MAG, "Phone information -> Country Code              : ", colors.HIGH_BLUE, struct2.CountryCode)
				fmt.Println(colors.RED, "|\033[90m+\033[31m|", colors.MAG, "Phone information -> Country Prefix            : ", colors.HIGH_BLUE, struct2.CountryPrefix)
				fmt.Println(colors.RED, "|\033[90m+\033[31m|", colors.MAG, "Phone information -> Location                  : ", colors.HIGH_BLUE, struct2.Location)
				fmt.Println(colors.RED, "|\033[90m+\033[31m|", colors.MAG, "Phone information -> Carrier                   : ", colors.HIGH_BLUE, struct2.Carrier)
				fmt.Println(colors.RED, "|\033[90m+\033[31m|", colors.MAG, "Phone information -> Line type                 : ", colors.HIGH_BLUE, struct2.LineType)
				fmt.Printf("\033[31m |\033[90m+\033[31m|  \033[0;35mPhone Information -> Country code to prefix    : \033[38;5;21m  %s\n", Watch_Dog_1)
				fmt.Printf("\033[31m |\033[90m+\033[31m|  \033[0;35mPhone Information -> Country code to name      : \033[38;5;21m  %s\n", Watch_Dog_3)
				fmt.Printf("\033[31m |\033[90m+\033[31m|  \033[0;35mPhone Information -> Country code to currency  : \033[38;5;21m  %s\n", Watch_Dog_2)
				fmt.Printf("\033[31m |\033[90m+\033[31m|  \033[0;35mPhone Information -> Country ISO2 to ISO3      : \033[38;5;21m  %s\n", Watch_Dog_4)
				fmt.Printf("\033[31m |\033[90m+\033[31m|  \033[0;35mPhone Information -> Country ISO3 to capital   : \033[38;5;21m  %s\n", Watch_Dog_5)
				switch is_us {
				case true: // true means its us
					vals := PHONE_SERVICE_GEO_CODE.Return_US(dial_code)
					for _, k := range vals {
						fmt.Printf("\033[39m |\033[90m+\033[31m|  \033[0;35mPhone Information -> Extra Value               : \033[38;5;21m  %s\n", k)
					}
				}
				switch is_be {
				case true:
					vals := PHONE_SERVICE_GEO_CODE.Return_BE(dial_code)
					for _, k := range vals {
						fmt.Printf("\033[39m |\033[90m+\033[31m|  \033[0;35mPhone Information -> Extra Value          : %s\n", k)
					}
				}
			}
		}
	}
}
