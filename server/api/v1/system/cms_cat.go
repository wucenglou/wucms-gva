package system

import (
	"fmt"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/pkg"

	"github.com/gin-gonic/gin"
)

type CmsCatApi struct{}

func (catApi *CmsCatApi) GetCmsCat(c *gin.Context) {
	var Term []pkg.Term
	err := global.GVA_DB.Model(&pkg.Term{}).Preload("TermMetas").Preload("TermTaxonomy").Find(&Term).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("********************")

	// var pageInfo request.PageInfo
	// err := c.ShouldBindJSON(&pageInfo)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	var TermTaxonomy []*pkg.TermTaxonomy
	err = global.GVA_DB.Model(&pkg.TermTaxonomy{}).Preload("Term").Find(&TermTaxonomy).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cat := catApi.BuildCmsCatTree(TermTaxonomy, 0)
	response.OkWithDetailed(gin.H{"list": cat, "Term": Term}, "查询成功", c)
}

func (catApi *CmsCatApi) BuildCmsCatTree(items []*pkg.TermTaxonomy, parentId int) []*pkg.TermTaxonomy {
	result := []*pkg.TermTaxonomy{}
	for _, item := range items {
		if item.ParentID == parentId {
			// itemitem.Term.Name
			item.Children = catApi.BuildCmsCatTree(items, int(item.TermTaxonomyId))
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
	response.OkWithMessage("创建成功", c)
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

// CreateTermStruct 创建TermStruct
// @Tags TermStruct
// @Summary 创建TermStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MyPkg.TermStruct true "创建TermStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /term/createTermStruct [post]
func (catApi *CmsCatApi) CreateCat(c *gin.Context) {
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
	response.OkWithMessage("创建成功", c)
	// if err := termService.CreateTermStruct(term); err != nil {
	// 	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	// 	response.FailWithMessage("创建失败", c)
	// } else {
	// 	response.OkWithMessage("创建成功", c)
	// }
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
// func (catApi *CmsCatApi) DeleteCat(c *gin.Context) {
// 	var term MyPkg.TermStruct
// 	err := c.ShouldBindJSON(&term)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err := termService.DeleteTermStruct(term); err != nil {
//         global.GVA_LOG.Error("删除失败!", zap.Error(err))
// 		response.FailWithMessage("删除失败", c)
// 	} else {
// 		response.OkWithMessage("删除成功", c)
// 	}
// }

// DeleteTermStructByIds 批量删除TermStruct
// @Tags TermStruct
// @Summary 批量删除TermStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TermStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /term/deleteTermStructByIds [delete]
// func (catApi *CmsCatApi) DeleteCatByIds(c *gin.Context) {
// 	var IDS request.IdsReq
//     err := c.ShouldBindJSON(&IDS)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err := termService.DeleteTermStructByIds(IDS); err != nil {
//         global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
// 		response.FailWithMessage("批量删除失败", c)
// 	} else {
// 		response.OkWithMessage("批量删除成功", c)
// 	}
// }

// UpdateTermStruct 更新TermStruct
// @Tags TermStruct
// @Summary 更新TermStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MyPkg.TermStruct true "更新TermStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /term/updateTermStruct [put]
// func (catApi *CmsCatApi) UpdateCat(c *gin.Context) {
// 	var term MyPkg.TermStruct
// 	err := c.ShouldBindJSON(&term)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err := termService.UpdateTermStruct(term); err != nil {
//         global.GVA_LOG.Error("更新失败!", zap.Error(err))
// 		response.FailWithMessage("更新失败", c)
// 	} else {
// 		response.OkWithMessage("更新成功", c)
// 	}
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
// func (catApi *CmsCatApi) FindCat(c *gin.Context) {
// 	var term MyPkg.Term
// 	err := c.ShouldBindQuery(&term)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// if reterm, err := termService.GetTermStruct(term.ID); err != nil {
//     global.GVA_LOG.Error("查询失败!", zap.Error(err))
// 	response.FailWithMessage("查询失败", c)
// } else {
// 	response.OkWithData(gin.H{"reterm": reterm}, c)
// }
// }

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
