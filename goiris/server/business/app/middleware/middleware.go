package middleware

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"goiris/business/app/middleware/jwt"
	api "goiris/business/conf"
	"goiris/common"
	"goiris/common/icasbin"
	"goiris/common/support"
)

type ApiMiddleware struct {
}

func ServeHTTPApi(ctx context.Context) {
	var err  error
	golog.Infof("==> host=[%s], method=[%s], path=[%s]", ctx.Host(), ctx.Method(), ctx.Path())

	// 不需要验证的特殊接口，如：登录
	if func(path string) bool {
		switch path {
		case api.API_USER_LOGOUT, api.API_USER_CODE:
			return  true
		default: break
		}
		return false
	}(ctx.Path()) {
		ctx.Next()
		return
	}
	// 检查回话
	if err = jwt.G_JWT_API.ServeHTTPApi(ctx); err != nil {
		golog.Errorf("中间件token检验失败，错误：%s", err)
		return
	}
	// 验证权限
	if !icasbin.G_Casbin.Enforce(ctx.Values().GetString(support.UID), common.G_AppConfig.Domain, ctx.Path(), ctx.Method(), ".*") {
		support.Error(ctx, iris.StatusForbidden, support.CODE_PERMISSION_NIL)
		return
	}
	// Pass to real API
	ctx.Next()
}
