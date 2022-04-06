package main

import (
	"github.com/micro/go-micro/config/cmd"
	"time"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

func sub() {
	_, err := broker.Subscribe("go.micro.learning.topic.log", func(p broker.Event) error {
		log.Infof("[sub] 收到消息: %s, 消息头: %s\n", string(p.Message().Body), p.Message().Header)
		return nil
	})
	if err != nil {
		log.Errorf("sub error: %s", err)
	}
}

func main() {
	cmd.Init()
	if err := broker.Init(); err != nil {
		log.Fatalf("broker.Init() error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("broker.Connect() error: %v", err)
	}

	go sub()

	<-time.After(time.Second * 10)
}
