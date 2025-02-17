package private

import (
	"encoding/json"
	"github.com/gorilla/websocket"

	"github.com/huangapple/go-okx/common"
	"github.com/huangapple/go-okx/rest/api/trade"
	"github.com/huangapple/go-okx/ws"
)

type HandlerOrders func(EventOrders)

type EventOrders struct {
	Arg  ws.Args  `json:"arg"`
	Data []*Order `json:"data"`
}

type Order struct {
	trade.OrderDetail
}

// default subscribe
func SubscribeOrders(args *ws.Args, auth *common.Auth, handler HandlerOrders, handlerError ws.HandlerError) (*websocket.Conn, error) {
	args.Channel = "orders"

	h := func(message []byte) {
		var event EventOrders
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPrivate(auth).Subscribe(args, h, handlerError)
}
