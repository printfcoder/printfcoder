package main

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
	proto "github.com/printfcoder/printfcoder/course-for-emc/micro/proto/api"
)

type Example struct{}

type Foo struct{}

// Example.Call 通过API向外暴露为/example/call，接收http请求
// 即：/example/call请求会调用go.micro.api.example服务的Example.Call方法
func (e *Example) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Example.Call接口收到请求")

	name, ok := req.Get["name"]

	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "参数不正确")
	}

	rsp.StatusCode = 200

	b, _ := json.Marshal(map[string]string{
		"message": "我们已经收到你的请求，" + strings.Join(name.Values, " "),
	})

	// 设置返回值
	rsp.Body = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)

	service.Init()

	// 注册 example handler
	proto.RegisterExampleHandler(service.Server(), new(Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
