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

type AStockBase interface {
	// DM 代码
	DM() string
	// MC 名称
	MC() string
	// JYS 交易所
	JYS() string
	// UT 更新时间
	UT() time.Time
}
