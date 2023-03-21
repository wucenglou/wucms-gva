package system

import (
	v1 "wucms-gva/server/api/v1"

	"github.com/gin-gonic/gin"
)

type CmsCatRouter struct{}

func (s *CmsCatRouter) InitCmsCatRouter(Router *gin.RouterGroup) {
	// apiRouter := Router.Group("cat").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("cat")
	CmsCatApi := v1.ApiGroupApp.SystemApiGroup.CmsCatApi
	{
		// apiRouter.POST("createApi", apiRouterApi.CreateApi)               // 创建Api
		// apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除Api
		// apiRouter.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条Api消息
		// apiRouter.POST("updateApi", apiRouterApi.UpdateApi)               // 更新api
		// apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouterWithoutRecord.GET("", CmsCatApi.GetCmsCat)
		apiRouterWithoutRecord.POST("cattest", CmsCatApi.CatTest)  // 获取所有api
		apiRouterWithoutRecord.POST("create", CmsCatApi.CreateCat) // 获取所有api
		// apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // 获取Api列表
	}
}
