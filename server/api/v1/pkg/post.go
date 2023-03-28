package pkg

import (
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/pkg"
	req "wucms-gva/server/model/pkg/request"
	"wucms-gva/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Post struct{}

func (post *Post) CreatePost(c *gin.Context) {
	var reqPost req.ReqPost
	err := c.ShouldBindJSON(&reqPost)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, _ := utils.GetUser(c)
	var Tem pkg.Post
	Tem.PostAuthor = user.ID
	reqPost.Post.PostAuthor = Tem.PostAuthor
	err = global.GVA_DB.Create(&reqPost.Post).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var termTaxonomy []pkg.TermTaxonomy
	for _, v := range reqPost.TermIds {
		t := pkg.TermTaxonomy{TermTaxonomyId: uint(v)}
		termTaxonomy = append(termTaxonomy, t)
	}
	global.GVA_DB.Model(&reqPost.Post).Association("TermTaxonomy").Append(&termTaxonomy)
	// var Tem pkg.Post
	// user, _ := utils.GetUser(c)
	// // fmt.Println("***************")
	// // fmt.Println(user.ID)
	// Tem.PostAuthor = user.ID
	// Post.PostAuthor = Tem.PostAuthor
	// // fmt.Println(Post)
	// err = global.GVA_DB.Create(&Post).Association("TermTaxonomy").Error
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
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

func (post *Post) GetPostList(c *gin.Context) {
	var posts []pkg.Post
	err := global.GVA_DB.Find(&posts).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"list": posts}, "查询成功", c)
}
