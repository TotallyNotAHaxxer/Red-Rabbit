package KNOXSS_RESPONSE

type Body struct {
	Cross_Site_Scripting string `json:"XSS"`
	Proof_Of_Concept     string `json:"PoC"`
	Target               string `json:"Target"`
	Posted_Data          string `json:"POST Data"`
	Knoxss_Error         string `json:"Error"`
	Call_api             string `json:"API Call"`
	Time_Passed          string `json:"Time Elapsed"`
	Time_Finish_Stamp    string `json:"Timestamp"`
}
