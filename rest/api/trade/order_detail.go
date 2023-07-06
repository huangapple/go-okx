package trade

import "github.com/huangapple/go-okx/common"

type OrderDetail struct {
	InstType        string              `json:"instType"`
	InstId          string              `json:"instId"`
	TgtCcy          string              `json:"tgtCcy"`
	Ccy             string              `json:"ccy"`
	OrdId           string              `json:"ordId"`
	ClOrdId         string              `json:"clOrdId"`
	Tag             string              `json:"tag"`
	Px              string              `json:"px"`
	Sz              string              `json:"sz"`
	Pnl             common.Float64Value `json:"pnl"`
	OrdType         string              `json:"ordType"`
	Side            string              `json:"side"`
	PosSide         string              `json:"posSide"`
	TdMode          string              `json:"tdMode"`
	AccFillSz       common.Float64Value `json:"accFillSz"`
	FillPx          string              `json:"fillPx"`
	TradeId         string              `json:"tradeId"`
	FillSz          string              `json:"fillSz"`
	FillTime        string              `json:"fillTime"`
	AvgPx           common.Float64Value `json:"avgPx"`
	State           string              `json:"state"`
	Lever           string              `json:"lever"`
	TpTriggerPx     string              `json:"tpTriggerPx"`
	TpTriggerPxType string              `json:"tpTriggerPxType"`
	SlTriggerPx     string              `json:"slTriggerPx"`
	SlTriggerPxType string              `json:"slTriggerPxType"`
	SlOrdPx         string              `json:"slOrdPx"`
	FeeCcy          string              `json:"feeCcy"`
	Fee             common.Float64Value `json:"fee"`
	RebateCcy       string              `json:"rebateCcy"`
	Source          string              `json:"source"`
	Rebate          string              `json:"rebate"`
	Category        string              `json:"category"`
	UTime           int64               `json:"uTime,string"`
	CTime           int64               `json:"cTime,string"`
}

func (od *OrderDetail) IsFilled() bool {

	return od.State == "filled"
}

func (od *OrderDetail) ActionDesc() string {

	if od.PosSide == "long" {
		if od.Side == "buy" {
			return "开多"
		} else if od.Side == "sell" {
			return "平多"
		}

	} else if od.PosSide == "short" {
		if od.Side == "buy" {
			return "平空"
		} else if od.Side == "sell" {
			return "开空"
		}
	}
	return "未知操作"
}
