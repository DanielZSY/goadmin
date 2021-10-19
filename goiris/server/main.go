package main

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"goiris/admin/app/middleware/auth"
	"goiris/admin/app/middleware/jwt"
	"goiris/admin/app/web/router"
	"goiris/common"
	"goiris/common/storage"
	"goiris/common/support"
	"goiris/sockets"
)

func main() {
	common.InitConfig()
	support.InitLog()
	support.InitValidator()
	storage.InitGorm()
	storage.InitRedis()
	auth.InitAuth()
	jwt.InitJWT()

	app := iris.New()
	// Websocket
	sockets.InitWebsocket(app)
	// HTTP
	router.Hub(app)
	// Start api server
	if err := app.Run(
		iris.Addr(fmt.Sprintf(":%d", common.G_AppConfig.Port)),
		iris.WithConfiguration(common.G_AppConfig.Configuration)); err != nil {
		golog.Fatalf("Start admin server failed. And err:%s", err.Error())
	}
}
