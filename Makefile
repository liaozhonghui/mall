# 检测操作系统类型
ifeq ($(OS),Windows_NT)
    BINARY_EXT := .exe
    RM := del /Q
    RM_DIR := rmdir /S /Q
else
    BINARY_EXT := 
    RM := rm -f
    RM_DIR := rm -rf
endif

DEFAULT := ./cmd
OUTPUT := ./bin/mall$(BINARY_EXT)

default: build

build: 
	go build -gcflags "-N -l" -o $(OUTPUT) $(DEFAULT)

clean:
ifeq ($(OS),Windows_NT)
	if exist .\bin\mall.exe del /Q .\bin\mall.exe
else
	$(RM) ./bin/mall$(BINARY_EXT)
endif

.PHONY: build clean default