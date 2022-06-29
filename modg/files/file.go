/*
Developer | ArkangeL43
Package   | files
Module    | file.go
File      | modg/files
Nest      | files

Does:
	Imports and defines all filepaths and constants
*/

package files

import (
	"fmt"
	"log"
	"os"
)

func Getfp() {
	fp, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fp)
	for i, s := range Files {
		fmt.Println(fp, s)
		fmt.Println("File Number -> ", i)
	}
}

var Files = []string{
	"--------------------- EXAMPLE NETWORK CAPTURES ------------------------",
	"examples/pcap/IMAP - Authenticate CRAM-MD5.cap",
	"examples/pcap/ospf-md5_key-1234.pcap",
	"examples/pcap/SMB - NTLMSSP (Windows 10).pcapng",
	"examples/pcap/ARP - Broadcast.pcap",
	"examples/pcap/IMAP - Authenticate CRAM-MD5.pcapng",
	"examples/pcap/ospf-simple.cap",
	"examples/pcap/SMTP - Auth Cram MD5.pcap",
	"examples/pcap/ARP - Broadcast.pcapng",
	"examples/pcap/OSPF_with_MD5_auth.cap",
	"examples/pcap/SMTP - Auth Login.pcapng",
	"examples/pcap/dot11.pcapng",
	"examples/pcap/IMAP - Authenticate Plain (Base64).pcap",
	"examples/pcap/POP3.pcap",
	"examples/pcap/SMTP - Auth Plain.pcap",
	"examples/pcap/DTMFsipinfo.pcap",
	"examples/pcap/IMAP - Authenticate Plain (Base64).pcapng",
	"examples/pcap/POP3.pcapng",
	"examples/pcap/SMTP - Auth Plain.pcapng",
	"examples/pcap/FAX-Call-t38-CA-TDM-SIP-FB-1.pcap",
	"examples/pcap/IMAP - Authenticate XYMPKI.pcap",
	"examples/pcap/SIP CALL.pcap",
	"examples/pcap/smtp-login.pcap",
	"examples/pcap/Ftp.pcap",
	"examples/pcap/IMAP - Authenticate XYMPKI.pcapng",
	"examples/pcap/SIP_DTMF2.pcap",
	"examples/pcap/smtp.pcap",
	"examples/pcap/Ftp.pcapng",
	"examples/pcap/IMAP - Login (Plaintext) 1.pcap",
	"examples/pcap/sip-rtp-dvi4.pcap",
	"examples/pcap/TCP - 5 Packets.pcap",
	"examples/pcap/h223-over-rtp.cap",
	"examples/pcap/IMAP - Login (Plaintext) 1.pcapng",
	"examples/pcap/sip-rtp-g711.pcap",
	"examples/pcap/TCP - 5 Packets.pcapng",
	"examples/pcap/h263-over-rtp.pcap",
	"examples/pcap/IMAP - Login (Plaintext) 2.pcap",
	"examples/pcap/sip-rtp-g722.pcap",
	"examples/pcap/TCP - File Downloads.pcap",
	"examples/pcap/HTTP - Basic Authentication.pcap",
	"examples/pcap/IMAP - Login (Plaintext) 2.pcapng",
	"examples/pcap/sip-rtp-g726.pcap",
	"examples/pcap/TCP - File Downloads.pcapng",
	"examples/pcap/HTTP - Basic Authentication.pcapng",
	"examples/pcap/IMAP - Login (Plaintext) 3.pcap",
	"examples/pcap/sip-rtp-g729a.pcap",
	"examples/pcap/TCP - Network.pcap",
	"examples/pcap/HTTP - Digest Authentication.pcap",
	"examples/pcap/IMAP - Login (Plaintext) 3.pcapng",
	"examples/pcap/sip-rtp-gsm.pcap",
	"examples/pcap/TCP - Network.pcapng",
	"examples/pcap/HTTP - Digest Authentication.pcapng",
	"examples/pcap/Kerberos-816.pcap",
	"examples/pcap/sip-rtp-ilbc.pcap",
	"examples/pcap/TCP - Tcpreplay.pcap",
	"examples/pcap/HTTP - Digest-MD5.pcap",
	"examples/pcap/Kerberos-816.pcapng",
	"examples/pcap/sip-rtp-l16.pcap",
	"examples/pcap/TCP - Tcpreplay.pcapng",
	"examples/pcap/HTTP - Digest-MD5.pcapng",
	"examples/pcap/Kerberos - v5 TCP.pcap",
	"examples/pcap/sip-rtp-lpc.pcap",
	"examples/pcap/Telnet - Char Mode2.pcap",
	"examples/pcap/HTTP - JPG Download.pcap",
	"examples/pcap/Kerberos - v5 TCP.pcapng",
	"examples/pcap/sip-rtp_only.pcap",
	"examples/pcap/Telnet - Char Mode2.pcapng",
	"examples/pcap/HTTP - JPG Download.pcapng",
	"examples/pcap/Kerberos v5 UDP 2.pcap",
	"examples/pcap/sip-rtp-opus.pcap",
	"examples/pcap/Telnet - Char Mode.pcap",
	"examples/pcap/HTTP - NTLM GSSAPI.pcap",
	"examples/pcap/Kerberos v5 UDP 2.pcapng",
	"examples/pcap/sip-rtp-speex.pcap",
	"examples/pcap/Telnet - Char Mode.pcapng",
	"examples/pcap/HTTP - NTLM GSSAPI.pcapng",
	"examples/pcap/Kerberos v5 - UDP 3.pcap",
	"examples/pcap/SMB - NTLM cifs_SessionSetupAndX_NTLM_Plain.pcap",
	"examples/pcap/Telnet - Line Mode.pcap",
	"examples/pcap/HTTP - NTLM.pcap",
	"examples/pcap/Kerberos v5 - UDP 3.pcapng",
	"examples/pcap/SMB - NTLM cifs_SessionSetupAndX_NTLM_Plain.pcapng",
	"examples/pcap/Telnet - Line Mode.pcapng",
	"examples/pcap/HTTP - NTLM.pcapng",
	"examples/pcap/Kerberos - v5 UDP.pcap",
	"examples/pcap/SMB - NTLMSSP Single Session (Windows 10).pcap",
	"examples/pcap/Telnet.pcap",
	"examples/pcap/HTTP - PDF file download.pcap",
	"examples/pcap/Kerberos - v5 UDP.pcapng",
	"examples/pcap/SMB - NTLMSSP Single Session (Windows 10).pcapng",
	"examples/pcap/Telnet.pcapng",
	"examples/pcap/HTTP - PDF file download.pcapng",
	"examples/pcap/MagicJack+_short_call.pcap",
	"examples/pcap/SMB - NTLMSSP (smb3 aes 128 ccm).pcap",
	"examples/pcap/udp_1_packet.pcap",
	"examples/pcap/HTTP - Small File.pcap",
	"examples/pcap/metasploit-sip-invite-spoof.pcap",
	"examples/pcap/SMB - NTLMSSP (smb3 aes 128 ccm).pcapng",
	"examples/pcap/udp_1_packet.pcapng",
	"examples/pcap/HTTP - Small File.pcapng",
	"examples/pcap/only_sip_sdp_rtp.pcap",
	"examples/pcap/SMB - NTLMSSP (Windows 10).pcap",
	"----------------------------- STEGONOGRAPHY -----------------------------",
	"examples/stego/cat1.gif",
	"examples/stego/download.jpeg",
	"examples/stego/example.pdf",
	"examples/stego/injected.png",
	"examples/stego/main.gif",
	"examples/stego/main.jpg",
	"examples/stego/main.webp",
	"-------------------------------- HASHING AND CRYPTOGRAPHY -----------------",
	"examples/hash/sha256.txt",
	"examples/hash/MD5.txt",
}
