package response

import "wucms-gva/server/model/MyPkg"

type UserResponse struct {
	User MyPkg.User `json:"user"`
}

type LoginResponse struct {
	User      MyPkg.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}
