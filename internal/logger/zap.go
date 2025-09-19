package logger

import (
	"context"
	"fmt"
	"io"
	"mall/internal/core"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func WithContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return logger
	}
	duration := (time.Now().UnixNano() - cast.ToInt64(ctx.Value("startTime"))) / int64(time.Microsecond)
	return logger.With("duration", duration).With("traceId", ctx.Value("traceId"))
}

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

func InitLogger() error {
	logWriter, err := GetWriter(core.GlobalConfig.Logger.LogFile)
	if err != nil {
		fmt.Printf("get log writer error: %v", err)
		return err
	}

	c := zapcore.NewCore(GetEncoder(), zapcore.AddSync(logWriter), LogLevel(core.GlobalConfig.Logger.LogLevel))
	log := zap.New(c, zap.AddCaller())
	logger = log.Sugar()

	return nil
}
