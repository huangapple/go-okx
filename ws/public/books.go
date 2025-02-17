package public

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/huangapple/go-okx/ws"
)

// 深度频道

type HandlerBooks func(interface{})

type EventBooks struct {
	Arg    ws.Args `json:"arg"`
	Data   []Book  `json:"data"`
	Action string  `json:"action"`
}

type Book struct {
	Asks     [][]string `json:"asks"`
	Bids     [][]string `json:"bids"`
	Ts       int64      `json:"ts,string"`
	Checksum int32      `json:"checksum"`
}

// default subscribe
func SubscribeBooks(args []*ws.Args, handler HandlerFunc, handlerError ws.HandlerError, simulated bool) (*websocket.Conn, error) {
	h := func(message []byte) { // convert raw data into EventBooks
		var event EventBooks
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
