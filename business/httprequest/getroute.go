package httprequest

import (
	"github.com/gin-gonic/gin"
	"godream/business/global"
)

func GetRouter(e *gin.Engine) {
	e.GET(GetTestRoute, GetTestRouteHandler)
	e.GET(GetTest2Route, GetTest2RouteHandler)
}

func GetTestRouteHandler(c *gin.Context) {
	var err error
	var request = "[Http-Get] test"
	var response string
	rpcC := global.GetRpcRdbClient(global.ConfigFile)
	err = rpcC.Call("RDB.Query", request, &response)
	if err != nil {
		panic(err)
	}
	c.JSON(200, response)
}

func GetTest2RouteHandler(c *gin.Context) {
	var err error
	var request = "[Http-Get] test"
	var response string
	rpcC := global.GetRpcRdbClient(global.ConfigFile)
	err = rpcC.Call("RDB.Query", request, &response)
	if err != nil {
		panic(err)
	}
	c.JSON(200, response)
}
