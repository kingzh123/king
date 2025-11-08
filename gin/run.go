package gin

import (
	"context"
	"fmt"
	m "king/gin/middlewares"
	"king/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	logger := m.CustomLogger("")
	// 三方日志库zap
	router := gin.New()
	router.Static("/asset", "./gin/asset")
	router.Use(m.ZapLogger(logger), m.ZapRecovery(logger, true))
	router.MaxMultipartMemory = 8 << 20
	//路由注册
	routers.User(router)
	routers.ParamSend(router)
	routers.Log(router)
	routers.Params(router)
	routers.Redirect(router)
	routers.Temp(router)
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
	// 启动web服务
	go func() {
		if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(err.Error())
		}
	}()
	// 创建一个信号通道
	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit //阻塞信号 当接收到上述信号时，才会通过
	// 创建一个超时5秒上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒后优雅的关闭web服务 如果没有执行完请求关闭服务会发生错误
	// Shutdown 为go1.18+版本的优雅关闭
	if err := serve.Shutdown(ctx); err != nil {
		fmt.Println("server shutdown")
	}
	fmt.Println("gin server exiting")
}
