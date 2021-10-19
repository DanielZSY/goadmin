package common

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

// global var
var (
	G_AppConfig AppConfig
	G_DBConfig  DBConfig
)

// 全局配置文件对应的结构体
type (
	// app
	AppConfig struct {
		iris.Configuration `yaml:"Configuration"`
		Own                `yaml: "own"`
	}
	Own struct {
		Separate      bool     `yaml:"separate"` // 是否前后端分离
		Port          int      `yaml:"port"`
		IgnoreURLs    []string `yaml:"ignoreUrls,flow"`
		InterceptURLs []string `yaml:"interceptUrls,flow"`
		JWTTimeout    int      `yaml:"jwtTimeout"`
		LogLevel      string   `yaml:"logLevel"`
		Secret        string   `yaml:"secret"`
		WebsocketPool int      `yaml:"websocketPool"`
		Domain        string   `yaml:"domain"`
		Debug		  bool	   `yaml:"Debug"`
	}

	// db
	DBConfig struct {
		Redis struct {
			Addr     string `yaml:"addr"`
			Password string `yaml:"password"`
			DB       int    `yaml:"db"`
			PoolSize int    `yaml:"poolSize"`
		}
		Mongodb struct {
			Addr           string `yaml:"addr"`
			Database       string `yaml:"database"`
			ConnectTimeout int    `yaml:"connectTimeout"`
			MaxPoolSize    int    `yaml:"maxPoolSize"`
		}
		Mysql struct {
			Dialect      string `yaml:"dialect"`
			User         string `yaml:"user"`
			Password     string `yaml:"password"`
			Host         string `yaml:"host"`
			Port         int    `yaml:"port"`
			Database     string `yaml:"database"`
			Charset      string `yaml:"charset"`
			ShowSql      bool   `yaml:"showSql"`
			LogLevel     string `yaml:"logLevel"`
			MaxOpenConns int    `yaml:"maxOpenConns"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
			//ParseTime       bool   `yaml:"parseTime"`
			//MaxIdleConns    int    `yaml:"maxIdleConns"`
			//MaxOpenConns    int    `yaml:"maxOpenConns"`
			//ConnMaxLifetime int64  `yaml:"connMaxLifetime: 10"`
			//Sslmode         string `yaml:"sslmode"`
		}
	}
)

func (conf DBConfig) DBConnUrl() string {
	var info = conf.Mysql
	//"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True", info.User, info.Password, info.Host, info.Port, info.Database, info.Charset)
	//return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", info.User, info.Password, info.Host, info.Port, info.Database, info.Charset)
}

func InitConfig() {
	var (
		app  = AppConfig{}
		db   = DBConfig{}
		err  error
	)
	var configViperConfig = viper.New()
	// app
	configViperConfig.SetConfigFile("admin/conf/app.yaml")
	//读取配置文件内容
	if err = configViperConfig.ReadInConfig(); err != nil {
		goto ERR
	}
	if err = configViperConfig.Unmarshal(&app); err != nil {
		goto ERR
	}
	G_AppConfig = app
	golog.Infof("[app config]=> %v", G_AppConfig)

	// db
	configViperConfig.SetConfigFile("admin/conf/db.yaml")
	//读取配置文件内容
	if err = configViperConfig.ReadInConfig(); err != nil {
		goto ERR
	}
	if err = configViperConfig.Unmarshal(&db); err != nil {
		goto ERR
	}
	G_DBConfig = db
	golog.Infof("[db  config]=> %v", G_DBConfig)
	return
ERR:
	golog.Fatalf("~~> 解析配置文件错误,原因:%s", err.Error())
}
