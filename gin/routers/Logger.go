package routers

import (
	c "king/gin/controllers"

	"github.com/gin-gonic/gin"
)

func Log(r *gin.Engine) {
	r.GET("/log", c.Logger())
	r.GET("/log2", c.Logger2())
}
