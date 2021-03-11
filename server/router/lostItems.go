package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitLostItemsRouter(Router *gin.RouterGroup) {
	LostItemsRouter := Router.Group("lostItems").Use(middleware.OperationRecord())
	{
		LostItemsRouter.POST("createLostItems", v1.CreateLostItems)   // 新建LostItems
		LostItemsRouter.DELETE("deleteLostItems", v1.DeleteLostItems) // 删除LostItems
		LostItemsRouter.DELETE("deleteLostItemsByIds", v1.DeleteLostItemsByIds) // 批量删除LostItems
		LostItemsRouter.PUT("updateLostItems", v1.UpdateLostItems)    // 更新LostItems
		LostItemsRouter.GET("findLostItems", v1.FindLostItems)        // 根据ID获取LostItems
		LostItemsRouter.GET("getLostItemsList", v1.GetLostItemsList)  // 获取LostItems列表
		LostItemsRouter.GET("getCurrentUserLostItemsList", v1.GetCurrentUserLostItemsList)
	}
}
