package system

import (
	v1 "wucms-gva/server/api/v1"
	"wucms-gva/server/middleware"

	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		sysRouter.POST("getSystemConfig", systemApi.GetSystemConfig)
		sysRouter.POST("setSystemConfig", systemApi.SetSystemConfig)
		sysRouter.POST("getServerInfo", systemApi.GetServerInfo)
	}
}
