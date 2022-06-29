package AES

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	cc "main/modg/constants"
)

func Enc(k, m []byte) ([]byte, uint32) {
	n, x := aes.NewCipher(k)
	if x != nil {
		fmt.Println("<RR6> Encryption Module -> AES: Got error when trying to make a new cipher with the key -> ", x)
	} else {
		a := len(m)
		t := make([]byte, aes.BlockSize+a)
		block := t[:aes.BlockSize]
		_, x := io.ReadFull(cc.READ, block)
		if x != nil {
			fmt.Println("<RR6> Encryption Module -> AES: Got error when trying to use the I/O to read the full enc block -> ", x)
		} else {
			mo := cipher.NewCFBEncrypter(n, block)
			mo.XORKeyStream(t[aes.BlockSize:], m)
			return t, 0x00
		}
	}
	return []byte{0, 0, 0, 0}, 0x00
}
