/*
Developer -> ArkAngeL43
File      -> lookup-files.go
Filepath  -> modg/scripts/file-forensics
Module    -> scripts
Package   -> forensicsfiles

Does:
	Digital forensics on images, filenames, files, filepaths
	will also execute premission, time, and timestamp stompers, this will also
	revolve around image forensics, and image injection / archive extraction

*/
package forensicsfiles

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	v "main/modg/colors"
	system "main/modg/system-runscript"
	"main/modg/warnings"
	warn "main/modg/warnings"
	"os"
	"os/exec"
	"syscall"
)

func Extract_ZIP(img_i, o string) {
	fileJpg := img_i
	file, err := os.Open(fileJpg)
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)
	fileStat, _ := file.Stat()
	for i := int64(0); i < fileStat.Size(); i++ {
		myByte, err := bufferedReader.ReadByte()
		warn.Warning_advanced("<RR6> Forensics Module: Could not read bytes -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")

		if myByte == '\x50' {
			byteSlice := make([]byte, 3)
			byteSlice, err = bufferedReader.Peek(3)
			warn.Warning_advanced("<RR6> Forensics Module: Could not peek buffer or read buffer with a peak of 3 -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
			if bytes.Equal(byteSlice, []byte{'\x4b', '\x03', '\x04'}) {
				log.Printf("Found zip signature at byte %d.", i)
			}
		}
	}
	for {
		var unz string
		fmt.Print("Unzip the archive? (Y/n) >>> ")
		_, err := fmt.Scanf("%s", &unz)
		if err != nil {
			fmt.Println("[!] Not an option")
			continue
		}
		switch {
		case unz == "Y" || unz == "y" || unz == "yes" || unz == "Yes":
			fmt.Println("OK")
			binary, err := exec.LookPath("/usr/bin/7z")
			warnings.Warning_advanced("<RR6> Syscall Module: Could not make a syscall for finding the binar path > ", v.REDHB, 1, false, false, false, err, 1, 255, "")
			command_parser := "-o" + o
			args := []string{"7z", "e", img_i, command_parser}
			env := os.Environ()
			err = syscall.Exec(binary, args, env)
			warnings.Warning_advanced("<RR6> Syscall Module: Could not make a syscall for execution of the binary path > ", v.REDHB, 1, false, false, false, err, 1, 255, "")
		case unz == "N":
			fmt.Println("Finished scan and data grabber")
			os.Exit(0)
		default:
			fmt.Println("[!] Continuing, this might not be the correct statement? ")
			continue
		}
	}
}

func Inject_ZIP(img_i, zip_file, Output string) {
	fileJpg := img_i
	fileZip := zip_file
	firstFile, err := os.Open(fileJpg)
	warn.Warning_advanced("<RR6> Forensics Module: Could not open file -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	fmt.Println(v.RED, "[INFO] ", v.BLU, system.FormatDate, v.RED, " -> File Opened and READ             \t", fileJpg)
	fmt.Println(v.RED, "[INFO] ", v.BLU, system.FormatDate, v.RED, " -> Preping output as injected file  \t", Output)
	defer firstFile.Close()
	secondFile, err := os.Open(fileZip)
	warn.Warning_advanced("<RR6> Forensics Module: Could not open the ZIP file -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	fmt.Println(v.RED, "[INFO] ", v.BLU, system.FormatDate, v.RED, " -> ZIP file has been opened         \t", zip_file)
	defer secondFile.Close()
	newFile, err := os.Create(Output)
	warn.Warning_advanced("<RR6> Forensics Module: Could not create output files -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	fmt.Println(v.RED, "[INFO] ", v.BLU, system.FormatDate, v.RED, " -> OUTPUT File with injection made  \t", Output)
	defer newFile.Close()
	_, err = io.Copy(newFile, firstFile)
	warn.Warning_advanced("<RR6> Forensics Module: Could not Copy file -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	_, err = io.Copy(newFile, secondFile)
	warn.Warning_advanced("<RR6> Forensics Module: Could not copy file -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	fmt.Println(v.RED, "[INFO] ", v.BLU, system.FormatDate, v.RED, " ->  Finished Injection of image at  \t", system.FormatDate, Output)
}

func Dump_boots(filepath string) {
	fmt.Println(v.BLUHB, "<RR6> Meta Module: Reading boot sector of < "+v.REDHB+filepath+" > \033[39m\033[49m")
	file, err := os.Open(filepath)
	warn.Warning_advanced("<RR6> Meta Module: Could not open the filepath, reason -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	byteSlice := make([]byte, 512)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	warn.Warning_advanced("<RR6> Could not read full byte slice, or bytes", v.REDHB, 1, false, false, true, err, 1, 233, "")
	fmt.Println(v.RED)
	fmt.Printf("Bytes read: %d\n\n", numBytesRead)
	fmt.Printf("Data as decimal:\n%d\n\n", byteSlice)
	fmt.Printf("Data as hex:\n%x\n\n", byteSlice)
	fmt.Printf("Data as string:\n%s\n\n", byteSlice)
}
