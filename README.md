# MentalHealth
毕设：心理健康平台

# MentalHealth-Platform API Documentation

## 概述

MentalHealth-Platform 是一个用于心理健康管理的平台，旨在帮助用户进行健康数据的管理、问卷调查、数据分析和预测等。本项目使用 Go 语言编写，并通过 Gin 框架提供 RESTful API 服务。

## 目录结构

项目目录结构如下：

```
MentalHealth-Platform/
├── app/
│   ├── logic/
│   │   └── ... (业务逻辑处理文件)
│   ├── middleware/
│   │   └── ... (中间件文件)
│   ├── model/
│   │   └── ... (数据模型文件)
│   └── view/
│       └── ... (HTML模板文件)
├── router/
│   └── router.go (路由配置文件)
└── main.go (主函数入口文件)
```

## 安装和运行

### 依赖安装

确保您已经安装了 Go 语言环境。然后，安装项目依赖：

```sh
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

### 运行项目

在项目根目录下运行以下命令启动服务：

```sh
go run main.go
```

或者，您也可以直接运行 `router.go` 文件：

```sh
go run router/router.go
```

服务将默认在端口 8087 上启动。

## API 路由

以下列出了项目中定义的 API 路由及其功能描述：

### 登录和用户管理

- `POST /user/login`
  - 描述：用户登录
  - 参数：用户信息
- `POST /user/email-login`
  - 描述：通过邮箱登录
  - 参数：邮箱信息
- `POST /user/update-login-status`
  - 描述：更新登录状态
- `GET /user/GetToken`
  - 描述：获取用户 Token
- `POST /user/SendEmailCaptcha`
  - 描述：发送邮箱验证码
- `POST /user/SendSMSCaptcha`
  - 描述：发送短信验证码
- `POST /user/VerifySMSCaptcha`
  - 描述：验证短信验证码
- `GET /user/wechat`
  - 描述：检查微信签名
- `GET /user/wechat/login`
  - 描述：微信登录跳转
- `GET /user/wechat/Callback`
  - 描述：微信登录回调
- `GET /user/wechat/check_login`
  - 描述：检查微信登录状态

### 问卷调查

- `POST /user/submitTeenSurvey`
  - 描述：提交青少年问卷表单
- `POST /user/submitCollegeSurvey`
  - 描述：提交大学生问卷表单
- `POST /user/submitWorkerSurvey`
  - 描述：提交社会工作者问卷表单
- `GET /surveys/SurveysFind`
  - 描述：根据用户ID查找所有问卷

### 健康预测

- `POST /pre/predict`
  - 描述：进行健康预测
- `GET /pre/teenPredict`
  - 描述：青少年健康预测
- `GET /pre/collegePredict`
  - 描述：大学生健康预测
- `GET /pre/workerPredict`
  - 描述：社会工作者健康预测

### 其他功能

- `GET /swagger/*any`
  - 描述：访问 Swagger API 文档
- `GET /captcha`
  - 描述：获取图片验证码
- `POST /captcha/verify`
  - 描述：验证图片验证码

## 注意事项

- 请确保您的数据库连接配置正确。
- 在生产环境中，您可能需要配置 HTTPS 以保证安全性。
- 请确保所有的依赖项都已经安装并配置正确。

## 贡献指南

如果您希望为本项目贡献代码或文档，请遵循以下步骤：

1. Fork 本仓库。
2. 创建您的功能分支：`git checkout -b my-new-feature`
3. 提交您的更改：`git commit -am 'Add some feature'`
4. 推送您的分支：`git push origin my-new-feature`
5. 创建一个新 Pull Request。

## 许可证

本项目采用 [Apache License 2.0](./LICENSE) 许可证。

```