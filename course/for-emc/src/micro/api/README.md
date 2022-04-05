## API

本示例介绍使用**API**。

## 使用方式

使用protoc生成相应的代码

```
protoc --proto_path=${GOPATH}/src:. --go_out=. --micro_out=. proto/api.proto 
```

运行**API**网关，我们传入api指令运行：

```
micro api --handler=api
```

再运行本api服务

```
go run api.go
```

## 调用服务

通过URL **/example/call**，就会调用**go.micro.api.example**服务的**Example.Call**接口

请求头的数据会被传到最终调用的接口

```
curl "http://localhost:8080/example/call?name=john"
```

## 设置命名空间

可以通过`--namespace`指定服务命令空间

```
micro api --handler=api --namespace=com.foobar.api
```

或者通过环境变量的方式

```
MICRO_API_NAMESPACE=com.foobar.api micro api --handler=api
```

切记，如果启动时指定命名空间，则代码中的服务名也要注意同步修改前缀，即把**micro.Name**的参数改成对应的命令空间前缀，以便**API**通过解析路由找到它。

```
service := micro.NewService(
        micro.Name("com.foobar.api.example"),
)
```   
