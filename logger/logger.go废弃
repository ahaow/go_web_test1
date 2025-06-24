package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 封装的 zap 日志实例
type Logger struct {
	*zap.Logger
}

// NewLogger 初始化 zap 日志库，根据环境选择配置
// env: "production" 或其他（如 "development"）
// logFile: 日志输出文件路径（仅生产环境使用）
func NewLogger(env, logFile string) (*Logger, error) {
	var cfg zap.Config
	var err error

	if env == "production" {
		// 生产环境：JSON 格式，输出到文件和控制台，带调用栈
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式为 ISO8601
		cfg.EncoderConfig.TimeKey = "timestamp"                   // 时间字段名
		cfg.OutputPaths = []string{logFile, "stdout"}             // 输出到文件和控制台
		cfg.ErrorOutputPaths = []string{logFile, "stderr"}        // 错误输出
	} else {
		// 开发环境：人类可读格式，仅输出到控制台
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 带颜色日志级别
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.OutputPaths = []string{"stdout"}
		cfg.ErrorOutputPaths = []string{"stderr"}
	}

	// 构建 Logger
	logger, err := cfg.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return &Logger{logger}, nil
}

// Close 关闭日志，同步缓冲区
func (l *Logger) Close() {
	_ = l.Logger.Sync() // 忽略 Sync 错误，通常在程序退出时调用
}

// WithFields 添加公共字段到日志
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return &Logger{l.Logger.With(zapFields...)}
}

// 封装的日志方法，简化调用
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.Logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.Logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.Logger.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.Logger.Fatal(msg, fields...)
}
