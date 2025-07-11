# 配置文件目录说明

## 目录结构

```
configs/
├── README.md                 # 本说明文件
├── config.dev.toml          # 开发环境配置
├── config.prod.toml         # 生产环境配置
├── config.test.toml         # 测试环境配置
├── shared/                  # 共享配置文件
│   ├── database.toml        # 数据库相关配置
│   └── security.toml        # 安全相关配置
├── local/                   # 本地开发配置
│   └── override.toml        # 本地覆盖配置（不提交到版本控制）
└── examples/                # 配置示例文件
    ├── config.example.toml  # 完整配置示例
    └── docker.toml          # Docker 环境配置示例
```

## 配置文件加载顺序

系统会按照以下顺序自动加载配置文件：

1. **环境配置**: 根据指定环境加载对应的配置文件

   - 开发环境: `config.dev.toml`
   - 生产环境: `config.prod.toml`
   - 测试环境: `config.test.toml`

2. **共享配置**: 自动加载 `shared/` 目录下的所有配置文件

   - `shared/database.toml`
   - `shared/security.toml`

3. **本地覆盖**: 自动加载 `local/` 目录下的所有配置文件（如果存在）
   - `local/override.toml`

## 环境变量

生产环境配置支持环境变量替换，使用 `${VAR_NAME}` 格式：

```toml
[Security]
JWTSecret = "${JWT_SECRET}"
```

## 配置优先级

后加载的配置会覆盖先加载的配置，优先级从低到高：

1. 环境配置文件 (`config.{env}.toml`)
2. 共享配置文件 (`shared/*.toml`)
3. 本地覆盖配置 (`local/*.toml`)

## 使用方法

### 启动服务器

```bash
# 使用开发环境配置
./code-push-server start -e dev

# 使用生产环境配置
./code-push-server start -e prod

# 使用测试环境配置
./code-push-server start -e test
```

### 配置文件格式

支持 JSON 和 TOML 格式：

- `.toml` 文件使用 TOML 格式
- `.json` 文件使用 JSON 格式

## 最佳实践

1. **环境隔离**: 不同环境使用不同的配置文件
2. **敏感信息**: 生产环境使用环境变量存储敏感信息
3. **本地开发**: 使用 `local/override.toml` 进行本地定制
4. **版本控制**:
   - 提交环境配置和共享配置
   - 不提交 `local/` 目录下的文件
   - 提供 `examples/` 目录作为参考

## 注意事项

- `local/` 目录下的配置文件不应提交到版本控制系统
- 生产环境的敏感配置应通过环境变量或密钥管理系统提供
- 修改配置后需要重启服务才能生效
- 系统会自动按文件名字母顺序加载同一目录下的配置文件
