package pkg

import (
	"fmt"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/pkg"
	"wucms-gva/server/utils"

	"github.com/gin-gonic/gin"
)

type MyApi struct{}

func (m *MyApi) GetPostList(c *gin.Context) {
	var pageInfo request.ModelPageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var term pkg.TermTaxonomy
	var total int64
	db := global.GVA_DB.Model(&pkg.TermTaxonomy{})
	db = db.Where("term_id = ?", pageInfo.TermId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Posts").Find(&term).Error
	if err != nil {
		return
	}
	// db := global.GVA_DB.Model(&pkg.Post{})
	// var posts []pkg.Post
	// var total int64

	// if len(pageInfo.Keyword) > 0 {
	// 	db = db.Where("post_title like ?", "%"+pageInfo.Keyword+"%")
	// }
	// err = db.Count(&total).Error
	// if err != nil {
	// 	return
	// }
	// err = db.Limit(limit).Offset(offset).Omit("post_content").Preload("User").Preload("TermTaxonomy.Term").Order("updated_at desc").Find(&posts).Error

	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	response.OkWithDetailed(response.PageResult{
		List:     term.Posts,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (m *MyApi) CreateApi(c *gin.Context) {
	var Term pkg.Term
	res, _ := utils.GetClaims(c)
	fmt.Println(res.ID)
	fmt.Println(res.Username)
	fmt.Println(res.NickName)
	err := global.GVA_DB.Model(&pkg.Term{}).Preload("TermMetas").Preload("TermTaxonomy.Posts").Find(&Term).Error
	if err != nil {
		return
	}
	fmt.Println(Term)
	for i := 0; i < len(Term.TermMetas); i++ {
		fmt.Println(Term.TermMetas[i].MetaKey)
		fmt.Println(Term.TermMetas[i].MetaValue)
		// for k, v := range Term.TermMeta[i] {
		// 	fmt.Println(k)
		// 	fmt.Println(v)
		// }
	}

	// var TermTaxonomy pkg.TermTaxonomy
	// err = global.GVA_DB.Model(&pkg.TermTaxonomy{}).Preload("Posts.PostMeta").Find(&TermTaxonomy).Error
	// if err != nil {
	// 	return
	// }
	// fmt.Println(TermTaxonomy)
	// var total int64
	fmt.Println("*********-***********")
	// db := global.GVA_DB.Model(&system.SysUser{})
	// err := db.Count(&total).Error
	// if err != nil {
	// 	return
	// }
	// var userList []system.SysUser
	// err = db.Preload("Authorities").Preload("Authority").Find(&userList).Error
	// if err != nil {
	// 	return
	// }
	// response.OkWithDetailed(response.PageResult{
	// 	List: userList,
	// }, "获取成功", c)
	response.Ok(c)
}
