package error

import "fmt"

func Return_error(err error, msg string, color string) {
	if err != nil {
		fmt.Println(color, msg, err)
	}
}

func Return_Warnings(msg, color string, code int) {
	fmt.Println(color, msg, code)
}
