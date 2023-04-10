package pkg

import (
	"fmt"
	"regexp"
	"strconv"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/pkg"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Reg struct{}

func (r *Reg) CreateReg(c *gin.Context) {
	var Reg pkg.Reg
	err := c.ShouldBindJSON(&Reg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	matched, _ := regexp.MatchString("^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$", strconv.Itoa(Reg.Phone))
	if !matched {
		response.FailWithMessage("手机号码有误", c)
		return
	}
	err = global.GVA_DB.Create(&Reg).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{}, "创建成功", c)
}

func (r *Reg) DeleteReg(c *gin.Context) {
	var request request.GetById
	err := c.ShouldBind(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var Reg pkg.Reg
	err = global.GVA_DB.Where("id = ?", request.ID).Delete(&Reg).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (r *Reg) DeleteRegByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Delete(&pkg.Reg{}, "id in ?", IDS.Ids).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (r *Reg) UpdateReg(c *gin.Context) {
	var Reg pkg.Reg
	err := c.ShouldBindJSON(&Reg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Save(&Reg).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (r *Reg) FindReg(c *gin.Context) {
	var request request.GetById
	err := c.ShouldBindQuery(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("pppppppppppppppppppppppp")
	fmt.Println(request.ID)
	var reg pkg.Reg
	err = global.GVA_DB.Where("id = ?", request.ID).Preload("Doctor").First(&reg).Error
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"Reg": reg}, c)
	}
}

func (r *Reg) GetRegList(c *gin.Context) {
	var pageInfo pkg.RegSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&pkg.Reg{})
	var Regs []pkg.Reg
	var total int64

	if len(pageInfo.Keyword) > 0 {
		db = db.Where("name like ?", "%"+pageInfo.Keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if pageInfo.StartCreatedAt != nil && pageInfo.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", pageInfo.StartCreatedAt, pageInfo.EndCreatedAt)
	}
	err = db.Limit(limit).Offset(offset).Preload("Doctor").Preload("User").Order("updated_at desc").Find(&Regs).Error

	// err = global.GVA_DB.Preload("User").Preload("TermTaxonomy.Term").Find(&Regs).Error

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     Regs,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
	// response.OkWithDetailed(gin.H{"list": Regs}, "查询成功", c)
}
