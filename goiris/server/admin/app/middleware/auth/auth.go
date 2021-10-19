package auth

import (
	"goiris/admin/conf"
	"goiris/common"
	"goiris/common/icasbin"
)

func InitAuth() {
	icasbin.InitCasbin(common.G_AppConfig.Domain)
	icasbin.G_Casbin.AddFunction("myFunc", myFunc)
}
func myFunc(args ...interface{}) (interface{}, error) {
	match := func(key string) bool {
		switch key {
		case api.ADMIN_USER_LOGOUT, api.ADMIN_USER_BASEINFO, api.ADMIN_MENU_ALL, api.ADMIN_MENU_OWN:
			return  true
		default: break
		}
		return false
	}(args[0].(string))
	return (bool)(match), nil
}
