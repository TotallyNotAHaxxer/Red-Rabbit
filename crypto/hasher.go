/*

Package       => main
Developer 	  => ArkAngeL43
Type		  => CLI
Does		  => General cryptography with go, hash a password or a list of passwords

*/
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"io"
)

var (
	password_list = flag.String("wordlist", "", "Set a wordlist for hashing")
	Generic_passq = flag.String("password", "", "Set a single password for hashing")
	hash_types    = "MD2, MD5, SH1, SHA256, SHA-384, SHA-512"
)

//(Go) Hash Algorithms: SHA-1, HAVAL, MD2, MD5, SHA-256, SHA-384, SHA-512

func hash_main(key string) string {
	hashing := md5.New()
	hashing.Write([]byte(key))
	return hex.EncodeToString(hashing.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(hash_main(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(hash_main(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
