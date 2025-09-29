# 使用更小的 Go 基础镜像
FROM golang:1.24-alpine AS builder

# 只安装必需的包
RUN apk add --no-cache ca-certificates make

WORKDIR /app

# 利用 Docker 缓存层，先复制 go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY configs/ ./configs/
COPY api/ ./api/
COPY Makefile ./

# 构建应用
RUN make prod

# 使用 scratch 作为最终镜像（最小化）
FROM scratch AS production

# 从 builder 阶段复制 CA 证书（用于 HTTPS 请求）
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 创建必要的目录结构
COPY --from=builder /app/bin/main /main
COPY --from=builder /app/configs/config.yaml /configs/config.yaml

ENTRYPOINT ["/main"]
CMD ["web", "--config", "configs/config.yaml"]
