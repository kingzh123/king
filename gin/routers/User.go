package routers

import (
	c "king/gin/controllers"

	"github.com/gin-gonic/gin"
)

// User 定义路由
func User(r *gin.Engine) {
	g := r.Group("/user") // 定义组
	g.GET("/:id", c.User())
	g.GET("/del/:id", c.UserDel())
}

func ParamSend(r *gin.Engine) {
	g := r.Group("/param")
	g.GET("/send", c.GetParamToStruct())
	g.POST("/send", c.PostParamToStruct())
	g.POST("/send/binding", c.PostBindingToStruct())
	g.POST("/upload", c.Upload())
	g.POST("/uploads", c.Uploads())
}
