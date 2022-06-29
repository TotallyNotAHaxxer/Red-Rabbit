package AES

import (
	"crypto/aes"
	"crypto/cipher"
)

// k = key
// t = text
func Dec(k, t []byte) ([]byte, uint16) {
	o, x := aes.NewCipher(k)
	if x != nil {
		return nil, 0x00
	}
	z := t[:aes.BlockSize]
	t = t[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(o, z)
	cfb.XORKeyStream(t, t)
	return t, 0x00
}
