package system

type ServiceGroup struct {
	JwtService
	MenuService
	UserService
	CasbinService
	InitDBService
	OperationRecordService
}
