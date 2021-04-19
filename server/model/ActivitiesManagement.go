// 自动生成模板ActivitiesManagement
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type ActivitiesManagement struct {
	global.GVA_MODEL
	Name               string     `json:"name" form:"name" gorm:"column:name;comment:活动名称;type:varchar(191);size:191;"`
	Start_time         time.Time  `json:"start_time" form:"start_time" gorm:"column:start_time;comment:活动开始时间"`
	End_time           time.Time  `json:"end_time" form:"end_time" gorm:"column:end_time;comment:活动开始时间"`
	Loaction           string     `json:"loaction" form:"loaction" gorm:"column:loaction;comment:活动位置;type:varchar(120);size:120;"`
	NeededPersonnel    int        `json:"neededPersonnel" form:"neededPersonnel" gorm:"column:needed_personnel;comment:需要人数;type:bigint;size:19;"`
	Budget             string     `json:"budget" form:"budget" gorm:"column:budget;comment:活动经费;type:varchar(1000);size:1000;"`
	Score              int        `json:"score" form:"score" gorm:"column:score;comment:活动分数;type:int;size:20;"`
	Description        string     `json:"description" form:"description" gorm:"column:description;comment:活动说明;type:varchar(1000);size:1000;"`
	CreatedBy          string     `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:申请人;type:varchar(191);size:191;"`
	ReqUnion           int        `json:"reqUnion" form:"reqUnion" gorm:"column:req_union;comment:申请部门;type:bigint;size:19;"`
	Approved           *bool      `json:"approved" form:"approved" gorm:"column:approved;comment:"`
	ManagementAudit    string     `json:"managementAudit" form:"managementAudit" gorm:"column:management_audit;comment:审核意见;type:varchar(191);size:191;"`
	CreatedUserUuid    string     `json:"createdUserUuid" form:"createdUserUuid" gorm:"column:created_user_uuid;comment:申请人ID;type:varchar(191);size:191;"`
	UserRefer          []*SysUser `gorm:"many2many:UserRefer"`
	UserAttendantRefer []*SysUser `gorm:"many2many:attendance_refer;"`
}

func (ActivitiesManagement) TableName() string {
	return "activities"
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
type ActivitiesManagementWorkflow struct {
	// 工作流操作结构体
	WorkflowBase         `json:"wf"`
	ActivitiesManagement `json:"business"`
}

// func (ActivitiesManagement) TableName() string {
// 	return "activities_mag"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["acm"] = func() model.GVA_Workflow {
//   return new(model.ActivitiesManagementWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["acm"] = func() interface{} {
// 	return new(model.ActivitiesManagement)
// }
