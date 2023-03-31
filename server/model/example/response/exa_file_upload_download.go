package response

import "wucms-gva/server/model/example"

type ExaFileResponse struct {
	File     example.ExaFileUploadAndDownload `json:"file"`
	Location string                           `json:"location"`
}
