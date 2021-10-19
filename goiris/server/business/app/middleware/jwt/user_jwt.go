package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12/context"
	"goiris/business/app/model"
	"goiris/common"
	"goiris/common/ijwt"
	"goiris/common/storage"
	"goiris/common/support"
	"strconv"
	"time"
)

var G_JWT_API *JWT

type (
	JWT struct {
		*ijwt.IJwt
	}
	Claims struct {
		Id     int64  `json:"id"`
		Phone  string `json:"phone"`
		Enable int    `json:"enable"`
		jwt.StandardClaims
	}
)

func InitJWTApi() {
	iJwt := ijwt.IJwt{}
	G_JWT_API = &JWT{iJwt.InitJwtConfig()}
}

// Serve the middleware's action
func (j *JWT) ServeHTTPApi(ctx context.Context) (err error) {
	var (
		token *jwt.Token
		user  *model.BUser
	)
	if token, err = j.Check(ctx); err != nil {
		return err
	}

	if user, err = j.Token2ModelApi(token); err != nil {
		return err
	}
	// 检查redis缓存
	if _, err = storage.G_Redis.GetToken(support.REDIS_USER_FORMAT, strconv.FormatInt(user.Id, 10)); err != nil {
		return err
	}
	// token校验通过，设置当前用户id到上下文
	ctx.Values().Set(support.UID, user.Id)
	return nil
}

// Serve the middleware's action
func (j *JWT) ServeWebsocketApi(ctx context.Context) {
	var (
		token *jwt.Token
		user  *model.BUser
		err error
	)
	if token, err = j.Check(ctx); err != nil {
		return
	}

	if user, err = j.Token2ModelApi(token); err != nil {
		return
	}
	if _, err = storage.G_Redis.GetToken(support.REDIS_USER_FORMAT, strconv.FormatInt(user.Id, 10)); err != nil {
		return
	}
	// If everything ok then call next.
	ctx.Next()
}

// 在登录成功的时候生成token
func (j *JWT) GenerateTokenApi(user *model.BUser) (string, error) {
	expireTime := time.Now().Add(time.Duration(common.G_AppConfig.JWTTimeout) * time.Minute)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		user.Id,
		user.Phone,
		user.Enable,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "bst-community-jwt",
		},
	})
	return tokenClaims.SignedString([]byte(common.G_AppConfig.Secret))
}

// 解析token的信息为用户
func (j *JWT) Token2ModelApi(token *jwt.Token) (*model.BUser, error) {
	//mapClaims := (jwt.Get(ctx).Claims).(jwt.MapClaims)
	var (
		mapClaims, ok = token.Claims.(jwt.MapClaims)
		id            int64
		phone         string
		enable        int
	)
	if !ok {
		return nil, fmt.Errorf("%s", support.CODE_TOKEN_INVALID.String())
	}

	id = int64(mapClaims["id"].(float64))
	phone = mapClaims["phone"].(string)
	enable = int(mapClaims["enable"].(float64))
	return &model.BUser{
		Id:     id,
		Phone:  phone,
		Enable: enable,
	}, nil
}

func (j *JWT) TokenString2ModelApi(tokenString string) (user *model.BUser, err error) {
	var token *jwt.Token
	if token, err = jwt.Parse(tokenString, j.Config.ValidationKeyGetter); err != nil {
		return nil, err
	}
	return j.Token2ModelApi(token)
}
