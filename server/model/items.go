// 自动生成模板Items
package model

import (
      "gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Items struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;type:varchar(191);size:191;"`
      Location  string `json:"location" form:"location" gorm:"column:location;comment:;type:varchar(191);size:191;"`
      Time  string `json:"time" form:"time" gorm:"column:time;comment:;type:varchar(191);size:191;"`
      IsFond  *bool `json:"isFond" form:"isFond" gorm:"column:is_fond;comment:;type:tinyint;"`
      Picture  string `json:"picture" form:"picture" gorm:"column:picture;comment:;type:varchar(1000);size:1000;"`
      QQ  string `json:"QQ" form:"QQ" gorm:"column:QQ;comment:;type:varchar(20);size:20;"`
      Wechat  string `json:"wechat" form:"wechat" gorm:"column:wechat;comment:;type:varchar(100);size:100;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:;type:varchar(20);size:20;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:;type:varchar(1000);size:1000;"`
      CreatedBy  string `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:;type:varchar(191);size:191;"`
      UUID  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;type:varchar(191);size:191;"`
}


func (Items) TableName() string {
  return "items"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ItemsWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Items   `json:"business"`
// }

// func (Items) TableName() string {
// 	return "items"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["itm"] = func() model.GVA_Workflow {
//   return new(model.ItemsWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["itm"] = func() interface{} {
// 	return new(model.Items)
// }
