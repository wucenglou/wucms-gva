package system

import "wucms-gva/server/service"

type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	AuthorityMenuApi
}

var (
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	menuService   = service.ServiceGroupApp.SystemServiceGroup.MenuService
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)
