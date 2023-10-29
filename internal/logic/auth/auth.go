package auth

import (
	"context"
	"github.com/iimeta/iim-api/internal/service"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{}
}

func (s *sAuth) Check(ctx context.Context) error {

	return nil
}
