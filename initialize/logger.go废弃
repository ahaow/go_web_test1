package initialize

import (
	"go_web_test1/logger"
	"time"

	"go.uber.org/zap"
)

func InitLogger(env string, logFile string) (*logger.Logger, error) {
	// 创建Logger
	l, err := logger.NewLogger(env, logFile)
	if err != nil {
		return nil, err
	}

	// 绑定一些全局字段
	l = l.WithFields(map[string]interface{}{
		"app":     "my-app",
		"version": "1.0.0",
		"env":     env,
	})

	// 日志准备好时可以打印一条启动信息
	l.Info("Logger initialized",
		zap.String("log_file", logFile),
	)

	// 这块可以按需要进行测试
	l.Error("Example error",
		zap.Error(nil),
		zap.Duration("retry_after", 5*time.Second),
	)

	return l, nil
}
