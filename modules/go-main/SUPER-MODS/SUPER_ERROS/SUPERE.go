package SUPER_ERRORS_AUTO_GEN

import "fmt"

func Erorr(x error, y string, c string) {
	if x != nil {
		fmt.Println(c, y, x)
	}
}
