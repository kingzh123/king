package controllers

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
)

type FromData struct {
	Name   string         `form:"name" json:"name"`
	Phone  string         `form:"phone" json:"phone,omitempty"`
	Pwd    string         `form:"pwd" json:"pwd"`
	RePwd  *FromDataRePwd `json:"r_pwd"`
	Email  FromDataEmail  `json:"email"`
	Colors []string       `form:"colors[]" json:"colors"`
}

type ShouldBindData struct {
	Name  string `form:"uname" json:"uname" validate:"required,min=4,max=10" label:"姓名"`
	Email string `form:"email" json:"email" validate:"required,email" label:"邮箱"`
}

type FromDataRePwd struct {
	RePwd string `form:"re_password" json:"re_password"`
}

type FromDataEmail struct {
	Email string `form:"email" json:"email"`
}

// 验证框架定义 name 中文名称
func customTagName(field reflect.StructField) string {
	label := field.Tag.Get("label")
	if label == "" {
		return field.Name
	}
	return label
}

func Redirect301() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	}
}

func Redirect302() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(302, gin.H{"code": 302, "msg": "redirect"})
	}
}

func RouterParams2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "msg": "success have * param"})
	}
}

func RouterParams1() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "msg": "success not * param"})
	}
}

// CheckData 验证框架 验证数据规则
func CheckData(data interface{}) string {
	// 设置语言为中文
	cn := zh.New()
	uni := ut.New(cn)
	trans, _ := uni.GetTranslator("zh")
	// 创建验证器
	validate := validator.New()
	// 对验证器进行中文注册
	err := zhtrans.RegisterDefaultTranslations(validate, trans)
	validate.RegisterTagNameFunc(customTagName)
	if err != nil {
		panic(err)
	}
	err = validate.Struct(data)
	if err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			return fmt.Sprintf("%s", fieldError.Translate(trans))
		}
	}
	return ""
}

// GetParamToStruct 绑定url参数到struct
func GetParamToStruct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data FromData
		err := c.Bind(&data)
		if err != nil {
			panic(err)
		}
		fmt.Println(data.Email)
		fmt.Println(data.RePwd)
		c.JSON(200, gin.H{"code": 200, "data": data})
	}
}

func PostParamToStruct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data FromData
		err := c.Bind(&data)
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{"code": 200, "data": data})
	}
}

func PostBindingToStruct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data ShouldBindData
		err := c.Bind(&data)
		if err != nil {
			panic(err)
		}
		con := CheckData(data)
		if con != "" {
			c.JSON(200, gin.H{"code": 500, "msg": con})
		} else {
			c.JSON(200, gin.H{"code": 200, "data": data})
		}
	}
}
