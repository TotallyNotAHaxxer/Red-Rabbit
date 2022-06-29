package errors_case

import (
	"fmt"
	"log"
	"os"
)

//   e = error, msg = print/output messages, exit/method is the exit and output method, si_exit, if you exit true
func Return_error(err uint, msg string, exit, method int, si_exit bool) {
	if err != 0x00 {
		switch method {
		case 1:
			fmt.Println("<RR6> IO SCANNER ERROR: -> ", msg, err)
			if si_exit {
				exiter(exit)
			}
		case 2:
			fmt.Printf("<RR6> IO SCANNER ERROR: [%v] ", err)
			if si_exit {
				exiter(exit)
			}
		case 3:
			log.Fatal(err)
		}
	}
}

func exiter(exit_code int) {
	os.Exit(exit_code)
}
