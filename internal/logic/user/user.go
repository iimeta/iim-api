package user

import (
	"context"
	"github.com/iimeta/iim-api/internal/dao"
	"github.com/iimeta/iim-api/internal/model"
	"github.com/iimeta/iim-api/internal/service"
	"github.com/iimeta/iim-api/utility/logger"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// 根据userId获取用户信息
func (s *sUser) GetUserById(ctx context.Context, userId int) (*model.User, error) {

	user, err := dao.User.FindUserByUserId(ctx, userId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.User{
		Id:        user.Id,
		UserId:    user.UserId,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Gender:    user.Gender,
		Mobile:    user.Mobile,
		Email:     user.Email,
		Birthday:  user.Birthday,
		Motto:     user.Motto,
		VipLevel:  user.VipLevel,
		SecretKey: user.SecretKey,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
