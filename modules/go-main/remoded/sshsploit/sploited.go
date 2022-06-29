package ssh_sploit

import (
	"fmt"
	Ssh_color "main/modg/colors"
	Key_loader "main/modules/go-main/remoded/sshsploit/key_loader"
	Key_signer "main/modules/go-main/remoded/sshsploit/key_reader"
	Ssh_const "main/modules/go-main/remoded/sshsploit/ssh-constants"
	"os"

	"golang.org/x/crypto/ssh"
)

// basically verifying that the host is alive
func Verify(user, host string) {
	u := Key_signer.Key_Sig(Ssh_const.ALGORITHM_SSH_KEY_FILE_LINUX)
	s := Key_loader.Init_Key(Ssh_const.ALGORITHM_SSH_KNOWN_HOSTS_FILE_LINUX)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Incident |>> User Key   -> ", u)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Incident |>> Server Key -> ", s)
	SSH_Conf := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(u),
		},
		HostKeyCallback: ssh.FixedHostKey(s),
		HostKeyAlgorithms: []string{
			Ssh_const.RSA_SSH_ALGORITHM,
			Ssh_const.DSA_SSH_ALGORITHM,
			Ssh_const.ECDSA256_SSH_ALGORITHM,
			Ssh_const.ECDSA384_SSH_ALGORITHM,
			Ssh_const.ECDSA521_SSH_ALGORITHM,
			Ssh_const.ED25519_SSH_ALGORITHM,
		},
	}
	c, x := ssh.Dial(Ssh_const.METHODIZED_PING_SSH_DIALER, host, SSH_Conf)
	if x != nil {
		fmt.Println(Ssh_color.REDHB, "<RR6> Was not able to make a proper dial to the target, exiting....")
		os.Exit(0)
	}
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Client Version | Was able to make a dial | -> ", string(c.ClientVersion()))
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Server Version | Was able to make a dial | -> ", string(c.ServerVersion()))
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Session ID     | Was able to make a dial | -> ", string(c.SessionID()))
}
