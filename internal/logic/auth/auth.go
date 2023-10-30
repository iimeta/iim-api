package auth

import (
	"context"
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

	uid := ctx.Value("uid")
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

	return true, nil
}
