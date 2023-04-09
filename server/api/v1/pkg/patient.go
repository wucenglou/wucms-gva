package pkg

import (
	"regexp"
	"strconv"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/pkg"
	"wucms-gva/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Patient struct{}

func (P *Patient) CreatePatient(c *gin.Context) {
	var Patient pkg.Patient
	err := c.ShouldBindJSON(&Patient)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	matched, _ := regexp.MatchString("^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$", strconv.Itoa(Patient.Phone))
	if !matched {
		response.FailWithMessage("手机号码有误", c)
		return
	}
	user, _ := utils.GetUser(c)
	ip := c.ClientIP()

	// var tmp pkg.Patient
	// tmp.UserId = user.ID
	Patient.UserId = user.ID
	Patient.Ip = ip

	err = global.GVA_DB.Create(&Patient).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{}, "创建成功", c)
}

func (P *Patient) DeletePatient(c *gin.Context) {
	var request request.GetById
	err := c.ShouldBind(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var Patient pkg.Patient
	err = global.GVA_DB.Where("id = ?", request.ID).Delete(&Patient).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (P *Patient) DeletePatientByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Delete(&pkg.Patient{}, "id in ?", IDS.Ids).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (P *Patient) UpdatePatient(c *gin.Context) {
	var Patient pkg.Patient
	err := c.ShouldBindJSON(&Patient)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Patient).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (P *Patient) FindPatient(c *gin.Context) {
	var request request.GetById
	err := c.ShouldBindQuery(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var Patient pkg.Patient
	err = global.GVA_DB.Where("id = ?", request.ID).First(&Patient).Error
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"Patient": Patient}, c)
	}
}

func (P *Patient) GetPatientList(c *gin.Context) {
	var pageInfo pkg.PatientSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&pkg.Patient{})
	var Patients []pkg.Patient
	var total int64

	if len(pageInfo.Keyword) > 0 {
		db = db.Where("name like ?", "%"+pageInfo.Keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Preload("User").Find(&Patients).Error

	// err = global.GVA_DB.Preload("User").Preload("TermTaxonomy.Term").Find(&Patients).Error

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     Patients,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
	// response.OkWithDetailed(gin.H{"list": Patients}, "查询成功", c)
}
