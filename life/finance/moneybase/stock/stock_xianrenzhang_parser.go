package stock

import "time"

/**
【中信建投】《仙人涨(天马行空)》北方国际(买入操作）：调入北方国际(000065)，建仓区间15.51元-15.95元，止损价13.45元，仓位4.81%调入理由：主营国际工程承包服务，前天大涨时调出，重新小仓位试探性配置。风险是调入后股价走势不及预期。（本建议仅供参考，请关注风险，审慎决策，分散投资，注意证券走势）
*/

type XianRenZhangAdvise struct {
	DM              string
	MC              string
	Name            string
	OperateRecordId int64
	JingLi          string
	RiQi            time.Time
	JieGeMin        float64
	JieGeMax        float64
	// 止损/止盈价格
	ZhiJiaGe           float64
	TuiJianCangWei     float64
	WoDeChengBen       float64
	WoDeChiGuShu       int
	WoDeCangWei        int
	WoDeChiGuShiZhi    int
	OperateRiQi        time.Time
	ChiCangZhangFu     float64
	WoDeChiCangZhangFu float64
	OperateJiaGe       float64
	// 操作理由
	LiYou string
}

type TianMaXingKong struct {
}
