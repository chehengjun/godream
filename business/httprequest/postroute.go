package httprequest

import (
	"github.com/gin-gonic/gin"
)

func PostRouter(r *gin.Engine) {
	r.POST("/post1", func(c *gin.Context) {
		c.String(200, "注册用户。")
	})

	r.POST("/post2", func(c *gin.Context) {
		c.String(200, "添加服务器。")
	})

	r.POST("/post3", func(c *gin.Context) {
		c.String(200, "发表帖子。")
	})
}
