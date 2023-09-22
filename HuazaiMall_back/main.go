package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"singo/conf"
	"singo/dal/query"
	"singo/dao"
	"singo/router"
)

func main() {

	// 从配置文件读取配置
	conf.Init()

	query.SetDefault(dao.DB)

	// 装载路由
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := router.SetupRouter()
	r.Run(":8082")
}
