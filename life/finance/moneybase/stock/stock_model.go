package stock

import "time"

// GuBen 股本类型接口
type GuBen interface {
	// DaiMa 股票代码
	DaiMa() string
	// ZongGuBen 总股本
	ZongGuBen() int64
	// LiuTongGuBen 流通股本
	LiuTongGuBen() int64
	// RenJunChiGu 人均持股
	RenJunChiGu() float64
	// BianDongRiQi 变动日期
	BianDongRiQi() time.Time
	// GuDongRenShu 股东人数
	GuDongRenShu() int64
}

// AStockBase A股股票基本数据结构，以mairui接口为基准
type AStockBase struct {
	DM  string    `json:"dm"`          // 代码
	MC  string    `json:"mc"`          // 名称
	JYS string    `json:"jys"`         // 交易所
	UT  time.Time `json:"update_time"` // 更新时间
}

/**
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
                "updatetime": "2023-04-03",
                "jjmc": "华夏能源革新股票A"
            },
*/

// FundChiGu 金融机构持股，以腾讯的字段为基准
type FundChiGu struct {
	// 基金代码
	JJDM string
	// 股票证券代码
	ZQDM string
	// 股票证券名称
	ZQMC string
	// 交易市场
	ZQSCLX string
	// 持仓市值
	CCSZ string
	// 净值占比（该股在基金中的占比）
	JZBL string
	// 本次变更日期
	BCBGRQ string
	// 本次持仓股数
	BCCCGS string
	// 上次变更日期
	SCBGRQ string
	// 上次持股股数
	SCCCGS string
	// 持仓变动
	CCBD string
	// 更新日期
	UPDATETIME string
	// 基金名称
	JJMC string
}

/**
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
            },
*/

// GuDong 股东，该接口以腾讯的字段为基准
type GuDong struct {
	// 股票代码
	Symbol string
	// 股东id
	Id string
	// 公司码
	CompanyCode string
	// 排名
	Rank int
	// 股东名称
	ShareholderName string
	// 股东类型
	ShareholderType string
	// 股东持股数
	ShareholderNum int64
	// 股东持股比率
	ShareholderPercent float64
	// 上报日期
	ReportDate time.Time
	// 持股变化
	ShareholderChange float64
}
