package pkg

import (
	"fmt"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/pkg"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Post struct{}

func (post *Post) CreatePost(c *gin.Context) {
	var Post pkg.Post
	err := c.ShouldBindJSON(&Post)
	fmt.Println("***************------------")
	fmt.Println(Post)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Create(&Post).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("***************------------")
	// fmt.Println(Post)
	response.OkWithDetailed(gin.H{}, "创建成功", c)
}

func (post *Post) FindPost(c *gin.Context) {
	var request request.GetById
	err := c.ShouldBindQuery(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var Post pkg.Post
	err = global.GVA_DB.Where("id = ?", request.ID).First(&Post).Error
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"post": Post}, c)
	}
}
