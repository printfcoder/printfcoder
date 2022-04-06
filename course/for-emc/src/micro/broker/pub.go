package main

import (
	"fmt"
	"time"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("第%d: 条消息", i)),
		}

		if err := broker.Publish("go.micro.learning.topic.log", msg); err != nil {
			log.Infof("[pub] 推送消息失败: %v", err)
		} else {
			log.Infof("[pub] 消息推送完成: %s", string(msg.Body))
		}
		i++
	}
}

func main() {
	if err := broker.Init(); err != nil {
		log.Fatalf("broker.Init() error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("broker.Connect() error: %v", err)
	}

	go pub()

	<-time.After(time.Second * 10)
}
