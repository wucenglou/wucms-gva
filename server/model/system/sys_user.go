package system

import (
	"wucms-gva/server/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"comment:用户UUID"`                                                                                                                                   // 用户UUID
	OpenId      string         `json:"openId" gorm:"comment:用户openid"`                                                                                                                               // 用户UUID
	Username    string         `json:"userName" gorm:"comment:用户登录名"`                                                                                                                                // 用户登录名
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                                                                                                     // 用户登录密码
	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                                                                                    // 用户昵称
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                                                                                                  // 用户侧边主题
	HeaderImg   string         `json:"headerImg" gorm:"default:https://mmbiz.qpic.cn/mmbiz/icTdbqWNOwNRna42FI242Lcia07jQodd2FJGIYQfG0LAJGFxM4FbnQP6yfMxBgJ0F3YRqJCJ1aPAK2dQagdusBZg/0;comment:用户头像"` // 用户头像
	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                                                                                                   // 基础颜色
	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                                                                                              // 活跃颜色
	AuthorityId uint           `json:"authorityId" gorm:"default:0;comment:用户角色ID888"`                                                                                                               // 用户角色ID
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone       string         `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email       string         `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable      int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结

	UserMeta []UserMeta `gorm:"foreignKey:UserId;"`
}

type UserMeta struct {
	UmetaId   *int   `gorm:"primarykey"` // 主键ID
	UserId    *int   `json:"user_id" form:"user_id" gorm:"index:user_id;column:user_id;comment:;default:0;"`
	MetaKey   string `json:"meta_key" form:"meta_key" gorm:"index:meta_key;column:meta_key;comment:;"`
	MetaValue string `json:"meta_value" form:"meta_value" gorm:"column:meta_value;comment:;"`
}

func (UserMeta) TableName() string {
	return "usermeta"
}

func (SysUser) TableName() string {
	return "sys_users"
}
