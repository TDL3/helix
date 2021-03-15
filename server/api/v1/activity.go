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

// @Tags Activity
// @Summary 创建Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activity true "创建Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /act/createActivity [post]
func CreateActivity(c *gin.Context) {
	var act model.Activity
	_ = c.ShouldBindJSON(&act)
	if err := service.CreateActivity(act); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Activity
// @Summary 删除Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activity true "删除Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /act/deleteActivity [delete]
func DeleteActivity(c *gin.Context) {
	var act model.Activity
	_ = c.ShouldBindJSON(&act)
	if err := service.DeleteActivity(act); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Activity
// @Summary 批量删除Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /act/deleteActivityByIds [delete]
func DeleteActivityByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteActivityByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Activity
// @Summary 更新Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activity true "更新Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /act/updateActivity [put]
func UpdateActivity(c *gin.Context) {
	var act model.Activity
	_ = c.ShouldBindJSON(&act)
	if err := service.UpdateActivity(act); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Activity
// @Summary 用id查询Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activity true "用id查询Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /act/findActivity [get]
func FindActivity(c *gin.Context) {
	var act model.Activity
	_ = c.ShouldBindQuery(&act)
	if err, react := service.GetActivity(act.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"react": react}, c)
	}
}

// @Tags Activity
// @Summary 分页获取Activity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ActivitySearch true "分页获取Activity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /act/getActivityList [get]
func GetActivityList(c *gin.Context) {
	var pageInfo request.ActivitySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetActivityInfoList(pageInfo); err != nil {
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
