package stock

import (
	"encoding/json"
	"time"
)

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
