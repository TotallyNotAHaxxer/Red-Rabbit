package AES

import (
	"crypto/rand"
	"log"
)

func Generate_AES_Key() []byte {
	randomBytes := make([]byte, 32) // 32 bytes, 256 bit
	numBytesRead, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Error generating random key.", err)
	}
	if numBytesRead != 32 {
		log.Fatal("Error generating 32 random bytes for key.")
	}
	return randomBytes
}
