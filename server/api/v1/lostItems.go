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

// @Tags LostItems
// @Summary 创建LostItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LostItems true "创建LostItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lostItems/createLostItems [post]
func CreateLostItems(c *gin.Context) {
	var lostItems model.LostItems
	_ = c.ShouldBindJSON(&lostItems)
	if err := service.CreateLostItems(lostItems); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags LostItems
// @Summary 删除LostItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LostItems true "删除LostItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lostItems/deleteLostItems [delete]
func DeleteLostItems(c *gin.Context) {
	var lostItems model.LostItems
	_ = c.ShouldBindJSON(&lostItems)
	if err := service.DeleteLostItems(lostItems); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LostItems
// @Summary 批量删除LostItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LostItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lostItems/deleteLostItemsByIds [delete]
func DeleteLostItemsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteLostItemsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags LostItems
// @Summary 更新LostItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LostItems true "更新LostItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lostItems/updateLostItems [put]
func UpdateLostItems(c *gin.Context) {
	var lostItems model.LostItems
	_ = c.ShouldBindJSON(&lostItems)
	if err := service.UpdateLostItems(lostItems); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags LostItems
// @Summary 用id查询LostItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LostItems true "用id查询LostItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lostItems/findLostItems [get]
func FindLostItems(c *gin.Context) {
	var lostItems model.LostItems
	_ = c.ShouldBindQuery(&lostItems)
	if err, relostItems := service.GetLostItems(lostItems.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relostItems": relostItems}, c)
	}
}

// @Tags LostItems
// @Summary 分页获取LostItems列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LostItemsSearch true "分页获取LostItems列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lostItems/getLostItemsList [get]
func GetLostItemsList(c *gin.Context) {
	var pageInfo request.LostItemsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetLostItemsInfoList(pageInfo); err != nil {
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

func GetCurrentUserLostItemsList(c *gin.Context) {
	var pageInfo request.LostItemsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	var uuid = getUserUuid(c)
	//service.FilterLostItemsInfoListByUUID(pageInfo, uuid)

	if err, list, total := service.FilterItemsInfoListByUUID(pageInfo, uuid); err != nil {
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