package a80211_cpp_iinterface_utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

var (
	out_buff bytes.Buffer
	err_buff bytes.Buffer
)

func Defualt_interface() string {
	GOOS := runtime.GOOS
	switch GOOS {
	case "linux":
		return "wlan0"
	default:
		return "unsupported"
	}
}

func Run_CPP(down_or_up, interface_name_or_IP, interface_switch_name, exe, yn_justip string) {
	prg := "./" + exe
	c := exec.Command(prg, down_or_up, interface_name_or_IP, interface_switch_name, yn_justip)
	fmt.Println("running -> ", c)
	s, x := c.Output()
	if x != nil {
		fmt.Println("<RR6> Got error when running executeable -> ", x)
	} else {
		fmt.Print(string(s))
	}
}
