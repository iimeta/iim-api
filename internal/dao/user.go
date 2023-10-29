package dao

import (
	"context"
	"github.com/iimeta/iim-api/internal/model/do"
	"github.com/iimeta/iim-api/internal/model/entity"
	"github.com/iimeta/iim-api/utility/db"
	"go.mongodb.org/mongo-driver/bson"
)

var User = NewUserDao()

type UserDao struct {
	*MongoDB[entity.User]
}

func NewUserDao(database ...string) *UserDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &UserDao{
		MongoDB: NewMongoDB[entity.User](database[0], do.USER_COLLECTION),
	}
}

// 根据userId查询用户
func (d *UserDao) FindUserByUserId(ctx context.Context, userId int) (*entity.User, error) {
	return d.FindOne(ctx, bson.M{"user_id": userId})
}

// 根据userIds查询用户列表
func (d *UserDao) FindUserListByUserIds(ctx context.Context, userIds []int) ([]*entity.User, error) {
	return d.Find(ctx, bson.M{"user_id": bson.M{"$in": userIds}})
}
