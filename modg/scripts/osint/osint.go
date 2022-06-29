package osint_utilities

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	c "main/modg/colors"
	simple_errors "main/modg/errors/serr"
	requests "main/modg/requests"
	aws "main/modg/scripts/cloud/aws/aws-constants"
	aws_tools "main/modg/scripts/cloud/aws/aws-json"
	cosint "main/modg/scripts/osint/osintc"
	osint_constants "main/modg/scripts/osint/osintc"
	SUPER_MAPS "main/modules/go-main/SUPER-MAPS"
	SUPER_TYPES_2 "main/modules/go-main/SUPER_TYPES"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

type IP_API struct {
	Service struct {
		Name string `json:"name"`
	} `json:"service"`
	Parameters struct {
		InverseLookup struct {
		} `json:"inverse-lookup"`
		TypeFilters struct {
		} `json:"type-filters"`
		Flags struct {
			Flag []struct {
				Value string `json:"value"`
			} `json:"flag"`
		} `json:"flags"`
		QueryStrings struct {
			QueryString []struct {
				Value string `json:"value"`
			} `json:"query-string"`
		} `json:"query-strings"`
		Sources struct {
			Source []struct {
				ID string `json:"id"`
			} `json:"source"`
		} `json:"sources"`
	} `json:"parameters"`
	Objects struct {
		Object []struct {
			Type string `json:"type"`
			Link struct {
				Type string `json:"type"`
				Href string `json:"href"`
			} `json:"link"`
			Source struct {
				ID string `json:"id"`
			} `json:"source"`
			PrimaryKey struct {
				Attribute []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"attribute"`
			} `json:"primary-key"`
			Attributes struct {
				Attribute []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
					Link  struct {
						Type string `json:"type"`
						Href string `json:"href"`
					} `json:"link,omitempty"`
					ReferencedType string `json:"referenced-type,omitempty"`
				} `json:"attribute"`
			} `json:"attributes"`
		} `json:"object"`
	} `json:"objects"`
	TermsAndConditions struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"terms-and-conditions"`
	Version struct {
		Version   string    `json:"version"`
		Timestamp time.Time `json:"timestamp"`
		CommitID  string    `json:"commit-id"`
	} `json:"version"`
}

type WhoisResources struct {
	XMLName xml.Name `xml:"whois-resources"`
	Text    string   `xml:",chardata"`
	Xlink   string   `xml:"xlink,attr"`
	Objects struct {
		Text   string `xml:",chardata"`
		Object struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Link struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Href string `xml:"href,attr"`
			} `xml:"link"`
			Source struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"source"`
			PrimaryKey struct {
				Text      string `xml:",chardata"`
				Attribute struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					Value string `xml:"value,attr"`
				} `xml:"attribute"`
			} `xml:"primary-key"`
			Attributes struct {
				Text      string `xml:",chardata"`
				Attribute []struct {
					Text           string `xml:",chardata"`
					Name           string `xml:"name,attr"`
					Value          string `xml:"value,attr"`
					ReferencedType string `xml:"referenced-type,attr"`
					Link           struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
						Href string `xml:"href,attr"`
					} `xml:"link"`
				} `xml:"attribute"`
			} `xml:"attributes"`
		} `xml:"object"`
	} `xml:"objects"`
	TermsAndConditions struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Href string `xml:"href,attr"`
	} `xml:"terms-and-conditions"`
	Version struct {
		Text      string `xml:",chardata"`
		Version   string `xml:"version,attr"`
		Timestamp string `xml:"timestamp,attr"`
		CommitID  string `xml:"commit-id,attr"`
	} `xml:"version"`
}

type WhoisResources_2_AUTO struct {
	XMLName xml.Name `xml:"whois-resources"`
	Text    string   `xml:",chardata"`
	Xlink   string   `xml:"xlink,attr"`
	Objects struct {
		Text   string `xml:",chardata"`
		Object struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Link struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Href string `xml:"href,attr"`
			} `xml:"link"`
			Source struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"source"`
			PrimaryKey struct {
				Text      string `xml:",chardata"`
				Attribute []struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					Value string `xml:"value,attr"`
				} `xml:"attribute"`
			} `xml:"primary-key"`
			Attributes struct {
				Text      string `xml:",chardata"`
				Attribute []struct {
					Text           string `xml:",chardata"`
					Name           string `xml:"name,attr"`
					Value          string `xml:"value,attr"`
					ReferencedType string `xml:"referenced-type,attr"`
					Link           struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
						Href string `xml:"href,attr"`
					} `xml:"link"`
				} `xml:"attribute"`
			} `xml:"attributes"`
		} `xml:"object"`
	} `xml:"objects"`
	TermsAndConditions struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Href string `xml:"href,attr"`
	} `xml:"terms-and-conditions"`
	Version struct {
		Text      string `xml:",chardata"`
		Version   string `xml:"version,attr"`
		Timestamp string `xml:"timestamp,attr"`
		CommitID  string `xml:"commit-id,attr"`
	} `xml:"version"`
}

func is_addr_IPv4_o_IPv6(addr string) string {
	for j := 0; j < len(addr); j++ {
		switch addr[j] {
		case '.':
			return "IPv4 Detected"
		case ':':
			return "IPv6 Detected"
		}
	}
	return "[-] Address is neither IPv6 or IPv4???????\n"
}

func Test_if_IP_Is_proton(ip string) {
	ipa := strings.Trim(ip, " ")
	regexcheck, _ := regexp.Compile(cosint.Regex__Ipaddr)
	if !regexcheck.MatchString(ipa) {
		fmt.Println("[-] ERROR: <RR6> | REGEXP | Could not verify IP address must match regex string -> ", cosint.Regex__Ipaddr)
		fmt.Print("Exiting.....")
		os.Exit(0)
	} else {
		response, e := http.Get(cosint.IP_url)
		if e != nil {
			fmt.Println("[ --!-- ] Warning: Could not make a GET request to the URL got error -> \033[39m", e)
		} else {
			if response.StatusCode >= 200 {
				out, err_cr := os.Create("out.txt")
				if err_cr != nil {
					fmt.Println("[ --!-- ] Error: Could not crate the output file of the response body to look for address, got error -> ", err_cr)
					os.Exit(0)
				}
				defer out.Close()
				io.Copy(out, response.Body)
				reader, e := ioutil.ReadFile("out.txt")
				if e != nil {
					fmt.Println("[ --!-- ] Error: Could not open the file to read, this is fatal, script now can not access the IP range to search -> ", e)
					os.Exit(0)
				} else {
					if strings.Contains(string(reader), ip) {
						fmt.Printf("++++>>> RR6: IP [%s] Is apart of the proton mail VPN...\n", ip)
					} else {
						fmt.Printf("---->>> RR6: IP [%s], is NOT apart of the proton mail VPN...\n", ip)
					}
				}
			} else {
				fmt.Println("Response code was not 200......")
				os.Exit(0)
			}
		}
	}
}

func Test_if_email_Is_proton(emails string) {
	mail := strings.Trim(emails, " ")
	valid := regexp.MustCompile(cosint.Regex__Email)
	if !valid.MatchString(mail) {
		fmt.Println("[-] ERROR: <RR6> | REGEXP | Could not verify Email address must match regex string -> ", cosint.Regex__Email)
		fmt.Print("Exiting.....")
		os.Exit(0)
	} else {
		email_string := "https://api.protonmail.ch/pks/lookup?op=index&search=" + emails
		response, e := http.NewRequest("GET", email_string, nil)
		if e != nil {
			fmt.Println("[-] HTTP Error: Couuld not make a new GET request to the url, got error -> ", e)
			os.Exit(0)
		} else {
			response.Header.Set("Accept", "application/json")
			httpresp, e := cosint.Client.Do(response)
			if e != nil {
				fmt.Println("[-] Requests HTTP ERROR: Was not able to make a client request to the http URL, got error -> ", e)
				os.Exit(0)
			} else {
				defer httpresp.Body.Close()
				body, e := io.ReadAll(httpresp.Body)
				if e != nil {
					fmt.Println("[-] READER ERROR: Could not read the http response body from the HTTP client > ", e)
					os.Exit(0)
				} else {
					if strings.Contains(string(body), cosint.Verify_true_proton) {
						fmt.Println("found it!")
					} else {
						fmt.Println("no data.....")
					}
				}

			}
		}
	}

}

func Test_if_IP_Is_CLOUDFLARE(ip string) {
	if is_addr_IPv4_o_IPv6(ip) == "IPv4 Detected" {
		for _, c := range osint_constants.CloudFlareCIDR_IP4 {
			_, sa, _ := net.ParseCIDR(c)
			ipa, _, _ := net.ParseCIDR(ip + "/0")
			if sa.Contains(ipa) {
				fmt.Printf("\033[31m<RR6> Address Range | %s | Does match to list  | %s | > | %s | Is a possible Cloudflare IP range \n", c, ip, c)
			} else {
				fmt.Printf("\033[32m<RR6> Address Range | %s | Does not match with | %s | > | %s | Is not a possible cloudflare IP range\n", c, ip, c)
			}
		}
	} else {
		for _, a := range osint_constants.CloudFlareCIDR_IP6 {
			_, sc, _ := net.ParseCIDR(a)
			ipa, _, _ := net.ParseCIDR(ip + "/0")
			if sc.Contains(ipa) {
				fmt.Printf("\033[31m<RR6> Address Range | %s | Does match to list  | %s | > | %s | Is a possible Cloudflare IP range \n", a, ip, a)
			} else {
				fmt.Printf("\033[32m<RR6> Address Range | %s | Does not match with | %s | > | %s | Is not a possible cloudflare IP range\n", a, ip, a)
			}
		}
	}
}

func Test_if_IP_Is_AWS(ip string) {
	if is_addr_IPv4_o_IPv6(ip) == "IPv4 Detected" || is_addr_IPv4_o_IPv6(ip) == "IPv6 Detected" {
		response_aws, e := http.Get(aws.IP_URL)
		simple_errors.See_errorbased(e, "\033[39m", " Was not able to make a GET request to the given URL for AWS matching...", false)
		if response_aws.StatusCode == 200 {
			defer response_aws.Body.Close()
			prefix, e := ioutil.ReadAll(response_aws.Body)
			simple_errors.See_errorbased(e, "\033[39m", "Could not read the entire response body from the url", false)
			u, err := url.Parse(aws.IP_URL)
			simple_errors.See_errorbased(err, "\033[39m", "Could not properly parse the AWS url, error -> ", false)
			domain := strings.Split(u.Hostname(), ".")
			d := domain[len(domain)-2] + "." + domain[len(domain)-1]
			filename := d + ".json"
			requests.Write(filename, string(prefix))
		} else {
			fmt.Println("<RR6> Requests module: Could not make a good decent 200 status code request to the url")
		}
	}
	aws_tools.Output_test(ip)
}

func Test_Public_addr() {
	uli := "https://api.ipify.org?format=text"
	response, e := http.Get(uli)
	simple_errors.See_errorbased(e, "\033[31m", "Could not make a get request to the API url", false)
	defer response.Body.Close()
	ip, e := ioutil.ReadAll(response.Body)
	simple_errors.See_errorbased(e, "\033[31m", "Could not read the entire response body from the url", false)
	fmt.Printf("[+] Address ~>  | %s | \n", ip)
}

func Test_Username_Site(username string, wordlist string) string {
	contents, err := os.Open(wordlist)
	if err != nil {
		log.Fatal("<RR6> OSINT Module: Could not open the wordlist for the website gathering.. -> ", err)
	} else {
		scanner := bufio.NewScanner(contents)
		for scanner.Scan() {
			parser := scanner.Text() + username
			response, err := http.Get(parser)
			if err != nil {
				log.Fatal("<RR6> OSINT Module: Could not make a proper request to the URL and host got error, -> ", err)
			} else {
				if response.StatusCode == 200 {
					fmt.Println(c.HIGH_BLUE)
					fmt.Printf("[+] Found username < %s > IS  on website < %s > \t| Stat code < 200 > \n", username, parser)
					fmt.Println(c.RET_RED)
				} else {
					fmt.Println(c.REDHB)
					fmt.Printf(c.REDHB, "[-] Error username < %s > NOT on website < %s > \t| Stat code < %v > \n", username, parser, response.StatusCode)
					fmt.Println(c.RET_RED)
				}
			}
		}
	}
	return "Error?"
}

func Test_IP_Address_Location(ip string) {
	parser := fmt.Sprintf("https://ipapi.co/%s/json/", ip)
	return_value, _, e := requests.Create_GET_Body(parser)
	if e != nil {
		fmt.Println("<RR6> OSINT Module talking to <- Requests module Caused -> error ->  ", e)
	} else {
		fmt.Println(return_value)
	}
}

func Get_IP(ip string) {
	link := "https://ipinfo.io/" + ip + "/json"
	f, x := http.Get(link)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> OSINT Module: Could not make a GET request to the URL, got error -> ", x)
	} else {
		defer f.Body.Close()
		b, x := ioutil.ReadAll(f.Body)
		if x != nil {
			fmt.Println(c.REDHB, "<RR6> OSINT Module: Was unable to properly read the response body -> ", x)
		} else {
			var results SUPER_TYPES_2.IP_Information
			if err := json.Unmarshal(b, &results); err != nil {
				fmt.Println(c.REDHB, "<RR6> OSINT Module: Was unable to unmarshal the JSON file -> ", x)
			} else {
				Watch_Dog_5 := SUPER_MAPS.Countr_Code_To_FULL_Country_Code_AKA_ISO2_To_ISO3_To_Country_Capital[results.Country]
				Watch_Dog_4 := SUPER_MAPS.Countr_Code_To_FULL_Country_Code_AKA_ISO2_To_ISO3[results.Country]
				Watch_Dog_3 := SUPER_MAPS.Countr_Code_To_Country_Name[results.Country]
				Watch_Dog_2 := SUPER_MAPS.International_Country_Code_To_Currency[results.Country]
				Watch_Dog_1 := SUPER_MAPS.Countr_Codes[results.Country]
				Watch_Dog_9 := SUPER_MAPS.COUNTR_NAME_TO_CONTINENT_NAME[results.Country]
				fmt.Println("-----------------------------------")
				fmt.Println(c.HIGH_BLUE)
				fmt.Printf("IPA Value | IP                                   | %s\n", results.IP)
				fmt.Printf("IPA Value | Hostname                             | %s\n", results.Hostname)
				fmt.Printf("IPA Value | Anycast?                             | %v\n", results.Anycast)
				fmt.Printf("IPA Value | City                                 | %v\n", results.City)
				fmt.Printf("IPA Value | Region / State                       | %s\n", results.Region)
				fmt.Printf("IPA Value | Country                              | %s\n", results.Country)
				fmt.Printf("IPA Value | Country code to prefix               | %s\n", Watch_Dog_1)
				fmt.Printf("IPA Value | Country code to currency             | %s\n", Watch_Dog_2)
				fmt.Printf("IPA Value | Country Code to Country name         | %s\n", Watch_Dog_3)
				fmt.Printf("IPA Value | Country ISO2 to Country ISO3         | %s\n", Watch_Dog_4)
				fmt.Printf("IPA Value | Country ISO2 to Country ISO3 Capital | %s\n", Watch_Dog_5)
				fmt.Printf("IPA Value | Country ISO2 to Continent ISO3       | %s\n", Watch_Dog_9)
				fmt.Printf("IPA Value | Country Timezone                     | %s\n", results.Timezone)
				fmt.Printf("IPA Value | GPS Lattitude and Longitude          | %s\n", results.Org)
				fmt.Printf("IPA Value | Postal Code                          | %s\n", results.Postal)
			}
		}
	}
}

func Trace_IP_Secondary_API(ip string) {
	ink := "https://rest.db.ripe.net/search.json?query-string=73.205.242.99&flags=no-referenced&flags=no-irt&source=RIPE"

	client := http.Client{}
	req, x := http.NewRequest("GET", ink, nil)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> API - Requests - filter - parser - reader - helper - request creation: Got error when making a new methodized GET request with the URL parsed -> ", x)
	} else {
		resp, X := client.Do(req)
		if X != nil {
			fmt.Println(c.REDHB, "<RR6> API - Requests - Clients - channels - readers - loggers - errors: Got error when making the HTTP client do the methodized GET request to the target URL -> ", X)
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				b, x := ioutil.ReadAll(resp.Body)
				if x != nil {
					fmt.Println("<RR6> API/I/O -> Requests -> Readers -> I/O -> FILES Modules caused a fatal error -> Got error when trying to decode the GET response body, please follow or try to fix this error -> ", x)
				} else {
					var res IP_API
					if X := json.Unmarshal(b, &res); X != nil {
						fmt.Println("<RR6> JSON/XML/GET/IO/REQUESTS/FORM -> Got error when trying to unmarshal the JSON response body, this could be due to a corrupt data tree or corrupt response body -> ", X)
					} else {
						//fmt.Println(string(b))
						for i := range b {
							fmt.Println("|+| Name     -> ", res.Objects.Object[i].Attributes.Attribute[i].Name)
							fmt.Println("|+| Value    ->", res.Objects.Object[i].Attributes.Attribute[i].Value)
						}
					}
				}
			}
		}
	}
}
