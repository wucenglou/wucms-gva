package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/pkg"
	"wucms-gva/server/model/system"
	systemRes "wucms-gva/server/model/system/response"
	"wucms-gva/server/utils"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

func (m *MyApi) GetPatientList(c *gin.Context) {
	var pageInfo pkg.PatientSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	user, _ := utils.GetUser(c)
	// 创建db
	db := global.GVA_DB.Model(&pkg.Patient{}).Where("user_id = ?", user.ID)
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
	var tmp pkg.Reg
	err = global.GVA_DB.Where("user_id = ?", user.ID).Where("time = ?", Reg.Time).First(&tmp).Error
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("已经挂号成功", c)
		return
	}
	err = global.GVA_DB.Create(&Reg).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{}, "创建成功", c)
}

func (m *MyApi) CreateUser(c *gin.Context) {}

func (m *MyApi) GetUser(c *gin.Context) {
	userInfo, _ := utils.GetUser(c)
	var user system.SysUser
	global.GVA_DB.First(&user, userInfo.ID)
	response.OkWithDetailed(gin.H{"user": user}, "查询成功", c)
}

type Wxinfo struct {
	Code string
}

func (m *MyApi) WxLogin(c *gin.Context) {
	var wxcode Wxinfo
	err := c.ShouldBindJSON(&wxcode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	resp, _ := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=wxd178a3abd041537c&secret=3ba16f2f82174472c92fbe32e4868d43&js_code=" + wxcode.Code + "&grant_type=authorization_code")
	closer := resp.Body
	bytes, _ := ioutil.ReadAll(closer)
	info := make(map[string]interface{})
	err = json.Unmarshal(bytes, &info)
	if err != nil {
		return
	}
	if info["openid"] == nil {
		response.FailWithMessage(err.Error(), c)
	}
	var user system.SysUser
	err = global.GVA_DB.Where("open_id = ?", info["openid"]).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 新建用户
		var newUser system.SysUser
		newUser.NickName = "微信用户"
		newUser.AuthorityId = 0
		newUser.UUID = uuid.NewV4()
		newUser.OpenId = info["openid"].(string)
		err = global.GVA_DB.Create(&newUser).Error
		if err != nil {
			return
		}
		m.TokenNext(c, newUser)
	} else {
		// 查询用户
		m.TokenNext(c, user)
	}

	// response.OkWithDetailed(gin.H{"user": info}, "查询成功", c)
}

// 登录以后签发jwt
func (m *MyApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(utils.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}
