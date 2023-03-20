package request

import (
	"time"
	"wucms-gva/server/model/common/request"
)

type TermStructSearch struct {
	// myPkg.TermStruct
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
