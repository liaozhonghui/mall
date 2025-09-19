package logs

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

func GetWriter(filename string) (logf io.Writer, err error) {
	logf, err = rotatelogs.New(filename+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	return
}
func GetEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,
	})
}

func LogLevel(level string) (Loglevel zapcore.Level) {
	switch level {
	case "debug", "Debug":
		Loglevel = zapcore.DebugLevel
	}
	return
}
