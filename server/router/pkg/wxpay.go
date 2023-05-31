package pkg

import (
	v1 "wucms-gva/server/api/v1"

	"github.com/gin-gonic/gin"
)

type WxPay struct {
}

// InitPkgRouter 初始化 Pkg 路由信息
func (s *WxPay) InitWxPay(Router *gin.RouterGroup) {
	// pkgRouter := Router.Group("wxpay").Use(middleware.OperationRecord())
	pkgRouterWithoutRecord := Router.Group("wxpay")
	var wxPayApi = v1.ApiGroupApp.PkgApiGroup.WxPay
	{
		// pkgRouter.POST("createDoctor", doctorApi.CreateDoctor)             // 新建Pkg
		// pkgRouter.DELETE("deleteDoctor", doctorApi.DeleteDoctor)           // 删除Pkg
		// pkgRouter.DELETE("deleteDoctorByIds", doctorApi.DeleteDoctorByIds) // 批量删除Pkg
		// pkgRouter.PUT("updateDoctor", doctorApi.UpdateDoctor)              // 更新Pkg
	}
	{
		pkgRouterWithoutRecord.GET("findDoctor", wxPayApi.FindDoctor) // 根据ID获取Pkg
		// pkgRouterWithoutRecord.GET("getDoctorList", wxPayApi.GetDoctorList) // 获取Pkg列表
	}
}
