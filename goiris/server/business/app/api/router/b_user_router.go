package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"goiris/business/app/api/mapper"
	"goiris/business/app/api/vo"
	"goiris/business/app/middleware/jwt"
	"goiris/business/app/model"
	"goiris/common"
	"goiris/common/icasbin"
	"goiris/common/storage"
	"goiris/common/support"
	"strconv"
	"strings"
	"time"
)

func BUserRouter(party iris.Party) {
	var (
		adinUser = party.Party("/user")
		aus      = BUserService{mapper: mapper.BUserMapper{}}
	)
	adinUser.Post("/register", hero.Handler(aus.create))
	adinUser.Post("/update", hero.Handler(aus.update))
	adinUser.Post("/delete", hero.Handler(aus.delete))
	adinUser.Post("/info", hero.Handler(aus.userInfo))
	adinUser.Post("/login", hero.Handler(aus.login))
	adinUser.Post("/logout", hero.Handler(aus.logout))
}

type BUserService struct {
	mapper mapper.BUserMapper
}

func (aus *BUserService) create(ctx iris.Context) {
	var (
		code      support.Code
		err       error
		vo = new(vo.AcceptBUserVO)
	)
	if err = ctx.ReadJSON(&vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if vo.BUser.Gender == 0 {
		vo.BUser.Gender = 1 // 默认性别设置为男
	}
	vo.BUser.Password = support.AESEncrypt([]byte(vo.BUser.Password))
	vo.BUser.Enable = 1
	vo.BUser.CreateTime = time.Now()
	if err = aus.mapper.Insert(vo); err != nil {
		code = support.CODE_USER_PHONE_DUPLICATE
		goto ERR
	}
	icasbin.G_Casbin.LoadPolicy()
	support.Ok_(ctx, support.CODE_USER_REGISTE_OK)
	return
ERR:
	golog.Errorf("用户[%s]注册失败。原因：%s，错误：%s", vo.BUser.Phone, code, err)
	support.InternalServerError(ctx, code)
}

func (aus *BUserService) delete(ctx iris.Context) {
	var (
		code support.Code
		err  error
		ids  []int64
	)
	if ids, err = func() (result []int64, err error) {
		idStr := strings.Split(ctx.URLParam("ids"), ",")
		for _, v := range idStr {
			var id int64
			if id, err = strconv.ParseInt(v, 10, 64); err != nil {
				code = support.CODE_SYS_PARSE_PARAMS_ERROR
				return nil, err
			}
			result = append(result, id)
		}
		return
	}(); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = aus.mapper.Delete(ids); err != nil {
		goto ERR
	}
	support.Ok_(ctx, support.CODE_OK)
	return
ERR:
	golog.Errorf("操作失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

func (aus *BUserService) update(ctx iris.Context) {
	var (
		code      support.Code
		err       error
		vo = new(vo.AcceptBUserVO)
	)
	if err = ctx.ReadJSON(&vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	vo.BUser.Password = support.AESEncrypt([]byte(vo.BUser.Password))
	if err = aus.mapper.UpdateOne(vo); err != nil {
		goto ERR
	}
	support.Ok_(ctx, support.CODE_OK)
	return
ERR:
	golog.Errorf("更新用户[%d]信息失败。原因：%s，错误：%s", vo.BUser.Id, code.String(), err)
	support.InternalServerError(ctx, code)
}

func (aus *BUserService) list(ctx iris.Context) {
	var (
		err    error
		code   support.Code
		vo     = new(vo.BUserVO)
		total  int64
		result []*model.BUser
	)
	if err = ctx.ReadJSON(vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	vo.Init()
	if total, result, err = aus.mapper.FindList(vo); err != nil {
		goto ERR
	}
	support.PaginationTableData(ctx, total, result)
	return
ERR:
	golog.Errorf("失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

// 登录获取token
func (aus *BUserService) login(ctx iris.Context) {
	var (
		code       support.Code
		err        error
		bUser  = new(model.BUser)
		mBUser = new(model.BUser)
		ckPassword bool
		token      string
	)
	// TODO 同时登陆一个账号，将已登陆的用户挤下线
	if err = ctx.ReadJSON(&bUser); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(bUser); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	mBUser.Username = bUser.Username
	mBUser.Enable = 1
	if mBUser, err = aus.mapper.FindOne(mBUser); err != nil {
		code = support.CODE_USER_PHONE_FAILUR
		goto ERR
	}
	ckPassword = support.CheckPWD(bUser.Password, mBUser.Password)
	if !ckPassword {
		code = support.CODE_USER_PASSWORD_FAILUR
		goto ERR
	}

	if token, err = jwt.G_JWT_API.GenerateTokenApi(mBUser); err != nil {
		code = support.CODE_TOKEN_CREATE_FAILUR
		goto ERR
	}
	if err = storage.G_Redis.SetToken(support.REDIS_USER_FORMAT, strconv.FormatInt(mBUser.Id, 10), token); err != nil {
		code = support.CODE_TOKEN_CACHE_ERROR
		goto ERR
	}
	support.Ok(ctx, support.CODE_USER_LOGIN_OK, iris.Map{
		"token":  token,
		"expire": common.G_AppConfig.Own.JWTTimeout,
	})
	return
ERR:
	golog.Errorf("~~> B用户[%s]登录失败。原因：%s，错误：%s", bUser.Username, code.String(), err)
	support.InternalServerError(ctx, code)
}

// 登录后获取用户信息
func (aus *BUserService) userInfo(ctx iris.Context) {
	var (
		code      support.Code
		err       error
		BUser *model.BUser
		uid       int64
	)
	if uid, err = ctx.Values().GetInt64(support.UID); err != nil {
		goto ERR
	}
	golog.Infof("userInfo, uid=%d", uid)
	if BUser, err = aus.mapper.FindOne(&model.BUser{Id: uid}); err != nil {
		code = support.CODE_SYS_FAILUR
		goto ERR
	}
	BUser.Password = ""
	BUser.Online = true
	support.Ok(ctx, support.CODE_USER_LOGIN_OK, BUser)
	return
ERR:
	golog.Errorf("~~> 获取用户[%d]信息失败。原因：%s， 错误：%s", uid, code.String(), err.Error())
	support.InternalServerError(ctx, code)
	return
}

func (aus *BUserService) logout(ctx iris.Context) {
	var (
		err  error
		code support.Code
		uid  int64
	)
	if uid, err = ctx.Values().GetInt64(support.UID); err != nil {
		goto ERR
	}
	if _, err = storage.G_Redis.DelToken(support.REDIS_USER_FORMAT, strconv.FormatInt(uid, 10)); err != nil {
		code = support.CODE_TOKEN_CACHE_ERROR
		goto ERR
	}
	support.Ok_(ctx, support.Code(200))
	return
ERR:
	golog.Errorf("~~> 用户[%d]退出失败。原因：%s，错误：%s", uid, code.String(), err)
	support.InternalServerError(ctx, code)
	return
}
