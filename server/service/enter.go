package service

import (
	"wucms-gva/server/service/example"
	"wucms-gva/server/service/pkg"
	"wucms-gva/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	PkgServiceGroup     pkg.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
