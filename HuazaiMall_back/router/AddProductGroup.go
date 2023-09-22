package router

import (
	"github.com/gin-gonic/gin"
	jwt "singo/module/auth"
	"singo/service"
)

func AddProductGroup(r *gin.RouterGroup) {
	//用户路由组
	userGroup := r.Group("product")
	{
		// 需要权限
		jwtGroup := userGroup.Group("/")
		jwtGroup.Use(jwt.JwtRequired())
		{
			// 查看指定商家的商品信息
			jwtGroup.GET("/info", service.ProductInfo)
			// 添加商品
			jwtGroup.POST("/add", service.AddProduct)
			// 删除商品
			jwtGroup.POST("/delete", service.DeleteProduct)
			// 修改商品信息
			jwtGroup.POST("/update", service.UpdateProduct)
		}
	}

}
