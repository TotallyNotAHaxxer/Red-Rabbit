package SUPER_Archive

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func Run(file string) {
	f, x := os.Open(file)
	if x != nil {
		fmt.Println("<RR6> GZIP LIB: Could not open the filename, got error -> ", x)
	} else {
		out, x := gzip.NewReader(f)
		if x != nil {
			fmt.Println("<RR6> GZIP LIB: Was not able to make a gzip new reader for the specified file, got error -> ", x)
		} else {
			defer out.Close()
			w, x := os.Create("unzipped.txt")
			if x != nil {
				fmt.Println("<RR6> GZIP Lib: Was not able to create the output filename... Got error, ", x)
			} else {
				defer w.Close()
				_, x = io.Copy(w, out)
				if x != nil {
					log.Fatal(x)
				}
			}
		}
	}
}
