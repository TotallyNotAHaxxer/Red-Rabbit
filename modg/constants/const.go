/*
Developer  | ArkAngeL43
Package    | constants
Module     | constants
File       | const
FP         | modules/constants/const.go

Does:
	Represents the constants and sub variables for main.go

high blue : \033[38;5;21m
high pink : \033[38;5;198m
*/

package constants

import (
	"crypto/rand"
	"crypto/tls"
	"debug/pe"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/google/gopacket/layers"
)

/*
============= HEX ==============

09 35 30 2C 20 35 33 2C 20 34 38 2C 20 34 35 2C 20 31 32 30 2C 20 39 39 2C 20 35 37 2C 20 34 38
2C 20 34 36 2C 20 31 31 39 2C 20 31 30 31 2C 20 39 38 2C 0A 09 31 31 35 2C 20 31 30 35 2C 20 31
31 36 2C 20 31 30 31 2C 20 31 31 39 2C 20 31 30 31 2C 20 31 30 38 2C 20 39 39 2C 20 31 31 31 2C
20 31 30 39 2C 0A 09 31 30 31 2C 20 34 36 2C 20 39 39 2C 20 31 31 31 2C 20 31 30 39 2C 20 33 32
2C 20 37 32 2C 20 31 30 31 2C 0A 09 31 30 38 2C 20 31 30 38 2C 20 31 31 31 2C 20 33 32 2C 20 37
31 2C 20 38 30 2C 20 33 32 2C 20 39 31 2C 0A 09 34 39 2C 20 35 30 2C 20 35 30 2C 20 34 36 2C 20
34 39 2C 20 35 34 2C 20 35 30 2C 20 34 36 2C 20 34 39 2C 0A 09 35 32 2C 20 35 31 2C 20 34 36 2C
20 34 39 2C 20 35 33 2C 20 35 35 2C 20 39 33 2C 20 31 33 2C 20 31 30 2C 0A 09 35 30 2C 20 35 33
2C 20 34 38 2C 20 34 35 2C 20 38 33 2C 20 37 33 2C 20 36 30 2C 20 36 39 2C 20 33 32 2C 0A 09 35
33 2C 20 35 30 2C 20 35 32 2C 20 35 30 2C 20 35 36 2C 20 35 36 2C 20 34 38 2C 20 34 38 2C 20 31
33 2C 0A 09 31 30 2C 20 35 30 2C 20 35 33 2C 20 34 38 2C 20 34 35 2C 20 38 30 2C 20 37 33 2C 20
38 30 2C 20 36 39 2C 0A 09 37 36 2C 20 37 33 2C 20 37 38 2C 20 37 33 2C 20 37 38 2C 20 37 31 2C
20 31 33 2C 20 31 30 2C 20 35 30 2C 0A 09 35 33 2C 20 34 38 2C 20 34 35 2C 20 36 35 2C 20 38 35
2C 20 38 34 2C 20 37 32 2C 20 33 32 2C 20 38 30 2C 0A 09 37 36 2C 20 36 35 2C 20 37 33 2C 20 37
38 2C 20 33 32 2C 20 37 36 2C 20 37 39 2C 20 37 31 2C 20 37 33 2C 0A 09 37 38 2C 20 31 33 2C 20
31 30 2C 20 35 30 2C 20 35 33 2C 20 34 38 2C 20 34 35 2C 20 38 33 2C 20 38 34 2C 0A 09 36 35 2C
20 38 32 2C 20 38 34 2C 20 38 34 2C 20 37 36 2C 20 38 33 2C 20 31 33 2C 20 31 30 2C 20 35 30 2C
0A 09 35 33 2C 20 34 38 2C 20 33 32 2C 20 37 32 2C 20 36 39 2C 20 37 36 2C 20 38 30 2C 20 31 33
2C 20 31 30 2C


============ END OF HEX ==============



| EXAMPLE OF BYTE PACKET PAYLOAD
*/
// this is used, just says it isnt
type Loco struct {
	XMLName  xml.Name  `xml:"location"`
	Location C_mac_XML `xml:"coordinates"`
}

type C_mac_XML struct {
	Latitude   string `xml:"latitude,attr"`
	Longitude  string `xml:"longitude,attr"`
	NLatitude  string `xml:"nlatitude,attr"`
	NLongitude string `xml:"nlongitude,attr"`
}

type Empty struct{}

type Result struct {
	IPADDR   string
	HOSTNAME string
}

//structureing and handeling functions

type UrlTitle struct {
	Idx   int
	Url   string
	Title string
}

var ()

const (
	Host                          = "localhost"
	Port_DB                       = "5432"
	Port_host                     = ":5501"
	Second_port                   = ":5502"
	Server_UI_URL                 = "127.0.0.1:8080"
	Method                        = "POST/GET"
	Powered_by                    = "Go"
	Processing                    = "r/perl5"
	Banner_file                   = "banner.txt"
	Defualt_route_test_connection = "https://www.google.com"
	Err_404                       = "html/error.html"
	Err_Passauth_failed_          = "html/auth-error.html"
	Filepath1                     = "/"
	Main_banner_rr                = "/text/banners/rr6.txt"
	Team_logo_rr                  = "/text/banners/team-logo.txt"
	Verticle_banner               = "/text/banners/screen-logo-small.txt"
	Help_flags_x_commands         = "/text/help/commands-flags.txt"
	Shark                         = "/text/banners/shark.txt"
	Sets                          = "/config/sets.json"
	Phone_codes                   = "/config/phone-code.json"
	Verified_commands             = "/config/verified.json"
	Flags                         = "/config/flags.json"
	Domain                        = `^((?!-)[A-Za-z0-9-]{1, 63}(?<!-)\\.)+[A-Za-z]{2, 6}$`
	Ip_port                       = `/(\\d{1,3}\.\\d{1,3}\.\\d{1,3}\.\\d{1,3}\:\\d{1,5})/`
	Port_num                      = `^([1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])`
	ValidIpAddrs                  = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	ValidHostname                 = `^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$`
	ValURL                        = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	EndChunkType                  = "IEND"
)

var (
	R         []string
	Mut       = &sync.Mutex{}
	Clear_hex = "\x1b[H\x1b[2J\x1b[3J"

	SIG_MP4 = []byte{
		'\x00', '\x00', '\x00',
		'\x18', '\x66', '\x74',
		'\x79', '\x70'}

	SIG_PNG = []byte{
		'\x89', '\x50', '\x4E',
		'\x47', '\x0D', '\x0A',
		'\x1A', '\x0A'}

	SIG_PDF = []byte{
		'\x25', '\x50', '\x44',
		'\x46', '\x2d'}

	SIG_BASE64 = []byte{
		'\x62', '\x61', '\x73',
		'\x65', '\x36', '\x34',
		'\x3a'}

	By = []byte{
		50, 53, 48, 45, 120, 99, 57, 48, 46, 119, 101, 98,
		115, 105, 116, 101, 119, 101, 108, 99, 111, 109,
		101, 46, 99, 111, 109, 32, 72, 101,
		108, 108, 111, 32, 71, 80, 32, 91,
		49, 50, 50, 46, 49, 54, 50, 46, 49,
		52, 51, 46, 49, 53, 55, 93, 13, 10,
		50, 53, 48, 45, 83, 73, 60, 69, 32,
		53, 50, 52, 50, 56, 56, 48, 48, 13,
		10, 50, 53, 48, 45, 80, 73, 80, 69,
		76, 73, 78, 73, 78, 71, 13, 10, 50,
		53, 48, 45, 65, 85, 84, 72, 32, 80,
		76, 65, 73, 78, 32, 76, 79, 71, 73,
		78, 13, 10, 50, 53, 48, 45, 83, 84,
		65, 82, 84, 84, 76, 83, 13, 10, 50,
		53, 48, 32, 72, 69, 76, 80, 13, 10,
	}

	SMTP_authPLAIN = []byte{
		65, 85, 84, 72, 32, 80, 76, 65, 73,
		78, 32, 100, 88, 78, 10, 89, 91, 9,
		53, 104, 98, 87, 85, 65, 65, 72, 66,
		10, 49, 95, 17, 85, 19, 85, 17, 41,
		07, 13, 10,
	}

	/*
		HELO
	*/
	SMTP_AUTH_PLAIN_HELO = []byte{
		72, 69, 76, 79, 32, 109, 117, 100,
		121, 110, 97, 109, 105, 99, 115,
		46, 99, 111, 109, 13, 10,
	}

	/*
		250 example message
	*/
	SMTP_AUTH_PLAIN_250 = []byte{
		50, 53, 48, 32, 78, 105, 99, 101, 32, 116,
		111, 32, 109, 101, 101, 116, 32, 121, 111,
		117, 13, 10,
	}

	// do not touch
	Pcap_Files = []string{
		"aaa.pcap", "IMAP - Authenticate CRAM-MD5.cap",
		"ospf-md5_key-1234.pcap", "SMB - NTLMSSP (Windows 10).pcapng",
		"ARP - Broadcast.pcap", "IMAP - Authenticate CRAM-MD5.pcapng",
		"ospf-simple.cap", "SMTP - Auth Cram MD5.pcap", "ARP - Broadcast.pcapng", "IMAP - Authenticate Digest-MD5.pcap",
		"ospf simple password authentication.cap", "SMTP - Auth Cram MD5.pcapng", "Asterisk_ZFONE_XLITE.pcap", "IMAP - Authenticate Digest-MD5.pcapng",
		"OSPF_with_MD5_auth.cap", "SMTP - Auth Login.pcapng", "dot11.pcapng", "IMAP - Authenticate Plain (Base64).pcap", "POP3.pcap",
		"SMTP - Auth Plain.pcap",
		"DTMFsipinfo.pcap", "IMAP - Authenticate Plain (Base64).pcapng", "POP3.pcapng", "SMTP - Auth Plain.pcapng",
		"FAX-Call-t38-CA-TDM-SIP-FB-1.pcap", "IMAP - Authenticate XYMPKI.pcap", "SIP CALL.pcap", "smtp-login.pcap",
		"Ftp.pcap", "IMAP - Authenticate XYMPKI.pcapng", "SIP_DTMF2.pcap", "smtp.pcap",
		"Ftp.pcapng", "IMAP - Login (Plaintext) 1.pcap", "sip-rtp-dvi4.pcap", "TCP - 5 Packets.pcap",
		"h223-over-rtp.cap", "IMAP - Login (Plaintext) 1.pcapng", "sip-rtp-g711.pcap", "TCP - 5 Packets.pcapng",
		"h263-over-rtp.pcap", "IMAP - Login (Plaintext) 2.pcap", "sip-rtp-g722.pcap", "TCP - File Downloads.pcap",
		"HTTP - Basic Authentication.pcap", "IMAP - Login (Plaintext) 2.pcapng", "sip-rtp-g726.pcap", "TCP - File Downloads.pcapng",
		"HTTP - Basic Authentication.pcapng", "IMAP - Login (Plaintext) 3.pcap", "sip-rtp-g729a.pcap", "TCP - Network.pcap",
		"HTTP - Digest Authentication.pcap", "IMAP - Login (Plaintext) 3.pcapng", "sip-rtp-gsm.pcap", "TCP - Network.pcapng",
		"HTTP - Digest Authentication.pcapng", "Kerberos-816.pcap", "sip-rtp-ilbc.pcap", "TCP - Tcpreplay.pcap",
		"HTTP - Digest-MD5.pcap", "Kerberos-816.pcapng", "sip-rtp-l16.pcap", "TCP - Tcpreplay.pcapng",
		"HTTP - Digest-MD5.pcapng", "Kerberos - v5 TCP.pcap", "sip-rtp-lpc.pcap", "Telnet - Char Mode2.pcap",
		"HTTP - JPG Download.pcap", "Kerberos - v5 TCP.pcapng", "sip-rtp_only.pcap", "Telnet - Char Mode2.pcapng",
		"HTTP - JPG Download.pcapng", "Kerberos v5 UDP 2.pcap", "sip-rtp-opus.pcap", "Telnet - Char Mode.pcap",
		"HTTP - NTLM GSSAPI.pcap", "Kerberos v5 UDP 2.pcapng", "sip-rtp-speex.pcap", "Telnet - Char Mode.pcapng",
		"HTTP - NTLM GSSAPI.pcapng", "Kerberos v5 - UDP 3.pcap", "SMB - NTLM cifs_SessionSetupAndX_NTLM_Plain.pcap", "Telnet - Line Mode.pcap",
		"HTTP - NTLM.pcap", "Kerberos v5 - UDP 3.pcapng", "SMB - NTLM cifs_SessionSetupAndX_NTLM_Plain.pcapng", "Telnet - Line Mode.pcapng",
		"HTTP - NTLM.pcapng", "Kerberos - v5 UDP.pcap", "SMB - NTLMSSP Single Session (Windows 10).pcap", "Telnet.pcap",
		"HTTP - PDF file download.pcap", "Kerberos - v5 UDP.pcapng", "SMB - NTLMSSP Single Session (Windows 10).pcapng", "Telnet.pcapng",
		"HTTP - PDF file download.pcapng", "MagicJack+_short_call.pcap", "SMB - NTLMSSP (smb3 aes 128 ccm).pcap", "udp_1_packet.pcap",
		"HTTP - Small File.pcap", "metasploit-sip-invite-spoof.pcap", "SMB - NTLMSSP (smb3 aes 128 ccm).pcapng", "udp_1_packet.pcapng",
		"HTTP - Small File.pcapng", "only_sip_sdp_rtp.pcap", "SMB - NTLMSSP (Windows 10).pcap",
	}

	Filepaths_made = []string{
		"examples/pcap/aaa.pcap",
		"examples/pcap/IMAP - Authenticate CRAM-MD5.cap",
		"examples/pcap/ospf-md5_key-1234.pcap",
		"examples/pcap/SMB - NTLMSSP (Windows 10).pcapng",
		"examples/pcap/ARP - Broadcast.pcap",
		"examples/pcap/IMAP - Authenticate CRAM-MD5.pcapng",
		"examples/pcap/ospf-simple.cap",
		"examples/pcap/SMTP - Auth Cram MD5.pcap",
		"examples/pcap/ARP - Broadcast.pcapng",
		"examples/pcap/IMAP - Authenticate Digest-MD5.pcap",
		"examples/pcap/ospf simple password authentication.cap",
		"examples/pcap/SMTP - Auth Cram MD5.pcapng",
		"examples/pcap/Asterisk_ZFONE_XLITE.pcap",
		"examples/pcap/IMAP - Authenticate Digest-MD5.pcapng",
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
		"examples/pcap/examples/pcap/HTTP - Basic Authentication.pcapng",
		"examples/pcap/IMAP - Login (Plaintext) 3.pcap",
		"examples/pcap/sip-rtp-g729a.pcap",
		"examples/pcap/TCP - Network.pcap",
		"examples/pcap/examples/pcap/HTTP - Digest Authentication.pcap",
		"examples/pcap/IMAP - Login (Plaintext) 3.pcapng",
		"examples/pcap/sip-rtp-gsm.pcap",
		"examples/pcap/TCP - Network.pcapng",
		"examples/pcap/HTTP - Digest Authentication.pcapng",
		"examples/pcap/Kerberos-816.pcap",
		"examples/pcap/sip-rtp-ilbc.pcap",
		"examples/pcap/TCP - Tcpreplay.pcap",
		"examples/pcap/examples/pcap/HTTP - Digest-MD5.pcap",
		"examples/pcap/Kerberos-816.pcapng",
		"examples/pcap/sip-rtp-l16.pcap",
		"examples/pcap/TCP - Tcpreplay.pcapng",
		"examples/pcap/HTTP - Digest-MD5.pcapng",
		"examples/pcap/Kerberos - v5 TCP.pcap",
		"examples/pcap/sip-rtp-lpc.pcap",
		"examples/pcap/Telnet - Char Mode2.pcap",
		"examples/pcap/HTTP - JPG Download.pcap",
		"examples/pcap/Kerberos - v5 TCP.pcapng", "sip-rtp_only.pcap",
		"examples/pcap/Telnet - Char Mode2.pcapng",
		"examples/pcap/HTTP - JPG Download.pcapng",
		"examples/pcap/Kerberos v5 UDP 2.pcap", "sip-rtp-opus.pcap",
		"examples/pcap/Telnet - Char Mode.pcap",
		"examples/pcap/examples/pcap/HTTP - NTLM GSSAPI.pcap",
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
		"examples/pcap/HTTP - PDF file download.pcap", "examples/pcap/Kerberos - v5 UDP.pcapng",
		"examples/pcap/SMB - NTLMSSP Single Session (Windows 10).pcapng", "examples/pcap/Telnet.pcapng",
		"examples/pcap/HTTP - PDF file download.pcapng", "examples/pcap/MagicJack+_short_call.pcap",
		"examples/pcap/SMB - NTLMSSP (smb3 aes 128 ccm).pcap",
		"examples/pcap/udp_1_packet.pcap",
		"examples/pcap/HTTP - Small File.pcap", "examples/pcap/metasploit-sip-invite-spoof.pcap",
		"examples/pcap/SMB - NTLMSSP (smb3 aes 128 ccm).pcapng",
		"examples/pcap/udp_1_packet.pcapng",
		"examples/pcap/HTTP - Small File.pcapng",
		"examples/pcap/only_sip_sdp_rtp.pcap", "examples/pcap/SMB - NTLMSSP (Windows 10).pcap",
	}

	Windows_data_directories = []string{
		"IMAGE_DIRECTORY_ENTRY_EXPORT",
		"IMAGE_DIRECTORY_ENTRY_IMPORT",
		"IMAGE_DIRECTORY_ENTRY_RESOURCE",
		"IMAGE_DIRECTORY_ENTRY_EXCEPTION",
		"IMAGE_DIRECTORY_ENTRY_SECURITY",
		"IMAGE_DIRECTORY_ENTRY_BASERELOC",
		"IMAGE_DIRECTORY_ENTRY_DEBUG",
		"IMAGE_DIRECTORY_ENTRY_COPYRIGHT",
		"IMAGE_DIRECTORY_ENTRY_GLOBALPTR",
		"IMAGE_DIRECTORY_ENTRY_TLS",
		"IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG",
		"IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT",
		"IMAGE_DIRECTORY_ENTRY_IAT",
		"IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT",
		"IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR",
		"IMAGE_NUMBEROF_DIRECTORY_ENTRIES",
	}

	SizeofOptionalHeader32 = uint16(binary.Size(pe.OptionalHeader32{}))
	SizeofOptionalHeader64 = uint16(binary.Size(pe.OptionalHeader64{}))
	X32                    pe.OptionalHeader32
	X64                    pe.OptionalHeader64

	Paramaters_SMTP = []string{
		"RCPT TO",
		"MAIL FROM",
		"USER",
		"AUTH",
	}

	Paramaters_SMTP_plain = []string{
		"AUTH PLAIN",
		"MAIL FROM",
		"RCPT TO",
		"HELO",
		"250",
	}

	Paramaters_IMAP_PLAIN_TEXT = []string{
		"LOGIN",
	}

	Paramaters_HTTP = []string{
		"GET",
		"POST",
		"HTTP/1.1",
		"Authorization",
		"Unauthorized",
	}

	Paramaters_HTTP_img = []string{
		"GIF",
		"IMG",
		"JPG",
		"JPEG",
		"",
	}

	Paramaters_SIP = []string{
		"+OK",
		"PASS",
		"USER",
		"POST",
		"GET",
		"INVITE",
		"REGISTER",
	}

	Hashinput string
	Hashkey   string
	Num       uint32
	UrlQueue  = make(chan string)
	config    = &tls.Config{InsecureSkipVerify: true}
	Transport = &http.Transport{
		TLSClientConfig: config,
	}
	HasCrawled                   = make(map[string]bool)
	NetClient                    *http.Client
	ETHER_                       = layers.LayerTypeEthernet
	TCP_                         = layers.LayerTypeTCP
	IP_                          = layers.LayerTypeIPv4
	IP_6                         = layers.LayerTypeIPv6
	DHCMP_                       = layers.LayerTypeDHCPv4
	DHCMP_6                      = layers.LayerTypeDHCPv6
	ARP_                         = layers.LayerTypeARP
	DNS_                         = layers.LayerTypeDNS
	ICMP_6_NA_                   = layers.LayerTypeICMPv6NeighborAdvertisement
	ICMP_6_SOL_                  = layers.LayerTypeICMPv6NeighborSolicitation
	ICP_4_                       = layers.LayerTypeICMPv4
	ICMP_6_                      = layers.LayerTypeICMPv6
	ICMP_ECH_                    = layers.LayerTypeICMPv6Echo
	ICMP_router_                 = layers.LayerTypeICMPv6RouterAdvertisement
	ICMP_REDI_                   = layers.LayerTypeICMPv6Redirect
	UDP_                         = layers.LayerTypeUDP
	DOT11_                       = layers.LayerTypeDot11
	DOT11_BEC_                   = layers.LayerTypeDot11MgmtBeacon
	DOT_11_DEIS_                 = layers.LayerTypeDot11MgmtDeauthentication
	DOT_11_DISASS_               = layers.LayerTypeDot11MgmtDisassociation
	DOT_11_PROBE_                = layers.Dot11InformationElementIDSSID
	DOT_11_PROBE_REQ_            = layers.LayerTypeDot11MgmtProbeReq
	DOT_11_PROBE_RESP_           = layers.LayerTypeDot11MgmtProbeResp
	DOT_11_WEP_                  = layers.LayerTypeDot11WEP
	DOT_11_RAD_TAP_              = layers.LayerTypeRadioTap
	ICMP_IGMP                    = layers.LayerTypeIGMP
	MLDv2MulticastListenerReport = layers.LayerTypeMLDv2MulticastListenerReport
	Multicase_listener           = layers.LayerTypeMLDv2MulticastListenerQuery
	IPv6HopByHop                 = layers.LayerTypeIPv6HopByHop
	IPv6_Router_advertisement    = layers.LayerTypeICMPv6RouterAdvertisement
	Perl_CPAN_Brute              = "modules/perl/files/cpan-brute.pl"
	READ                         = rand.Reader
)

func Parse_filepath(filepath string) (string, error) {
	cw, err := os.Getwd()
	if err != nil {
		fmt.Println("<RR6> Constants Module: Could not get or read users working directory -> ", err)
	}
	fp := cw + filepath
	return fp, nil
}
