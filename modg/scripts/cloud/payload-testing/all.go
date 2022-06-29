// package will test wether or not the payload past the list of hosts / IDS's
package payloadtesting_403

import (
	"fmt"
	"log"
	Cloud_Requests "main/modg/requests"
	Cloud_Constants "main/modg/scripts/cloud/types"
	"strings"
)

func Payload(payload string) {
	parser := "trippedIDS?=" + payload
	for _, detected := range Cloud_Constants.TRIPPED_IDS_CONFIG {
		ret := Cloud_Requests.Remove_Params(detected.URL, parser)
		response, stat, x := Cloud_Requests.GET_RESP(ret)
		if x != nil {
			log.Fatal(x)
		} else {
			if strings.Contains(response, detected.BLOCKED_MSG) {
				fmt.Printf(" | Payload has been detected | < STAT > | Blocked | < Service > |  %s  | < Code > |  %v  | \n", detected.IDS_NAME, detected.BLOCKED_CODE)
				continue
			}
			if stat == detected.BLOCKED_CODE {
				fmt.Printf(" | Payload has been detected | < STAT > | Blocked | < Service > |  %s  | < Code > |  %v  | \n", detected.IDS_NAME, detected.BLOCKED_CODE)
				continue
			}
			fmt.Printf(" | Payload has sucessfully bypassed | < STAT > | Passed | < Service > |  %s  | < Code > |  %v  | \n", detected.IDS_NAME, detected.BLOCKED_CODE)
		}

	}
}
