package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitActivitiesRouter(Router *gin.RouterGroup) {
	ActivitiesRouter := Router.Group("acti").Use(middleware.OperationRecord())
	{
		ActivitiesRouter.POST("createActivities", v1.CreateActivities)   // 新建Activities
		ActivitiesRouter.DELETE("deleteActivities", v1.DeleteActivities) // 删除Activities
		ActivitiesRouter.DELETE("deleteActivitiesByIds", v1.DeleteActivitiesByIds) // 批量删除Activities
		ActivitiesRouter.PUT("updateActivities", v1.UpdateActivities)    // 更新Activities
		ActivitiesRouter.GET("findActivities", v1.FindActivities)        // 根据ID获取Activities
		ActivitiesRouter.GET("getActivitiesList", v1.GetActivitiesList)  // 获取Activities列表
	}
}
