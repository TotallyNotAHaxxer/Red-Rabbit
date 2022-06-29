package ServicesEtcUtils

import (
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func Agent() ssh.AuthMethod {
	if l, x := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); x == nil {
		return ssh.PublicKeysCallback(agent.NewClient(l).Signers)
	}
	return nil
}
