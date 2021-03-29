package response

import "github.com/helix/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
