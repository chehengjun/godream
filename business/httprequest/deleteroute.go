package httprequest

import (
	"github.com/gin-gonic/gin"
)

func DeleteRouter(r *gin.Engine) {

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "HTTP请求\033[32m[DELETE]\033[0m方法测试成功。")
	})

}
