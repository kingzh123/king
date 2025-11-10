package controllers

import (
	"github.com/gin-gonic/gin"
)

func T1() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := map[string]interface{}{
			"name": "lee",
			"sex":  "男",
			"age":  18,
		}
		c.HTML(200, "u.html", gin.H{
			"title": "hello world",
			"user":  u,
			"name":  "东方甄选",
		})
	}
}

func T2(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 自定义函数
		type User struct {
			Name    string
			Age     int
			Address string
		}
		var users []User
		users = append(users, User{
			Name:    "lee",
			Age:     18,
			Address: "address",
		})
		users = append(users, User{
			Name:    "king",
			Age:     20,
			Address: "address",
		})
		users = append(users, User{
			Name:    "mary",
			Age:     19,
			Address: "sldkfj;alskdjf;asdkjf",
		})
		list := []int{1, 2, 3, 4}
		c.HTML(200, "index.tmpl", gin.H{
			"title": "hello world",
			"name":  "lee",
			"u":     users,
			"list":  list,
		})
	}
}
