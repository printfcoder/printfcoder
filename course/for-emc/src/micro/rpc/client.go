package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	proto "github.com/printfcoder/printfcoder/course-for-emc/micro/proto/greeter"
)

func main() {
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// 创建客户端
	greeter := proto.NewGreeterService("greeter.service", service.Client())

	// 调用greeter服务
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Printfcoder"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印响应结果
	fmt.Println(rsp.Greeting)
}
