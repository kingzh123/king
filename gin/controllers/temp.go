package controllers

import "github.com/gin-gonic/gin"

func T1(r *gin.Engine) gin.HandlerFunc {
	r.LoadHTMLGlob("gin/template/**/*")
	return func(c *gin.Context) {
		c.HTML(200, "u.html", gin.H{
			"title": "hello world",
		})
	}
}

func T2(r *gin.Engine) gin.HandlerFunc {
	r.LoadHTMLGlob("gin/template/**/*")
	return func(c *gin.Context) {
		c.HTML(200, "home.html", gin.H{
			"title": "hello world",
		})
	}
}
