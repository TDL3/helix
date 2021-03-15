package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateActivity
//@description: 创建Activity记录
//@param: act model.Activity
//@return: err error

func CreateActivity(act model.Activity) (err error) {
	err = global.GVA_DB.Create(&act).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteActivity
//@description: 删除Activity记录
//@param: act model.Activity
//@return: err error

func DeleteActivity(act model.Activity) (err error) {
	err = global.GVA_DB.Delete(&act).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteActivityByIds
//@description: 批量删除Activity记录
//@param: ids request.IdsReq
//@return: err error

func DeleteActivityByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Activity{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateActivity
//@description: 更新Activity记录
//@param: act *model.Activity
//@return: err error

func UpdateActivity(act model.Activity) (err error) {
	err = global.GVA_DB.Save(&act).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetActivity
//@description: 根据id获取Activity记录
//@param: id uint
//@return: err error, act model.Activity

func GetActivity(id uint) (err error, act model.Activity) {
	err = global.GVA_DB.Where("id = ?", id).First(&act).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetActivityInfoList
//@description: 分页获取Activity记录
//@param: info request.ActivitySearch
//@return: err error, list interface{}, total int64

func GetActivityInfoList(info request.ActivitySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Activity{})
    var acts []model.Activity
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if !info.Time.IsZero() {
         db = db.Where("`time` <> ?",info.Time)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&acts).Error
	return err, acts, total
}