package auth

import (
	"goiris/business/conf"
	"goiris/common"
	"goiris/common/icasbin"
)

func InitAuthApi() {
	icasbin.InitCasbin(common.G_AppConfig.Domain)
	icasbin.G_Casbin.AddFunction("myApiFunc", myApiFunc)
}
func myApiFunc(args ...interface{}) (interface{}, error) {
	match := func(key string) bool {
		switch key {
		case api.API_USER_LOGOUT, api.API_USER_CODE:
			return  true
		default: break
		}
		return false
	}(args[0].(string))
	return (bool)(match), nil
}
