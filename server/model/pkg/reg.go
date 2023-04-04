package pkg

import (
	"time"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/system"
)

type RegSearch struct {
	Reg
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type DoctorSearch struct {
	Doctor
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type PatientSearch struct {
	Patient
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

// Pkg 结构体
type Reg struct {
	global.GVA_MODEL
	Name       string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Phone      *int   `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
	Bz         string `json:"bz" form:"bz" gorm:"column:bz;comment:;"`
	Gender     string `json:"gender" form:"gender" gorm:"column:gender;comment:;"`
	From       string `json:"from" form:"from" gorm:"column:from;comment:;"`
	Time       string `json:"time" form:"time" gorm:"column:time;comment:;"`
	Desc       string `json:"desc" form:"desc" gorm:"type:longtext;column:desc;comment:;"`
	GptSuggest string `json:"gptSuggest" form:"gptSuggest" gorm:"type:longtext;column:gptSuggest;comment:;"`
	Ip         string `json:"ip" form:"ip" gorm:"column:ip;comment:;"`
	IpDesc     string `json:"ipDesc" form:"ipDesc" gorm:"column:ipDesc;comment:;"`
	DoctorName string `json:"doctor_name" form:"doctor_name" gorm:"column:doctor_name;comment:;"`

	PatientId *int           `json:"patientid" form:"patientid" gorm:"foreignKey:ID;index:patientid;column:patientid;"`
	Patient   system.SysUser `gorm:"foreignKey:PatientId;"`

	DoctorId *int           `json:"doctorid" form:"doctorid" gorm:"foreignKey:ID;index:doctorid;column:doctorid;"`
	Doctor   system.SysUser `gorm:"foreignKey:DoctorId;"`
}

type Doctor struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:名字;"`
	Title     string `json:"title" form:"title" gorm:"column:title;comment:头衔;"`
	Specialty string `json:"specialty" form:"specialty" gorm:"type:longtext;column:specialty;comment:专长;"`
	Status    string `json:"status" form:"status" gorm:"size:20;index:status;column:status;default:'';comment:Publish发布Draft草稿Pending Review待审核Private私有Future未来Trash垃圾箱reject拒绝Auto-Draft自动草稿"`
	MiniDesc  string `json:"mini_desc" form:"mini_desc" gorm:"type:longtext;column:mini_desc;comment:简短介绍;"`
	Desc      string `json:"desc" form:"desc" gorm:"type:longtext;column:desc;comment:个人介绍;"`

	ProfileImg string `json:"profile_img" gorm:"comment:头像"`
	DescImg    string `json:"desc_img" gorm:"comment:描述大图"`
	RegNum     *int   `json:"reg_num" form:"reg_num" gorm:"column:reg_num;comment:预约量;"`
	StarNum    *int   `json:"star_num" form:"star_num" gorm:"column:star_num;comment:关注量;"`
	Phone      *int   `json:"phone" form:"phone" gorm:"column:phone;comment:联系方式手机号;"`
	Ext        string `json:"ext" form:"ext" gorm:"column:ext;default:''"`
}

type Patient struct {
	global.GVA_MODEL
	Name   string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Phone  *int   `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
	Gender string `json:"gender" form:"gender" gorm:"column:gender;comment:;"`
	From   string `json:"from" form:"from" gorm:"column:from;comment:;"`
	Desc   string `json:"desc" form:"desc" gorm:"type:longtext;column:desc;comment:;"`
	Ip     string `json:"ip" form:"ip" gorm:"column:ip;comment:;"`

	UserId *int           `json:"user_id" form:"user_id" gorm:"size:200;foreignKey:ID;index:user_id;column:user_id;"`
	User   system.SysUser `gorm:"foreignKey:UserId;"`
}

// TableName Pkg 表名
func (Reg) TableName() string {
	return "reg"
}

// TableName Pkg 表名
func (Doctor) TableName() string {
	return "doctor"
}

// TableName Pkg 表名
func (Patient) TableName() string {
	return "patient"
}
