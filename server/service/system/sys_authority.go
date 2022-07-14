package system

import (
	"errors"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/system"

	"gorm.io/gorm"
)

var ErrRoleExistence = errors.New("存在相同角色id")

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}
	err = global.GVA_DB.Create(&auth).Error
	return auth, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysAuthority{})
	err = db.Where("parent_id = ?", "0").Count(&total).Error
	var authority []system.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = ?", "0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = authorityService.findChildrenAuthority(&authority[k])
		}
	}
	return authority, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.SysAuthority
//@return: error

func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority) error {
	var s system.SysAuthority
	global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.SysAuthority
//@return: err error

func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
