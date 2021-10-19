package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	rcover "github.com/kataras/iris/v12/middleware/recover"
	"goiris/business/app/middleware"
	"goiris/common"
	"goiris/common/support"
)

// Api所有的路由
func HubApi(app *iris.Application) {
	party := preSettring(app)

	BUserRouter(party)
}

func preSettring(app *iris.Application) (party iris.Party) {
	app.Logger().SetLevel(common.G_AppConfig.LogLevel)
	logger, close := support.NewRequestLogger()
	defer close()
	app.Use(
		rcover.New(), 		// 恢复
		logger,       		// 记录请求
		//middleware.ServeHTTP,
	)
	// 定义错误处理
	app.OnErrorCode(iris.StatusNotFound, logger, func(ctx iris.Context) {
		support.Error(ctx, iris.StatusNotFound, support.CODE_NOTFOUNT)
	})
	app.OnErrorCode(iris.StatusInternalServerError, logger, func(ctx iris.Context) {
		support.Error(ctx, iris.StatusInternalServerError, support.CODE_FAILUR)
	})
	//app.OnErrorCode(iris.StatusForbidden, logger, func(ctx iris.Context) {
	//	support.Error(ctx, iris.StatusForbidden, support.CODE_PERMISSION_NIL)
	//})
	////捕获所有http错误:
	//app.OnAnyErrorCode(logger, func(ctx iris.Context) {
	//	//这应该被添加到日志中，因为`logger.Config＃MessageContextKey`
	//	ctx.Values().Set("logger_message", "a dynamic message passed to the logs")
	//	support.Error(ctx, iris.StatusInternalServerError, support.CODE_SYS_ERROR)
	//})
	// 设置跨域
	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST"},//{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug: common.G_AppConfig.Debug,
	})
	party = app.Party("/api/business", crs).AllowMethods(iris.MethodOptions)
	party.Use(middleware.ServeHTTPApi)
	return party
}
