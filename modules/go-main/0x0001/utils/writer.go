package utils

import (
	"bytes"
	"fmt"
	"io"
	co "main/modg/colors"
	rr6_opts "main/modg/copt"
	e "main/modules/go-main/0x0001/steg-err"
	"os"
	"strconv"
)

func J(byter_ *bytes.Reader, opt_ *rr6_opts.RR6_options, b []byte) {
	offset, err := strconv.ParseInt(opt_.Image_offset, 10, 64)
	e.Return_error(err, "<RR6> Stego Module: Could not parse integer based value -> ", co.REDHB)
	w, err := os.OpenFile(opt_.Output, os.O_RDWR|os.O_CREATE, 0777)
	e.Return_error(err, "<RR6> Stego Module: Could not open the file got error -> ", co.REDHB)
	byter_.Seek(0, 0)
	var buffer = make([]byte, offset)
	byter_.Read(buffer)
	w.Write(buffer)
	w.Write(b)
	if opt_.Payload_Decode {
		byter_.Seek(int64(len(b)), 1)
	}
	_, err = io.Copy(w, byter_)
	if err == nil {
		fmt.Printf("<RR6> Stego Module: %s has been sucessfully created\n", opt_.Output)
	}
}
