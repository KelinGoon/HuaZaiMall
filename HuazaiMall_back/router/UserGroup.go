package router

import (
	"github.com/gin-gonic/gin"
	jwt "singo/module/auth"
	"singo/service"
)

func AddUserGroup(r *gin.RouterGroup) {
	//用户路由组
	userGroup := r.Group("user")
	{
		userGroup.POST("/login", service.UserLogin)       //  /user/login
		userGroup.POST("/register", service.UserRegister) //  /user/login
		//userGroup.GET("/index", service.Userme)

		// 需要权限
		jwtGroup := userGroup.Group("/")
		jwtGroup.Use(jwt.JwtRequired())
		{
			// 查看个人信息
			jwtGroup.GET("/info", service.Userme)
			// 	_ "information_collection_back_end/docs"

			//修改个人信息
			//jwtGroup.POST("/update", service.UpdateUser)
			// 退出登录
			//jwtGroup.POST("/user/logout", service.UserLogout)
			// 修改密码
			//jwtGroup.POST("/update_pass", service.UpdatePassword)
		}
	}

}
