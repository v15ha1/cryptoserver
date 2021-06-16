package hitbtc

import (
	"bytes"
	"net/url"
	"encoding/json"
	//"io/ioutil"
	data "cryptoserver-clean-app/data"
	"cryptoserver-clean-app/model"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/websocket"
)

type HitBTCClient struct {
	symbol  string
	baseUrl string
	
}

//{ "method":"subscribeTicker", "params": {"symbol": "ETHBTC"}, "id": 123 }
//{ "method":"subscribeTicker", "params": {"symbol": "BTCUSD"}, "id": 123 }
type HitBTCRequest struct {
	Method 			string  `json:"method"`
	Params  struct {
		Symbol string `json:"symbol"`
	} `json:"params"`
	Id int  `json:"id"`
}

func NewHitBTCClient(config *data.Config, symbol string) *HitBTCClient {

	u := url.URL{Scheme: "wss", Host: config.HitBTC.Endpoint, Path: "/api/2/ws"}

	return &HitBTCClient{
		symbol: symbol,
		baseUrl: u.String(),
	}

}

func (svc *HitBTCClient) Start(log *logrus.Entry) error {

	log.Printf("connecting to %s", svc.baseUrl)

	c, _, err := websocket.DefaultDialer.Dial(svc.baseUrl, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	var hitBTCReq HitBTCRequest
	hitBTCReq.Method = "subscribeTicker"
	hitBTCReq.Params.Symbol = svc.symbol 
	hitBTCReq.Id = 123

	jsonVal, _ := json.Marshal(hitBTCReq)

	log.Infoln("sending to Host : ", svc.baseUrl, ", Request : ", bytes.NewBuffer(jsonVal))
	c.WriteMessage(websocket.TextMessage, []byte(jsonVal)) 
	
	for {
		mtype, message, err := c.ReadMessage()
		if err != nil {
			log.Println("WS read error:", err)
			return err
		}

		if mtype != websocket.TextMessage {
			continue
		}

		//log.Printf("recv: %s", message)

		var hitBTCResp model.HitBTCResponse

		err = json.Unmarshal(message, &hitBTCResp)
		if err != nil {
				log.Errorln("JSON Unmarshal failed with error ", err)
				return err
		}

		log.Printf("Caching Data : %v ", hitBTCResp)

		MyCache.Set(svc.symbol, hitBTCResp)
	}

	return nil
}