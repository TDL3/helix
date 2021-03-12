package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)


func FilterItemsInfoListByUUID(info request.ItemsSearch, uuid string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Items{})
	var Items []model.Items
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Where("uuid = ?", uuid).Error
	err = db.Where("is_fond = 0").Error
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Items).Error
	return err, Items, total
}



//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateItems
//@description: 创建Items记录
//@param: itm model.Items
//@return: err error

func CreateItems(itm model.Items) (err error) {
	err = global.GVA_DB.Create(&itm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteItems
//@description: 删除Items记录
//@param: itm model.Items
//@return: err error

func DeleteItems(itm model.Items) (err error) {
	err = global.GVA_DB.Delete(&itm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteItemsByIds
//@description: 批量删除Items记录
//@param: ids request.IdsReq
//@return: err error

func DeleteItemsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Items{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateItems
//@description: 更新Items记录
//@param: itm *model.Items
//@return: err error

func UpdateItems(itm model.Items) (err error) {
	err = global.GVA_DB.Save(&itm).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetItems
//@description: 根据id获取Items记录
//@param: id uint
//@return: err error, itm model.Items

func GetItems(id uint) (err error, itm model.Items) {
	err = global.GVA_DB.Where("id = ?", id).First(&itm).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetItemsInfoList
//@description: 分页获取Items记录
//@param: info request.ItemsSearch
//@return: err error, list interface{}, total int64

func GetItemsInfoList(info request.ItemsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Items{})
    var itms []model.Items
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Title != "" {
        db = db.Where("`title` LIKE ?","%"+ info.Title+"%")
    }
    if info.Location != "" {
        db = db.Where("`location` LIKE ?","%"+ info.Location+"%")
    }
    if info.Time != "" {
         db = db.Where("`time` = ?",info.Time)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&itms).Error
	return err, itms, total
}

func GetItemsInfoListUser(info request.ItemsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Items{})
	err = db.Where("is_fond = 0").Error
	var itms []model.Items
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Title != "" {
		db = db.Where("`title` LIKE ?","%"+ info.Title+"%")
	}
	if info.Location != "" {
		db = db.Where("`location` LIKE ?","%"+ info.Location+"%")
	}
	if info.Time != "" {
		db = db.Where("`time` = ?",info.Time)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&itms).Error
	return err, itms, total
}