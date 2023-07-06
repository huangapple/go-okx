package public

import (
	"github.com/gorilla/websocket"
	"github.com/huangapple/go-okx/ws"
)

type Public struct {
	C *ws.Client
}

func NewPublic(simulated bool) *Public {
	public := &Public{
		C: ws.DefaultClientPublic,
	}
	if simulated {
		public.C = ws.DefaultClientPrivateSimulated
	}
	return public
}

// subscribe
func (p *Public) Subscribe(args []*ws.Args, handler ws.Handler, handlerError ws.HandlerError) (*websocket.Conn, error) {

	tempArr := make([]interface{}, len(args))
	for i := range args {
		tempArr[i] = args[i]
	}

	subscribe := ws.NewOperateSubscribe(tempArr, handler, handlerError)
	return p.C.Operate(subscribe, nil)
}
