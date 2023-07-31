package market

import "github.com/huangapple/go-okx/rest/api"

func NewGetHistoryTrades(param *GetHistoryTradesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/history-trades",
		Method: api.MethodGet,
		Param:  param,
	}, &GetHistoryTradesResponse{}
}

type GetHistoryTradesParam struct {
	InstId string `url:"instId"`
	Limit  int    `url:"limit"` // 分页返回的结果集数量, 最大为100, 默认100
	Type   string `url:"type"`
	After  string `url:"after"`
	Before string `url:"before"`
}

type GetHistoryTradesResponse struct {
	api.Response
	Data []TradesData `json:"data"`
}
