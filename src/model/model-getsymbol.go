package model

/*

	"id": "ETH",
	"fullName": "Ethereum",
	"ask": "0.054464",
	"bid": "0.054463",
	"last": "0.054463",
	"open": "0.057133",
	"low": "0.053615",
	"high": "0.057559",
	"feeCurrency": "BTC"
*/

/*
{
  "jsonrpc":"2.0",
  "method":"ticker",
  "params":
     {
      "ask":"0.064165",
      "bid":"0.064144",
      "last":"0.064162",
      "open":"0.063657",
      "low":"0.062452",
      "high":"0.064583",
      "volume":"18764.7030",
      "volumeQuote":"1193.1946416800",
      "timestamp":"2021-06-15T04:08:29.377Z",
      "symbol":"ETHBTC"
     }
}
*/

type GetSymbolResponseBody struct {
	Id 			string  `json:"id"`
	FullName 	string  `json:"fullName"`
	Ask			float64 `json:"Ask"`
	Bid 		float64 `json:"Bid"`
	Last 		float64 `json:"Last"`
	Open 		float64 `json:"Open"`
	Low 		float64 `json:"Low"`
	High 		float64 `json:"High"`
	FeeCurrency string  `json:"feeCurrency"`
}

type GetSymbolResponse struct {
	Body GetSymbolResponseBody
	Err  error
}

type HitBTCResponse struct {
	Jsonrpc		string    `json:"jsonrpc"`
	Method      string    `json:"method"`
	Params struct {
		Ask  string   `json:"ask"`
		Bid  string   `json:"bid"`
		Last string   `json:"last"`
		Open string   `json:"open"`
		Low  string   `json:"low"`
		High string   `json:"high"`
		Volume string `json:"volume"`
		VolumeQuote  string  `json:"volumeQuote"`
		Timestamp   string  `json:"timestamp"`
		Symbol   string  `json:"symbol"`

	} `json:"params"`
}