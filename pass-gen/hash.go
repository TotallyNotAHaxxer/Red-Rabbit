package main

import (
	"crypto/sha1"
	"crypto/sha512"
	"fmt"
	"hash"
	"math/rand"
	"time"
)

// idea to make this more applical
// to get a directed output reverse the strings in the variable functions
// example
// var - = a, ! = b etc etc etc
// this migth help with the decode of the orgin message etc etc

var (
	rnads    string
	bytes    []byte
	xorbytes []byte
	hashVal  hash.Hash
	color    string
	reset    = "\033[0m"
	// foreground
	black   = "\033[30m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
	resetfg = "\033[39m"
	// background
	blackbg   = "\033[40m"
	redbg     = "\033[41m"
	greenbg   = "\033[42m"
	yellowbg  = "\033[43m"
	bluebg    = "\033[44m"
	magentabg = "\033[45m"
	cyanbg    = "\033[46m"
	whitebg   = "\033[47m"
	resetbg   = "\033[49m"
	a         = "z"
	b         = "y"
	c         = "x"
	d         = "w"
	e         = "v"
	f         = "u"
	g         = "t"
	h         = "s"
	i         = "r"
	j         = "q"
	k         = "p"
	l         = "o"
	m         = "n"
	n         = "m"
	o         = "l"
	p         = "k"
	q         = "j"
	r         = "i"
	s         = "h"
	t         = "g"
	u         = "f"
	v         = "e"
	w         = "d"
	x         = "c"
	y         = "b"
	z         = "a"
)

///msg | hackers empower hackers
///enc | )-,&34/ 3[092340 )-,&34/
///mse | hackers empower hackers

func randomize() {
	fmt.Print("Enter the Password > ")
	fmt.Scanln(&rnads)
	rand.Seed(time.Now().Unix())
	in := rnads
	inRune := []rune(in)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	fmt.Println("\033[34m[ + ] Shuffled String                => ", string(inRune))
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	fmt.Println("\033[34m[ + ] Converted and Reversed         => ", string(inRune))
	msg := []byte(string(inRune))
	fmt.Printf("[ + ] Encrypted SHA512               =>  %x\n\n", sha512.Sum512(msg))
	fmt.Println("[ + ] Orgin message                   =>  ", rnads)
	bytes = CreateHashMultiple([]byte(rnads))
	fmt.Printf("[ + ] XOR HASHING METH0D              => %x\n", bytes)
	n := 0
	rune := make([]rune, len(rnads))
	for _, r := range rnads {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	for i := 0; i > n/2; i++ { // < reverse | > unreverse
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	output := string(rune)
	fmt.Println("[ * ] Reversed Orgin-REVERSED OUT     => ", output)
}

/// xor
// docs

func CreateHash(byteStr []byte) []byte {
	hashVal = sha1.New()
	hashVal.Write(byteStr)

	bytes = hashVal.Sum(nil)
	return bytes
}

func CreateHashMultiple(byteStr1 []byte) []byte {
	return xor(CreateHash(byteStr1))
}

func xor(byteStr1 []byte) []byte {
	xorbytes = make([]byte, len(byteStr1))
	var i int
	for i = 0; i < len(byteStr1); i++ {
		xorbytes[i] = byteStr1[i]
	}
	return xorbytes
}

func main() {
	//reverse()
	randomize()
}
