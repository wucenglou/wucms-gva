package router

import (
	"wucms-gva/server/router/backend"
	"wucms-gva/server/router/example"
	"wucms-gva/server/router/frontend"
	"wucms-gva/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Backend  backend.RouterGroup
	Frontend frontend.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
