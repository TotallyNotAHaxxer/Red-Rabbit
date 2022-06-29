package SSH_errors

import "fmt"

// simplistic error
func Normal(x error, y string, c string) error {
	if x != nil {
		fmt.Println(c, y)
		return x
	}
	return nil
}
