package AES

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	c "main/modg/colors"
)

func Enc_Msg(k, m []byte, filename string) {
	r, x := aes.NewCipher(k)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> Encryptiopn module -> AES: Could not make or create a new cipher got a error -> ", x)
	} else {
		p, x := cipher.NewGCM(r)
		if x != nil {
			fmt.Println("<RR6> Encryption Module -> AES: Got an error when switching to make new GCM function in function runner, error given to STDOUT -> ", x)
		}
		nonce := make([]byte, p.NonceSize())
		if _, x = io.ReadFull(rand.Reader, nonce); x != nil {
			fmt.Println(x)
		}
		fmt.Println("Data -> ", p.Seal(nonce, nonce, m, nil))
		x = ioutil.WriteFile(filename, p.Seal(nonce, nonce, m, nil), 0777)
		if x != nil {
			fmt.Println(c.REDHB, "<RR6> Encryption module -> AES: Got error when trying to write the input data to the output file -???? ", x)
		} else {
			fmt.Println(c.BLUHB, "<RR6> Information: Data written to filename -> ", filename)
		}
	}
}
