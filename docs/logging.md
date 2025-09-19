# Zap 日志使用指南

## 概述

项目已经配置好了基于 Zap 的日志系统，支持结构化日志和日志轮转。

## 配置

在 `configs/config.yaml` 中配置日志：

```yaml
logger:
  logFile: ./logs/mall.log
  logLevel: debug # debug, info, warn, error, panic, fatal
```

## 使用方式

### 1. 在主程序中

在 `cmd/mall/web.go` 中已经初始化了全局 logger：

```go
if err := InitLogger(); err != nil {
    fmt.Printf("Error initializing logger: %v\n", err)
    os.Exit(1)
}

// 使用全局 Logger 变量
Logger.Info("Starting web server", zap.String("address", core.GlobalConfig.Server.Addr))
```

### 2. 在其他包中使用全局 logger

可以使用 `zap.L()` 访问全局 logger：

```go
import "go.uber.org/zap"

zap.L().Info("This is an info message")
zap.L().Error("This is an error", zap.Error(err))
```

### 3. 使用便捷的包装函数

导入 `mall/internal/logger` 包，使用更简洁的接口：

```go
import "mall/internal/logger"

// 结构化日志
logger.Info("User login successful", zap.String("userId", "123"))
logger.Error("Database connection failed", zap.Error(err))

// 格式化日志（类似 fmt.Printf）
logger.Infof("User %s logged in at %s", userId, time.Now())
logger.Errorf("Failed to process request: %v", err)
```

### 4. 在控制器中的使用示例

```go
func SetUserInfo(c *gin.Context) {
    logger.Info("SetUserInfo request started")

    var req entity.SetUserInfoReq{}
    if err := c.ShouldBindJSON(&req); err != nil {
        logger.Error("Failed to bind JSON request", zap.Error(err))
        c.JSON(http.StatusOK, httputils.Error(err))
        return
    }

    logger.Debug("Request data", zap.Any("request", req))

    if err := service.SetUserInfo(c, req); err != nil {
        logger.Error("Service error", zap.Error(err))
        c.JSON(http.StatusOK, httputils.Error(err))
        return
    }

    logger.Info("SetUserInfo completed successfully")
    c.JSON(http.StatusOK, httputils.SuccessWithData(nil))
}
```

### 5. 在服务层中的使用示例

```go
func SetUserInfo(ctx context.Context, req entity.SetUserInfoReq) error {
    logger.Info("Starting user info update", zap.String("operation", "SetUserInfo"))

    // 业务逻辑
    if err := validateUserInfo(req); err != nil {
        logger.Warn("User info validation failed", zap.Error(err))
        return err
    }

    if err := database.UpdateUser(req); err != nil {
        logger.Error("Database update failed",
            zap.Error(err),
            zap.String("userId", req.UserId))
        return err
    }

    logger.Info("User info updated successfully", zap.String("userId", req.UserId))
    return nil
}
```

## 日志级别

- **Debug**: 调试信息，详细的程序执行信息
- **Info**: 一般信息，程序正常运行的关键节点
- **Warn**: 警告信息，可能的问题但不影响程序运行
- **Error**: 错误信息，程序出现错误但仍能继续运行
- **Panic**: 严重错误，程序会 panic
- **Fatal**: 致命错误，程序会退出

## 字段类型

Zap 提供了类型安全的字段：

```go
zap.String("key", "value")
zap.Int("count", 42)
zap.Bool("success", true)
zap.Error(err)
zap.Any("object", someStruct)
zap.Duration("elapsed", time.Since(start))
```

## 最佳实践

1. 在请求开始和结束时记录日志
2. 记录重要的业务操作
3. 错误发生时记录详细信息
4. 不要记录敏感信息（如密码、token）
5. 使用结构化字段而不是字符串拼接
6. 在生产环境中使用 info 级别，开发环境使用 debug 级别

## 日志轮转

日志会自动按天轮转，保留 7 天的历史日志：

- 当前日志：`./logs/mall.log`
- 历史日志：`./logs/mall.log.202501191200` (年月日时分)
