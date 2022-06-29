package ssh_Key

import (
	"io/ioutil"

	sshc "main/modg/colors"
	sshe "main/modules/go-main/remoded/sshsploit/ssh-errors"

	"golang.org/x/crypto/ssh"
)

func Key_Sig(file string) ssh.Signer {
	Key, e := ioutil.ReadFile(file)
	sshe.Error_And_Exit("<RR6> SSH Module: Could not get the key file or read the filename, error at -> ", sshc.REDHB, e)
	ctxp, x := ssh.ParsePrivateKey(Key)
	sshe.Error_And_Exit("<RR6> SSH Module: Could not parse the SSH Private key, error at -> ", sshc.REDHB, x)
	return ctxp
}
