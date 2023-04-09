package pkg

import (
	"fmt"
	"regexp"
	"strconv"
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
	err = db.Limit(limit).Offset(offset).Preload("Posts.User").Find(&term).Error
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

// 患者管理
func (m *MyApi) CreatePatient(c *gin.Context) {
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
	matchedName, _ := regexp.MatchString("^[\u4E00-\u9FA5|A-Z|a-z].+$", Patient.Name)
	if !matchedName {
		response.FailWithMessage("姓名不能含特殊字符", c)
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

func (m *MyApi) DeletePatient(c *gin.Context) {
	var request request.GetById
	err := c.ShouldBind(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, _ := utils.GetUser(c)
	var Patient pkg.Patient
	db := global.GVA_DB.Where("id = ?", request.ID)
	err = db.Find(&Patient).Error
	if Patient.UserId == user.ID {
		err = db.Delete(&Patient).Error
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	} else {
		response.OkWithMessage("没有权限", c)
	}
	// err = global.GVA_DB.Where("id = ?", request.ID).Delete(&Patient).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (m *MyApi) GetRegList(c *gin.Context) {
	var pageInfo pkg.RegSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, _ := utils.GetUser(c)
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&pkg.Reg{}).Where("user_id =?", user.ID)
	var Regs []pkg.Reg
	var total int64
	if len(pageInfo.Keyword) > 0 {
		db = db.Where("status = ?", pageInfo.Keyword)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&Regs).Error

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

func (m *MyApi) CreateReg(c *gin.Context) {
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
	user, _ := utils.GetUser(c)
	ip := c.ClientIP()
	ipInfo, _ := utils.GetIpInfo("113.121.36.199")
	fmt.Println(ipInfo)
	Reg.UserId = user.ID
	Reg.Ip = ip
	Reg.IpDesc = ipInfo
	err = global.GVA_DB.Create(&Reg).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{}, "创建成功", c)
}
