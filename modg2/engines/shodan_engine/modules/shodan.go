package Engine_Shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Host(query, key string) (*T, uint32) {
	response, X := http.Get(fmt.Sprintf("https://api.shodan.io/shodan/host/search?key=%s&query=%s", key, query))
	if X != nil {
		fmt.Println("<RR6> Engine - SHODAN -> Got error when trying to make a GET request to the new formatted shodan URL using the query and key -> ", X)
		return nil, 0x01
	}
	defer response.Body.Close()
	if X := json.NewDecoder(response.Body).Decode(&Results); X != nil {
		fmt.Println("<RR6> Engine - SHODAN - JSON - Requests - DECODER: Got error when trying to decode the response body from the HOST -> ", X)
		return nil, 0x01
	}
	return &Results, 0x00
}
