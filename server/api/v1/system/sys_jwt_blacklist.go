package system

import (
	"fmt"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JwtApi struct{}

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "jwt加入黑名单"
// @Router /jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	err := jwtService.JsonInBlacklist(jwt)
	if err != nil {
		fmt.Println("jwt作废失败!")
		global.GVA_LOG.Error("jwt作废失败!", zap.Error(err))
		response.FailWithMessage("jwt作废失败", c)
		return
	}
	fmt.Println("jwt作废成功!")
	response.OkWithMessage("jwt作废成功", c)
}
