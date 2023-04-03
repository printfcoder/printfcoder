package stock

import "time"

type GuBenTencent struct {
	// 股票代码
	GPDM string `json:"gpdm,omitempty"`
	// 变动时间
	BDRQ time.Time `json:"bdrq"`
	// 总股本
	ZGB int64 `json:"zgb,omitempty"`
	// 流通股本
	LTGB int64 `json:"ltgb,omitempty"`
	// 股东人数
	GDRS int64 `json:"gdrs,omitempty"`
	// 人均持股
	RJCG float64 `json:"rjcg,omitempty"`
}

func (g *GuBenTencent) DaiMa() string {
	return g.GPDM
}

func (g *GuBenTencent) ZongGuBen() int64 {
	return g.ZGB * 10000
}

func (g *GuBenTencent) LiuTongGuBen() int64 {
	return g.LTGB * 10000
}

func (g *GuBenTencent) RenJunChiGu() float64 {
	return g.RJCG
}

func (g *GuBenTencent) BianDongRiQi() time.Time {
	return g.BDRQ
}

func (g *GuBenTencent) GuDongRenShu() int64 {
	return g.GDRS
}
