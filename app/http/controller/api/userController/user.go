package userController

import (
	"develop-template/app/entity/req/userReqEntity"
	"develop-template/app/error/common"
	baseController "develop-template/app/http/controller"
	"develop-template/app/logic/userLogic"
	"github.com/gin-gonic/gin"
)

type controller struct {
	baseController.BaseController
}

func NewController(ctx *gin.Context) *controller {
	return &controller{baseController.NewBaseController(ctx)}
}
func (c *controller) UserList() {
	req := userReqEntity.List{}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Fail(common.ReqParamErr, c.GetValidMsg(err, &req))
		return
	}

	logic := userLogic.NewLogic(c.GCtx)
	resp, err := logic.GetUserList(req)
	if err != nil {
		c.Fail(err.Code(), err.Message())
		return
	}
	c.Success(resp)
}
