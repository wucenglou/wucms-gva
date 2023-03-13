package v1

import (
	"wucms-gva/server/api/v1/backend"
	"wucms-gva/server/api/v1/example"
	"wucms-gva/server/api/v1/frontend"
	"wucms-gva/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	BackendApiGroup  backend.ApiGroup
	FrontendApiGroup frontend.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
