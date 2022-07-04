package main

import (
	"wucms-gva/server/core"
	"wucms-gva/server/global"
	"wucms-gva/server/initialize"

	"github.com/gin-gonic/gin"
)

func main() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("127.0.0.1:8010") // listen and serve on 0.0.0.0:8080
}
