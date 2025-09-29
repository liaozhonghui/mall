DEFAULT := ./cmd
OUTPUT := ./bin/main

# 自动检测当前系统
GOOS := $(strip $(shell go env GOOS))

# Windows 平台输出 .exe 文件
ifeq ($(GOOS),windows)
	OUTPUT := ./bin/main.exe
endif

# 默认目标
default:
	go build -gcflags "-N -l" -o $(OUTPUT) $(DEFAULT)

# 开发模式构建（禁用优化，内联）
dev:
	go build -gcflags "-N -l" -o $(OUTPUT) $(DEFAULT)

# 生产模式构建（优化大小, 体积能减少一半）
prod:
	go build -a -installsuffix cgo -ldflags="-w -s" -o $(OUTPUT) $(DEFAULT)

# 清理构建产物
clean:
ifeq ($(GOOS),windows)
	-rmdir /s /q bin 2>nul
else
	rm -rf ./bin
endif