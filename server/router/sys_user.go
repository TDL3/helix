package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.POST("register", v1.Register)
		UserRouter.POST("changePassword", v1.ChangePassword)     // 修改密码
		UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
		UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser", v1.DeleteUser)           // 删除用户
		UserRouter.PUT("setUserInfo", v1.SetUserInfo)            // 设置用户信息
		UserRouter.GET("getUserInfo", v1.GetUserInfo) 			 // 获取当前登录用户的信息
		UserRouter.POST("getStudentUserScoreList", v1.GetStudentUserScoreList) 			 // 获取所有学生的成绩列表
		UserRouter.POST("getUnionScoreList", v1.GetUnionScoreList) 			 // 获取所有系部的成绩列表

	}
}
