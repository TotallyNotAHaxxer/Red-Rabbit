package B80211_errors

import (
	"fmt"
	"os"
)

func Warn(x error, y, c string, l int) {
	if x != nil {
		fmt.Println(c, y, l)
	}
}

// y = message
// x = error
// c = color
// l = logger
// e = exit code
// f = fatal yes or no?
func Error(x error, y, c string, l, e int, f bool) {
	switch f {
	case true:
		if x != nil {
			fmt.Println(c, y, x)
			os.Exit(e)
		}
	case false:
		if x != nil {
			fmt.Println(c, y, x)
		}
	}
}
