package IOSCAN

import (
	"bufio"
	"fmt"
	v "main/modg/colors"
	"main/modg/system-runscript"
	"os"
)

func Scanv(method int, IO_STRING string, color string, exit string, variable string) (string, uint) {
	fmt.Print(color, IO_STRING)
	if method == 1 {
		fmt.Scanf("%s", &variable)
		return variable, 0x00
	}
	return variable, 0x00
}

func ScanV2(method int, IO_STRING string, color string, exit string, variable string, variable2 string, varible3 string) (string, string, uint) {
	fmt.Print(color, IO_STRING)
	values := make([]string, 2)
	scanner := bufio.NewScanner(os.Stdin)
	system.Sep("\n\n")
	if method == 1 {
		for {
			fmt.Print(v.BLKHB, "Enter all 2 values> ", v.RET_RED)
			scanner.Scan()
			txt := scanner.Text()
			values = append(values, txt)
			break
		}
	}
	return values[1], values[2], 0x00
}
