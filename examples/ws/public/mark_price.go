package main

import (
	"log"

	"github.com/huangapple/go-okx/ws/public"
)

func main() {
	handler := func(c public.EventMarkPrice) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeMarkPrice("BTC-USDT", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
