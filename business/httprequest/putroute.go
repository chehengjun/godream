package httprequest

import (
	"github.com/gin-gonic/gin"
)

func PutRouter(r *gin.Engine) {

	r.PUT("/put", func(c *gin.Context) {
		c.String(200, "HTTP请求\033[32m[PUT]\033[0m方法测试成功。")
	})

}
