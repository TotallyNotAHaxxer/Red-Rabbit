/*

This script is for base64 encoding a string apart of the
png image injection payload script and experimentation

*/
package main

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
)

var (
	fdata = flag.String("data", "", "set a string or payload to base 64 encode")
)

func main() {
	flag.Parse()
	data := *fdata
	enc_m := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("[+] Payload string -> ", data)
	fmt.Println("[+] Base 64 encode -> ", enc_m)
}
