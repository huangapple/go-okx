package main

import (
	"log"

	"github.com/huangapple/go-okx/ws"
	"github.com/huangapple/go-okx/ws/public"
)

func main() {
	args := &ws.Args{
		Channel: "books5",
		InstId:  "BTC-USDT",
	}
	handler := func(c interface{}) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if _, err := public.SubscribeBooks(args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
