package stock

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	log "github.com/stack-labs/stack/logger"
)

type methodWrapperKey struct{}
type syncerWrapperKey struct{}

func MethodNameInjectWrapper(c context.Context, ctx *app.RequestContext) {
	log.Infof("[MethodNameInjectWrapper] 收到请求，请求接口：[%s], syncer: [%s]", ctx.URI(), ctx.Query("syncer"))
	methodName := string(ctx.URI().Path())
	split := strings.Split(methodName, "/")
	methodName = split[len(split)-1]
	c = context.WithValue(c, methodWrapperKey{}, methodName)
	if ctx.Query("syncer") != "" {
		c = context.WithValue(c, syncerWrapperKey{}, ctx.Query("syncer"))
	}

	ctx.Next(c)
}

// 获取注入的方法名
func getMethodNameFromHTTP(ctx context.Context) (string, string, error) {
	if method, ok := ctx.Value(methodWrapperKey{}).(string); ok {
		if syncerName, alsoOK := ctx.Value(syncerWrapperKey{}).(string); alsoOK {
			return method, syncerName, nil
		}

		return method, "", nil
	}

	err := fmt.Errorf("同步方法未找到")
	log.Errorf("获取方法名异常：%s", err)

	return "", "", err
}
