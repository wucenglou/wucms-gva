package system

import "wucms-gva/server/service"

type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	SystemApi
	SystemApiApi
	AuthorityApi
	AuthorityMenuApi
}

var (
	apiService          = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService          = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService         = service.ServiceGroupApp.SystemServiceGroup.UserService
	menuService         = service.ServiceGroupApp.SystemServiceGroup.MenuService
	initDBService       = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	baseMenuService     = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	authorityService    = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
)
