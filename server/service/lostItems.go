package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateLostItems
//@description: 创建LostItems记录
//@param: lostItems model.LostItems
//@return: err error

func CreateLostItems(lostItems model.LostItems) (err error) {
	err = global.GVA_DB.Create(&lostItems).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteLostItems
//@description: 删除LostItems记录
//@param: lostItems model.LostItems
//@return: err error

func DeleteLostItems(lostItems model.LostItems) (err error) {
	err = global.GVA_DB.Delete(&lostItems).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteLostItemsByIds
//@description: 批量删除LostItems记录
//@param: ids request.IdsReq
//@return: err error

func DeleteLostItemsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.LostItems{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateLostItems
//@description: 更新LostItems记录
//@param: lostItems *model.LostItems
//@return: err error

func UpdateLostItems(lostItems model.LostItems) (err error) {
	err = global.GVA_DB.Save(&lostItems).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetLostItems
//@description: 根据id获取LostItems记录
//@param: id uint
//@return: err error, lostItems model.LostItems

func GetLostItems(id uint) (err error, lostItems model.LostItems) {
	err = global.GVA_DB.Where("id = ?", id).First(&lostItems).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetLostItemsInfoList
//@description: 分页获取LostItems记录
//@param: info request.LostItemsSearch
//@return: err error, list interface{}, total int64

func GetLostItemsInfoList(info request.LostItemsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
   // 创建db
	db := global.GVA_DB.Model(&model.LostItems{})
   var lostItems []model.LostItems
   // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&lostItems).Error
	return err, lostItems, total
}

//func FilterLostItemsInfoListByUUID(info request.LostItemsSearch, uuid string) (err error, list interface{}, total int64) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	// 创建db
//	db := global.GVA_DB.Model(&model.LostItems{})
//	var lostItems []model.LostItems
//	// 如果有条件搜索 下方会自动创建搜索语句
//	err = db.Where("uuid = ?", uuid).Error
//	err = db.Count(&total).Error
//	err = db.Limit(limit).Offset(offset).Find(&lostItems).Error
//	return err, lostItems, total
//}