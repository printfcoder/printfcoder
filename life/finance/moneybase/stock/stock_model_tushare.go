package stock

import (
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
		/**
		"ts_code",
		"trade_date",
		"open",
		"high",
		"low",
		"close",
		"pre_close",
		"change",
		"pct_chg",
		"vol",
		"amount"

		"000001.SZ",
		"20230504",
		12.4,
		12.83,
		12.19,
		12.73,
		12.55,
		0.18,
		1.4343,
		1478264.38, 1478264
		1863146.144 186315
		*/
		for _, item := range r.Data.Items {
			d := StockQTDataTencent{}

			code, ok := item[0].(string)
			if !ok {
				log.Errorf("[ToStdQTData] code不正常：%v", item[0])
			}
			d.DaiMa = strings.Split(code, ".")[0]

			// 20230504 -> 20230504161406
			date, ok := item[1].(string)
			if !ok {
				log.Errorf("[ToStdQTData] date不正常：%v", item[1])
			}
			d.ShiJian, _ = strconv.ParseInt(date+"161406", 10, 64)

			// 今开
			open := item[2].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] open不正常：%v", item[2])
			}
			d.JinKai = open

			high := item[3].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] high不正常：%v", item[3])
			}
			d.Max = high

			low := item[4].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] low不正常：%v", item[4])
			}
			d.Min = low

			cls := item[5].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] close不正常：%v", item[5])
			}
			d.DangQianJiaGe = cls

			preClose := item[6].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] pre_close不正常：%v", item[6])
			}
			d.ZuoShou = preClose

			change := item[7].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] change不正常：%v", item[7])
			}
			d.ZhangDie = change

			pctChg := item[8].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] pctChg不正常：%v", item[8])
			}
			d.ZhangDiePercent = pctChg

			vol := item[9].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] vol不正常：%v", item[9])
			}
			d.ChengJiaoLiangShou = vol

			amount := item[10].(float64)
			if !ok {
				log.Errorf("[ToStdQTData] amount不正常：%v", item[10])
			}
			d.ChengJiaoEWan = amount / 10

			ret = append(ret, d)
		}
	}

	return
}
