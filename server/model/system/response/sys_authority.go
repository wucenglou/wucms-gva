package response

import "wucms-gva/server/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}
