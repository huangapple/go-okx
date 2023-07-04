package main

import (
	"log"

	"github.com/huangapple/go-okx/examples/rest"
	"github.com/huangapple/go-okx/rest/api/subaccount"
)

func main() {
	param := &subaccount.GetBillsParam{}
	req, resp := subaccount.NewGetBills(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*subaccount.GetBillsResponse))
}
