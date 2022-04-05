# Service

本篇演示如何使用rpc service。

## 内容

- main.go - 服务端
- client.go - 客户端

## 运行

使用protoc生成相应的代码

```
protoc --go_out=. --micro_out=. proto/greeter/greeter.proto
```

> proto原型文件统一放到micro/proto目录下，故执行protoc指令时，要到micro目录

运行服务端

```shell
go run server.go
```

运行客户端

```shell
go run client.go
```
