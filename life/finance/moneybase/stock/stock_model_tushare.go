package stock

import (
	"fmt"
	log "github.com/stack-labs/stack/logger"
	"strconv"
	"strings"
)

type tushareReqBody struct {
	APIName string      `json:"api_name,omitempty"`
	Token   string      `json:"token,omitempty"`
	Params  interface{} `json:"params,omitempty"`
}

type tushareDailyParams struct {
	TsCode    string `json:"ts_code,omitempty"`
	TradeDate string `json:"trade_date,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type tushareTop10GuDongParams struct {
	TsCode    string `json:"ts_code,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

/**
{
    "request_id": "09f14e3ef24111edb2194fdb4a7b8cb3",
    "code": 0,
    "msg": "",
    "data": {
        "fields": [
            "ts_code",
            "ann_date",
            "end_date",
            "holder_name",
            "hold_amount"
        ],
        "items": [
            [
                "000786.SZ",
                "20230429",
                "20230331",
                "阿布达比投资局",
                11627831.0
            ],
*/

type tushareGuDong struct {
	TsCode     string  `json:"ts_code"`
	AnnDate    string  `json:"ann_date"`
	EndDate    string  `json:"end_date"`
	HolderName string  `json:"holder_name"`
	HoldAmount float64 `json:"hold_amount"`
}

func (t *tushareGuDong) ToTop10GuDong() StockTop10GuDong {
	st10 := StockTop10GuDong{
		DM:         t.TsCode,
		AnnDate:    t.AnnDate,
		EndDate:    t.EndDate,
		HolderName: t.HolderName,
		HoldAmount: t.HoldAmount,
	}

	return st10
}

type tushareRsp struct {
	RequestId string `json:"request_id,omitempty"`
	Code      int    `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
	Data      *struct {
		Fields []string        `json:"fields" json:"fields,omitempty"`
		Items  [][]interface{} `json:"items" json:"items,omitempty"`
	} `json:"data"`
	Fields  []string        `json:"fields,omitempty"`
	Items   [][]interface{} `json:"items,omitempty"`
	HasMore bool            `json:"has_more,omitempty"`
}

func (r *tushareRsp) ToStdQTData() (ret []StockQTDataTencent) {
	if r.Data != nil {
		for _, v := range r.Data.Items {
			d := StockQTDataTencent{}

			code, ok := v[0].(string)
			if !ok {
				log.Errorf("[ToStdQTData] code不正常：%v", v[0])
			}
			d.DaiMa = strings.Split(code, ".")[0]

			// 20230504 -> 20230504161406
			date, ok := v[1].(string)
			if !ok {
				log.Errorf("[ToStdQTData] date不正常：%v", v[1])
			}
			d.ShiJian, _ = strconv.ParseInt(date+"161406", 10, 64)

			// 今开
			open, ok := v[2].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] open不正常：%v", v[2])
			}
			d.JinKai = open

			high, ok := v[3].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] high不正常：%v", v[3])
			}
			d.Max = high

			low, ok := v[4].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] low不正常：%v", v[4])
			}
			d.Min = low

			cls, ok := v[5].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] close不正常：%v", v[5])
			}
			d.DangQianJiaGe = cls

			preClose, ok := v[6].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] pre_close不正常：%v", v[6])
			}
			d.ZuoShou = preClose

			change, ok := v[7].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] change不正常：%v", v[7])
			}
			d.ZhangDie = change

			pctChg, ok := v[8].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] pctChg不正常：%v", v[8])
			}
			d.ZhangDiePercent = pctChg

			vol, ok := v[9].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] vol不正常：%v", v[9])
			}
			d.ChengJiaoLiangShou = vol

			amount, ok := v[10].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] amount不正常：%v", v[10])
			}
			d.ChengJiaoEWan = amount / 10

			ret = append(ret, d)
		}
	}
	return
}

func (r *tushareRsp) ToTop10GuDong(guDongType int) (ret []StockTop10GuDong) {
	if r.Data != nil {
		for _, v := range r.Data.Items {
			s := StockTop10GuDong{
				DM:         fmt.Sprintf("%v", v[0]),
				AnnDate:    fmt.Sprintf("%v", v[1]),
				EndDate:    fmt.Sprintf("%v", v[2]),
				HolderName: fmt.Sprintf("%v", v[3]),
				GuDongType: guDongType,
			}

			s.DM = strings.Split(s.DM, ".")[0]

			holdAmount, ok := v[4].(float64)
			if !ok {
				log.Errorf("[ToTop10GuDong] holdAmount不正常：%v", v[4])
			}

			s.HoldAmount = holdAmount

			holdRatio, ok := v[5].(float64)
			if !ok {
				log.Errorf("[ToTop10GuDong] holdRatio不正常：%v", v[4])
			}

			s.HoldRatio = holdRatio

			ret = append(ret, s)
		}
	}
	return
}
