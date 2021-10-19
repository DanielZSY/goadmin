package support

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// common error define
func Error(ctx iris.Context, status int, code Code) {
	ctx.StatusCode(status)
	ctx.JSON(json(code, nil))
}

// 200 无数据
func Ok_(ctx iris.Context, code Code) {
	Ok(ctx, code, nil)
}
// 200 有数据
func Ok(ctx iris.Context, code Code, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(json(code, data))
}

// 401 error 无授权
func Unauthorized(ctx iris.Context, code Code) {
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.JSON(json(code, nil))
}
// 自定义错误
func InternalServerError(ctx iris.Context, code Code) {
	ctx.StatusCode(iris.StatusInternalServerError)
	ctx.JSON(json(code, nil))
}
// 分页数据结构
func PaginationTableData(ctx iris.Context, total int64, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(json(CODE_OK, iris.Map{
		"total": total,
		"data":  data,
	}))
}
// 返回json组装
func json(code Code, data interface{}) context.Map {
	return iris.Map{
		"code": fmt.Sprintf("%d", code),
		"msg":  code.String(),
		"data": data,
	}
}
