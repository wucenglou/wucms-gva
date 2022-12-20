package system

import (
	"github.com/gin-gonic/gin"
)

type AuthorityMenuApi struct{}

func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	print("+++++++++++++++++++")
	// if menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c)); err != nil {
	// 	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	// 	response.FailWithMessage("获取失败", c)
	// } else {
	// 	if menus == nil {
	// 		menus = []system.SysMenu{}
	// 	}
	// 	response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
	// }
}
