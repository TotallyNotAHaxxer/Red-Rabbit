package ssh_errors

import (
	"fmt"
	"log"
)

func Error_And_Exit(g, c string, x error) bool {
	if x != nil {
		log.Fatal(c, g, x)
		return true
	}
	return false
}

func Error_And_Continue(g, c string, x error) bool {
	if x != nil {
		fmt.Println(c, g, x)
		return true
	}
	return false
}
