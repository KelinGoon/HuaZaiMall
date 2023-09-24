package router

import (
	"github.com/gin-gonic/gin"
	"singo/module/middleware"
)

// SetupRouter @Security	ApiKeyAuth
func SetupRouter() *gin.Engine {
	// SetupRouter是一个函数，用于在gin框架中设置路由和HTTP请求的处理程序。
	// 通定义路由和相应的处理程序，
	// 使得不同的HTTP请求的处理程序可以被正确地调用。
	r := gin.Default()
	r.Use(middleware.Cors()) //配置跨域
	//加载静态资源
	r.Static("/static", "static")
	//api服务
	apiGroup := r.Group("/")
	{
		//访问各个路由组，展开程序
		AddUserGroup(apiGroup)
		AddProductGroup(apiGroup)
		AddOrderGroup(apiGroup)
		//AddRecordGroup(apiGroup)
		//AddTickerGroup(apiGroup)
		//AddRecordTagGroup(apiGroup)
		//AddCenterApiGroup(apiGroup)
	}
	return r
}
