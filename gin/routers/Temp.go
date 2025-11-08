package routers

import (
	c "king/gin/controllers"

	"github.com/gin-gonic/gin"
)

func Temp(r *gin.Engine) {
	r.GET("/t1", c.T1(r))
	r.GET("/t2", c.T2(r))
}
