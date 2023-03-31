package pkg

import (
	v1 "wucms-gva/server/api/v1"
	"wucms-gva/server/middleware"

	"github.com/gin-gonic/gin"
)

type DoctorRouter struct {
}

// InitPkgRouter 初始化 Pkg 路由信息
func (s *DoctorRouter) InitDoctorRouter(Router *gin.RouterGroup) {
	pkgRouter := Router.Group("doctor").Use(middleware.OperationRecord())
	pkgRouterWithoutRecord := Router.Group("doctor")
	var doctorApi = v1.ApiGroupApp.PkgApiGroup.Doctor
	{
		pkgRouter.POST("createDoctor", doctorApi.CreateDoctor)             // 新建Pkg
		pkgRouter.DELETE("deleteDoctor", doctorApi.DeleteDoctor)           // 删除Pkg
		pkgRouter.DELETE("deleteDoctorByIds", doctorApi.DeleteDoctorByIds) // 批量删除Pkg
		pkgRouter.PUT("updateDoctor", doctorApi.UpdateDoctor)              // 更新Pkg
	}
	{
		pkgRouterWithoutRecord.GET("findDoctor", doctorApi.FindDoctor)       // 根据ID获取Pkg
		pkgRouterWithoutRecord.GET("getDoctorList", doctorApi.GetDoctorList) // 获取Pkg列表
	}
}
