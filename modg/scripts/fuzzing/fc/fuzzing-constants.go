package fuzzer_constants

import "time"

var (
	Threading_Channel = make(chan bool)
	Response_buffer   = make([]byte, 4096)
	Fuzz_Dial_method  = "tcp"
	Fuzz_Dial_timeout = time.Duration(10)
	AT                = 0
	Max_Bytes         = 1024
)
