package main

import (
	"context"

	"github.com/printfcoder/printfcoder/life/finance/moneybase/db"
	"github.com/stack-labs/stack"

	_ "github.com/lib/pq"
)

func main() {
	s := stack.NewWebService(stack.Name("prinfcoder.money.base"))
	err := s.Init(
		stack.BeforeStart(func() error {
			err := db.Init(context.Background())
			return err
		}))
	if err != nil {
		panic(err)
	}

	err = s.Run()
	if err != nil {
		panic(err)
	}
}
