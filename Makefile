.PHONY: dev build run clean install-tools fmt test lint help

# 默认目标
help:
	@echo "可用的命令："
	@echo "  dev          - 开发模式运行（热重载，需要先安装air）"
	@echo "  run          - 直接运行（不编译到文件）"
	@echo "  build        - 构建项目"
	@echo "  install-run  - 安装并运行"
	@echo "  clean        - 清理临时文件"
	@echo "  install-tools- 安装开发工具"
	@echo "  fmt          - 格式化代码"
	@echo "  test         - 运行测试"
	@echo "  lint         - 检查代码"

# 开发模式运行（热重载）
dev:
	@echo "启动开发模式（热重载）..."
	air

# 直接运行（不编译到文件）
run:
	@echo "直接运行项目..."
	go run main.go start -c configs -e dev

# 构建项目（跨平台兼容）
build:
	@echo "构建项目..."
ifeq ($(OS),Windows_NT)
	go build -o bin/code-push-server.exe main.go
else
	go build -o bin/code-push-server main.go
endif

# 安装并运行
install-run:
	@echo "安装并运行项目..."
	go install && code-push-server start -c configs -e dev

# 清理临时文件（跨平台兼容）
clean:
	@echo "清理临时文件..."
ifeq ($(OS),Windows_NT)
	rm -rf tmp/ bin/ build-errors.log
else
	@powershell -Command "if (Test-Path tmp) { Remove-Item -Recurse -Force tmp }"
	@powershell -Command "if (Test-Path bin) { Remove-Item -Recurse -Force bin }"
	@powershell -Command "if (Test-Path build-errors.log) { Remove-Item -Force build-errors.log }"
endif

# 安装开发工具
install-tools:
	@echo "安装开发工具..."
	go install github.com/air-verse/air@latest

# 格式化代码
fmt:
	@echo "格式化代码..."
	go fmt ./...

# 运行测试
test:
	@echo "运行测试..."
	go test ./...

# 检查代码（需要先安装golangci-lint）
lint:
	@echo "检查代码..."
	golangci-lint run

# 生产环境构建
build-prod:
	@echo "生产环境构建..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/code-push-server main.go

# 运行生产环境
run-prod:
	@echo "运行生产环境..."
	go run main.go start -c configs -e prod

# 运行测试环境
run-test:
	@echo "运行测试环境..."
	go run main.go start -c configs -e test 