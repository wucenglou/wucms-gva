package pkg

// type User struct {
// 	global.GVA_MODEL
// 	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                                                           // 用户UUID
// 	UserName  string    `json:"userName" gorm:"index:user_name;comment:用户登录名"`                                        // 用户登录名
// 	Password  string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
// 	NickName  string    `json:"nickName" gorm:"index:nick_name;default:系统用户;comment:用户昵称"`                            // 用户昵称
// 	HeaderImg string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
// 	Phone     string    `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
// 	Email     string    `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
// 	Enable    int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                      //用户是否被冻结 1正常 2冻结
// 	UserMeta  []UserMeta
// }

// type UserMeta struct {
// 	UmetaId   *int   `gorm:"primarykey"` // 主键ID
// 	UserId    *int   `json:"user_id" form:"user_id" gorm:"index:user_id;column:user_id;comment:;default:0;"`
// 	MetaKey   string `json:"meta_key" form:"meta_key" gorm:"index:meta_key;column:meta_key;comment:;"`
// 	MetaValue string `json:"meta_value" form:"meta_value" gorm:"column:meta_value;comment:;"`
// }

// func (User) TableName() string {
// 	return "users"
// }

// func (UserMeta) TableName() string {
// 	return "usermeta"
// }
