package system

import (
	"fmt"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/pkg"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CmsCatApi struct{}

// func (catApi *CmsCatApi) getModel(c *gin.Context) {
// 	var TermTaxonomy []*pkg.TermTaxonomy
// 	err := global.GVA_DB.Model(&pkg.TermTaxonomy{}).Where("taxonomy = ?", "model").Preload("Term").Find(&TermTaxonomy).Error
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	termTaxonomy := catApi.BuildCmsCatTree(TermTaxonomy, 0)
// 	response.OkWithDetailed(gin.H{"list": termTaxonomy}, "查询成功", c)
// }

func (catApi *CmsCatApi) GetCmsCat(c *gin.Context) {
	var pageInfo request.ModelPageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("9999999999999")
	fmt.Println(pageInfo)
	var Term []pkg.Term
	err = global.GVA_DB.Model(&pkg.Term{}).Preload("TermMetas").Preload("TermTaxonomy", "taxonomy = ?", pageInfo.Model).Find(&Term).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var TermTaxonomy []*pkg.TermTaxonomy
	// err = global.GVA_DB.Model(&pkg.TermTaxonomy{}).Not("parent_id = ? AND taxonomy <> ?", 0, pageInfo.Model).Preload("Term").Find(&TermTaxonomy).Error
	err = global.GVA_DB.Model(&pkg.TermTaxonomy{}).Where("taxonomy = ?", pageInfo.Model).Preload("Term").Find(&TermTaxonomy).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cat := catApi.BuildCmsCatTree(TermTaxonomy, 0, 0, 0)
	response.OkWithDetailed(gin.H{"list": cat, "Term": Term}, "查询成功", c)
}

func (catApi *CmsCatApi) BuildCmsCatTree(items []*pkg.TermTaxonomy, parentId int, level int, endlevel int) []*pkg.TermTaxonomy {
	result := []*pkg.TermTaxonomy{}
	for _, item := range items {
		if item.ParentID == parentId {
			item.Children = catApi.BuildCmsCatTree(items, int(item.TermTaxonomyId), level, endlevel)
			result = append(result, item)
		}
	}
	return result
}

func (catApi *CmsCatApi) CreateCmsCat(c *gin.Context) {
	fmt.Println("---------c1111111at+++++++++++")
	var Term pkg.Term
	fmt.Println("---------cat+++++++++++")
	err := c.ShouldBindJSON(&Term)
	fmt.Println(Term)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Create(&Term).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("***************------------")
	fmt.Println(Term)
	response.OkWithDetailed(gin.H{}, "创建成功", c)
}

func (catApi *CmsCatApi) CatTest(c *gin.Context) {
	// var term MyPkg.Term
	// err := c.ShouldBindQuery(&term)
	// fmt.Println(term)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	response.OkWithMessage("创建成功", c)
}

// DeleteTermStruct 删除TermStruct
// @Tags TermStruct
// @Summary 删除TermStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MyPkg.TermStruct true "删除TermStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /term/deleteTermStruct [delete]
func (catApi *CmsCatApi) DeleteCmsCat(c *gin.Context) {
	var request request.GetById
	err := c.ShouldBind(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var termTaxonomy pkg.TermTaxonomy
	err = global.GVA_DB.Where("parent_id = ?", request.ID).First(&termTaxonomy).Error
	if err == nil {
		response.FailWithMessage("先删除子目录", c)
		return
	}
	err = global.GVA_DB.Select(clause.Associations).Delete(&pkg.Term{TermId: uint(request.ID)}).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteTermStructByIds 批量删除TermStruct
// @Tags TermStruct
// @Summary 批量删除TermStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TermStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /term/deleteTermStructByIds [delete]
func (catApi *CmsCatApi) DeleteCmsCatByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Delete(&pkg.Term{}, "term_id in ?", IDS.Ids).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Delete(&pkg.TermTaxonomy{}, "term_id in ?", IDS.Ids).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Delete(&pkg.TermMeta{}, "term_id in ?", IDS.Ids).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateTermStruct 更新TermStruct
// @Tags TermStruct
// @Summary 更新TermStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MyPkg.TermStruct true "更新TermStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /term/updateTermStruct [put]
func (catApi *CmsCatApi) UpdateCmsCat(c *gin.Context) {
	var term pkg.Term
	err := c.ShouldBindJSON(&term)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&term).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("+++++++++++++++++++++")
	fmt.Println(term)
	response.OkWithMessage("更新成功", c)
}

// type test struct {
// 	TermID int `json:"term_id" form:"term_id"` // 主键ID
// 	TermSD int `json:"term_sd" form:"term_sd"` // 主键ID
// }

// FindTermStruct 用id查询TermStruct
// @Tags TermStruct
// @Summary 用id查询TermStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MyPkg.TermStruct true "用id查询TermStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /term/findTermStruct [get]
func (catApi *CmsCatApi) FindCmsCat(c *gin.Context) {
	var request request.GetById
	err := c.ShouldBindQuery(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var term pkg.Term
	err = global.GVA_DB.Where("term_id = ?", request.ID).Preload("TermMetas").Preload("TermTaxonomy").First(&term).Error
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"term": term}, c)
	}
}

// GetTermStructList 分页获取TermStruct列表
// @Tags TermStruct
// @Summary 分页获取TermStruct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MyPkgReq.TermStructSearch true "分页获取TermStruct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /term/getTermStructList [get]
// func (catApi *CmsCatApi) GetCatList(c *gin.Context) {
// 	var pageInfo MyPkgReq.TermStructSearch
// 	err := c.ShouldBindQuery(&pageInfo)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if list, total, err := termService.GetTermStructInfoList(pageInfo); err != nil {
// 	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
//         response.FailWithMessage("获取失败", c)
//     } else {
//         response.OkWithDetailed(response.PageResult{
//             List:     list,
//             Total:    total,
//             Page:     pageInfo.Page,
//             PageSize: pageInfo.PageSize,
//         }, "获取成功", c)
//     }
// }
