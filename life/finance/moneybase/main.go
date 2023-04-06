package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
	_ "github.com/lib/pq"
	"github.com/printfcoder/printfcoder/life/finance/moneybase/common"
	"github.com/printfcoder/printfcoder/life/finance/moneybase/stock"
	"github.com/stack-labs/stack/config"
	"github.com/stack-labs/stack/pkg/config/source/file"
)

func init() {
	wkDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	c := config.NewConfig(config.Source(file.NewSource(file.WithPath(fmt.Sprintf("%s%s%s", wkDir, string(os.PathSeparator), "stock-biz.yml")))))
	err = c.Init()
	if err != nil {
		panic(err)
	}
}

func main() {
	err := stock.Init(context.Background())
	if err != nil {
		panic(err)
	}

	h := server.Default(
		server.WithBasePath("/api/money-base"),
		server.WithHostPorts(":8899"),
	)
	h.Use(stock.MethodNameInjectWrapper)

	handlers(h)
	h.Name = "printfcoder.money.base"

	h.Spin()
}

func handlers(h *server.Hertz) {
	var hdls []common.HandlerFunc
	hdls = append(stock.Handlers())

	for _, hdl := range hdls {
		h.Handle(hdl.Method, hdl.Path, hdl.HandlerFunc)
	}

	return
}
