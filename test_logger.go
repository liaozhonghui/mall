package main

import (
	"mall/internal/core"
	"mall/internal/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func testLogger() {
	// 初始化配置
	if err := core.InitConfig(""); err != nil {
		panic(err)
	}

	// 初始化日志
	logWriter, err := logger.GetWriter(core.GlobalConfig.Logger.LogFile)
	if err != nil {
		panic(err)
	}

	c := zapcore.NewCore(
		logger.GetEncoder(),
		zapcore.AddSync(logWriter),
		logger.LogLevel(core.GlobalConfig.Logger.LogLevel),
	)

	log := zap.New(c, zap.AddCaller())
	zap.ReplaceGlobals(log)

	// 测试各种日志级别
	logger.Debug("This is a debug message", zap.String("key", "value"))
	logger.Info("This is an info message", zap.Int("count", 42))
	logger.Warn("This is a warning message", zap.Bool("flag", true))
	logger.Error("This is an error message", zap.String("error", "test error"))

	// 测试格式化日志
	logger.Infof("User %s logged in at %d", "john", 1234567890)
	logger.Errorf("Failed to process request: %v", "connection timeout")

	println("日志测试完成，请检查 ./logs/mall.log 文件")
}
