package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	v "main/modg/colors"
	"main/modg/constants"
	jsa "main/modg/json"
	"main/modg/system-runscript"
	ec "main/modg/warnings"
	"os"
	"os/exec"

	opc "main/modg/copt"
	gen "main/modg/helpers/id-gen"

	"github.com/spf13/pflag"
)

var (
	flags = pflag.FlagSet{SortFlags: false}
	rr6f  opc.RR6_options
)

type NIL struct {
}

func Usage(file, color string) {
	content, err := ioutil.ReadFile(file)
	ec.Warning_advanced("<RR6> File Module: Could not open or read file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	fmt.Println(v.RED, constants.Clear_hex, string(content))
}

func Readall() {
	fmt.Println(" =========== VERIFIED")
	fp := constants.Verified_commands
	newfp := jsa.Read_filepath(fp)
	jsa.Open_json_verified_commands(newfp)
	fmt.Println(" =========== Commands General")
	fp1 := constants.Verified_commands
	newfp1 := jsa.Read_filepath(fp1)
	jsa.Open_json_verified_commands(newfp1)
	fmt.Println(" =========== Commands flags ")
	fp2 := constants.Flags
	newfp2 := jsa.Read_filepath(fp2)
	fmt.Println(newfp)
	jsa.Help_usage_and_menus(true, newfp2, v.RED, "\n", 3)
	system.Sep("\n\n\n")
}

func Write(body1 string) {
	if !system.Isexisting("robots.txt") {
		d1 := []byte(body1)
		err := os.WriteFile("robots.txt", d1, 0644)
		if err != nil {
			log.Fatal(err)
		}
		path, err3 := os.Getwd()
		ec.Warning_advanced("<RR6> OS ERRORS Module: Could not get the working directory, THIS IS A FATAL ERROR", v.REDHB, 1, false, false, true, err3, 1, 233, "")
		fmt.Println("FILE PATH OF GENERATED FILE => ", path)
	}
	if !system.Isexisting("robots.txt") {
		fmt.Println(v.REDHB, "<RR6> Download Module: For some reason this file already exists, creating new one\033[49m", v.RED)
		s := ".txt"
		a := gen.ShortID(4) + s
		f, err := os.Create(a)
		ec.Warning_advanced("<RR6> OS ERRORS Module: Could not create the file, THIS IS A FATAL ERR -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		defer f.Close()
		d1 := []byte(body1)
		err1 := os.WriteFile(a, d1, 0644)
		ec.Warning_advanced("<RR6> OS ERRORS Module: Could not write to the file, THIS IS A FATAL ERROR", v.REDHB, 1, false, false, true, err1, 1, 233, "")
		path, err3 := os.Getwd()
		ec.Warning_advanced("<RR6> OS ERRORS Module: Could not get the working directory, THIS IS A FATAL ERROR", v.REDHB, 1, false, false, true, err3, 1, 233, "")
		fmt.Println(v.BLKHB, "<RR6> Download Module: New file output is in < ", path+"/"+a+" > \033[49m\033[39m", v.RED)
	}

}

func Call_sql(wordlist, sqlfile string) {
	prg := "perl"
	arg1 := "r6.pl"
	arg2 := "-o"
	arg3 := "ql-hash-dumper"
	arg8 := "-f"
	arg9 := wordlist
	arg10 := "-q"
	arg11 := sqlfile
	fmt.Println(prg, arg1, arg2, arg3, arg8, arg9, arg10, arg11)
	cmd := exec.Command(prg, arg1, arg2, arg3, arg8, arg9, arg10, arg11)
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	system.Sep("\n")
	fmt.Print(v.RED, string(stdout))
	system.Sep("\n")
}

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
	} else if Screen_rotation == "none" {
		fmt.Println(constants.Clear_hex)
		fmt.Println(v.BLKHB, "\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| <<< Starting red rabbit console >>>")
		fmt.Println(v.BLKHB, "[!] <<< Setting: Screen resolution set to `none` no output banner or format id specified >>>")
	}
}

func Banner(file, color string) {
	content, err := ioutil.ReadFile(file)
	ec.Ce(err, v.RED, "Could not read file", 1)
	fmt.Println(constants.Clear_hex, color, string(content))
}
