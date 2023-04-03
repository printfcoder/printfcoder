package stock

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	log "github.com/stack-labs/stack/logger"
)

type methodWrapperKey struct{}

func MethodNameInjectWrapper(c context.Context, ctx *app.RequestContext) {
	log.Infof("[MethodNameInjectWrapper] 收到请求，请求接口：%s", ctx.URI())
	methodName := string(ctx.URI().Path())
	split := strings.Split(methodName, "/")
	methodName = split[len(split)-1]
	c = context.WithValue(c, methodWrapperKey{}, methodName)
	ctx.Next(c)
}

// 获取注入的方法名
func getMethodNameFromHTTP(ctx context.Context) (string, error) {
	if method, ok := ctx.Value(methodWrapperKey{}).(string); ok {
		return method, nil
	}

	err := fmt.Errorf("同步方法未找到")
	log.Errorf("获取方法名异常：%s", err)

	return "", err
}
