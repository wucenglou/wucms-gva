package pkg

import (
	v1 "wucms-gva/server/api/v1"

	"github.com/gin-gonic/gin"
)

type Post struct{}

func (s *Post) InitPostRouter(Router *gin.RouterGroup) {
	// apiRouter := Router.Group("cat").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("post")
	Post := v1.ApiGroupApp.PkgApiGroup.Post
	{
		// apiRouter.POST("createApi", apiRouterApi.CreateApi)               // 创建Api
		// apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除Api
		// apiRouter.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条Api消息
		// apiRouter.POST("updateApi", apiRouterApi.UpdateApi)               // 更新api
		// apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouterWithoutRecord.DELETE("", Post.DeletePost)
		apiRouterWithoutRecord.DELETE("ByIds", Post.DeletePostByIds)
		apiRouterWithoutRecord.PUT("", Post.UpdatePost)
		apiRouterWithoutRecord.GET("ById", Post.FindPost)
		apiRouterWithoutRecord.GET("", Post.GetPostList)
		// apiRouterWithoutRecord.POST("cattest", Post.CatTest)
		apiRouterWithoutRecord.POST("", Post.CreatePost)
		// apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // 获取Api列表
	}
}
