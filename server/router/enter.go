package router

import (
	"wucms-gva/server/router/example"
	"wucms-gva/server/router/pkg"
	"wucms-gva/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Pkg     pkg.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
