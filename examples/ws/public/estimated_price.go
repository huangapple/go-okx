package main

import (
	"log"

	"github.com/huangapple/go-okx/ws"
	"github.com/huangapple/go-okx/ws/public"
)

func main() {
	args := &ws.Args{
		InstType:   "FUTURES",
		InstFamily: "BTC-USDT",
	}
	handler := func(c public.EventEstimatedPrice) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeEstimatedPrice(args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
