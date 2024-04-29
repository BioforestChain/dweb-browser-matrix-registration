package userRespEntity

import (
	"develop-template/app/entity"
	"develop-template/app/entity/db/userDbEntity"
)

// List 列表
type List struct {
	entity.Pagination
	List []*userDbEntity.UserInfo `json:"list"`
}

type OnlineList struct {
	//List []userDbEntity.TestData `json:"list"`
	List []*userDbEntity.UserInfo `json:"list"`
}

// Info 信息
type Info struct {
	userDbEntity.UserInfo
}
