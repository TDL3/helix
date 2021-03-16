package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitActivitiesManagementRouter(Router *gin.RouterGroup) {
	ActivitiesManagementRouter := Router.Group("acm").Use(middleware.OperationRecord())
	{
		ActivitiesManagementRouter.POST("createActivitiesManagement", v1.CreateActivitiesManagement)   // 新建ActivitiesManagement
		ActivitiesManagementRouter.DELETE("deleteActivitiesManagement", v1.DeleteActivitiesManagement) // 删除ActivitiesManagement
		ActivitiesManagementRouter.DELETE("deleteActivitiesManagementByIds", v1.DeleteActivitiesManagementByIds) // 批量删除ActivitiesManagement
		ActivitiesManagementRouter.PUT("updateActivitiesManagement", v1.UpdateActivitiesManagement)    // 更新ActivitiesManagement
		ActivitiesManagementRouter.GET("findActivitiesManagement", v1.FindActivitiesManagement)        // 根据ID获取ActivitiesManagement
		ActivitiesManagementRouter.GET("getActivitiesManagementList", v1.GetActivitiesManagementList)  // 获取ActivitiesManagement列表
	}
}
