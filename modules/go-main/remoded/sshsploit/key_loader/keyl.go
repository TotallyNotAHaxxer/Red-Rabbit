package keyloader

import (
	"io/ioutil"

	sshc "main/modg/colors"
	sshe "main/modules/go-main/remoded/sshsploit/ssh-errors"

	"golang.org/x/crypto/ssh"
)

func Init_Key(File string) ssh.PublicKey {
	p, x := ioutil.ReadFile(File)
	sshe.Error_And_Exit("<RR6> SSH Module: Could not get the key file or read the filename, error at -> ", sshc.REDHB, x)
	_, _, Key, _, _, z := ssh.ParseKnownHosts(p)
	sshe.Error_And_Exit("<RR6> SSH Module: Could not parse the known hosts list, error at -> ", sshc.REDHB, z)
	return Key
}
