# UniqueRecruitmentBackend2023-Remake

Backend of recruitment system for Unique Studio 

------

### External packages in project 

- gin
- gorm
- go-redis
- zap + lumberjack
  - log 
- swag : converts 
  - Go annotations to Swagger Documentation
- viper: 
  - configuration management 

**Get user's uid from http header field "X-UID"**

------

### Directory Structure

```bash
uniqueRecruitmentBackend2023-Remake
├── configs
├── docs
├── global
├── internal
│   ├── constants
│   ├── controllers
│   ├── middlewares
│   ├── models
│   ├── request
│   ├── response
│   └── router
├── pkg
│   ├── msg
│   └── utils
├── config.yaml
├── Docker-compose.yml
├── Dockerfile
└── main.go
```

<br>
<br>
<br>

**<h2>暑假待办 （尽量不鸽</h2>**
- 尽快与unique SSO连接，方便后续的工作！
- 调整model tables的结构，与sso user/服务器上的postgres适配
- 开始完成api工作