package system

import (
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/system"
	systemRes "wucms-gva/server/model/system/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApi struct{}

// @Tags System
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=systemRes.SysConfigResponse,msg=string} "获取配置文件内容,返回包括系统配置"
// @Router /system/getSystemConfig [post]
func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	config, err := systemConfigService.GetSystemConfig()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", c)
	}
}

// @Tags System
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body system.System true "设置配置文件内容"
// @Success 200 {object} response.Response{data=string} "设置配置文件内容"
// @Router /system/setSystemConfig [post]
func (s *SystemApi) SetSystemConfig(c *gin.Context) {
	var sys system.System
	_ = c.ShouldBindJSON(&sys)
	err := systemConfigService.SetSystemConfig(sys)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

// @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取服务器信息"
// @Router /system/getServerInfo [post]
func (s *SystemApi) GetServerInfo(c *gin.Context) {
	server, err := systemConfigService.GetServerInfo()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}
}
