package main

import (
	"fmt"
	"wucms-gva/server/core"
	"wucms-gva/server/global"
	"wucms-gva/server/initialize"

	"go.uber.org/zap"
)

func main() {
	fmt.Println("1*******************************")
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	fmt.Println("2*******************************")
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	fmt.Println("3*******************************")
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	fmt.Println("4*******************************")

	// Router := gin.Default()
	// Router.GET("/ping2", func(ctx *gin.Context) {
	// 	ctx.String(200, "ping2")
	// })

	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
