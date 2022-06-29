package SUPER_Unarchive

import (
	"archive/zip"
	"fmt"
	"io"
	c "main/modg/colors"
	"os"
	"path/filepath"
)

func Run(filename string) {
	r, x := zip.OpenReader(filename)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> SUPER CLASS ZIP: Could not make a new reader for the current file")
	} else {
		defer r.Close()
		for _, k := range r.Reader.File {
			f, x := k.Open()
			if x != nil {
				fmt.Println("<RR6> SUPER CLASS ZIP: Could not read or open the file inside of the archives/zipped file, got error -> ", x)
			} else {
				defer f.Close()
				outdir := "./"
				extract := filepath.Join(outdir, k.Name)
				if k.FileInfo().IsDir() {
					fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Extracting       : ", k.Name)
					fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Creating file    : ", extract)
					os.MkdirAll(extract, k.Mode())
				} else {
					fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Extracting       : ", k.Name)
					opener, x := os.OpenFile(extract, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, k.Mode())
					if x != nil {
						fmt.Println("<RR6> Reader I/O SUPER: Could not create or open the filename, got error -> ", x)
					} else {
						defer opener.Close()
						_, x = io.Copy(opener, f)
						if x != nil {
							fmt.Println("<RR6> Writer I/O SUPER: Could not copy the contents of the unextracted file to a new output file, got error -> ", x)
						} else {
							fmt.Println("<RR6> Write I/O SUPER: Was able to copy data to the new file")
						}
					}
				}
			}
		}
	}
}
