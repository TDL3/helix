package model

import "gin-vue-admin/global"

type Attendance struct {
	global.GVA_MODEL
	UserId uint `gorm:"column:user_id; "`
	ActivityID int `gorm:"column:activity_id; "`
	Present []ActivitiesManagement `gorm:"many2many:attendance_refer;"`
}