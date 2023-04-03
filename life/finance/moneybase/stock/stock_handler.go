package stock

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/printfcoder/printfcoder/life/finance/moneybase/common"
)

func Handlers() []common.HandlerFunc {
	return []common.HandlerFunc{
		{
			Method:      "POST",
			Path:        "sync-stock-base",
			HandlerFunc: SyncStocks,
		},
	}
}

func SyncStocks(ctx context.Context, r *app.RequestContext) {
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
