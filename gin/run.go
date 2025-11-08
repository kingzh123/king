package gin

import (
	"fmt"
	m "king/gin/middlewares"
	"king/routers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	logger := m.CustomLogger("")
	// 三方日志库zap
	router := gin.New()
	router.Use(m.ZapLogger(logger), m.ZapRecovery(logger, true))
	router.MaxMultipartMemory = 8 << 20
	//路由注册
	routers.User(router)
	routers.ParamSend(router)
	routers.Log(router)
	routers.Params(router)
	routers.Redirect(router)
	gin.SetMode(gin.DebugMode) // gin.:DeBugMode(开发环境模式)、ReleaseMode(生产环境模式)、TestMode(测试模式)
	router.GET("/", func(c *gin.Context) {
		//Gin的上下文设计为‌非并发安全‌（non-concurrent safe），若在多个Goroutine中共享原始上下文，
		//可能导致竞态条件如多个Goroutine同时修改响应内容
		cCp := c.Copy()
		go func() {
			time.Sleep(2 * time.Second)
			fmt.Println(cCp.Request.URL.Path)
		}()
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "okk"})
	})
	serve := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}
	fmt.Println(serve.ListenAndServe())
}
