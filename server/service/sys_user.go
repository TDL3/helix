package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser

func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysUser{})
	var userList []model.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func DeleteUser(id float64) (err error) {
	var user model.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func SetUserInfo(reqUser model.SysUser) (err error, user model.SysUser) {
	err = global.GVA_DB.Updates(&reqUser).Error
	return err, reqUser
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func FindUserById(id int) (err error, user *model.SysUser) {
	var u model.SysUser
	err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func FindUserByUuid(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil{
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}


func GetStudentUserScoreList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysUser{})
	var studentList []response.UserScoreInfo
	err = db.Count(&total).Error
	// Query students only
	//err = db.Limit(limit).Offset(offset).Preload("Authority").Where("authority_id = ?", 1000).Find(&studentList).Error

	rows, err := db.Limit(limit).Offset(offset).Preload("Authority").Where("authority_id = ?", 1000).Rows()

	if err == nil {
		defer rows.Close()

		for rows.Next() {
			// In this case user is student, since all users with an authority id of 1000 is considered student
			var user model.SysUser
			var activities []model.ActivitiesManagement
			var userScoreTotal int
			// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
			err = db.ScanRows(rows, &user)
			//fmt.Println("user: ", user)
			err = global.GVA_DB.Model(&user).Association("Activities").Find(&activities)
			for _, activity := range activities {
				userScoreTotal += activity.Score
				//fmt.Println("activity score: ", activity.Score)
			}
			student := response.UserScoreInfo{
				HeaderImg: user.HeaderImg,
				NickName: user.NickName,
				Score: userScoreTotal,
				Username: user.Username,
			}
			studentList = append(studentList, student)
		}
		//fmt.Println(studentList)
	}
	return err, studentList, total
}


func GetUnionScoreList(info request.PageInfo) (err error, list interface{}, total int64) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
	//db := global.GVA_DB.Model(&model.SysUser{})
	db := global.GVA_DB.Table("activities")

	err = db.Count(&total).Error
	// Query students only
	//err = db.Limit(limit).Offset(offset).Preload("Authority").Where("authority_id = ?", 1000).Find(&studentList).Error

	rows, err := db.Where("deleted_at is null").Rows()
	var scores []response.UnionScoreInfo
	//err = db.Select().Scan(&scores).Error()
	//fmt.Println(rows)
	if err == nil {
		defer rows.Close()

		for rows.Next() {
			// activities stores all activities' union id and score for individual activities
			var activities []model.ActivitiesManagement
			scoreTotal := make(map[int]int)
			//fmt.Println(rows)
			// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
			err = db.ScanRows(rows, &activities)
			////fmt.Println("user: ", user)
			//err = global.GVA_DB.Model(&user).Association("Activities").Find(&activities)
			for _, activity := range activities {
				unionId := activity.ReqUnion
				score := activity.Score
				//fmt.Println(unionId, score)
				//fmt.Println(scoreTotal[unionId])
				if scoreTotal[unionId] == 0 {
					scoreTotal[unionId] = score
					continue
				}
				scoreTotal[unionId] += score
				//fmt.Println("activity score: ", activity.Score)

			}
			for key, value := range scoreTotal {
				score := response.UnionScoreInfo{
					Score:		value,
					ReqUnion:   key,
				}
				scores = append(scores, score)
			}

			//fmt.Println(scores)
		}
	}
	return err, scores, total
}
