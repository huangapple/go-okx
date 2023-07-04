package main

import (
	"log"

	"github.com/huangapple/go-okx/examples/rest"
	"github.com/huangapple/go-okx/rest/api/trade"
)

func main() {
	param := &trade.GetOrderParam{
		InstId: "OKB-USDT",
		OrdId:  "501163171776954368",
	}
	req, resp := trade.NewGetOrder(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetOrderResponse))
}
