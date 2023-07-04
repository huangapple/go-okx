package main

import (
	"log"

	"github.com/huangapple/go-okx/examples/rest"
	"github.com/huangapple/go-okx/rest/api/account"
)

func main() {
	param := &account.GetBillsParam{}
	req, resp := account.NewGetBillsArchive(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetBillsResponse))
}
