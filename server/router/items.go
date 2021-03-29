package router

import (
	"github.com/helix/api/v1"
	"github.com/helix/middleware"
	"github.com/gin-gonic/gin"
)

func InitItemsRouter(Router *gin.RouterGroup) {
	ItemsRouter := Router.Group("itm").Use(middleware.OperationRecord())
	{
		ItemsRouter.POST("createItems", v1.CreateItems)   // 新建Items
		ItemsRouter.DELETE("deleteItems", v1.DeleteItems) // 删除Items
		ItemsRouter.DELETE("deleteItemsByIds", v1.DeleteItemsByIds) // 批量删除Items
		ItemsRouter.PUT("updateItems", v1.UpdateItems)    // 更新Items
		ItemsRouter.GET("findItems", v1.FindItems)        // 根据ID获取Items
		ItemsRouter.GET("getItemsList", v1.GetItemsList)  // 获取Items列表
		ItemsRouter.GET("getItemsListUser", v1.GetItemsListUser)  // 获取Items列表 (User version, same as above but fonded items are not send)
		ItemsRouter.GET("getCurrentUserItemsList", v1.GetCurrentUserItemsList)
	}
}
