package main

import (
	"context"

	"github.com/printfcoder/printfcoder/life/finance/moneybase/stock"
	"github.com/stack-labs/stack"
	"github.com/stack-labs/stack/service/web"

	_ "github.com/lib/pq"
)

func main() {
	s := stack.NewWebService(
		stack.Name("prinfcoder.money.base"),
		stack.Address("localhost:8899"),
	)

	err := s.Init(
		stack.BeforeStart(func() error {
			err := stock.Init(context.Background())
			if err != nil {
				return err
			}

			return nil
		}),
		stack.WebRootPath("/api/money-base"),
		stack.WebHandleFuncs(
			handlers()...,
		),
	)
	if err != nil {
		panic(err)
	}

	err = s.Run()
	if err != nil {
		panic(err)
	}
}

func handlers() []web.HandlerFunc {
	var ret []web.HandlerFunc
	ret = append(stock.Handlers())

	return ret
}
