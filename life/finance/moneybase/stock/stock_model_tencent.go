package stock

import (
	"encoding/json"
	"time"
)

// region guben
/**
{
    "code": 0,
    "msg": "",
    "data": {
        "fundchigu": [
            {
                "jjdm": "003834",
                "zqdm": "601633",
                "zqmc": "长城汽车",
                "zqsclx": "sh",
                "ccsz": "957291238.46",
                "jzbl": "6.1000",
                "bcbgrq": "2022-12-31",
                "bcccgs": "32319083",
                "scbgrq": "2022-09-30",
                "scccgs": "32319083",
                "ccbd": "0",
                "updatetime": "2023-04-04",
                "jjmc": "华夏能源革新股票A"
            }
        ],
        "guben": [
            {
                "GPDM": "sh601633",
                "BDRQ": "2023-03-30",
                "GGRQ": "2023-04-01",
                "ZGB": "848656",
                "LTGB": "613250",
                "GDRS": "205372",
                "SQ_GDRS": "233141",
                "RJCG": "41323"
            }
        ],
        "gudong": [
            {
                "symbol": "sh601633",
                "id": "2315115",
                "company_code": "10002291",
                "rank": "1",
                "shareholder_name": "保定创新长城资产管理有限公司",
                "shareholder_type": "其它",
                "shareholder_num": "5115000000.00",
                "shareholder_percent": "58.3603",
                "report_date": "2022-12-31",
                "shareholder_change": "0"
            }
        ]
    }
}
*/

type GuBenTencentRsp struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data GuBenInfoTencent `json:"data"`
}

type GuDongTencent struct {
	Symbol              string      `json:"symbol"`
	Id                  string      `json:"id"`
	Company_code        string      `json:"company_code"`
	Rank                json.Number `json:"rank"`
	Shareholder_name    string      `json:"shareholder_name"`
	Shareholder_type    string      `json:"shareholder_type"`
	Shareholder_num     json.Number `json:"shareholder_num"`
	Shareholder_percent json.Number `json:"shareholder_percent"`
	Report_date         string      `json:"report_date"`
	Shareholder_change  json.Number `json:"shareholder_change"`
}

type FundChiGuTencent struct {
	Jjdm       string      `json:"jjdm"`
	Zqdm       string      `json:"zqdm"`
	Zqmc       string      `json:"zqmc"`
	Zqsclx     string      `json:"zqsclx"`
	Ccsz       json.Number `json:"ccsz"`
	Jzbl       json.Number `json:"jzbl"`
	Bcbgrq     string      `json:"bcbgrq"`
	Bcccgs     json.Number `json:"bcccgs"`
	Scbgrq     string      `json:"scbgrq"`
	Scccgs     json.Number `json:"scccgs"`
	Ccbd       json.Number `json:"ccbd"`
	Updatetime string      `json:"updatetime"`
	Jjmc       string      `json:"jjmc"`
}

type GuBenTencent struct {
	GPDM    string      `json:"GPDM"`
	BDRQ    string      `json:"BDRQ"`
	GGRQ    string      `json:"GGRQ"`
	ZGB     json.Number `json:"ZGB"`
	LTGB    json.Number `json:"LTGB"`
	GDRS    json.Number `json:"GDRS"`
	SQ_GDRS json.Number `json:"SQ_GDRS"`
	RJCG    json.Number `json:"RJCG"`
}

type GuBenInfoTencent struct {
	GuBen     []GuBenTencent     `json:"guben"`
	FundChiGu []FundChiGuTencent `json:"fundchigu"`
	GuDong    []GuDongTencent    `json:"gudong"`
}

func (g *GuBenInfoTencent) ToGuBenInfo() GuBenInfo {
	gg := GuBenInfo{}

	for _, v := range g.GuBen {
		gb := GuBen{}

		gb.GPDM = v.GPDM
		gb.BDRQ, _ = time.Parse("2006-01-02", v.BDRQ)
		gb.GGRQ, _ = time.Parse("2006-01-02", v.GGRQ)
		gb.ZGB, _ = v.ZGB.Int64()
		gb.LTGB, _ = v.LTGB.Int64()
		gb.GDRS, _ = v.GDRS.Int64()
		gb.SQ_GDRS, _ = v.SQ_GDRS.Int64()
		gb.RJCG, _ = v.RJCG.Float64()

		gg.GuBen = append(gg.GuBen, gb)
	}

	for _, v := range g.FundChiGu {
		fcg := FundChiGu{}
		fcg.JJDM = v.Jjdm
		fcg.ZQDM = v.Zqdm
		fcg.ZQMC = v.Zqmc
		fcg.ZQSCLX = v.Zqsclx
		fcg.CCSZ, _ = v.Ccsz.Float64()
		fcg.JZBL, _ = v.Jzbl.Float64()
		fcg.BCBGRQ, _ = time.Parse("2006-01-02", v.Bcbgrq)
		fcg.BCCCGS, _ = v.Bcccgs.Int64()
		fcg.SCBGRQ, _ = time.Parse("2006-01-02", v.Scbgrq)
		fcg.SCCCGS, _ = v.Scccgs.Int64()
		fcg.CCBD, _ = v.Ccbd.Int64()
		fcg.UPDATETIME, _ = time.Parse("2006-01-02", v.Updatetime)
		fcg.JJMC = v.Jjmc

		gg.FundChiGu = append(gg.FundChiGu, fcg)
	}

	for _, v := range g.GuDong {
		gc := GuDong{}
		gc.Symbol = v.Symbol
		gc.Id = v.Id
		gc.CompanyCode = v.Company_code
		gc.Rank, _ = v.Rank.Int64()
		gc.ShareholderName = v.Shareholder_name
		gc.ShareholderType = v.Shareholder_type
		gc.ShareholderNum, _ = v.Shareholder_num.Int64()
		gc.ShareholderPercent, _ = v.Shareholder_percent.Float64()
		gc.ReportDate, _ = time.Parse("2006-01-02", v.Report_date)
		gc.ShareholderChange, _ = v.Shareholder_change.Int64()

		gg.GuDong = append(gg.GuDong, gc)
	}

	return gg
}

// endregion

// region 股票行情

/**
 "sz000610": [
                    "51",
                    "西安旅游",
                    "000610",
                    "16.78",
                    "17.03",
                    "16.91",
                    "204946",
                    "94025",
                    "110920",
                    "16.78",
                    "1894",
                    "16.77",
                    "250",
                    "16.76",
                    "112",
                    "16.75",
                    "210",
                    "16.74",
                    "35",
                    "16.79",
                    "1461",
                    "16.80",
                    "951",
                    "16.81",
                    "90",
                    "16.82",
                    "362",
                    "16.83",
                    "146",
                    "",
                    "20230406150748",
                    "-0.25",
                    "-1.47",
                    "17.09",
                    "16.40",
                    "16.78/204946/344370860",
                    "204946",
                    "34437",
                    "8.71",
                    "-33.05",
                    "",
                    "17.09",
                    "16.40",
                    "4.05",
                    "39.50",
                    "39.73",
                    "5.39",
                    "18.73",
                    "15.33",
                    "0.62",
                    "-509",
                    "16.80",
                    "-36.84",
                    "-54.37",
                    "",
                    "",
                    "1.23",
                    "34437.0860",
                    "0.0000",
                    "0",
                    " ",
                    "GP-A",
                    "-18.42",
                    "0.48",
                    "0.00",
                    "-16.30",
                    "-6.44",
                    "21.77",
                    "7.15",
                    "8.33",
                    "0.90",
                    "-5.30",
                    "235421894",
                    "236747901",
                    "-9.24",
                    "83.39",
                    "235421894",
                    "",
                    "",
                    "73.53",
                    "-0.12",
                    "",
                    "CNY"
                ],

0: 未知
      1: 股票名字
      2: 股票代码
      3: 当前价格
      4: 昨收
      5: 今开
      6: 成交量（手）
      7: 外盘
      8: 内盘
      9: 买一
     10: 买一量（手）
     11-18: 买二 买五
     19: 卖一
     20: 卖一量
     21-28: 卖二 卖五
     29: 最近逐笔成交
     30: 时间
     31: 涨跌
     32: 涨跌%
     33: 最高
     34: 最低
     35: 价格/成交量（手）/成交额
     36: 成交量（手）
     37: 成交额（万）
     38: 换手率
     39: 市盈率
     40:
     41: 最高
     42: 最低
     43: 振幅
     44: 流通市值
     45: 总市值
     46: 市净率
     47: 涨停价
     48: 跌停价
————————————————
*/

// StockQTData 当前行情
type StockQTData struct {
	MC                                string `json:"mc"`
	DaiMa                             string `json:"dai_ma"`
	DangQianJiaGe                     string `json:"dang_qian_jia_ge"`
	ZuoShou                           string `json:"zuo_shou"`
	JinKai                            string `json:"jin_kai"`
	ChengJiaoLiangShou                string `json:"cheng_jiao_liang"`
	WaiPan                            string `json:"wai_pan"`
	NaiPan                            string `json:"nai_pan"`
	Mai3Yi                            string `json:"mai_3_yi"`
	Mai3YiShou                        string `json:"mai_3_yi_shou"`
	Mai4Yi                            string `json:"mai_4_yi"`
	Mai4YiShou                        string `json:"mai_4_yi_shou"`
	ZuiJinZhuBiChengJiao              string `json:"zui_jin_zhu_bi_cheng_jiao"`
	ShiJian                           string `json:"shi_jian"`
	ZhangDie                          string `json:"zhang_die"`
	ZhangDiePercent                   string `json:"zhang_die_percent"`
	Max                               string `json:"max"`
	Min                               string `json:"min"`
	JiaGeChengJiaoLiangShouChengJiaoE string `json:"jia_ge_cheng_jiao_liang_shou_cheng_jiao_e"`
	ChengJiaoLiangShou2               string `json:"cheng_jiao_liang_shou_2"`
	ChengJiaoEWan                     string `json:"cheng_jiao_e_wan"`
	HuanShouLv                        string `json:"huan_shou_lv"`
	ShiYingLv                         string `json:"shi_ying_lv"`
	ZuiGao2                           string `json:"zui_gao_2"`
	ZuiDi2                            string `json:"zui_di_2"`
	ZhenFu                            string `json:"zhen_fu"`
	LiuTongShiZhi                     string `json:"liu_tong_shi_zhi"`
	ZongShiZhi                        string `json:"zong_shi_zhi"`
	ShiJingLv                         string `json:"shi_jing_lv"`
	ZhangTingJia                      string `json:"zhang_ting_jia"`
	DieTingJia                        string `json:"die_ting_jia"`
}

// endregion
