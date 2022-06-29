package imgvars

var Extensions = []string{
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

var Paramaters_SMTP = []string{
	"RCPT TO",
	"MAIL FROM",
	"USER",
	"AUTH",
}

var Paramaters_SMTP_plain = []string{
	"AUTH PLAIN",
	"MAIL FROM",
	"RCPT TO",
	"HELO",
	"250",
}

var Paramaters_IMAP_PLAIN_TEXT = []string{
	"LOGIN",
}

var Paramaters_HTTP = []string{
	"GET",
	"POST",
	"HTTP/1.1",
	"Authorization",
	"Unauthorized",
}

var Paramaters_HTTP_img = []string{
	"GIF",
	"IMG",
	"JPG",
	"JPEG",
	"",
}

var Paramaters_SIP = []string{
	"+OK",
	"PASS",
	"USER",
	"POST",
	"GET",
	"INVITE",
	"REGISTER",
}

var MassSign = []FileSig{
	{`474946`, `*.gif`, `GIF files`}, {`GIF89a`, `*.gif`, `GIF files`}, {`FFD8FF`, `*.jpg`, `JPEG files`}, {`JFIF`, `*.jpg`, `JPEG files`}, {`504B03`, `*.zip`, `ZIP files`}, {`25504446`, `*.pdf`, `PDF files`},
	{`%PDF`, `*.pdf`, `PDF files`}, {`006E1EF0`, `*.ppt`, `PPT`}, {`A0461DF0`, `*.ppt`, `PPT`}, {`ECA5C100`, `*.doc`, `Doc file`}, {`000100005374616E64617264204A6574204442`, `*.mdb`, `Microsoft database`},
	{`Standard Jet DB`, `*.mdb`, `Microsoft database`}, {`2142444E`, `*.pst`, `PST file`}, {`!BDN`, `*.pst`, `PST file`}, {`0908100000060500`, `*.xls`, `XLS file`}, {`D0CF11E0A1B11AE1`, `*.msi`, `MSI file`},
	{`D0CF11E0A1B11AE1`, `*.doc`, `DOC`}, {`D0CF11E0A1B11AE1`, `*.xls`, `Excel`}, {`D0CF11E0A1B11AE1`, `*.vsd`, `Visio`}, {`D0CF11E0A1B11AE1`, `*.ppt`, `PPT`}, {`0A2525454F460A`, `*.pdf`, `PDF file`}, {`.%%EOF.`, `*.pdf`, `PDF file`},
	{`4040402000004040`, `*.hlp`, `HLP file`}, {`465753`, `*.swf`, `SWF file`}, {`FWS`, `*.swf`, `SWF file`}, {`CWS`, `*.swf`, `SWF file`}, {`494433`, `*.mp3`, `MP3 file`}, {`ID3`, `*.mp3`, `MP3 file`}, {`MSCF`, `*.cab`, `Cab file`},
	{`0x4D534346`, `*.cab`, `Cab file`}, {`ITSF`, `*.chm`, `Compressed Help`}, {`49545346`, `*.chm`, `Compressed Help`}, {`4C00000001140200`, `*.lnk`, `Link file`}, {`4C01`, `*.obj`, `OBJ file`}, {`4D4D002A`, `*.tif`, `TIF graphics`},
	{`MM`, `*.tif`, `TIF graphics`}, {`000000186674797033677035`, `*.mp4`, `MP4 Video`}, {`ftyp3gp5`, `*.mp4`, `MP4 Video`}, {`0x00000100`, `*.ico`, `Icon file`}, {`300000004C664C65`, `*.evt`, `Event file`}, {`LfLe`, `*.evt`, `Event file`},
	{`38425053`, `*.psd`, `Photoshop file`}, {`8BPS`, `*.psd`, `Photoshop file`}, {`4D5A`, `*.ocx`, `Active X`}, {`4D6963726F736F66742056697375616C2053747564696F20536F6C7574696F6E2046696C65`, `*.sln`, `Microsft SLN file`},
	{`Microsoft Visual Studio Solution File`, `*.sln`, `Microsft SLN file`}, {`504B030414000600`, `*.docx`, `Microsoft DOCX file`}, {`504B030414000600`, `*.pptx`, `Microsoft PPTX file`}, {`504B030414000600`, `*.xlsx`, `Microsoft XLSX file`},
	{`504B0304140008000800`, `*.xlsx`, `Java JAR file`}, {`415649204C495354`, `*.avi`, `AVI file`}, {`AVI LIST`, `*.avi`, `AVI file`}, {`57415645666D7420`, `*.wav`, `WAV file`}, {`WAVEfmt`, `*.wav`, `WAV file`}, {`Rar!`, `*.rar`, `RAR file`},
	{`526172211A0700`, `*.rar`, `RAR file`}, {`52657475726E2D506174683A20`, `*.eml`, `EML file`}, {`Return-Path:`, `*.eml`, `EML file`}, {`6D6F6F76`, `*.mov`, `MOV file`}, {`moov`, `*.mov`, `MOV file`}, {`7B5C72746631`, `*.rtf`, `RTF file`},
	{`{\rtf1`, `*.rtf`, `RTF file`}, {`89504E470D0A1A0A`, `*.png`, `PNG file`}, {`PNG`, `*.png`, `PNG file`}, {`C5D0D3C6`, `*.eps`, `EPS file`}, {`CAFEBABE`, `*.class`, `Java class file`}, {`D7CDC69A`, `*.WMF`, `WMF file`},
}

type FileSig struct {
	Sign       string
	SuffixFile string
	FileFormat string
}
