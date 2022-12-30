package system

import (
	v1 "wucms-gva/server/api/v1"
	"wucms-gva/server/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register)
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList)
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo) // 获取自身信息
	}
}
