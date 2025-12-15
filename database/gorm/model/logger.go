package model

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm/logger"
)

// DateFileLogger 带日期的文件日志
type DateFileLogger struct {
	zapLogger                 *zap.Logger
	sugarLogger               *zap.SugaredLogger
	logLevel                  logger.LogLevel
	slowThreshold             time.Duration
	ignoreRecordNotFoundError bool
	basePath                  string // 日志基础路径
	currentDate               string // 当前日期
	writer                    *lumberjack.Logger
}

// NewDateFileLogger 创建带日期的文件日志器
func NewDateFileLogger(logDir, logName string, level logger.LogLevel) (*DateFileLogger, error) {
	// 确保日志目录存在
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}

	// 获取当前日期
	currentDate := time.Now().Format("2006-01-02")

	// 构建日志文件路径
	logPath := filepath.Join(logDir, fmt.Sprintf("%s_%s.log", logName, currentDate))

	// 配置 lumberjack
	writer := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    100,  // 每个日志文件最大100MB
		MaxBackups: 30,   // 保留30个备份
		MaxAge:     90,   // 保留90天
		Compress:   true, // 压缩旧文件
		LocalTime:  true, // 使用本地时间
	}

	// 创建 zap core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(writer),
		zapcore.InfoLevel,
	)

	// 创建 logger
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))

	return &DateFileLogger{
		zapLogger:     zapLogger,
		sugarLogger:   zapLogger.Sugar(),
		logLevel:      level,
		slowThreshold: 200 * time.Millisecond,
		basePath:      logPath,
		currentDate:   currentDate,
		writer:        writer,
	}, nil
}

// checkAndRotate 检查是否需要轮转日志文件
func (l *DateFileLogger) checkAndRotate() {
	today := time.Now().Format("2006-01-02")
	if today != l.currentDate {
		// 日期变更，创建新文件
		l.currentDate = today
		newPath := strings.Replace(l.basePath,
			time.Now().AddDate(0, 0, -1).Format("2006-01-02"),
			today, 1)
		l.writer.Filename = newPath

		// 重新创建 logger
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(l.writer),
			zapcore.InfoLevel,
		)

		newLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
		l.zapLogger = newLogger
		l.sugarLogger = newLogger.Sugar()
	}
}

// GORM Logger 接口实现
func (l *DateFileLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.logLevel = level
	return l
}

func (l *DateFileLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Info {
		l.checkAndRotate()
		l.sugarLogger.Infof(msg, data...)
	}
}

func (l *DateFileLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Warn {
		l.checkAndRotate()
		l.sugarLogger.Warnf(msg, data...)
	}
}

func (l *DateFileLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Error {
		l.checkAndRotate()
		l.sugarLogger.Errorf(msg, data...)
	}
}

func (l *DateFileLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.logLevel <= logger.Silent {
		return
	}

	l.checkAndRotate()

	elapsed := time.Since(begin)
	sql, rows := fc()

	// 获取调用者信息
	var callerInfo string
	if _, file, line, ok := runtime.Caller(4); ok {
		callerInfo = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}

	fields := []zap.Field{
		zap.String("sql", sql),
		zap.Int64("rows", rows),
		zap.Duration("elapsed", elapsed),
		zap.String("time", time.Now().Format("2006-01-02 15:04:05.000")),
	}

	if callerInfo != "" {
		fields = append(fields, zap.String("caller", callerInfo))
	}

	switch {
	case err != nil && !(l.ignoreRecordNotFoundError && logger.ErrRecordNotFound == err):
		l.zapLogger.Error("SQL Error", append(fields, zap.Error(err))...)
	case elapsed > l.slowThreshold && l.slowThreshold != 0:
		l.zapLogger.Warn("Slow SQL", fields...)
	case l.logLevel >= logger.Info:
		l.zapLogger.Info("SQL", fields...)
	}
}
