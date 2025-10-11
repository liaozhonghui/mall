# go-mall-backend

这是一个基于 Go 的电商后台服务示例（轻量级架构），使用了常见的中间件、配置管理、路由框架与模块化分层设计。

本文档概述项目目录结构、主要设计思想、配置与启动方式，以及常见开发/调试步骤，帮助开发者快速上手并扩展功能。

## 关键特性

- 使用 Cobra 提供命令行入口（`mall` 二进制，支持子命令，例如 `web`）。
- 使用 Gin 作为 HTTP 框架，组织路由与中间件。
- 使用 Viper 加载 YAML 配置（`internal/core.InitConfig`）。
- 结构化日志（项目内有 `logger` 模块，基于 zap 实现）。
- 分层组织（controller/api、service、repo/dao、internal/core、entity 等）。

## 项目结构（摘要）

- `cmd/`：程序入口（CLI）。
	- `main.go`：程序入口，调用 `mall.Execute()`。
	- `mall/`：cobra 命令集合；包含 `web` 命令用于启动 HTTP 服务。
- `api/`
	- `controller/`：HTTP 控制器（处理请求、返回响应），例如 `healthCheck.go`、`user.go`。
	- `router/`：路由注册（`RegisterRouter`），把中间件和路由绑定到 `*gin.Engine` 上。
	- `httputils/`：通用的 HTTP 帮助/错误处理工具。
- `middleware/`：Gin 中间件（例如上下文注入 `context.go`、访问日志 `accessLogger.go`、恢复 `recover.go`、登录校验 `checkLogin.go`）。
- `internal/`：内部模块
	- `core/`：核心配置结构定义（`MallConfig` 与 `GlobalConfig`），以及 `InitConfig` 的辅助（`viper` 配置加载）。
	- `logger/`：日志初始化及配置（基于 zap）。
	- `dao/`：与数据源相关的封装（`db`、`redis`）。
	- `repo/`、`service/`：仓储层与业务服务层实现。
- `entity/`：领域模型（例如 `goods.go`、`user.go`）。
- `configs/`：默认配置文件（`config.yaml`、`config.docker.yaml`）。
- `test/`：单元/集成测试示例。

（仓库中还有 `Makefile`、`Dockerfile`、`docker-compose.yaml` 等运维辅助文件。）

## 主要模块说明

1) 启动与命令行

- 程序入口在 `cmd/main.go`：调用 `mall.Execute()`，由 Cobra 管理命令。
- `cmd/mall/web.go` 提供 `web` 子命令：
	- 在启动时会执行：
		- `core.InitConfig(config)`：用 Viper 加载 YAML 配置文件（如果未指定 `-c`，会使用默认 `./configs/config.yaml`）。
		- `logger.InitLogger()`：初始化全局日志器。
		- 创建 `gin.Engine`，并调用 `router.RegisterRouter(engine)` 注册中间件与路由。
		- 启动 `http.Server`（参数来自 `core.GlobalConfig.Server`，例如 Addr、Read/Write/Idle Timeout）。

示例：在项目根目录运行（PowerShell）：

```powershell
# 直接运行（开发）
go run ./cmd/main.go mall web -c ./configs/config.yaml

# 或先构建二进制再运行
go build -o mall ./cmd
./mall web -c ./configs/config.yaml
```

2) 配置

- 配置读取入口：`internal/core/viper.go` 中的 `InitConfig(configFile string)`。行为：
	- 若 `configFile==""`，使用内置默认路径 `./configs/config.yaml`。
	- 使用 `viper.SetConfigFile` + `viper.ReadInConfig()` 读取 YAML 配置。
	- 使用 `viper.Unmarshal(&GlobalConfig)` 将配置映射到 `core.GlobalConfig`（类型定义在 `internal/core/config.go`）。

- `core.GlobalConfig`（部分字段）：
	- `Server`（`ServerConfig`）：Addr、ReadTimeout、WriteTimeout、IdleTimeout。
	- `Logger`（`LoggerConfig`）：日志文件路径、日志级别。
	- `Mysql`（[]MysqlConfig）：多实例 MySQL 配置（DSN、trace 等）。
	- `Redis`（[]RedisConfig）：Redis 实例配置。
	- `Jwt`（`JwtConfig`）：API/管理员 token secret 与过期时间。

注意：`mapstructure` tag 已在结构体上声明，确保 viper.Unmarshal 正确绑定字段名。

3) 路由与中间件

- 路由注册在 `api/router/router.go`：
	- 全局中间件：`middleware.Context`、`middleware.AccessLogger`。
	- 管理后台路由组 `/admin`（`registerAdminRoutes`）。
	- 对外 API 路由组 `/api`（`registerAPIRoutes`），当前示例中注册了：
		- `/api/panic`：用于测试 panic 与 recover 中间件。
		- `/api/healthCheck`、`/api/healthCheckV1`：健康检查接口。
		- `/api/users`（PUT）：设置用户信息。
		- `/api/login`（POST）：登录接口。

- 主要中间件：
	- `Context`：在请求生命周期中注入请求上下文或跟踪信息。
	- `AccessLogger`：请求访问日志记录（通常结合 zap）。
	- `Recover`：捕获 panic 并返回 5xx 错误。
	- `CheckLogin`：认证校验（用于需要登录保护的路由）。

4) 服务层、仓储层与 DAO

- `service/`：封装业务逻辑，controller 层只负责解析请求与处理响应。
- `repo/`：仓储接口定义与实现，隐藏具体数据源细节。
- `internal/dao`（`db`、`redis`）：对底层 DB/Redis 连接与简单 CRUD 的封装。

这种分层有助于：测试、替换实现（例如把 MySQL 换成别的存储）、单元测试 mock。

## 常见操作

- 运行测试：

```powershell
go test ./... -v
```

- 代码格式化与静态检查：

```powershell
gofmt -w .
go vet ./...
# 可选: golangci-lint run
```

- 构建与运行（开发机器）：

```powershell
go build -o mall ./cmd
./mall web -c ./configs/config.yaml
```

## 如何添加新路由

1. 在 `api/controller/` 新建控制器处理函数（签名满足 gin.HandlerFunc）。
2. 在 `api/router/router.go` 的 `registerAPIRoutes` 或 `registerAdminRoutes` 中注册路由。
3. 如需鉴权，在路由注册时添加 `middleware.CheckLogin` 或在控制器中手动校验。

示例（伪代码）：

```go
// api/controller/foo.go
func FooHandler(c *gin.Context) {
		// 解析请求 -> 调用 service -> 返回 JSON
}

// api/router/router.go
rg.GET("/foo", middleware.CheckLogin, controller.FooHandler)
```

## 配置示例（要点）

- `configs/config.yaml`（示例说明）：
	- `server.addr`：监听地址（例：":8080"或"0.0.0.0:8080"）。
	- `server.readTimeOut` / `server.writeTimeOut` / `server.idleTimeOut`：超时配置（解析为 time.Duration）。
	- `mysql`：数组，可配置多个实例（每个包含 `instance`、`dsn` 等）。
	- `redis`：数组，配置连接参数与超时。

注意：viper 在解析 duration 字段时，YAML 中可以使用字符串（例如 "30s"）。如果使用整数字段，需按项目中类型要求设置。

## 开发者规范与风格

- 代码组织：按功能分层（controller -> service -> repo/dao -> entity）。
- 错误处理：controller 层应把业务错误转换为 HTTP 响应码与统一的错误结构（参考 `api/httputils`）。
- 日志：使用结构化日志（zap），日志级别与文件路径由配置控制。
- 配置：优先使用配置文件，必要时支持环境变量覆盖（可扩展 Viper 的读取策略）。

## 调试与排错建议

- 如果启动失败，先检查配置文件路径是否正确：`mall web -c ./configs/config.yaml`。
- 查看日志文件（`core.GlobalConfig.Logger.LogFile`）或控制台输出。
- 使用 `curl` 或浏览器访问健康检查：

```powershell
curl http://localhost:8080/ping
curl http://localhost:8080/api/healthCheck
```

## 扩展建议（短期优先）

- 在 `router` 中把路由注册拆成多个文件（按资源拆分），提高可维护性。
- 提供示例 `configs/config.local.yaml` 与 `configs/config.docker.yaml`，并在 README 中给出 docker-compose 启动示例。
- 增加更多单元测试覆盖 `service` 与 `repo` 层，使用接口 mock 依赖。

## 总结

该项目采用了典型的 Go 后端分层结构，使用 Cobra + Gin + Viper + Zap 的组合，易于扩展与维护。文档、路由与配置清晰，可作为一个电商后台服务的骨架。

如果你希望我把 README 再细化为：
- 完整的配置文件示例（逐字段注释）
- Docker / docker-compose 的启动示例
- 新增一个示例 API 的端到端实现（controller -> service -> repo -> dao）

我可以继续为你生成这些内容或将它们加入到仓库中。
