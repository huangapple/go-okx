package main

import (
	"log"

	"github.com/huangapple/go-okx/examples/rest"
	"github.com/huangapple/go-okx/rest/api"
	"github.com/huangapple/go-okx/rest/api/account"
)

func main() {
	param := &account.GetTradeFeeParam{
		InstId:   "BTC-USDT",
		InstType: api.InstTypeSPOT,
	}
	req, resp := account.NewGetTradeFee(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetTradeFeeResponse))
}
