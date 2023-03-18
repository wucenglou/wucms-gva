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
	var api = v1.ApiGroupApp.PkgApiGroup.MyApi
	{
		MyRouterWithoutRecord.POST("create", api.CreateApi) // 新建TestStrutc
		// MyRouter.DELETE("deleteTestStrutc", testStrutcApi.DeleteTestStrutc)           // 删除TestStrutc
		// MyRouter.DELETE("deleteTestStrutcByIds", testStrutcApi.DeleteTestStrutcByIds) // 批量删除TestStrutc
		// MyRouter.PUT("updateTestStrutc", testStrutcApi.UpdateTestStrutc)              // 更新TestStrutc
	}
	{
		// MyRouterWithoutRecord.GET("findTestStrutc", testStrutcApi.FindTestStrutc)       // 根据ID获取TestStrutc
		// MyRouterWithoutRecord.GET("getTestStrutcList", testStrutcApi.GetTestStrutcList) // 获取TestStrutc列表
	}
}
