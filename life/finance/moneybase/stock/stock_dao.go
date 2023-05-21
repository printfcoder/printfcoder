package stock

import "database/sql"

type Dao interface {
	// SetDB 设置DB
	SetDB(db *sql.DB)

	// ReadAllAStockBases 读取所有A股股本信息到数据库
	ReadAllAStockBases() (list []AStockBase, err error)

	// ReadAllAStockBasesForSyncGuDong 读取所有A股股本信息到数据库
	ReadAllAStockBasesForSyncGuDong(guDongType int) (list []AStockBase, err error)

	// ReadAllAStockBasesForSyncQT 读取所有A股股本信息到数据库
	ReadAllAStockBasesForSyncQT(dateStart, dateEnd int64) (list []AStockBase, err error)

	// WriteAllAStockBases 写入所有A股信息到数据库
	WriteAllAStockBases(aStockBases ...AStockBase) error

	// WriteAStockGuBen 写入所有A股股本信息到数据库
	WriteAStockGuBen(guBenInfo GuBenInfo) error

	// WriteStockQTDaily 写入当天盘信息到数据库
	WriteStockQTDaily(qt StockQTData) error

	// WriteStockTop10GuDong 写入top10的股东
	WriteStockTop10GuDong(gudongs []StockTop10GuDong) error

	// InsertIgnoreStock 操作可忽略的股
	InsertIgnoreStock(dm string, ignoreType int) error

	// QueryXiaDieTrend 指定代码与日期还在下跌的股票
	QueryXiaDieTrend(daiMa []StockXiaDie, dateStart, dateEnd int64) (trend []StockXiaDie, err error)

	// QueryAllXiaDieTrendByDay 指定日期的下跌股票
	QueryAllXiaDieTrendByDay(dateStart, dateEnd int64) (trend []StockXiaDie, err error)

	// QueryBenchmarkTradeDays 查询某支代码的基准交易日
	QueryBenchmarkTradeDays(daiMa string, benchmarkDay int64, days int) (tradeDays []int64, err error)
}
