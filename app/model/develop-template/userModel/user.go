package userModel

import (
	"context"
	"develop-template/app/entity/db/userDbEntity"
	myError "develop-template/app/error"
	"develop-template/app/model/common"
	"gorm.io/gorm"
)

type model struct {
	Ctx context.Context
	DB  *gorm.DB
}

func NewModel(ctx context.Context) *model {
	return &model{ctx, common.ConnectionObject(ctx).DB}
}

func (m *model) GetUserList(cond userDbEntity.Condition) (list []*userDbEntity.UserInfo, total int64, err myError.Error) {
	userModule := userDbEntity.UserInfo{}
	db := m.DB.Model(&userModule)

	if err := db.Count(&total).Error; err != nil {
		return list, total, common.CheckMysqlError(err)
	}

	if err := db.Limit(cond.Limit).Offset(cond.Offset).Order("id desc").Find(&list).Error; err != nil {
		return list, total, common.CheckMysqlError(err)
	}
	return list, total, nil
}
