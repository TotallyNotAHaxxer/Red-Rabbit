package OSINT_Engine

import "fmt"

func STDOUT(data map[string][]string) {
	for k, v := range data {
		for _, t := range v {
			if k != "" {
				if t != "" {
					fmt.Printf("\033[38;5;88m[TAG] > \033[38;5;55m%s -> \033[38;5;55m\t|+| -> \033[38;5;50m%s\n", k, t)
				}
			}
		}
	}
}
