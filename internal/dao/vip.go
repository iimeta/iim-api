package dao

import (
	"github.com/iimeta/iim-api/internal/model/entity"
	"github.com/iimeta/iim-api/utility/db"
)

var Vip = NewVipDao()

type VipDao struct {
	*MongoDB[entity.Vip]
}

func NewVipDao(database ...string) *VipDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &VipDao{
		MongoDB: NewMongoDB[entity.Vip](database[0], entity.VIP_COLLECTION),
	}
}
