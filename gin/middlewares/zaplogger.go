package middlewares

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func CustomLogger(fileName string) *zap.Logger {
	var fName string
	if fileName == "" {
		fName = viper.GetString("server.env")
	} else {
		fName = fileName
	}
	encoder := getEncoder()
	writeSyncer := getLumberJackWriter(fName)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(zapcore.NewTee(core))
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)
	return logger
}

func InitZapLoggerToFile() {
	var (
		allCore []zapcore.Core
		core    zapcore.Core
	)
	InitConfig()
	encoder := getEncoder()
	fileName := viper.GetString("server.env")
	writeSyncer := getLumberJackWriter(fileName)
	allCore = append(allCore, zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel))

	core = zapcore.NewTee(allCore...)
	logger = zap.New(core, zap.AddCaller())
	defer logger.Sync()
}

func getLumberJackWriter(fileName string) zapcore.WriteSyncer {
	file := fmt.Sprintf("./logs/%s.log", fileName)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件位置
		MaxSize:    1,    // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: 5,    // 保留旧文件的最大个数
		MaxAge:     1,    // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger))
	//return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		InitConfig()
		logger = CustomLogger("")
		defer logger.Sync()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		start := time.Now()
		c.Next()

		elapsed := time.Since(start)
		logger.Debug(fmt.Sprintf("%s %s", path, "basic logging"),
			zap.String("method", c.Request.Method),
			zap.String("query", query),
			zap.Int("status", c.Writer.Status()),
			zap.String("ip", c.ClientIP()),
			zap.Duration("elapsed", elapsed),
		)
	}
}

func ZapRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
