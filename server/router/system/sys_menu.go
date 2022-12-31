package system

import (
	v1 "wucms-gva/server/api/v1"
	"wucms-gva/server/middleware"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("updateBaseMenu", authorityMenuApi.UpdateBaseMenu) // 更新菜单
	}
	{
		menuRouterWithoutRecord.POST("getMenuList", authorityMenuApi.GetMenuList)
		menuRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu)
		menuRouterWithoutRecord.POST("getBaseMenuById", authorityMenuApi.GetBaseMenuById) // 根据id获取菜单
	}
	return menuRouter
}
