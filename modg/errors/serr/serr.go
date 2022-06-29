package simple_error

import (
	"fmt"
	"os"
)

func See_intbased(err uint, color, msg string, exit bool) {
	if err != 0x00 {
		fmt.Println(color, "<RR6> Simple errors module: ", msg, err)
		switch exit {
		case exit:
			os.Exit(1)
		case !exit:
			fmt.Println(color, "<RR6> Somple errors module: Will not be exiting")
		}
	}
}

func See_errorbased(err error, color, msg string, exit bool) {
	if err != nil {
		fmt.Println(color, "<RR6> Simple errors module: ", msg, err)
		switch exit {
		case exit:
			os.Exit(1)
		case !exit:
			fmt.Println(color, "<RR6> Somple errors module: Will not be exiting")
		}
	}
}
