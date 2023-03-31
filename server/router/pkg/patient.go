package pkg

import (
	v1 "wucms-gva/server/api/v1"
	"wucms-gva/server/middleware"

	"github.com/gin-gonic/gin"
)

type PatientRouter struct {
}

// InitPkgRouter 初始化 Pkg 路由信息
func (s *PatientRouter) InitPatientRouter(Router *gin.RouterGroup) {
	pkgRouter := Router.Group("patient").Use(middleware.OperationRecord())
	pkgRouterWithoutRecord := Router.Group("patient")
	var patientApi = v1.ApiGroupApp.PkgApiGroup.Patient
	{
		pkgRouter.POST("createPatient", patientApi.CreatePatient)             // 新建Pkg
		pkgRouter.DELETE("deletePatient", patientApi.DeletePatient)           // 删除Pkg
		pkgRouter.DELETE("deletePatientByIds", patientApi.DeletePatientByIds) // 批量删除Pkg
		pkgRouter.PUT("updatePatient", patientApi.UpdatePatient)              // 更新Pkg
	}
	{
		pkgRouterWithoutRecord.GET("findPatient", patientApi.FindPatient)       // 根据ID获取Pkg
		pkgRouterWithoutRecord.GET("getPatientList", patientApi.GetPatientList) // 获取Pkg列表
	}
}
