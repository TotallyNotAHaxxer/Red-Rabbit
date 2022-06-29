package Testing

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// gets filepath and parses the files with the working filepath
func Getfp() {
	fp, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fp)
	test_file_visi("")
	testd("")
	for i, s := range Files {
		fmt.Println(fp, "/", s)
		fmt.Println("TEST: --------------------------- File Number -> ", i)
		var name = fp + "/" + s
		Ouput(name)
	}
}

// test if the file is real and exists, rather than the filepath it makes the os stat the file
func test_file_visi(filename string) {
	for i, s := range Files {
		time.Sleep(60 * time.Millisecond)
		if _, err := os.Stat(s); err == nil {
			fmt.Println("TEST: --------------------------- \033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File exists | ", i)
			fmt.Println("[    +     ] TEST HAS PASSED!!!!")
		} else {
			log.Fatal(err)
		}
	}
}

// output and test all files, this is like an EOF checker or file output error
func Ouput(filename string) {
	content, err := os.Open(filename)
	if err != nil {
		fmt.Println("TEST: --------------------------- [-] File DOES NOT exists | ", filename)
		fmt.Println("[    -     ] TEST HAS FAILED!!!!")
		panic(err)
	} else {
		scanner := bufio.NewScanner(content)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			fmt.Println("TEST: --------------------------- [-] File DOES EXIST!!!!!!!!!!!!!!!!!!| ", filename)
			fmt.Println("\033[32m[    +     ] TEST HAS PASSED!!!! -> OUTPUT of line -> UNKNOWN")
		}
	}
}

// test is a dir dir is needed, if it is not a dir it needs to be changed
func testd(filename string) {
	for i, s := range Files {
		time.Sleep(60 * time.Millisecond)
		f, err := os.Open(s)
		if err != nil {
			log.Fatal(err)
		} else {
			defer f.Close()
		}
		l, err := f.Stat()
		if err != nil {
			log.Fatal(err)
		} else {
			if l.IsDir() {
				fmt.Println("\033[38m[     +    ] FILE -> ", filename, i, " <- is a directory")
			}
		}
	}
}

// all files that need to be checked and outputted
var Files = []string{
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
	"config/be.json",
	"config/flags.json",
	"config/phone-code.json",
	"config/sets.json",
	"config/verified.json",
	"config/oui.txt",
	"config/cmd.txt",
	"text/banners/rr6.txt",
	"text/banners/screen-logo-small.txt",
	"text/banners/team-logo.txt",
	"text/help/advance.txt",
	"text/help/commands-flags.txt",
}
