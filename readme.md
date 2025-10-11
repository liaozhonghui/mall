# Go Mall Backend

一个基于 Go 语言开发的现代化商城后端管理系统，采用清洁架构设计模式，提供高性能、可扩展的电商解决方案。

## 🚀 项目特性

- **清洁架构设计**：遵循DDD（领域驱动设计）和分层架构原则
- **高性能**：基于 Gin 框架构建的高性能 HTTP 服务
- **多数据库支持**：支持 MySQL 和 PostgreSQL
- **缓存机制**：集成 Redis 和 BigCache 双重缓存策略
- **安全认证**：基于 JWT 的身份认证和授权机制
- **容器化部署**：支持 Docker 容器化部署
- **配置管理**：灵活的多环境配置管理
- **日志系统**：结构化日志记录和轮转
- **中间件支持**：完善的中间件生态

## 📁 项目架构

```
go-mall-backend/
├── api/                    # API 层
│   ├── controller/         # 控制器层，处理HTTP请求
│   ├── httputils/          # HTTP工具类
│   ├── middleware/         # 中间件
│   └── router/             # 路由配置
├── internal/               # 内部业务逻辑（不对外暴露）
│   ├── core/               # 核心配置
│   ├── dao/                # 数据访问层
│   │   ├── cache/          # 缓存层
│   │   ├── db/             # 数据库操作
│   │   ├── pg/             # PostgreSQL操作
│   │   └── redis/          # Redis操作
│   ├── entity/             # 实体模型
│   ├── logger/             # 日志组件
│   ├── repo/               # 仓储层接口和实现
│   └── service/            # 业务逻辑层
├── cmd/                    # 命令行工具
├── configs/                # 配置文件
├── docs/                   # 文档
├── test/                   # 测试文件
├── web/                    # 前端静态资源
└── bin/                    # 编译输出目录
```

## 🏗️ 架构设计

### 分层架构

本项目采用经典的分层架构模式，从外到内分为：

1. **API Layer (api/)**
   - **Controller**: 处理HTTP请求，参数验证，响应格式化
   - **Middleware**: 认证、日志、错误处理、链路追踪等横切关注点
   - **Router**: 路由配置和注册

2. **Service Layer (internal/service/)**
   - 业务逻辑实现
   - 业务规则验证
   - 事务管理

3. **Repository Layer (internal/repo/)**
   - 数据访问抽象
   - 接口定义与实现分离
   - 支持多种数据源

4. **Data Access Layer (internal/dao/)**
   - 具体的数据访问实现
   - 缓存策略
   - 数据库操作

### 核心组件

#### 配置管理 (internal/core/)
- 使用 Viper 进行配置管理
- 支持多环境配置（开发/测试/生产）
- 热配置重载

#### 数据访问 (internal/dao/)
- **Database**: GORM ORM 框架，支持 MySQL/PostgreSQL
- **Cache**: Redis + BigCache 双重缓存
- **Connection Pool**: 数据库连接池管理

#### 身份认证 (internal/service/auth.go)
- JWT Token 生成和验证
- 用户登录状态管理
- 权限控制

#### 日志系统 (internal/logger/)
- 基于 Zap 的高性能日志
- 日志轮转和归档
- 结构化日志输出

## 🛠️ 技术栈

### 核心框架
- **[Gin](https://gin-gonic.com/)**: 高性能 HTTP Web 框架
- **[GORM](https://gorm.io/)**: Go 语言 ORM 库
- **[Viper](https://github.com/spf13/viper)**: 配置管理
- **[Cobra](https://cobra.dev/)**: CLI 命令行工具

### 数据存储
- **MySQL**: 主数据库
- **PostgreSQL**: 可选数据库
- **Redis**: 缓存和会话存储
- **BigCache**: 内存缓存

### 工具库
- **[Zap](https://github.com/uber-go/zap)**: 高性能日志库
- **[JWT](https://github.com/golang-jwt/jwt)**: JSON Web Token
- **[UUID](https://github.com/hashicorp/go-uuid)**: UUID 生成

## 🚀 快速开始

### 环境要求

- Go 1.24.3+
- MySQL 8.0+ 或 PostgreSQL 12+
- Redis 6.0+
- Docker & Docker Compose (可选)

### 本地开发

1. **克隆项目**
```bash
git clone https://github.com/liaozhonghui/go-mall-backend.git
cd go-mall-backend
```

2. **安装依赖**
```bash
go mod download
```

3. **配置数据库**
```bash
# 执行数据库初始化脚本
mysql -u root -p < docs/basic.sql
```

4. **修改配置文件**
```bash
cp configs/config.yaml configs/config.local.yaml
# 编辑 configs/config.local.yaml 配置数据库连接信息
```

5. **编译和运行**
```bash
# 开发模式编译
make dev

# 运行 Web 服务
./bin/main web -c configs/config.local.yaml
```

### Docker 部署

1. **使用 Docker Compose**
```bash
# 构建并启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f mall-backend
```

2. **单独构建 Docker 镜像**
```bash
# 构建镜像
docker build -t go-mall-backend .

# 运行容器
docker run -d \
  --name mall-backend \
  -p 9080:9080 \
  -v $(pwd)/configs:/app/configs \
  go-mall-backend
```

## 📖 API 文档

### 公共 API (需要认证)

| 方法 | 路径 | 描述 |
|------|------|------|
| ANY | `/api/healthCheck` | 健康检查 |
| ANY | `/api/healthCheckV1` | 健康检查 V1 |
| PUT | `/api/users` | 更新用户信息 |
| POST | `/api/login` | 用户登录 |

### 管理后台 API

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/admin/*` | 管理后台路由 (待完善) |

### 认证说明

- 除登录接口外，所有 `/api/*` 路径都需要 JWT Token 认证
- Token 通过 `Authorization: Bearer <token>` 头部传递
- Token 默认有效期为 24 小时

## 🔧 配置说明

### 主要配置项

```yaml
server:
  addr: 0.0.0.0:9080          # 服务监听地址
  readTimeOut: 3s             # 读超时
  writeTimeOut: 3s            # 写超时
  idleTimeOut: 100s           # 空闲超时

mysql:
  - instance: default
    dsn: "root:password@tcp(localhost:3306)/mall?charset=utf8mb4&loc=Local&parseTime=True"
    trace_log: true           # SQL 执行日志
    slow_threshold: 100       # 慢查询阈值(ms)

redis:
  addr: localhost:6379        # Redis 地址
  password: ""                # Redis 密码
  db: 0                       # Redis 数据库

jwt:
  api_secret: "mall_api"      # API JWT 密钥
  admin_secret: "mall_admin"  # 管理后台 JWT 密钥
  expireTime: 86400           # Token 过期时间(秒)

logger:
  logFile: logs/mall.log      # 日志文件路径
  logLevel: debug             # 日志级别
```

## 🧪 测试

### 运行测试
```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./test/

# 运行测试并显示覆盖率
go test -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### 测试结构
- `test/`: 集成测试和端到端测试
- `*_test.go`: 单元测试文件

## 📝 开发规范

### 代码风格
- 遵循 Go 官方代码规范
- 使用 `gofmt` 和 `golint` 进行代码格式化
- 变量和函数使用驼峰命名法
- 包名使用小写字母

### 目录结构规范
- `api/`: 对外暴露的 API 接口
- `internal/`: 内部业务逻辑，不对外暴露
- `cmd/`: 命令行工具和程序入口
- `configs/`: 配置文件
- `docs/`: 项目文档
- `test/`: 测试文件

### Git 提交规范
```
feat: 新功能
fix: 修复bug
docs: 文档更新
style: 代码格式调整
refactor: 代码重构
test: 测试相关
chore: 构建过程或辅助工具的变动
```

## 🤝 贡献指南

1. Fork 本仓库
2. 创建特性分支: `git checkout -b feature/new-feature`
3. 提交更改: `git commit -am 'Add new feature'`
4. 推送分支: `git push origin feature/new-feature`
5. 提交 Pull Request

## 📄 许可证

本项目采用 MIT 许可证。详情请参见 [LICENSE](LICENSE) 文件。

## 📞 联系方式

如有问题或建议，请通过以下方式联系：

- 项目地址: [https://github.com/liaozhonghui/go-mall-backend](https://github.com/liaozhonghui/go-mall-backend)
- Issues: [提交问题](https://github.com/liaozhonghui/go-mall-backend/issues)

---

⭐ 如果这个项目对你有帮助，请给个 Star 支持一下！
