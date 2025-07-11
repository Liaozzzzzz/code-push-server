# 开发环境使用指南

本项目提供了多种运行方式来改善开发体验，避免每次都需要 `go install` 再执行。

## 快速开始

### 方式一：使用批处理文件（推荐 Windows 用户）

直接双击运行：

- `dev.bat` - 启动开发环境
- `build.bat` - 构建项目
- `install-air.bat` - 安装 Air 热重载工具

### 方式二：使用 Makefile（推荐 Linux/Mac 用户）

```bash
# 查看所有可用命令
make help

# 直接运行（推荐）
make run

# 构建项目
make build

# 安装并运行
make install-run

# 格式化代码
make fmt

# 运行测试
make test
```

### 方式三：使用 go run 命令

```bash
# 开发环境
go run main.go start -c configs -e dev

# 生产环境
go run main.go start -c configs -e prod

# 测试环境
go run main.go start -c configs -e test
```

## 热重载开发（可选）

如果你想要代码自动重载功能：

1. 安装 Air 工具：

   ```bash
   go install github.com/cosmtrek/air@latest
   ```

   或者运行 `install-air.bat`

2. 启动热重载：
   ```bash
   air
   ```
   或者 `make dev`

## 项目结构

```
.
├── main.go              # 主入口文件
├── cmd/                 # 命令行相关
├── internal/            # 内部包
├── configs/             # 配置文件
├── Makefile            # Make 命令
├── dev.bat             # Windows 开发脚本
├── build.bat           # Windows 构建脚本
├── install-air.bat     # Windows Air 安装脚本
└── .air.toml           # Air 配置文件
```

## 常用命令对比

| 功能     | 传统方式                                                 | 新方式                          |
| -------- | -------------------------------------------------------- | ------------------------------- |
| 开发运行 | `go install && code-push-server start -c configs -e dev` | `make run` 或双击 `dev.bat`     |
| 构建项目 | `go build -o bin/code-push-server main.go`               | `make build` 或双击 `build.bat` |
| 热重载   | 无                                                       | `air` 或 `make dev`             |
| 格式化   | `go fmt ./...`                                           | `make fmt`                      |
| 测试     | `go test ./...`                                          | `make test`                     |

## 配置说明

- 默认使用 `configs` 目录下的配置文件
- 默认环境为 `dev`
- 可以通过命令行参数修改：
  ```bash
  go run main.go start -c /path/to/configs -e production
  ```

## 注意事项

1. 确保 Go 环境已正确安装
2. 首次运行前请确保依赖已安装：`go mod download`
3. 如果使用热重载，请先安装 Air 工具
4. Windows 用户可以直接使用 `.bat` 文件，无需安装 Make
