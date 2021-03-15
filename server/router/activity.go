package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitActivityRouter(Router *gin.RouterGroup) {
	ActivityRouter := Router.Group("act").Use(middleware.OperationRecord())
	{
		ActivityRouter.POST("createActivity", v1.CreateActivity)   // 新建Activity
		ActivityRouter.DELETE("deleteActivity", v1.DeleteActivity) // 删除Activity
		ActivityRouter.DELETE("deleteActivityByIds", v1.DeleteActivityByIds) // 批量删除Activity
		ActivityRouter.PUT("updateActivity", v1.UpdateActivity)    // 更新Activity
		ActivityRouter.GET("findActivity", v1.FindActivity)        // 根据ID获取Activity
		ActivityRouter.GET("getActivityList", v1.GetActivityList)  // 获取Activity列表
	}
}
