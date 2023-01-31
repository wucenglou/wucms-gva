package request

import (
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/system"
)

type SearchApiParams struct {
	system.SysApi
	request.PageInfo
	OrderKey string `json:"orderKey"` //排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
