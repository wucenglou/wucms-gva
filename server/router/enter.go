package router

import (
	"wucms-gva/server/router/example"
	"wucms-gva/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
