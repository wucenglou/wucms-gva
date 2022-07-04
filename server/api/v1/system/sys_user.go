package system

import (
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/system"
	systemReq "wucms-gva/server/model/system/request"
	"wucms-gva/server/utils"

	"github.com/gin-gonic/gin"
)

// systemReq "wucms-gva/server/model/system/request"
// systemRes "wucms-gva/server/model/system/response"

func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	_ = c.ShouldBindJSON(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		if err, user := userService.Login(u); err != nil {

		}
	}
}
