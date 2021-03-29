package v1

import (
	"github.com/helix/global"
    "github.com/helix/model"
    "github.com/helix/model/request"
    "github.com/helix/model/response"
    "github.com/helix/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)


func GetCurrentUserItemsList(c *gin.Context) {
	var pageInfo request.ItemsSearch
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


// @Tags Items
// @Summary 创建Items
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Items true "创建Items"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /itm/createItems [post]
func CreateItems(c *gin.Context) {
	var itm model.Items
	_ = c.ShouldBindJSON(&itm)
	if err := service.CreateItems(itm); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Items
// @Summary 删除Items
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Items true "删除Items"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /itm/deleteItems [delete]
func DeleteItems(c *gin.Context) {
	var itm model.Items
	_ = c.ShouldBindJSON(&itm)
	if err := service.DeleteItems(itm); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Items
// @Summary 批量删除Items
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Items"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /itm/deleteItemsByIds [delete]
func DeleteItemsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteItemsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Items
// @Summary 更新Items
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Items true "更新Items"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /itm/updateItems [put]
func UpdateItems(c *gin.Context) {
	var itm model.Items
	_ = c.ShouldBindJSON(&itm)
	if err := service.UpdateItems(itm); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Items
// @Summary 用id查询Items
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Items true "用id查询Items"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /itm/findItems [get]
func FindItems(c *gin.Context) {
	var itm model.Items
	_ = c.ShouldBindQuery(&itm)
	if err, reitm := service.GetItems(itm.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reitm": reitm}, c)
	}
}

// @Tags Items
// @Summary 分页获取Items列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ItemsSearch true "分页获取Items列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /itm/getItemsList [get]
func GetItemsList(c *gin.Context) {
	var pageInfo request.ItemsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetItemsInfoList(pageInfo); err != nil {
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

func GetItemsListUser(c *gin.Context) {
	var pageInfo request.ItemsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetItemsInfoListUser(pageInfo); err != nil {
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