package system

type ServiceGroup struct {
	JwtService
	ApiService
	MenuService
	UserService
	CasbinService
	InitDBService
	BaseMenuService
	DictionaryService
	SystemConfigService
	AuthorityService
	OperationRecordService
}
