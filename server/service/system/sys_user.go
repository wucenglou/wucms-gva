package system

import (
	"fmt"
	"wucms-gva/server/global"
	"wucms-gva/server/model/system"
	"wucms-gva/server/utils"
)

type UserService struct{}

func (userService *UserService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init"), nil
	}

	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
	return err, &user
}
