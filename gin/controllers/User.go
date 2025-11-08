package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "msg": "create ok"})
	}
}

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"code": 200, "msg": fmt.Sprintf("user is %s", id)})
	}
}

func UserDel() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"code": 200, "msg": fmt.Sprintf("delete user is %s", id)})
	}
}
