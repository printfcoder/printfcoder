package stock

import (
	"net/http"

	"github.com/printfcoder/printfcoder/life/finance/moneybase/common"
	"github.com/stack-labs/stack/service/web"
)

func Handlers() []web.HandlerFunc {
	return []web.HandlerFunc{
		{
			"sync-stock-base",
			SyncStocks,
		},
	}
}

func SyncStocks(w http.ResponseWriter, r *http.Request) {
	rsp := &common.HTTPRsp{}
	nation := r.URL.Query().Get("nation")
	if nation != "cn" {
		nation = "cn"
	}

	var err error
	switch nation {
	case "cn":
		err = SyncAllStockBases()
	default:
		// do nothing
	}

	if err != nil {
		common.WriteFailHTTP(w, rsp, err)
		return
	}

	common.WriteSuccessHTTP(w, rsp)
}
