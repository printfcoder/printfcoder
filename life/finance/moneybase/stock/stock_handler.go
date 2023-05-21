package stock

import (
	"context"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/printfcoder/printfcoder/life/finance/moneybase/common"
)

func Handlers() []common.HandlerFunc {
	return []common.HandlerFunc{
		{
			Method:      "POST",
			Path:        "sync-stock-base",
			HandlerFunc: syncStocks,
		},
		{
			Method:      "POST",
			Path:        "sync-single-stock-guben",
			HandlerFunc: syncSingleGuBen,
		},
		{
			Method:      "POST",
			Path:        "sync-all-stock-guben",
			HandlerFunc: syncAllGuBen,
		},
		{
			Method:      "POST",
			Path:        "sync-top10-gudong",
			HandlerFunc: syncTop10Gudong,
		},
		{
			Method:      "GET",
			Path:        "get-current-value",
			HandlerFunc: getCurrentValue,
		},
		{
			Method:      "GET",
			Path:        "get-top10-gudong",
			HandlerFunc: getStockTop10GuDong,
		},
		{
			Method:      "POST",
			Path:        "write-single-qt-daily",
			HandlerFunc: writeSingleQTDaily,
		},
		{
			Method:      "POST",
			Path:        "write-qt-daily",
			HandlerFunc: writeQTDaily,
		},
		{
			Method:      "GET",
			Path:        "query-xia-die-by-days",
			HandlerFunc: queryXiaDieByDays,
		},
	}
}

func syncStocks(ctx context.Context, r *app.RequestContext) {
	rsp := &common.HTTPRsp{}
	nation := r.Query("nation")
	if nation != "cn" {
		nation = "cn"
	}

	var err error
	switch nation {
	case "cn":
		err = SyncAllStockBases(ctx)
	default:
		// do nothing
	}

	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	common.WriteSuccessHTTP(r, rsp)
}

func syncSingleGuBen(ctx context.Context, r *app.RequestContext) {
	rsp := &common.HTTPRsp{}
	nation := r.Query("nation")
	if nation != "cn" {
		nation = "cn"
	}

	code := r.Query("code")
	if code == "" {
		common.WriteFailHTTP(r, rsp, common.ErrorStockInvalidCode)
		return
	}

	var err error
	switch nation {
	case "cn":
		err = SyncSingleStockGuBen(ctx, code)
	default:
		// do nothing
	}

	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	common.WriteSuccessHTTP(r, rsp)
}

func syncAllGuBen(ctx context.Context, r *app.RequestContext) {
	rsp := &common.HTTPRsp{}
	nation := r.Query("nation")
	if nation != "cn" {
		nation = "cn"
	}

	var err error
	switch nation {
	case "cn":
		err = SyncAllStockGuBen(ctx)
	default:
		// do nothing
	}

	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	common.WriteSuccessHTTP(r, rsp)
}

func getCurrentValue(ctx context.Context, r *app.RequestContext) {
	rsp := &common.HTTPRsp{}

	symbolStr := r.Query("symbols")
	if symbolStr == "" {
		common.WriteFailHTTP(r, rsp, common.ErrorStockInvalidCode)
		return
	}

	symbols := strings.Split(symbolStr, ",")

	date := r.Query("date")
	if date == "" {
		date = common.TodayStr()
	}

	stockQTs, err := GetStockQT(ctx, date, symbols...)
	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	rsp.Data = stockQTs
	common.WriteSuccessHTTP(r, rsp)
}

func getStockTop10GuDong(ctx context.Context, r *app.RequestContext) {
	rsp := &common.HTTPRsp{}

	symbol := r.Query("symbol")
	if symbol == "" {
		common.WriteFailHTTP(r, rsp, common.ErrorStockInvalidCode)
		return
	}

	startDate := r.Query("startDate")
	if startDate == "" {
		startDate = "20210101"
	}

	endDate := r.Query("endDate")
	if endDate == "" {
		endDate = common.TodayStr()
	}

	guDongType := r.Query("guDongType")
	if guDongType == "" || (guDongType != "1" && guDongType != "2") {
		common.WriteFailHTTP(r, rsp, common.ErrorStockTop10GuDongInvalidGuDongTypeError)
		return
	}

	guDongTypeInt, _ := strconv.ParseInt(guDongType, 10, 64)

	stockQTs, err := GetStockTop10GuDong(ctx, symbol, startDate, endDate, int(guDongTypeInt))
	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	rsp.Data = stockQTs
	common.WriteSuccessHTTP(r, rsp)
}

func syncTop10Gudong(ctx context.Context, r *app.RequestContext) {
	rsp := &common.HTTPRsp{}

	startDate := r.Query("startDate")
	if startDate == "" {
		startDate = "20210101"
	}

	endDate := r.Query("endDate")
	if endDate == "" {
		endDate = common.TodayStr()
	}

	guDongType := r.Query("guDongType")
	if guDongType == "" || (guDongType != "1" && guDongType != "2") {
		common.WriteFailHTTP(r, rsp, common.ErrorStockTop10GuDongInvalidGuDongTypeError)
		return
	}

	guDongTypeInt, _ := strconv.ParseInt(guDongType, 10, 64)
	err := SyncTop10Gudong(ctx, startDate, endDate, int(guDongTypeInt))
	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	common.WriteSuccessHTTP(r, rsp)
}

func writeSingleQTDaily(ctx context.Context, r *app.RequestContext) {
	rsp := &common.HTTPRsp{}

	symbol := r.Query("symbol")
	if symbol == "" {
		common.WriteFailHTTP(r, rsp, common.ErrorStockInvalidCode)
		return
	}

	date := r.Query("date")
	if date == "" {
		date = common.TodayStr()
	}

	err := WriteSingleStockQTDaily(ctx, date, symbol)
	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	common.WriteSuccessHTTP(r, rsp)
}

func writeQTDaily(ctx context.Context, r *app.RequestContext) {
	rsp := &common.HTTPRsp{}

	date := r.Query("date")
	if date == "" {
		date = common.TodayStr()
	}

	err := WriteStockQTDaily(ctx, date)
	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	common.WriteSuccessHTTP(r, rsp)
}

func queryXiaDieByDays(ctx context.Context, r *app.RequestContext) {
	rsp := &common.HTTPRsp{}

	date := r.Query("date")
	if date == "" {
		date = common.TodayStr()
	}

	daysStr := r.Query("days")
	if date == "" {
		date = common.TodayStr()
	}

	days, err := strconv.ParseInt(daysStr, 10, 10)
	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	tradeTrends, err := QueryXiaDieByDays(ctx, date, int(days))
	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	rspData := StockXiaDieRsp{
		Trend:      tradeTrends,
		TrendCount: map[int64]int{},
	}

	for k, v := range rspData.Trend {
		rspData.TrendCount[k] = len(v)
	}

	rsp.Data = rspData

	common.WriteSuccessHTTP(r, rsp)
}
