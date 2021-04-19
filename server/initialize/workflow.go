package initialize

import "gin-vue-admin/model"

func initWorkflowModel() {
	model.WorkflowBusinessStruct = make(map[string]func() model.GVA_Workflow)
	model.WorkflowBusinessStruct["leave"] = func() model.GVA_Workflow {
		return new(model.ExaWfLeaveWorkflow)
	}

	model.WorkflowBusinessStruct["acm"] = func() model.GVA_Workflow {
	  return new(model.ActivitiesManagementWorkflow)
	}
}

func initWorkflowTable() {
	model.WorkflowBusinessTable = make(map[string]func() interface{})
	model.WorkflowBusinessTable["leave"] = func() interface{} {
		return new(model.ExaWfLeave)
	}

	model.WorkflowBusinessTable["acm"] = func() interface{} {
		return new(model.ActivitiesManagement)
	}
}

func InitWkMode() {
	initWorkflowModel()
	initWorkflowTable()
}
