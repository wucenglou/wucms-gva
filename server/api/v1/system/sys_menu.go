package system

import (
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/system"
	systemRes "wucms-gva/server/model/system/response"
	"wucms-gva/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct{}

func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	if menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		if menus == nil {
			menus = []system.SysMenu{}
		}
		response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}
