package pkg

import (
	v1 "wucms-gva/server/api/v1"

	"github.com/gin-gonic/gin"
)

type MyApi struct {
}

// InitTestStrutcRouter 初始化 TestStrutc 路由信息
func (s *MyApi) InitMyApiRouter(Router *gin.RouterGroup) {
	// MyRouter := Router.Group("myapi").Use(middleware.OperationRecord())
	MyRouterWithoutRecord := Router.Group("myapi")

	doctorApi := v1.ApiGroupApp.PkgApiGroup.Doctor
	patientApi := v1.ApiGroupApp.PkgApiGroup.Patient
	RegApi := v1.ApiGroupApp.PkgApiGroup.Reg
	Post := v1.ApiGroupApp.PkgApiGroup.Post
	CmsCatApi := v1.ApiGroupApp.SystemApiGroup.CmsCatApi
	var api = v1.ApiGroupApp.PkgApiGroup.MyApi
	{

		// MyRouter.DELETE("deleteTestStrutc", testStrutcApi.DeleteTestStrutc)           // 删除TestStrutc
		// MyRouter.DELETE("deleteTestStrutcByIds", testStrutcApi.DeleteTestStrutcByIds) // 批量删除TestStrutc
		// MyRouter.PUT("updateTestStrutc", testStrutcApi.UpdateTestStrutc)              // 更新TestStrutc
	}
	{

		MyRouterWithoutRecord.DELETE("deletePatient", api.DeletePatient)
		MyRouterWithoutRecord.PUT("updatePatient", patientApi.UpdatePatient)

		MyRouterWithoutRecord.POST("create", api.CreateApi)                 // 新建TestStrutc
		MyRouterWithoutRecord.GET("findDoctor", doctorApi.FindDoctor)       // 根据ID获取Pkg
		MyRouterWithoutRecord.GET("getDoctorList", doctorApi.GetDoctorList) // 获取Pkg列表
		// 患者管理
		MyRouterWithoutRecord.GET("findPatient", patientApi.FindPatient) // 根据ID获取Pkg
		MyRouterWithoutRecord.GET("getPatientList", api.GetPatientList)  // 获取Pkg列表
		MyRouterWithoutRecord.POST("createPatient", api.CreatePatient)   // 新建Pkg

		MyRouterWithoutRecord.GET("findReg", RegApi.FindReg)    // 根据ID获取Pkg
		MyRouterWithoutRecord.GET("getRegList", api.GetRegList) // 获取Pkg列表
		MyRouterWithoutRecord.POST("createReg", api.CreateReg)

		MyRouterWithoutRecord.GET("post/ById", Post.FindPost)
		MyRouterWithoutRecord.GET("post", api.GetPostList)
		MyRouterWithoutRecord.PUT("", CmsCatApi.UpdateCmsCat)
		MyRouterWithoutRecord.GET("cat/ById", CmsCatApi.FindCmsCat)
		MyRouterWithoutRecord.GET("cat", CmsCatApi.GetCmsCat)
		MyRouterWithoutRecord.POST("cattest", CmsCatApi.CatTest) // 获取所有api
		MyRouterWithoutRecord.POST("", CmsCatApi.CreateCmsCat)

		MyRouterWithoutRecord.GET("user", api.GetUser)
		MyRouterWithoutRecord.POST("user", api.CreateUser)

		//登录,用户管理
		MyRouterWithoutRecord.POST("wxlogin", api.WxLogin)
		MyRouterWithoutRecord.PUT("setUserInfo", api.SetUserInfo)
		// MyRouterWithoutRecord.GET("findTestStrutc", testStrutcApi.FindTestStrutc)       // 根据ID获取TestStrutc
		// MyRouterWithoutRecord.GET("getTestStrutcList", testStrutcApi.GetTestStrutcList) // 获取TestStrutc列表
	}
}
