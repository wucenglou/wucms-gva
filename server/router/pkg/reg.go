package pkg

import (
	v1 "wucms-gva/server/api/v1"
	"wucms-gva/server/middleware"

	"github.com/gin-gonic/gin"
)

type RegRouter struct {
}

// InitPkgRouter 初始化 Pkg 路由信息
func (s *RegRouter) InitRegRouter(Router *gin.RouterGroup) {
	pkgRouter := Router.Group("reg").Use(middleware.OperationRecord())
	pkgRouterWithoutRecord := Router.Group("reg")
	var RegApi = v1.ApiGroupApp.PkgApiGroup.Reg
	{
		pkgRouter.POST("createReg", RegApi.CreateReg)             // 新建Pkg
		pkgRouter.DELETE("deleteReg", RegApi.DeleteReg)           // 删除Pkg
		pkgRouter.DELETE("deleteRegByIds", RegApi.DeleteRegByIds) // 批量删除Pkg
		pkgRouter.PUT("updateReg", RegApi.UpdateReg)              // 更新Pkg
	}
	{
		pkgRouterWithoutRecord.GET("findReg", RegApi.FindReg)       // 根据ID获取Pkg
		pkgRouterWithoutRecord.GET("getRegList", RegApi.GetRegList) // 获取Pkg列表
	}
}
