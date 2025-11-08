package controllers

import (
	"fmt"
	"log"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadFileToDst(file *multipart.FileHeader, dst string, c *gin.Context) error {
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		return err
	}
	return nil
}

// Upload 上传单个文件
func Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if file == nil {
			c.JSON(200, gin.H{"code": 200, "msg": "file is nil"})
			return
		}
		if err != nil {
			c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		dst := "./uploads/" + file.Filename
		err = UploadFileToDst(file, dst, c)
		if err != nil {
			c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		log.Println(file.Filename)
	}
}

// Uploads 批量上传文件
func Uploads() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get context from
		from, err := c.MultipartForm()
		if err != nil {
			c.JSON(200, gin.H{"code": 200, "msg": err.Error()})
		}
		// get files
		files := from.File["file[]"]
		var errs []string
		for _, file := range files {
			dst := "./uploads/" + file.Filename
			err = UploadFileToDst(file, dst, c)
			if err != nil {
				errs = append(errs, file.Filename)
			}
		}
		if len(errs) > 0 {
			c.JSON(200, gin.H{"code": 500, "msg": fmt.Sprintf("%s 上传失败", strings.Join(errs, ","))})
			return
		}
		c.JSON(200, gin.H{"code": 200})
	}
}
