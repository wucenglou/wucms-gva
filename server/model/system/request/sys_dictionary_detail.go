package request

import (
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
