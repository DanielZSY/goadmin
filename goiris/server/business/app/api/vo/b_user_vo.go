package vo

import (
	"goiris/business/app/model"
	"goiris/common/support"
)

type BUserVO struct {
	support.Pagination
}

// 更新和创建用户时需要的vo
type AcceptBUserVO struct {
	// 用户本身的信息
	BUser  *model.BUser    `json:"bUser"`
}
