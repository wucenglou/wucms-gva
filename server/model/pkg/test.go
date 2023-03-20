package pkg

// type User struct {
// 	gorm.Model
// 	Name string
// 	// Roles []Role `gorm:"many2many:user_roles;"`
// }

// type Role struct {
// 	gorm.Model
// 	Name  string
// 	Users []User `gorm:"many2many:user_roles;foreignKey:ID;joinForeignKey:role_id;References:ID;joinReferences:user_id"`
// }

// type UserRole struct {
// 	gorm.Model
// 	RoleId uint
// 	UserId uint
// }

// func (UserRole) TableName() string {
// 	return "user_roles"
// }
