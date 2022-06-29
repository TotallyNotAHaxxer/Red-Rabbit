package fuzzing

import (
	"fmt"
	fuzzer_constants "main/modg/scripts/fuzzing/fc"
	"math/rand"
	"net"
	"time"
)

func writer(conn net.Conn, msg []byte, size int, Chan chan bool) {
	write_byte, x := conn.Write(msg)
	if x != nil {
		fmt.Println("<RR6> Fuzzer Writer: Could not write the byte to the server, got error -> ", x)
		Chan <- true
		return
	} else {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| STAT: Was able to send the message and packet to the server")
		if write_byte != size {
			fmt.Println("[-] STAT: Could not write the full size and payload to the server......")
			return
		}
	}
}

func reader(conn net.Conn, Chan chan bool) {
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	read, x := conn.Read(fuzzer_constants.Response_buffer)
	if x != nil {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| STAT: Could not read the 4K buffer, got error -> ", x)
		Chan <- true
		return
	}
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| STAT: (Ignore) Reader response -> ", read)
}

func Dial(port, host string, buffer int, Chan chan bool) {
	cons := host + ":" + port
	connect, x := net.DialTimeout(fuzzer_constants.Fuzz_Dial_method, cons, time.Second*fuzzer_constants.Fuzz_Dial_timeout)
	if x != nil {
		fmt.Println("<RR6> Fuzzing Module: Could not make a proper connection to the service, got error -> ", x)
		Chan <- true
		return
	}
	pack := make([]byte, buffer)
	rand.Read(pack)
	t := time.Now()
	connect.SetWriteDeadline(t.Add(time.Second * 5))
	writer(connect, pack, buffer, Chan)
	// read the data back
	reader(connect, Chan)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| STAT: (Ignore) Sent -> ", buffer)
}

func Main(port, host string, fuzz_ int) {
	Chan := make(chan bool)
	for fuzz_ := 1; fuzz_ <= fuzzer_constants.Max_Bytes; fuzz_ = fuzz_ * 2 {
		go Dial(port, host, fuzz_, Chan)
		fuzzer_constants.AT++
	}
	for fuzzer_constants.AT > 0 {
		<-Chan
		fuzzer_constants.AT--
	}
}
