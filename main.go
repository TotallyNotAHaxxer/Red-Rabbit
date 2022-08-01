package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	v "main/modg/colors"
	"main/modg/constants"
	opc "main/modg/copt"
	helpers "main/modg/helpers"
	sig "main/modg/sig"
	uio "main/modg/switch"
	system "main/modg/system-runscript"
	steglib "main/modules/go-main/0x0001/lib"
	stegutils "main/modules/go-main/0x0001/utils"
	MODULES_AND_PATHNAMES_SIMPLE_TXT_DIG_LIV_NONEDIT_AUTO_TMPL_BASED_GENERATED_FILE "main/text/modules"

	ec "main/modg/warnings"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

var (
	RETURN_ALL_COLOR = "\033[39m\033[49m"
	flags            = pflag.FlagSet{SortFlags: false}
	rr6f             opc.RR6_options
	hw_chan          = make(chan bool)
	ac_threads       = 0
	mt               = 100
	png              steglib.MetaChunk
)

func SSB(Screen_rotation string) {
	if Screen_rotation == "landscape" || Screen_rotation == "Landscape" {
		newfilepath, err := constants.Parse_filepath(constants.Team_logo_rr)
		ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		Banner(newfilepath, v.RED)
	}
	if Screen_rotation == "verticle" {
		newfilepath, err := constants.Parse_filepath(constants.Verticle_banner)
		ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		Banner(newfilepath, v.RED)
	}
	if Screen_rotation == "" {
		newfilepath, err := constants.Parse_filepath(constants.Verticle_banner)
		ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		Banner(newfilepath, v.RED)
	} else if Screen_rotation == "shark" {
		newfilepath, err := constants.Parse_filepath(constants.Shark)
		ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		Banner(newfilepath, v.RED)
	} else if Screen_rotation == "none" || Screen_rotation == "no" {
		fmt.Println(constants.Clear_hex)
	}
}

// make help maps
var Module_help_names = map[string]string{
	"help search": "search",
	"help brute":  "brute",
	"help ping":   "ping",
	"help dump":   "dump",
	"help check":  "check",
	"help crack":  "crack",
	"help inject": "inject",
	"help sniff":  "sniff",
	"help utils":  "utils",
	"help parse":  "parse",
	"help encode": "encode",
	"help engine": "engine",
}

func Banner(file, color string) {
	content, err := ioutil.ReadFile(file)
	ec.Ce(err, v.RED, "Could not read file", 1)
	fmt.Println(constants.Clear_hex, color, string(content))
}

func open_out(filename, color_out string) {
	s, x := os.Open(filename)
	if x != nil {
		log.Fatal(x)
	} else {
		defer s.Close()
		scanner := bufio.NewScanner(s)
		for scanner.Scan() {
			fmt.Println(color_out, scanner.Text())
		}
	}
}

func main() {
	system.Check_parser()
	fmt.Print(v.RET_RED)
	helpers.SSB(rr6f.Screen_rotation)
	terminal("RR6> ", v.RED)

}

func terminal(input_statement, color string) {
	if rr6f.H {
		open_out("text/help/flags.txt", v.HIGH_BLUE)
		os.Exit(0)
	}
	if rr6f.Extreme_Help {
		fmt.Println("\033[39m")
		f, x := os.Open("text/menus/advanced.txt")
		if x != nil {
			fmt.Println("<RR6> File -> I/O -> ERROR -> FATAL -> Could not locate text file this could be a MASSIVE issue, especially if this directory does not exist -> ", x)
			os.Exit(1)
		} else {
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				fmt.Println(string(scanner.Text()))
			}
		}
		os.Exit(0)
	}
	if rr6f.HH {
		open_out("text/help/flag-commands.txt", v.HIGH_BLUE)
		os.Exit(0)
	}
	if rr6f.HELP {
		open_out("text/help/flags-nocmd.txt", v.HIGH_BLUE)
		os.Exit(0)
	}
	ac_threads++
	rr6t := bufio.NewReader(os.Stdin)
	fmt.Print(color, "RR6> ")
	for {
		go sig.Handelreturncon(make(chan os.Signal, 1))
		t, _ := rr6t.ReadString('\n')
		t = strings.Replace(t, "\n", "", -1)
		if strings.Compare("inject png", t) == 0 {
			if rr6f.Input == "" || rr6f.Output == "" || rr6f.Payload == "" {
				fmt.Println("Sorry but you cant run this command right now, try using --payload, -i, -o, --offset to specify a payload, png image to inject, image, output file, and the injectable offset")
				terminal("RR6> ", v.RED)
			} else {
				if rr6f.Payload != "" && rr6f.Image_offset != "" && rr6f.Input != "" && rr6f.Output != "" {
					fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: All flags are correct")
					var png steglib.MetaChunk
					t, e := os.Open(rr6f.Input)
					ec.Warning_simple("<RR6> UIO Module: Could not make chan to update or open image", v.REDHB, e)
					defer t.Close()
					r, e := stegutils.PreProcessImage(t)
					ec.Warning_simple("<RR6> UIO Module: Could not process image", v.REDHB, e)
					png.Run_Functions(true, r, &rr6f, false)
					terminal("RR6> ", v.RED)
				}
			}
		}
		if strings.Compare("dump png meta", t) == 0 {
			fmt.Println(v.BLKHB, "WARNING: For some reason the meta png module will only work once unless the script is restarted, this can not be used in one session", v.RET_RED)
			t, e := os.Open(rr6f.Input)
			ec.Warning_simple("<RR6> UIO Module: Could not make chan to update or open image", v.REDHB, e)
			defer t.Close()
			r, e := stegutils.PreProcessImage(t)
			ec.Warning_simple("<RR6> UIO Module: Could not process image", v.REDHB, e)
			png.Run_Functions(false, r, &opc.RR6_options{}, true)
			terminal("RR6> ", v.RED)
		} else {
			if val, ok := Module_help_names[t]; ok {
				a := MODULES_AND_PATHNAMES_SIMPLE_TXT_DIG_LIV_NONEDIT_AUTO_TMPL_BASED_GENERATED_FILE.Return_File(val)
				open_out(a, v.HIGH_BLUE)
			}
		}
		if t != "" {
			uio.M_TTY(t, &rr6f)
		} else {
			terminal("RR6> ", v.RED)
		}
		terminal("RR6> ", v.RED)
		fmt.Println(RETURN_ALL_COLOR)
	}
}

func init() {

	flags.StringVar(&rr6f.Pheight, "ph", "1200", "Set the pixel height for stego functions")
	flags.StringVar(&rr6f.Pwidth, "pw", "800", "Set the pixel width for stego functions")
	flags.StringVarP(&rr6f.Jpgchunk, "jpgc", "p", "COM", "Set a location or chunk to inject data into a JPG image")
	flags.StringVarP(&rr6f.Screen_rotation, "reso", "v", "verticle", "Set the display rotation type <Verticle|Landscape|> For banner type")
	flags.StringVarP(&rr6f.Input, "input", "i", "", "Path to the original image file")
	flags.StringVarP(&rr6f.Output, "output", "o", "", "Path to output the new image file")
	flags.BoolVarP(&rr6f.S_M, "meta", "m", false, "Display the actual image meta details")
	flags.BoolVarP(&rr6f.Suppress, "suppress", "s", false, "Suppress the chunk hex data (can be large)")
	flags.StringVar(&rr6f.Image_offset, "offset", "", "The offset location to initiate data injection")
	flags.BoolVar(&rr6f.Inject, "inject", false, "Enable this to inject data at the offset location specified")
	flags.StringVar(&rr6f.Payload, "payload", "", "Payload is data that will be read as a byte stream")
	flags.StringVar(&rr6f.Type, "type", "rNDm", "Type is the name of the Chunk header to inject")
	flags.StringVar(&rr6f.Key, "key", "", "The enryption key for payload")
	flags.BoolVarP(&rr6f.H, "he", "h", false, "General help on flags and help commands")
	flags.BoolVar(&rr6f.HH, "hh", false, "General help on help commands")
	flags.BoolVar(&rr6f.HELP, "help", false, "General help on flags")
	flags.BoolVar(&rr6f.Extreme_Help, "ehelp", false, "Help on everything such as command flags, commands, help commands, help lists, and examples")
	flags.BoolVar(&rr6f.Payload_Encode, "encode", false, "XOR encode the payload")
	flags.BoolVar(&rr6f.Payload_Decode, "decode", false, "XOR decode the payload")
	flags.BoolVar(&rr6f.Extract_ZIP, "JPGEXTRACT", false, "Scan for ZIP files and extract them from images | JPEG FORMAT ONLY")
	flags.BoolVar(&rr6f.INJECT_ZIP, "JPGINJECT", false, " Start / Activate ZIP file injection")
	flags.StringVar(&rr6f.Filepath_general, "filepath", "", "path to the ZIP FILE")
	flags.BoolVar(&rr6f.Hexdump, "hexd", false, "Hex dump a image")
	flags.BoolVar(&rr6f.Geo, "geo", false, "Get the GEO location of a JPG/JPEG Info, of which has GPS location ")
	flags.BoolVar(&rr6f.Walk, "walk", false, "Walk a filepath for images and EXIF dump all data to all images")
	flags.BoolVar(&rr6f.Walkerfp, "walkf", false, "Walk a filepath for images")
	flags.BoolVar(&rr6f.Discover, "discover", false, "Determin the type of file of an unknown file")
	flags.StringVar(&rr6f.Hashlist, "hashl", "", "Set a hash list for hash brute forcing")
	flags.StringVar(&rr6f.Userlist, "userl", "", "Set a user list for user brute forcing")
	flags.StringVar(&rr6f.Brute_list, "wordl", "/usr/share/wordlists/rockyou.txt", "Set a wordlist for brute forcing")
	flags.IntVar(&rr6f.Workers, "workers", 200, "Set the amount of workers for brute forcing -> defualt 200")
	flags.StringVarP(&rr6f.Packet_t, "packet", "", "", "")
	flags.StringVarP(&rr6f.Sniffc, "interface", "", "", "")
	flags.StringVarP(&rr6f.Target_mac, "targetm", "", "", "Set the target's mac address for arp poisoning")
	flags.StringVarP(&rr6f.Target_spoof, "targetip", "", "", "Gateway IP address")
	flags.StringVarP(&rr6f.Gateway_mac, "gatemac", "", "", "Set the targets gateway mac for arp poisoning")
	flags.StringVarP(&rr6f.Iprange, "CIDR", "z", "192.168.1.8/24", "Set a Network or IP range to scan for hosts")
	flags.IntVar(&rr6f.Pass_length, "passlen", 16, "Set a password length for pass generation")
	flags.StringVar(&rr6f.Per_mode, "perm", "", "Set the premissions of a file for permission stomper / changing")
	flags.StringVar(&rr6f.Url, "target", "", "Set the target url for testig or injecting or recon")
	flags.StringVar(&rr6f.PayloadList, "payloadl", "", "Set a payload list for XSS, SQLI, Admin panel finding, vuln finding, recon, or subdomains")
	flags.StringVarP(&rr6f.XML_file, "XMLF", "X", "", "Set a XML file for nmap parsing")
	flags.StringVarP(&rr6f.JSON_file, "JSONF", "J", "", "Set a JSON file for parsing")
	flags.StringVarP(&rr6f.PCAP_file, "PCAP", "P", "", "Set a PCAP file for parsing")
	flags.StringVarP(&rr6f.Search_Query, "SQ", "S", "", "Set the google search query or thing you would like to search for the engine module")
	flags.IntVarP(&rr6f.Results_per_Page, "RPP", "R", 10, "Set the amount of results per google page scraped for the engine module (Defualt = 10)")
	flags.IntVarP(&rr6f.Pages_To_Crawl, "PTC", "T", 1, "Set the amount of google pages to crawl for the engine module (Defualt = 1)")
	flags.Lookup("type").NoOptDefVal = "rNDm"
	if rr6f.Image_offset != "" {
		byteOffset, _ := strconv.ParseInt(rr6f.Image_offset, 0, 64)
		rr6f.Image_offset = strconv.FormatInt(byteOffset, 10)
	}
	if rr6f.Suppress && (!rr6f.S_M) {
		log.Fatal("Fatal: The --meta flag is required when using --suppress")
	}
	if rr6f.S_M && (rr6f.Image_offset != "") {
		log.Fatal("Fatal: The --meta flag is mutually exclusive with --offset")
	}
	if rr6f.Inject && (rr6f.Image_offset == "") {
		log.Fatal("Fatal: The --offset flag is required when using --inject")
	}
	if rr6f.Inject && (rr6f.Payload == "") {
		log.Fatal("Fatal: The --payload flag is required when using --inject")
	}
	if rr6f.Inject && rr6f.Key == "" {
		fmt.Println("Warning: No key provided. Payload will not be encrypted")
	}
	if rr6f.Payload_Encode && rr6f.Key == "" {
		log.Fatal("Fatal: The --encode flag requires a --key value")
	}
	if rr6f.INJECT_ZIP && rr6f.Filepath_general == "" && rr6f.Output == "" {
		log.Fatal("the --JPGINJECT requires the --filepath flag to specify the zip files as well as the -o flag to specify a file or image output name | EXAMPLE BELOW \n\n")
		log.Fatal("| -> +++++? | go run main.go -i image.jpg -o main.jpg --JPGINJECT --filepath file.zip")
	}
	flags.BoolVar(&rr6f.Rr6_help_flags, "rrh", false, "Spawn a help menu for flags")
	flags.Parse(os.Args[1:])
}
