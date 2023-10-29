package middleware

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/iim-api/utility/logger"
	"strings"
)

const JWTSessionConst = "__JWT_SESSION__"
const UID_KEY = "uid"

var (
	ErrorNoLogin = errors.New("请登录后操作")
)

type IStorage interface {
	// 判断是否是黑名单
	IsBlackList(ctx context.Context, token string) bool
}

type JSession struct {
	Uid       int    `json:"uid"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

func Auth(r *ghttp.Request) {

	//r.SetCtxVar(UID_KEY, uid)

	if gstr.HasPrefix(r.GetHeader("Content-Type"), "application/json") {
		logger.Debugf(r.GetCtx(), "url: %s, request body: %s", r.GetUrl(), r.GetBodyString())
	} else {
		logger.Debugf(r.GetCtx(), "url: %s, Content-Type: %s", r.GetUrl(), r.GetHeader("Content-Type"))
	}

	r.Middleware.Next()
}

func AuthHeaderToken(r *ghttp.Request) string {

	token := r.GetHeader("Authorization")
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

	// Headers 中没有授权信息则读取 url 中的 token
	if token == "" {
		token = r.Get("token", "").String()
	}

	return token
}
