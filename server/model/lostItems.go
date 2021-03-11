// 自动生成模板LostItems
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type LostItems struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;type:varchar(191);size:191;"`
      Location  string `json:"location" form:"location" gorm:"column:location;comment:;type:varchar(191);size:191;"`
      Time  time.Time `json:"time" form:"time" gorm:"column:time;comment:;type:datetime;"`
      IsFond  *bool `json:"isFond" form:"isFond" gorm:"column:is_fond;comment:;type:tinyint;"`
      Contact  string `json:"contact" form:"contact" gorm:"column:contact;comment:;type:varchar(191);size:191;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:;type:varchar(191);size:191;"`
      Created_by  string `json:"created_by" form:"created_by" gorm:"column:created_by;comment:"`
      Uuid  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:"`
}


func (LostItems) TableName() string {
  return "lost_items"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type LostItemsWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	LostItems   `json:"business"`
// }

// func (LostItems) TableName() string {
// 	return "lost_items"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["lostItems"] = func() model.GVA_Workflow {
//   return new(model.LostItemsWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["lostItems"] = func() interface{} {
// 	return new(model.LostItems)
// }
