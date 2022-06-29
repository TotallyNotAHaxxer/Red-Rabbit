package A80211_constants

import (
	"time"

	l "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	DOT11_plain                                  = l.LayerTypeDot11
	DOT11_Control                                = l.LayerTypeDot11Ctrl
	DOT11_Data                                   = l.LayerTypeDot11Data
	DOT11_Data_CFAK                              = l.LayerTypeDot11DataCFAck
	DOT11_Data_CFPOLL                            = l.LayerTypeDot11DataCFPoll
	DOT11_Data_CFACKPOLL                         = l.LayerTypeDot11DataCFAckPoll
	DOT11_Data_NULL                              = l.LayerTypeDot11DataNull
	DOT11_MGMT_Association_request               = l.LayerTypeDot11MgmtAssociationReq
	DOT11_MGMT_Association_response              = l.LayerTypeDot11MgmtAssociationResp
	DOT11_MGMT_Reassociation_request             = l.LayerTypeDot11MgmtReassociationReq
	DOT11_MGMT_Reassociation_response            = l.LayerTypeDot11MgmtReassociationResp
	DOT11_MGMT_Probe_Request                     = l.LayerTypeDot11MgmtProbeReq
	DOT11_MGMT_Probe_Response                    = l.LayerTypeDot11MgmtProbeResp
	DOT11_MGMT_Measurement_Pilot_main            = l.LayerTypeDot11MgmtMeasurementPilot
	DOT11_MGMT_BEACON                            = l.LayerTypeDot11MgmtBeacon
	DOT11_MGMT_ATIM                              = l.LayerTypeDot11MgmtATIM
	DOT11_MGMT_Disassociation                    = l.LayerTypeDot11MgmtDisassociation
	DOT11_MGMT_Authentication                    = l.LayerTypeDot11MgmtAuthentication
	DOT11_MGMT_Deauthentication                  = l.LayerTypeDot11MgmtDeauthentication
	DOT11_MGMT_ACTION                            = l.LayerTypeDot11MgmtAction
	DOT11_MGMT_Action_No_Ack                     = l.LayerTypeDot11MgmtActionNoAck
	DOT11_MGMT_Aruba_Wireless_Local_Area_Network = l.LayerTypeDot11MgmtArubaWLAN
	DOT11_MGMT_WEP                               = l.LayerTypeDot11WEP
	DOT11_RadioTap                               = l.LayerTypeRadioTap
	DOT11_EAPOL                                  = l.LayerTypeEAPOL
	TCP_                                         = l.LayerTypeTCP
	MDNS_                                        = l.LayerTypeDNS
	ICMPV6_                                      = l.LayerTypeICMPv6
)

var (
	Interface      string
	Snaplen        int32
	X              error
	Monitor        bool
	Timeout        time.Duration
	Controller     *pcap.Handle
	PB             []byte
	PB2            []byte
	PROBE_DATA     []string
	CIPHER         = ""
	AUTHENTICATION = ""
	F              = false
	ENCRYPTION     = ""
)

var WPAS = []byte{
	0, 0x50,
	0xf2, 1,
}
