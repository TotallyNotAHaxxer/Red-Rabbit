// filepath walking script with regex
// author => ArkAngeL43
// lang   => go
// Is     => Filepath pillaging

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	filepath_walk = flag.String("fp", "/", "set filepath for pillaging")
	strin         = flag.String("regexc", "admin", "set a regex paramater")
)

func WALKER_MAIN_FUNC(path string, f os.FileInfo, err error) error {
	flag.Parse()
	parser_regex_format := "(?i)" + *strin
	var regex_scanner = []*regexp.Regexp{
		regexp.MustCompile(parser_regex_format),
	}
	for _, r := range regex_scanner {
		if r.MatchString(path) {
			fmt.Printf("\033[31m FOUND \033[32mFILEPATH -> %s\n", path)
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if err := filepath.Walk(*filepath_walk, WALKER_MAIN_FUNC); err != nil {
		log.Panicln(err)
	}
}
