package AES

import (
	"fmt"
	"io/ioutil"
	c "main/modg/colors"
	"os"
)

type AES_Data struct {
	Key_File string // Key_File
	File     string // file to encrypt
	Decrypt  bool   // decrypt yes or no if true yes else no
	Gen      bool   // generate key yes or no if true yes if else no
}

func (q *AES_Data) FUNC(file_to_enc_or_dec, key_file string, de bool, gen bool) {
	q.File = file_to_enc_or_dec
	q.Key_File = key_file
	q.Decrypt = de
	q.Gen = gen
	if gen {
		msg := Generate_AES_Key()
		fmt.Printf("[*] Random Generated Key formatted -> %x", msg)
		os.Exit(0)
	}
	kd, x := ioutil.ReadFile(q.Key_File)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> Encryption -> AES Module: Could not open key file, got error when using i/o to open and read all contents of the given file -> ", x)
		os.Exit(0)
	}
	kted, x := ioutil.ReadFile(q.File)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> Encryption -> AES module: Could not open the contents of the file you would like to encrypt, got error when using the i/o utils to read file -> ", x, " [ EXITING WITH STATUS 0, RE RUN RED RABBIT ] ")
		os.Exit(0)
	}

	if q.Decrypt {
		msg, x := Dec(kd, kted)
		if x != 0x00 {
			fmt.Println(x)
		} else {
			fmt.Printf("[*] Decrypted text - %s\n", msg)
		}
	} else {
		msg, x := Enc(kd, kted)
		if x != 0x00 {
			fmt.Println(c.REDHB, "<RR6> Encryption -> AES Module: Could not load the key, output the data, or format error and bytes, got error when attempting to run all 3 functions or 1 of the 3 functions -> ", x, " [ NON EXIT RELATION STATUS CODE, NOT EXITING RETURNING TO USER IO ]")
		} else {
			fmt.Printf("%s", msg)
		}
	}
}
