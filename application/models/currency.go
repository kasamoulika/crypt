package models

// Currency represents the response structure of /v1/currency/:symbol
type Currency struct {
	Id           string  `json:"id"`
	AskPrice     float64 `json:"ask"`
	BidPrice     float64 `json:"bid"`
	HighestPrice float64 `json:"high"`
	LowestPrice  float64 `json:"low"`
	OpenPrice    float64 `json:"open"`
	LastPrice    float64 `json:"last"`
}
