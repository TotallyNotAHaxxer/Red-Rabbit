package SUPER_Compression

import (
	"compress/gzip"
	"fmt"
	"log"
	"os"
)

func File(filename string, encoding string) {
	var end_encode string
	switch encoding {
	case "gzip":
		end_encode = ".gz"
	case "bzip2":
		end_encode = ".bz2"
	case "lzw":
		end_encode = ".lzw"
	case "zlib":
		end_encode = ".zl"
	}
	encpparse := filename + end_encode
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File save     : ", encpparse)
	f, x := os.Create(encpparse)
	if x != nil {
		fmt.Println("<RR6> SUPER UTILS: Could not create the new filename, got error -> ", x)
	} else {
		w := gzip.NewWriter(f)
		defer f.Close()
		_, x = w.Write([]byte("some data"))
		if x != nil {
			log.Fatal(x)
		} else {
			fmt.Println("data written to file....")
		}
	}
}
