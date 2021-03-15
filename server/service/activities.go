package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateActivities
//@description: 创建Activities记录
//@param: acti model.Activities
//@return: err error

func CreateActivities(acti model.Activities) (err error) {
	err = global.GVA_DB.Create(&acti).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteActivities
//@description: 删除Activities记录
//@param: acti model.Activities
//@return: err error

func DeleteActivities(acti model.Activities) (err error) {
	err = global.GVA_DB.Delete(&acti).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteActivitiesByIds
//@description: 批量删除Activities记录
//@param: ids request.IdsReq
//@return: err error

func DeleteActivitiesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Activities{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateActivities
//@description: 更新Activities记录
//@param: acti *model.Activities
//@return: err error

func UpdateActivities(acti model.Activities) (err error) {
	err = global.GVA_DB.Save(&acti).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetActivities
//@description: 根据id获取Activities记录
//@param: id uint
//@return: err error, acti model.Activities

func GetActivities(id uint) (err error, acti model.Activities) {
	err = global.GVA_DB.Where("id = ?", id).First(&acti).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetActivitiesInfoList
//@description: 分页获取Activities记录
//@param: info request.ActivitiesSearch
//@return: err error, list interface{}, total int64

func GetActivitiesInfoList(info request.ActivitiesSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Activities{})
    var actis []model.Activities
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if !info.Time.IsZero() {
         db = db.Where("`time` <> ?",info.Time)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&actis).Error
	return err, actis, total
}