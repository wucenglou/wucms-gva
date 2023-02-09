package system

type ServiceGroup struct {
	JwtService
	ApiService
	MenuService
	UserService
	CasbinService
	InitDBService
	AutoCodeService
	BaseMenuService
	DictionaryService
	SystemConfigService
	AutoCodeHistoryService
	AuthorityService
	OperationRecordService
	DictionaryDetailService
	AuthorityBtnService
}
