package controllers

import (
	m "king/gin/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger = m.CustomLogger("kk")

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 0, "msg": "ok"})
		query := c.Request.URL.Query()
		logger.Debug("debug log1:",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", query.Encode()))
	}
}

func Logger2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 0, "msg": "logger 2"})
		logger.Debug("debug log2:",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path))
	}
}
