package log

import (
	"fmt"
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

func GetWriter(filename string) (logf io.Writer, err error) {
	fmt.Printf("Initializing log file: %s\n", filename)
	logf, err = rotatelogs.New(filename+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	return
}
func GetEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

func LogLevel(level string) zapcore.Level {
	switch level {
	case "debug", "Debug", "DEBUG":
		return zapcore.DebugLevel
	case "info", "Info", "INFO":
		return zapcore.InfoLevel
	case "warn", "Warn", "WARN":
		return zapcore.WarnLevel
	case "error", "Error", "ERROR":
		return zapcore.ErrorLevel
	case "panic", "Panic", "PANIC":
		return zapcore.PanicLevel
	case "fatal", "Fatal", "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
