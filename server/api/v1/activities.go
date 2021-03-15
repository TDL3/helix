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

// @Tags Activities
// @Summary 创建Activities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activities true "创建Activities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /acti/createActivities [post]
func CreateActivities(c *gin.Context) {
	var acti model.Activities
	_ = c.ShouldBindJSON(&acti)
	if err := service.CreateActivities(acti); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Activities
// @Summary 删除Activities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activities true "删除Activities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /acti/deleteActivities [delete]
func DeleteActivities(c *gin.Context) {
	var acti model.Activities
	_ = c.ShouldBindJSON(&acti)
	if err := service.DeleteActivities(acti); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Activities
// @Summary 批量删除Activities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Activities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /acti/deleteActivitiesByIds [delete]
func DeleteActivitiesByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteActivitiesByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Activities
// @Summary 更新Activities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activities true "更新Activities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /acti/updateActivities [put]
func UpdateActivities(c *gin.Context) {
	var acti model.Activities
	_ = c.ShouldBindJSON(&acti)
	if err := service.UpdateActivities(acti); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Activities
// @Summary 用id查询Activities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activities true "用id查询Activities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /acti/findActivities [get]
func FindActivities(c *gin.Context) {
	var acti model.Activities
	_ = c.ShouldBindQuery(&acti)
	if err, reacti := service.GetActivities(acti.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reacti": reacti}, c)
	}
}

// @Tags Activities
// @Summary 分页获取Activities列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ActivitiesSearch true "分页获取Activities列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /acti/getActivitiesList [get]
func GetActivitiesList(c *gin.Context) {
	var pageInfo request.ActivitiesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetActivitiesInfoList(pageInfo); err != nil {
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
