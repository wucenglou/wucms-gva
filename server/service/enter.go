package service

import (
	"wucms-gva/server/service/backend"
	"wucms-gva/server/service/example"
	"wucms-gva/server/service/frontend"
	"wucms-gva/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	BackendServiceGroup  backend.ServiceGroup
	FrontendServiceGroup frontend.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
