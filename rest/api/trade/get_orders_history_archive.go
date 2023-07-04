package trade

import (
	"github.com/huangapple/go-okx/rest/api"
)

func NewGetOrdersHistoryArchive(param *GetOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-history-archive",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrderResponse{}
}
