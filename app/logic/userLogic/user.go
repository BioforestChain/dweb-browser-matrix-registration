package userLogic

import (
	"context"
	"develop-template/app/entity/db/userDbEntity"
	"develop-template/app/entity/req/userReqEntity"
	"develop-template/app/entity/resp/userRespEntity"
	myError "develop-template/app/error"
	commonLogic "develop-template/app/logic"
	"develop-template/app/model/develop-template/userModel"
	"fmt"
	"github.com/gin-gonic/gin"
)

type logic struct {
	Ctx  context.Context
	GCtx *gin.Context
}

func NewLogic(ctx *gin.Context) *logic {
	return &logic{Ctx: ctx, GCtx: ctx}
}

func (l *logic) GetUserList(req userReqEntity.List) (res userRespEntity.List, err myError.Error) {
	defer func() {
		r := recover()
		fmt.Println("============panic GetUserList============", r)
	}()
	cond := userDbEntity.Condition{}

	if req.Page >= 0 {
		cond.Page = req.Page
	}

	if req.Limit >= 0 {
		cond.Limit = req.Limit
	}

	cond.Page, cond.Limit, cond.Offset = commonLogic.InitCondition(cond.Page, cond.Limit)

	list, total, err := userModel.NewModel(l.GCtx).GetUserList(cond)
	if err != nil {
		return res, err
	}

	res.Total = total
	res.Page = cond.Page
	res.List = list
	res.LastPage = commonLogic.GetLastPage(total, cond.Limit)
	return res, nil
}
