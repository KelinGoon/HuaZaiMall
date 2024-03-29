package conf

import (
	"github.com/joho/godotenv"
	"os"
	"singo/dao"
	"singo/module/util"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	dao.Database("root:123456@tcp(localhost:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local")
	// cache.Redis()
}
