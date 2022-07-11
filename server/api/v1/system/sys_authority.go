package system

import (
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityApi struct{}

// @Tags SysUser
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取用户信息"
// @Router /user/getUserInfo [get]

func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	if ReqUser, err := userService.GetUserInfo(uuid); err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
	}
}
