# 调试指南

本项目支持多种调试方式，以下是详细的使用说明。

## 准备工作

首先安装必要的调试工具：

```bash
make install-tools
```

这会安装：

- `air` - 热重载工具
- `dlv` - Delve 调试器

## 调试方式

### 1. 直接调试模式

启动支持 delve 调试器的服务：

```bash
make debug
```

这会：

- 启动 delve 调试器
- 监听端口 `:2345`
- 支持多客户端连接
- 使用开发环境配置

### 2. 热重载调试模式

启动支持热重载的调试服务：

```bash
make debug-air
```

这会：

- 使用 air 进行热重载
- 自动重新编译并启动 delve 调试器
- 文件变更时自动重启

### 3. VS Code 调试

在 VS Code 中有三种调试配置：

#### 3.1 直接调试 (Debug Server)

- 按 `F5` 或使用调试面板
- 选择 "Debug Server" 配置
- 直接在 VS Code 中启动调试

#### 3.2 附加调试 (Attach to Debug Server)

- 先运行 `make debug` 或 `make debug-air`
- 在 VS Code 中选择 "Attach to Debug Server"
- 附加到已运行的调试器

#### 3.3 测试调试 (Debug Test)

- 用于调试单元测试
- 选择 "Debug Test" 配置

## 调试技巧

### 设置断点

- 在 VS Code 中点击行号左侧设置断点
- 或在代码中使用 `runtime.Breakpoint()`

### 查看变量

- 使用 VS Code 的变量面板
- 或在调试控制台中使用 delve 命令

### 常用 delve 命令

```bash
# 连接到调试器
dlv connect :2345

# 设置断点
break main.main
break internal/service/dept_service.go:39

# 继续执行
continue

# 单步执行
step

# 查看变量
print variableName

# 查看调用栈
stack

# 退出
quit
```

### 远程调试

如果需要远程调试，可以修改调试器监听地址：

```bash
dlv debug --headless --listen=0.0.0.0:2345 --api-version=2 --accept-multiclient -- start -c configs -e dev
```

## 配置说明

### 编译选项

调试模式使用以下编译选项：

- `-gcflags='all=-N -l'` - 禁用优化和内联，保留调试信息

### 环境配置

- 默认使用 `dev` 环境
- 可以通过修改命令参数更改环境
- 调试模式会自动启用详细日志

## 故障排除

### 端口冲突

如果端口 `:2345` 被占用：

1. 查找占用进程：`netstat -ano | findstr :2345`
2. 结束进程或更改端口

### 调试器连接失败

1. 确保防火墙允许端口访问
2. 检查调试器是否正常启动
3. 确认 VS Code Go 扩展已安装

### 断点不生效

1. 确保使用调试模式编译
2. 检查源码路径映射
3. 重新启动调试会话

## 性能分析

### CPU 分析

```bash
go tool pprof http://localhost:8040/debug/pprof/profile
```

### 内存分析

```bash
go tool pprof http://localhost:8040/debug/pprof/heap
```

### 查看运行时信息

```bash
curl http://localhost:8040/debug/pprof/
```

## 日志调试

开发环境默认启用详细日志：

- 日志级别：`debug`
- 输出格式：`json`
- 输出位置：`stdout` 和 `logs/dev.log`

可以通过修改 `configs/config.dev.toml` 中的日志配置来调整日志行为。
