package request

import "github.com/helix/model"

type WorkflowProcessSearch struct {
	model.WorkflowProcess
	PageInfo
}
