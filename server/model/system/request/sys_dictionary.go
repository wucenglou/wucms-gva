package request

import (
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
