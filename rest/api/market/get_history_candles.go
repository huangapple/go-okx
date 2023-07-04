package market

import "github.com/huangapple/go-okx/rest/api"

func NewGetHistoryCandles(param *GetCandlesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/history-candles",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCandlesResponse{}
}
