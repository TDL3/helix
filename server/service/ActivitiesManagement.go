package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateActivitiesManagement
//@description: 创建ActivitiesManagement记录
//@param: acm model.ActivitiesManagement
//@return: err error

func CreateActivitiesManagement(acm model.ActivitiesManagement) (err error) {
	err = global.GVA_DB.Create(&acm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteActivitiesManagement
//@description: 删除ActivitiesManagement记录
//@param: acm model.ActivitiesManagement
//@return: err error

func DeleteActivitiesManagement(acm model.ActivitiesManagement) (err error) {
	err = global.GVA_DB.Delete(&acm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteActivitiesManagementByIds
//@description: 批量删除ActivitiesManagement记录
//@param: ids request.IdsReq
//@return: err error

func DeleteActivitiesManagementByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ActivitiesManagement{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateActivitiesManagement
//@description: 更新ActivitiesManagement记录
//@param: acm *model.ActivitiesManagement
//@return: err error

func UpdateActivitiesManagement(acm model.ActivitiesManagement) (err error) {
	err = global.GVA_DB.Save(&acm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetActivitiesManagement
//@description: 根据id获取ActivitiesManagement记录
//@param: id uint
//@return: err error, acm model.ActivitiesManagement

func GetActivitiesManagement(id uint) (err error, acm model.ActivitiesManagement) {
	err = global.GVA_DB.Where("id = ?", id).First(&acm).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetActivitiesManagementInfoList
//@description: 分页获取ActivitiesManagement记录
//@param: info request.ActivitiesManagementSearch
//@return: err error, list interface{}, total int64

func GetActivitiesManagementInfoList(info request.ActivitiesManagementSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ActivitiesManagement{})
    var acms []model.ActivitiesManagement
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if !info.Start_time.IsZero() {
         db = db.Where("`start_time` > ?",info.Start_time)
    }
    if !info.End_time.IsZero() {
         db = db.Where("`end_time` < ?",info.End_time)
    }
    if info.Approved != nil {
        db = db.Where("`approved` = ?",info.Approved)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&acms).Error
	return err, acms, total
}



func GetAttendedActivitiesInfoList(user model.SysUser) (err error, list interface{}, total int64) {
	db := global.GVA_DB.Model(&user).Association("Activities")
	//db.Association("Activities")
	var acms []model.ActivitiesManagement
	//err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Association("Activities").Find(&acms)
	//total = db.Count()
	err = db.Find(&acms)
	//fmt.Println("acms:", acms)
	//fmt.Println("err: ", err)
	return err, acms, total
}



func UpdateAttendedActivities(user model.SysUser, activity model.ActivitiesManagement) (err error) {
	err = global.GVA_DB.Model(&user).Association("Activities").Append(&activity)
	//err := global.GVA_DB.Migrator().CreateConstraint(&model.SysUser{}, "fk_users_activities").Error()
	//global.GVA_LOG.Error(err)
	//var adminUser model.SysUser
	////err := global.GVA_DB.Where("username = ?", "admin").Preload("Authority").First(&adminUser).Error
	//err = global.GVA_DB.Where("username = ?", "admin").First(&adminUser).Error
	//print(err)
	////fmt.Println(adminUser)
	//var act1 model.ActivitiesManagement
	//err = global.GVA_DB.Where("id = ?", 8).First(&act1).Error
	//print(err)
	////global.GVA_DB.Model(model.SysUser{}).Association("Activities").Find(&act1)
	////fmt.Println(act1)
	//err = global.GVA_DB.Model(&adminUser).Association("Activities").Append(&act1)
	//
	////fmt.Println(act1)
	//var acts []model.ActivitiesManagement
	//////
	//err = global.GVA_DB.Model(&user).AssociatLimit(limit).Offset(offset)ion("Activities").Find(&acts)
	//fmt.Println("acts", acts)



	// 创建db
	//db := global.GVA_DB.Model(&user).Association("Activities")
	////db.Association("Activities")
	//var acms []model.ActivitiesManagement
	////err = db.Count(&total).Error
	////err = db.Limit(limit).Offset(offset).Association("Activities").Find(&acms)
	////total = db.Count()
	//err = db.Find(&acms)
	//fmt.Println("acms:", acms)
	//fmt.Println("err: ", err)
	//return err, acms, total

	return err
}