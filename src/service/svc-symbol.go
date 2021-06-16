package service

import (
	"context"
	"encoding/json"
	"strconv"
	"cryptoserver-clean-app/model"
	"cryptoserver-clean-app/util"
	hb "cryptoserver-clean-app/hitbtc"
)
/*
func WorkerThread(log *logrus.Entry) {

	fmt.Println("Starting worker thread ...")

	for {
		log.Infoln("Processing Symbol  # ")
		time.Sleep(10 * time.Second)
		timeStart := time.Now()
		log.Printf("Refreshed in %2.2f secs", time.Now().Sub(timeStart).Seconds())
 
	}
}
*/
func (svc *CryptoServerSvc) GetSymbol(ctx context.Context, symbol string) (model.GetSymbolResponseBody, error) {

	var response model.GetSymbolResponseBody
	log := util.Logger(ctx)

	log.Infoln("Invoked API")

	b, err := hb.MyCache.Get(symbol)

	// Cache miss 
	if err != nil {
		log.Errorln("Cache miss ", err)
		return response, err
	}

	var hitBTCResp model.HitBTCResponse

	err = json.Unmarshal(b, &hitBTCResp)
	if err != nil {
			log.Errorln("JSON Unmarshal failed with error ", err)
			return response, err
	}

	log.Printf("Cache Data : %v", hitBTCResp)

	response.FullName = symbol
	response.Ask, _ = strconv.ParseFloat(hitBTCResp.Params.Ask, 64)
	response.Bid, _ = strconv.ParseFloat(hitBTCResp.Params.Bid, 64)
	response.Last, _ = strconv.ParseFloat(hitBTCResp.Params.Last, 64)
	response.Open, _ = strconv.ParseFloat(hitBTCResp.Params.Open, 64)
	response.Low, _ = strconv.ParseFloat(hitBTCResp.Params.Low, 64)
	response.High, _ = strconv.ParseFloat(hitBTCResp.Params.High, 64)
	
	return response, nil

}

