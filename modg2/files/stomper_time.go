package stomper

import (
	"fmt"
	l "main/modg/colors"
	e "main/modg/errors/serr"
	"os"
	"time"
)

var x error

// timestamp stomping

func Change_Timestamp(filename string, hour, minute time.Duration) {
	c := time.Now().Add(hour * time.Hour).Add(minute * time.Minute)
	at := c
	am := c
	x = os.Chtimes(filename, at, am)
	e.See_errorbased(x, l.REDHB, "<RR6> File I/O Errors: Could not stomp or change both the modification time, and the last accessed time of the given file, got error -> ", false)
	if x == nil {
		fmt.Println("P | Was able to change modification time, and access time of the file")
	}
}
