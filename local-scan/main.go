package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func isOpen(host string, port int, timeout time.Duration) bool {
	con, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		_ = con.Close()
		return true
	}
	return false
}

func main() {
	PIP := flag.String("hostname", "", "Private IP -> EX: 10.0.0.0")
	SPort := flag.Int("start-port", 80, "Range to start with  -> EX: 80")
	EPort := flag.Int("end-port", 100, "Range where scan ends -> EX: 100")
	timeout := flag.Duration("timeout", time.Millisecond*200, "timeout")

	flag.Parse()

	ports := []int{}
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	for port := *SPort; port <= *EPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*PIP, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	fmt.Printf("Ports Returned as True -> %v\n", ports)
}
