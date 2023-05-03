package stock

import (
	"context"
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

	stockQTs, err := GetStockQT(ctx, symbols...)
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

	err := WriteSingleStockQTDaily(ctx, symbol)
	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	common.WriteSuccessHTTP(r, rsp)
}

func writeQTDaily(ctx context.Context, r *app.RequestContext) {
	rsp := &common.HTTPRsp{}

	err := WriteStockQTDaily(ctx)
	if err != nil {
		common.WriteFailHTTP(r, rsp, err)
		return
	}

	common.WriteSuccessHTTP(r, rsp)
}
