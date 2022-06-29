package ssh_constants

import "golang.org/x/crypto/ssh"

const (
	RSA_SSH_ALGORITHM                    = ssh.KeyAlgoRSA       // Allow the RSA SSG Key algorithm
	DSA_SSH_ALGORITHM                    = ssh.KeyAlgoDSA       // Allow the DSA SSH Key algorithm
	ECDSA256_SSH_ALGORITHM               = ssh.KeyAlgoECDSA256  // Allow the ECDSA256 SSH Key algorithm
	ECDSA384_SSH_ALGORITHM               = ssh.KeyAlgoECDSA384  // Allow the ECDSA384 SSH Key algorithm
	ECDSA521_SSH_ALGORITHM               = ssh.KeyAlgoECDSA521  // Allow the ECDSA521 SSH Key algorithm
	ED25519_SSH_ALGORITHM                = ssh.KeyAlgoED25519   // Allow the ED25519 SSH Key algorithm
	ALGORITHM_SSH_KEY_FILE_LINUX         = "~/.ssh/id_rsa"      // this should be the standard path, if you have ssh installed it should be here at least im sure.
	ALGORITHM_SSH_KNOWN_HOSTS_FILE_LINUX = "~/.ssh/known_hosts" // this again should be the standard path for known hosts, but then again its just experimental and not tested
	METHODIZED_PING_SSH_DIALER           = "tcp"                // this is a defualt setting, but i kept it here in the case i do not want to use tcp or i will be uising a different method of SSH Ping
)
