package v1

import (
	"wucms-gva/server/api/v1/example"
	"wucms-gva/server/api/v1/pkg"
	"wucms-gva/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	PkgApiGroup     pkg.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
