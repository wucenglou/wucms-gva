package v1

import (
	"wucms-gva/server/api/v1/example"
	"wucms-gva/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
