package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ActivitiesManagement
// @Summary 创建ActivitiesManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActivitiesManagement true "创建ActivitiesManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /acm/createActivitiesManagement [post]
func CreateActivitiesManagement(c *gin.Context) {
	var acm model.ActivitiesManagement
	_ = c.ShouldBindJSON(&acm)
	if err := service.CreateActivitiesManagement(acm); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ActivitiesManagement
// @Summary 删除ActivitiesManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActivitiesManagement true "删除ActivitiesManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /acm/deleteActivitiesManagement [delete]
func DeleteActivitiesManagement(c *gin.Context) {
	var acm model.ActivitiesManagement
	_ = c.ShouldBindJSON(&acm)
	if err := service.DeleteActivitiesManagement(acm); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ActivitiesManagement
// @Summary 批量删除ActivitiesManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActivitiesManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /acm/deleteActivitiesManagementByIds [delete]
func DeleteActivitiesManagementByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteActivitiesManagementByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ActivitiesManagement
// @Summary 更新ActivitiesManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActivitiesManagement true "更新ActivitiesManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /acm/updateActivitiesManagement [put]
func UpdateActivitiesManagement(c *gin.Context) {
	var acm model.ActivitiesManagement
	_ = c.ShouldBindJSON(&acm)
	if err := service.UpdateActivitiesManagement(acm); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ActivitiesManagement
// @Summary 用id查询ActivitiesManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActivitiesManagement true "用id查询ActivitiesManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /acm/findActivitiesManagement [get]
func FindActivitiesManagement(c *gin.Context) {
	var acm model.ActivitiesManagement
	_ = c.ShouldBindQuery(&acm)
	if err, reacm := service.GetActivitiesManagement(acm.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reacm": reacm}, c)
	}
}

// @Tags ActivitiesManagement
// @Summary 分页获取ActivitiesManagement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ActivitiesManagementSearch true "分页获取ActivitiesManagement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /acm/getActivitiesManagementList [get]
func GetActivitiesManagementList(c *gin.Context) {
	var pageInfo request.ActivitiesManagementSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetActivitiesManagementInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}


func GetAttendedActivitiesInfoList(c *gin.Context) {
	var pageInfo request.ActivitiesManagementSearch
	_ = c.ShouldBindQuery(&pageInfo)
	var user model.SysUser
	uuid := getUserUuid(c)
	global.GVA_DB.Model(&model.SysUser{}).Where("uuid = ?", uuid).First(&user)
	if err, list, total := service.GetAttendedActivitiesInfoList(user); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func UpdateAttendedActivitiesInfoList(c *gin.Context) {
	var user model.SysUser
	uuid := getUserUuid(c)
	global.GVA_DB.Model(&model.SysUser{}).Where("uuid = ?", uuid).First(&user)

	var acm model.ActivitiesManagement
	_ = c.ShouldBindQuery(&acm)

	if err := service.UpdateAttendedActivities(user, acm); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func GetAttendant(c *gin.Context) {
	var pageInfo request.ActivitiesManagementSearch
	_ = c.ShouldBindQuery(&pageInfo)
	var user model.SysUser
	uuid := getUserUuid(c)
	global.GVA_DB.Model(&model.SysUser{}).Where("uuid = ?", uuid).First(&user)
	if err, list, total := service.GetAttendance(user); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func UpdateAttendant(c *gin.Context) {
	var user model.SysUser
	uuid := getUserUuid(c)
	global.GVA_DB.Model(&model.SysUser{}).Where("uuid = ?", uuid).First(&user)

	var acm model.ActivitiesManagement
	_ = c.ShouldBindQuery(&acm)
	//fmt.Println("c ", c)
	//fmt.Println("acm ", acm)
	if err := service.UpdateAttendance(user, acm); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func GetAttendantList(c *gin.Context) {
	var pageInfo request.ActivitiesManagementSearch
	_ = c.ShouldBindQuery(&pageInfo)
	var user model.SysUser
	uuid := getUserUuid(c)
	var acm model.ActivitiesManagement
	_ = c.ShouldBindQuery(&acm)
	actID := acm.ID
	global.GVA_DB.Model(&model.SysUser{}).Where("uuid = ?", uuid).First(&user)
	global.GVA_DB.Model(&model.ActivitiesManagement{}).Where("id = ?", actID).First(&acm)
	if err, list, total := service.GetAttendantList(user, acm); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}