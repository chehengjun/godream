package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"godream/business/global"
	"godream/business/httprequest"
	"log"
)

func main() {
	log.Println("\n\t\t\t\033[31mA Simple Web App Frame For Golang Of Back-End.\n\033[0m")

	// 读取配置文件
	fp := flag.String("c", global.CONFIG_FILE, global.USAGE_01)
	global.ConfigFile = global.LoadConfigFile(*fp)
	global.SelfConfig = global.LoadSelfConfig(global.ConfigFile)

	// 处理Http(s)[get|post|put|delete]等请求
	router := gin.Default()
	httprequest.GetRouter(router)
	httprequest.PostRouter(router)
	httprequest.PutRouter(router)
	httprequest.DeleteRouter(router)

	// 运行web后端服务
	router.Run(":" + global.SelfConfig.PORT)
}
