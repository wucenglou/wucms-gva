package request

import "wucms-gva/server/model/pkg"

// TermStruct 结构体
type ReqPost struct {
	pkg.Post
	TermIds []int `json:"term_id"`
}
