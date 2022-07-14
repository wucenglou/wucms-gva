package response

import "wucms-gva/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
