package AES

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	c "main/modg/colors"
)

func Dec_data(k []byte, f string) {
	t, x := ioutil.ReadFile(f)
	if x != nil {
		fmt.Println(c.REDHB, "<RR6> Encryption module -> AES: Got error when trying to use the IO utils to read the file or data in the file -> ", x)
	} else {
		l, x := aes.NewCipher(k)
		if x != nil {
			fmt.Println(c.REDHB, "<RR6> Encryption Module -> AES: Got error when trying to create or call a new AES function to cipher method -> ", x)
		} else {
			p, x := cipher.NewGCM(l)
			if x != nil {
				fmt.Println(c.REDHB, "<RR6> Encryption Module -> AES: Got error when trying to create a new cipher method for GCM -> ", x)
			} else {
				nonceSize := p.NonceSize()
				if len(t) < nonceSize {
					fmt.Println(x)
				}

				nonce, ciphertext := t[:nonceSize], t[nonceSize:]
				text, x := p.Open(nil, nonce, ciphertext, nil)
				if x != nil {
					fmt.Printf(c.REDHB, "<RR6> Encryption Module -> AES: Got error when trying to call ct open func, -> ", x)
				} else {
					fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Plain Text Data -> ", string(text))
				}
			}
		}
	}
}
