package request

import (
	"wucms-gva/server/global"
	"wucms-gva/server/model/system"
)

type AddMenuAuthorityInfo struct {
	Menus       []system.SysBaseMenu `json:"menus"`
	AuthorityId uint                 `json:"authorityId"` // 角色ID
}

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		GVA_MODEL: global.GVA_MODEL{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index",
		Sort:      1,
		Meta: system.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
