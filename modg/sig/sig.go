/*
Package     | Sig
Developer   | ArkAngeL43
Module      | sig.go
File        | modg/sig
Nest        | sig

Does:
	Holds all interupt functions
*/
package sig

import (
	"fmt"
	"os"
	"os/signal"
)

const (
	Interupt1 = "<RR6> Signal Module: Interupt has been deteted, exiting with code 0"
	Kil1      = "<RR6> Signal Module: Operating System Kill Call has been recived, exiting...."
)

// sig handel defualt
func Handelreturncon(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println(Interupt1)
			os.Exit(0)
		case os.Kill:
			fmt.Println(Kil1)
			os.Exit(0)
		}
	}
}
