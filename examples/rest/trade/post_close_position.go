package main

import (
	"log"

	"github.com/huangapple/go-okx/examples/rest"
	"github.com/huangapple/go-okx/rest/api"
	"github.com/huangapple/go-okx/rest/api/trade"
)

func main() {
	param := &trade.PostClosePositionParam{
		InstId:  "OKB-USDT",
		MgnMode: api.MgnModeCross,
		Ccy:     "OKB",
	}
	req, resp := trade.NewPostClosePosition(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostClosePositionResponse))
}
