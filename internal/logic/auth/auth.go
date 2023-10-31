package auth

import (
	"context"
	"github.com/iimeta/iim-api/internal/consts"
	"github.com/iimeta/iim-api/internal/errors"
	"github.com/iimeta/iim-api/internal/service"
	"github.com/iimeta/iim-api/utility/logger"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{}
}

func (s *sAuth) GetUid(ctx context.Context) int {

	uid := ctx.Value(consts.UID_KEY)
	if uid == nil {
		logger.Error(ctx, "uid is nil")
		return 0
	}

	return uid.(int)
}

func (s *sAuth) VerifyToken(ctx context.Context, token string) (bool, error) {

	if token == "" {
		return false, errors.New("token is nil")
	}

	return s.CheckUsage(ctx), nil
}

func (s *sAuth) GetToken(ctx context.Context) string {

	sk := ctx.Value(consts.SECRET_KEY)
	if sk == nil {
		logger.Error(ctx, "sk is nil")
		return ""
	}

	return sk.(string)
}

func (s *sAuth) CheckUsage(ctx context.Context) bool {

	usedTokens, err := service.Common().GetUsedTokens(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	totalTokens, err := service.Common().GetTotalTokens(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	return usedTokens < totalTokens
}
