package F_Options

func Return_Packet_To_parse(opt string) string {
	switch opt {
	case "wep": // WEP
		return "wep"
	case "aruba": // Aruba WLAN
		return "Aruba"
	case "actionnk": // Action no ack
		return "ActionACK"
	case "action": // Action
		return "Action"
	case "deauth": // DOT11 MGMT Deauthentication
		return "Deauthentication"
	case "auth": // DOT11 MGMT Authentication
		return "Authentication"
	case "disass": // DOT11 Disassociation
		return "Disassociation"
	case "atim": // DOT11 ATIM
		return "ATIM"
	case "beacon": // DOT11 MGMT Beacon
		return "Beacon"
	case "discover":
		return "table"
	case "measurement": // DOT11 MGMT Measurement Plot
		return "Measurement_Plot"
	case "proberesp": // Dot11 MGMT Probe Response
		return "Probe_response"
	case "probereq": // DOT11 MGMT Probe Request
		return "Probe_request"
	case "reassresp": // Dot11 MGMT Reassociation Response
		return "ReassociationResp"
	case "reassreq": // Dot11 MGMT Reassociation response
		return "ReassociationRequest"
	case "assoresp": // Dot11 MGMT Association Response
		return "AssociationResp"
	case "assoreq":
		return "AssociationReq"
	case "tap":
		return "RadioTap"
	case "eopal":
		return "eopal"
	case "test":
		return "IPV6"
	case "dot11":
		return "dot11"
	case "normal":
		return "*"
	case "airodump":
		return "air"
	}
	return ""
}
