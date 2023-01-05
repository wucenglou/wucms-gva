package system

type ServiceGroup struct {
	JwtService
	ApiService
	MenuService
	UserService
	CasbinService
	InitDBService
	BaseMenuService
	SystemConfigService
	AuthorityService
	OperationRecordService
}
