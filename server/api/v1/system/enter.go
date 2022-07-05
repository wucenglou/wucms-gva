package system

import "wucms-gva/server/service"

type ApiGroup struct {
	DBApi
	BaseApi
}

var (
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)
