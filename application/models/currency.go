package models

// Currency represents the response structure of /v1/currency/:symbol
type Currency struct {
	Id           string `json:"id"`
	AskPrice     string `json:"ask"`
	BidPrice     string `json:"bid"`
	HighestPrice string `json:"high"`
	LowestPrice  string `json:"low"`
	OpenPrice    string `json:"open"`
	LastPrice    string `json:"last"`
}
