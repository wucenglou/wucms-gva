package initialize

import (
	_ "wucms-gva/server/source/example"
	_ "wucms-gva/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
