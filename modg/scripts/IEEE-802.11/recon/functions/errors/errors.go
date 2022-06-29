//https://cs.opensource.google/go/go/+/refs/tags/go1.18.1:src/errors/errors.go
// the reason i wanted to create my own errors package based on the one from the OFFICAL GO
// source code was to test the modifcation limits of this script, in fact the whole reasonn for the modules outside of organization was to see how flexible this code can be based on how it was built
// hence the credits
// make sure to read :D
package errors

import (
	"fmt"
	ARP_CONSTANTS "main/modg/scripts/IEEE-802.11/IEEE-802.11-c"
	"net"
)

type Error_Return struct {
	s string
}

func (e *Error_Return) Error() string {
	return e.s
}

func Error(msg string) error {
	if msg != "" {
		return &Error_Return{msg}
	}
	return nil
}

func Sanity(card *net.IPNet) error {
	if card == nil {
		return Error("[-] Address was nil or empty  | Skipping network card\n")
	} else if card.IP[0] == ARP_CONSTANTS.LOADDR {
		return Error("[-] Address was 127 or local  | Skipping network card\n")
	} else if card.Mask[0] != ARP_CONSTANTS.NULLMARK_1 || ARP_CONSTANTS.Addresses.Mask[1] != ARP_CONSTANTS.NULLMARK_2 {
		return Error("[-] Address mask was to large | Skipping network card\n")
	}
	return nil
}

func Return_err(x error, msg string, color string) {
	if x != nil {
		fmt.Println(color, msg, x)
	}
}
