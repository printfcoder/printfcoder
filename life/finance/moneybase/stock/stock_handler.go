package stock

import (
	"context"
	"fmt"
	log "github.com/stack-labs/stack/logger"
	"strings"
	"time"

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
			Method:      "GET",
			Path:        "get-current-value",
			HandlerFunc: getCurrentValue,
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
		date = todayStr()
	}

	stockQTs, err := GetStockQT(ctx, date, symbols...)
	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	rsp.Data = stockQTs
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
		date = todayStr()
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
		date = todayStr()
	}

	err := WriteStockQTDaily(ctx, date)
	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	common.WriteSuccessHTTP(r, rsp)
}

func todayStr() string {
	now := time.Now()
	year := now.Year()
	mon := now.Month()
	day := now.Day()

	monS := fmt.Sprintf("%d", mon)
	if mon < 10 {
		monS = "0" + monS
	}

	dayS := fmt.Sprintf("%d", day)
	if day < 10 {
		dayS = "0" + dayS
	}
	ret := fmt.Sprintf("%d%s%s", year, monS, dayS)

	log.Infof("[todayStr] use today: %s", ret)
	return ret
}
