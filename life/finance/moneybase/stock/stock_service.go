package stock

import (
	"context"
	"fmt"

	"github.com/printfcoder/printfcoder/life/finance/moneybase/common"
	log "github.com/stack-labs/stack/logger"
)

// SyncAllStockBases 从A股中同步所有股票基本信息
func SyncAllStockBases(ctx context.Context) error {
	return SyncerAdapter(ctx).SyncAllStockBases()
}

// SyncSingleStockGuBen 从A股中同步单个股票股本
func SyncSingleStockGuBen(ctx context.Context, code string) error {
	return SyncerAdapter(ctx).SyncSingleStockGuBen(code)
}

// SyncAllStockGuBen 从A股中同步所有股票股本
func SyncAllStockGuBen(ctx context.Context) error {
	return SyncerAdapter(ctx).SyncAllStockGuBen()
}

// GetStockQT 从A股获取股票QT信息
func GetStockQT(ctx context.Context, date string, symbols ...string) ([]StockQTDataTencent, error) {
	return SyncerAdapter(ctx).GetStockQT(date, symbols...)
}

// GetStockTop10GuDong GetStockTop10GuDong 获取 Top 10股东
func GetStockTop10GuDong(ctx context.Context, symbol, startDate string, endDate string, guDongType int) ([]StockTop10GuDong, error) {
	return SyncerAdapter(ctx).GetStockTop10GuDong(symbol, startDate, endDate, guDongType)
}

// SyncTop10Gudong 写入 Top 10股东
func SyncTop10Gudong(ctx context.Context, startDate string, endDate string, guDongType int) error {
	return SyncerAdapter(ctx).SyncTop10GuDong(startDate, endDate, guDongType)
}

// WriteSingleStockQTDaily 写入单个每天QT
func WriteSingleStockQTDaily(ctx context.Context, date string, symbol string) (err error) {
	return SyncerAdapter(ctx).WriteSingleStockQTDaily(date, symbol)
}

// WriteStockQTDaily 写入每天QT
func WriteStockQTDaily(ctx context.Context, date string) (err error) {
	return SyncerAdapter(ctx).WriteStockQTDaily(date)
}

// QueryXiaDieByDays 查询一直在下跌的
func QueryXiaDieByDays(ctx context.Context, date string, days int) (trends map[int64][]StockXiaDie, err error) {
	dao := GetDao(ctx)

	start, end := common.ParseDayStartAndEnd(date)

	stocks, err := dao.QueryAllXiaDieTrendByDay(start, end)
	if err != nil {
		err = fmt.Errorf("[QueryXiaDieByDays] 查询所有stock异常：%s", err)
		log.Error(err)
		return
	}

	tradeDays, err := dao.QueryBenchmarkTradeDays("000001", end, days)
	if err != nil {
		err = fmt.Errorf("[QueryXiaDieByDays] 查询所有基准交易日异常：%s", err)
		log.Error(err)
		return
	}

	trends = make(map[int64][]StockXiaDie, len(tradeDays))
	tempStocks := stocks
	for _, tradeDay := range tradeDays {
		startIn, endIn := common.ParseDayStartAndEnd(fmt.Sprintf("%d", tradeDay))
		xiaDies, errIn := dao.QueryXiaDieTrend(tempStocks, startIn, endIn)
		if errIn != nil {
			err = fmt.Errorf("[QueryXiaDieByDays] 查询指定日期下跌股票异常：%s", errIn)
			log.Error(err)
			return
		}

		trends[tradeDay] = xiaDies

		var tempStocksIn []StockXiaDie
		for i := range xiaDies {
			tempStocksIn = append(tempStocksIn, StockXiaDie{DaiMa: xiaDies[i].DaiMa})
		}

		log.Infof("[QueryXiaDieByDays] day[%d] stocks current[%d] next[%d]", tradeDay, len(tempStocks), len(tempStocksIn))
		tempStocks = tempStocksIn
	}

	return
}
