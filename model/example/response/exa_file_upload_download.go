package response

import "github.com/lhxlnsy/pap-go-server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
