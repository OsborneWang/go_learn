## 项目简介

本示例项目演示了如何使用 **Gin + GORM + gomail** 搭建一个简单且结构清晰的用户注册/登录服务端。适合刚接触 Go 的同学阅读学习。

```
goServer
├── cmd/server/main.go          // 程序入口
├── internal
│   ├── config/config.go        // 配置加载
│   ├── database/database.go    // GORM 初始化
│   ├── models/user.go          // User 模型及密码方法
│   ├── services
│   │   ├── auth_service.go     // 注册登录业务逻辑
│   │   └── mail_service.go     // 邮件发送封装
│   ├── handlers/auth_handler.go// HTTP Handler
│   └── routes/routes.go        // Gin 路由注册
└── docs/README.md
```

## 配置

通过环境变量控制运行参数，常用变量如下：

| 变量名 | 默认值 | 说明 |
| --- | --- | --- |
| `APP_PORT` | 8080 | 启动端口 |
| `DB_HOST`/`DB_PORT`/`DB_USER`/`DB_PASSWORD`/`DB_NAME` | | MySQL 连接信息 |
| `MAIL_HOST`/`MAIL_PORT`/`MAIL_USERNAME`/`MAIL_PASSWORD` | | SMTP 信息 |
| `MAIL_FROM` | Demo Service <demo@example.com> | 邮件展示发件人 |

本地开发可复制 `docs/env.example` 为根目录下的 `.env` 并根据实际情况修改，也可以直接通过 shell 导出环境变量。

## 启动步骤

1. 启动或连接到一个 MySQL 实例，并创建数据库（如 `go_demo`）。
2. 设置环境变量（示例）：
   ```bash
   export DB_USER=root
   export DB_PASSWORD=yourpassword
   export DB_NAME=go_demo
   export MAIL_HOST=smtp.example.com
   export MAIL_USERNAME=demo@example.com
   export MAIL_PASSWORD=app-token
   ```
3. 安装依赖并启动：
   ```bash
   go mod tidy
   go run ./cmd/server
   ```

服务启动后默认监听 `http://localhost:8080`。

## API 测试

- 注册：
  ```bash
  curl -X POST http://localhost:8080/api/v1/register \
    -H "Content-Type: application/json" \
    -d '{"email":"user@example.com","password":"secret12","name":"Alice"}'
  ```
- 登录：
  ```bash
  curl -X POST http://localhost:8080/api/v1/login \
    -H "Content-Type: application/json" \
    -d '{"email":"user@example.com","password":"secret12"}'
  ```
- 发送测试邮件：
  ```bash
  curl -X POST http://localhost:8080/api/v1/mail/test \
    -H "Content-Type: application/json" \
    -d '{"to":"user@example.com","subject":"Hello","message":"这是一封测试邮件"}'
  ```

响应成功时返回用户基础信息；失败会返回 `error` 字段描述。

## 模块说明

- `config`：集中读取、校验配置，便于在 main 中统一初始化。
- `database`：封装 GORM 初始化及连接池配置。
- `models`：定义数据表结构及与业务强相关的实体方法（如密码加密）。
- `services`：编排业务流程。`AuthService` 与数据库交互并调用 `MailService`。
- `handlers`：处理请求参数、响应格式及错误码。
- `routes`：管理 Gin 路由与版本前缀，方便后续扩展更多模块。

> 通过这种分层方式，可以清晰地将配置、基础设施、业务逻辑与 HTTP 层解耦，方便日后扩展与维护。

