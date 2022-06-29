package SUPER_regex

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	color "main/modg/colors"
)

func Regex(regex_string string, data_ string) bool {
	a := regexp.MustCompile(regex_string)
	return a.MatchString(data_)
}

func Compare_contents(find_data, regex_string string) {
	re := regexp.MustCompile(regex_string)
	ctx := re.FindAllStringIndex(find_data, -1)
	for _, k := range ctx {
		h := find_data[k[0]:k[1]]
		fmt.Println(color.HIGH_PINK)
		fmt.Println("###### NEW DATA SEGMENT....")
		fmt.Print("\n")
		fmt.Println(color.HIGH_BLUE)
		fmt.Printf("Found Data -> |  %s   \n", h)
		fmt.Printf("Line       -> |  %d    \n", k[0])
		fmt.Printf("Column     -> |  %d    \n", k[1])
	}
}

func Open_file(filename string, regexstr string) string {
	f, x := os.Open(filename)
	if x != nil {
		log.Fatal("<RR6> File I/O: Could not open, locate, or describe filename -> ", x)
	} else {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			if Regex(regexstr, scanner.Text()) {
				Compare_contents(scanner.Text(), regexstr)
			}
		}
	}
	return ""
}
