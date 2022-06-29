package stomper

import (
	"fmt"
	"os"
	"strconv"
)

func Permissions(mode, filename string) {
	v, x := strconv.ParseUint(mode, 8, 32)
	if x != nil {
		fmt.Println("<RR6> File I/O Module -> Parsing UINT: Got error when trying to parse the file mode, please try to correct this error -> ", x)
	} else {
		fm := os.FileMode(v)
		x = os.Chmod(filename, fm)
		if x != nil {
			fmt.Println("<RR6> File I/O Module -> CHMOD: Got error when trying to remod this filename with a new mode -> ", x)
		} else {
			fmt.Println("[*] File premissions changed on file -> ", filename)
			fmt.Println("[*] File premissions changed as      -> ", mode)
			fmt.Println("[*] Operation on file sucess         -> PASS")
		}
	}
}
