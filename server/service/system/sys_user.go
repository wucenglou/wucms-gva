package system

import (
	"errors"
	"fmt"
	"wucms-gva/server/global"
	"wucms-gva/server/model/system"
	"wucms-gva/server/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter system.SysUser, err error

type UserService struct{}

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		var SysAuthorityMenus []system.SysAuthorityMenu
		err = global.GVA_DB.Where("sys_authority_authority_id = ?", user.AuthorityId).Find(&SysAuthorityMenus).Error
		if err != nil {
			return
		}

		fmt.Println("+++++++++++")
		fmt.Println(user.AuthorityId)

		var MenuIds []string

		for i := range SysAuthorityMenus {
			MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
		}

		var am system.SysBaseMenu
		ferr := global.GVA_DB.First(&am, "name = ? and id in (?)", user.Authority.DefaultRouter, MenuIds).Error
		if errors.Is(ferr, gorm.ErrRecordNotFound) {
			user.Authority.DefaultRouter = "404"
		}
	}
	return &user, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var reqUser system.SysUser
	err = global.GVA_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}

	var SysAuthorityMenus []system.SysAuthorityMenu
	err = global.GVA_DB.Where("sys_authority_authority_id = ?", reqUser.AuthorityId).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}

	var MenuIds []string

	for i := range SysAuthorityMenus {
		MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
	}

	var am system.SysBaseMenu
	ferr := global.GVA_DB.First(&am, "name = ? and id in (?)", reqUser.Authority.DefaultRouter, MenuIds).Error
	if errors.Is(ferr, gorm.ErrRecordNotFound) {
		reqUser.Authority.DefaultRouter = "404"
	}
	return reqUser, err
}
