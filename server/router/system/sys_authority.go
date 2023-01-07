package system

import (
	v1 "wucms-gva/server/api/v1"
	"wucms-gva/server/middleware"

	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority) // 创建角色
		authorityRouter.PUT("updateAuthority", authorityApi.UpdateAuthority)  // 更新角色
	}
	{
		authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}
}
