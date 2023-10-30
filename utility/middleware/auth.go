package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/iim-api/internal/consts"
	"github.com/iimeta/iim-api/internal/service"
	"github.com/iimeta/iim-api/utility/logger"
	"net/http"
	"strings"
)

type JSession struct {
	Uid       int    `json:"uid"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

func Auth(r *ghttp.Request) {

	token := AuthHeaderToken(r)

	pass, err := service.Auth().VerifyToken(r.GetCtx(), token)
	if err != nil || !pass {
		r.Response.Header().Set("Content-Type", "application/json")
		r.Response.WriteStatus(http.StatusUnauthorized, g.Map{"code": 401, "message": "Unauthorized"})
		r.Exit()
		return
	}

	r.SetCtxVar(consts.SECRET_KEY, token)

	uid, err := gregex.ReplaceString("[a-zA-Z-]*", "", token)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError, g.Map{"code": 500, "message": "解析 sk 失败"})
		r.Exit()
		return
	}

	r.SetCtxVar(consts.UID_KEY, gconv.Int(uid))

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
