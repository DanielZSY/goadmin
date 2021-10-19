package storage

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"goiris/common"
	"time"
	"xorm.io/core"
)

var G_Xorm *xorm.Engine

// http://gobook.io/read/github.com/go-xorm/manual-en-US/chapter-02/1.mapping.html
func InitXorm() {
	var (
		err  error
		engine *xorm.Engine
		config = common.G_DBConfig
	)
	if engine, err = xorm.NewEngine(config.Mysql.Dialect, config.DBConnUrl()); err != nil {
		goto ERR
	}
	if err = engine.Ping(); err != nil {
		goto ERR
	}
	engine.ShowSQL(config.Mysql.ShowSql)
	engine.SetMapper(core.GonicMapper{})
	engine.SetMaxOpenConns(config.Mysql.MaxOpenConns)
	engine.SetMaxIdleConns(config.Mysql.MaxIdleConns)
	engine.SetTZLocation(time.Local)
	// 性能优化的时候才考虑，加上本机的SQL缓存
	//engine.SetDefaultCacher(xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000))
	G_Xorm = engine
	return
ERR:
	golog.Fatalf("~~> Mysql初始化错误,原因:%s", err.Error())
	return
}
