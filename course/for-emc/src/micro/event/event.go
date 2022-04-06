package main

import (
	"context"
	"github.com/micro/go-micro"
	proto "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/util/log"
)

// 切记，事件订阅结构的所有公有方法都会被执行 方法名没有限制，但是方法一定要接收ctx，event
type Event struct{}

func (e *Event) Process(ctx context.Context, event *proto.Event) error {
	log.Info("Process 收到事件，", event.Name)
	log.Info("Process 数据", event.Data)
	return nil
}

func main() {
	service := micro.NewService(
		// 服务名可以随意
		micro.Name("go.micro.evt.user"),
	)
	service.Init()

	// register subscriber
	// 注意topic不能随意，会与micro api配合，找到user命名空间
	micro.RegisterSubscriber("go.micro.evt.user", service.Server(), new(Event))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
