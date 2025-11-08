package routers

import (
	c "king/gin/controllers"

	"github.com/gin-gonic/gin"
)

func Params(r *gin.Engine) {
	g := r.Group("/p")
	g.POST("/:user", c.RouterParams1())
	g.POST("/:user/*action", c.RouterParams2())
}

func Redirect(r *gin.Engine) {
	g := r.Group("/r")
	g.GET("/301", c.Redirect301())
	g.GET("/302", c.Redirect302())
}
