package request

import (
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
