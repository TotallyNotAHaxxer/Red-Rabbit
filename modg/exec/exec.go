/*
Developer   | ArkAngeL43
Package     | execg
Module      | exec
Nest        | exec
FP          | modg/exec/exec

Does:
	This module is written to automate system commands, or like calls
	or execute perl files such as the GUI's, perl modules, and ruby scripts
	as well as other payload generators written in any other language not written
	to port straight into golang, mostly this automates the perl caller, since
	current version of red rabbit will not include a perl wrapper for languages like
	go


*/

package execg

import (
	"fmt"
	"log"
	v "main/modg/colors"
	opc "main/modg/copt"
	"main/modg/system-runscript"
	ec "main/modg/warnings"
	"os/exec"
)

const (
	rr6pl = "r6.pl"
)

// perl_straight_caller_call_module_offensive_security
func Perl_straight_caller_call_module_offensive_security(options, color, filename string) {
	prg := "perl"
	arg1 := rr6pl
	arg00 := "-o"
	arg01 := options
	arg5 := "-f"
	arg2 := filename
	cmd := exec.Command(prg, arg1, arg00, arg01, arg5, arg2)
	stdout, err := cmd.Output()
	ec.Warning_simple("<RR6> EXEC Module: Could not execute file? error: -> ", v.REDHB, err)
	fmt.Print(color, string(stdout))
}

// perl caller
func Call_perl_s(image string) {
	prg := "perl"
	arg1 := rr6pl
	arg00 := "-o"
	arg01 := "opt1"
	arg5 := "-f"
	arg2 := image
	cmd := exec.Command(prg, arg1, arg00, arg01, arg5, arg2)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -<> ", 1)
	fmt.Print(string(stdout))

}

// perl caller for GUI's
func Call_perl_GUI_host_scanner(ip string) {
	prg := "perl"
	arg1 := rr6pl

	arg5 := ip
	cmd := exec.Command(prg, arg1, arg5)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -> ", 1)
	system.Sep("\n")
	fmt.Print(v.RED, string(stdout))
	system.Sep("\n")
}

// call R6 perl GUI
func Call_perl_RR6_GUI() {
	prg := "perl"
	arg1 := "gui-multi.pl"
	cmd := exec.Command(prg, arg1)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -> ", 1)
	system.Sep("\n")
	fmt.Print(v.RED, string(stdout))
	system.Sep("\n")
}

// call perl GIF injection
func Run_GIF(input_file, payload, pixel_height, pixel_width string) {
	prg := "perl"
	arg1 := rr6pl

	arg2 := "-o"
	arg3 := "gif"
	arg4 := "-f"
	arg5 := input_file
	arg6 := "-p"
	arg7 := payload
	arg8 := "-h"
	arg9 := pixel_height
	arg10 := "-w"
	arg11 := pixel_width
	cmd := exec.Command(prg, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -> ", 1)
	system.Sep("\n")
	fmt.Print(v.RED, string(stdout))
	system.Sep("\n")
}

// call webp
func Run_Webp(input_file, payload string) {
	prg := "perl"
	arg1 := rr6pl

	arg2 := "-o"
	arg3 := "webp"
	arg4 := "-f"
	arg5 := input_file
	arg6 := "-p"
	arg7 := payload
	cmd := exec.Command(prg, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -> ", 1)
	system.Sep("\n")
	fmt.Print(v.RED, string(stdout))
	system.Sep("\n")
}

// call jpg
func Run_JPG(chunk, payload, input string) {
	prg := "perl"
	arg1 := rr6pl
	arg2 := "-o"
	arg3 := "jpg"
	arg4 := "-f"
	arg5 := input
	arg6 := "-p"
	arg7 := payload
	arg8 := "-c"
	arg9 := chunk
	cmd := exec.Command(prg, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -> ", 1)
	system.Sep("\n")
	fmt.Print(v.RED, string(stdout))
	system.Sep("\n")
}

func Run_bmp(payload, input string) {
	prg := "perl"
	arg1 := rr6pl
	arg2 := "-o"
	arg3 := "bmp"
	arg4 := "-f"
	arg5 := input
	arg6 := "-p"
	arg7 := payload
	cmd := exec.Command(prg, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -> ", 1)
	system.Sep("\n")
	fmt.Print(v.RED, string(stdout))
	system.Sep("\n")
}

// call traceroute
func Traceroute() {
	var domain string
	fmt.Printf("Enter a domain> ")
	fmt.Scanf("%s", &domain)
	prg := "perl"
	arg1 := rr6pl
	arg2 := "-o"
	arg3 := "traceroute"
	arg4 := "-r"
	arg5 := domain
	cmd := exec.Command(prg, arg1, arg2, arg3, arg4, arg5)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -> ", 1)
	system.Sep("\n")
	fmt.Print(v.RED, string(stdout))
	system.Sep("\n")
}

// run perl r6.pl file for stego options

func Perl_run_r6pl(options string, po *opc.RR6_options) {
	prg := "perl"
	arg1 := rr6pl
	arg2 := "-o"
	arg3 := options
	arg4 := "-f"
	arg5 := po.Input
	arg6 := "-p"
	arg7 := po.Payload
	arg8 := "-c"    // chunk call ARG
	arg9 := po.Type // chunk if injecting JPG images
	fmt.Println("Execute: ", prg, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	cmd := exec.Command(prg, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -> ", 1)
	fmt.Print(string(stdout))
}

// push file help to perl table module
func Push(filename string) {
	prg := "perl"
	arg1 := rr6pl
	arg2 := "-o"
	arg3 := "tablef"
	arg4 := "-f"
	arg5 := filename
	cmd := exec.Command(prg, arg1, arg2, arg3, arg4, arg5)
	stdout, err := cmd.Output()
	ec.Che(err, "Could not run perl file -> ", 1)
	system.Sep("\n")
	fmt.Print(v.WHT, string(stdout))
	system.Sep("\n")
}

// BOF

func BOF(filepath string) {
	filep := "modg2/ported-from-beta/BOF.pl"

	prg := "perl"
	arg0 := filep
	arg := filepath
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| RR6> LOG: Settings: CONF -> ", prg, arg0, arg)
	cmd := exec.Command(prg, arg0, arg)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	system.Sep("\n")
	fmt.Print(v.HIGH_PINK, string(stdout))
	system.Sep("\n")
}

// deface

func Defacer_(filepath string) {
	filep := "modg2/ported-from-beta/defacer.pl"
	prg := "perl"
	prg2 := filep
	arg0 := "-f"
	arg1 := filepath
	cmd := exec.Command(prg, prg2, arg0, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	system.Sep("\n")
	fmt.Print(v.RED, string(stdout))
	system.Sep("\n")
}

// host discovery
// call perl r6.pl host discover

func Host_discover_r6(method, CIDR string) {
	prg := "perl"
	arg1 := "r6.pl"
	arg2 := "-o"
	arg3 := method
	arg8 := "-t"
	arg9 := CIDR
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| RR6> LOG: Settings: CONF -> ", CIDR)
	cmd := exec.Command(prg, arg1, arg2, arg3, arg8, arg9)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	system.Sep("\n")
	fmt.Print(v.RED, string(stdout))
	system.Sep("\n")
}

// run normal any file

func Run(language, filename string) {
	prog := language
	filepath := filename
	cmd := exec.Command(prog, filepath)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	system.Sep("\n")
	fmt.Print(v.RED, string(stdout))
	system.Sep("\n")
}

// run with args
func RunA(lang, file string, args string) string {
	prog := lang
	prog2 := file
	cmd := exec.Command(prog, prog2, args)
	fmt.Println("Running -> ", prog, prog2, args)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(stdout)
}
