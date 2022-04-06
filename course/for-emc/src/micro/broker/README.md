# Broker

Micro中的[broker](https://godoc.org/github.com/micro/go-micro/broker#Broker)用于代理消息的发布与订阅。

## 内容

- broker.go - main程序运行两个协程20秒钟，各自负责发布与订阅消息。

## 运行程序

如果使用默认的http broker，请运行：

```bash
go run sub.go
```

```bash
go run pub.go
```

如果想使用其他消息队列服务，例如nats，请运行：

```bash
export MICRO_BROKER=nats
go run sub.go
```

或者：

```bash
go run sub.go --broker=nats
```

或者：

```bash
go run sub.go --broker=nats --broker_address=127.0.0.1:4222
```

> 注意，要先显式调用cmd.Init()，让cmd识别参数，否则指定Broker中间件会无效，如下

```go
func main() {
    cmd.Init()
    if err := broker.Init(); err != nil {
    //... 省略的代码
```
