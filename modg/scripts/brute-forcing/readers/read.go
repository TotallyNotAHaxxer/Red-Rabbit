package readers

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Read_filepath(filepath string) []string {
	data, e := ioutil.ReadFile(filepath)
	if e != nil {
		fmt.Println("<RR6> Errors module: Could not read filepath")
	}
	// split the file into new lines
	return strings.Split(string(data), "\n")
}
