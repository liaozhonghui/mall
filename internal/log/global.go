package log

import "go.uber.org/zap"

// Logger 包装函数，提供更方便的日志接口

func Debug(msg string, fields ...zap.Field) {
	zap.L().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zap.L().Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	zap.L().Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zap.L().Fatal(msg, fields...)
}

// Sugar logger 风格的便捷函数
func Debugf(template string, args ...interface{}) {
	zap.S().Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	zap.S().Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	zap.S().Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	zap.S().Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	zap.S().Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	zap.S().Fatalf(template, args...)
}
