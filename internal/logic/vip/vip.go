package vip

import (
	"context"
	"github.com/iimeta/iim-api/internal/service"
)

type sVip struct{}

func init() {
	service.RegisterVip(New())
}

func New() service.IVip {
	return &sVip{}
}

func (s *sVip) CheckUserVipModelPermission(ctx context.Context, userId int, model string) bool {

	return false
}
