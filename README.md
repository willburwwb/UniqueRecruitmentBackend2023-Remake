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




**!!  Get user's uid from http header field "X-UID"**

------

### Docker-compose

#### local:

1. Create  container `unique_recruitment_psql_dev`  by image `postgres:latest` , add password `mysecretpassword` , create network `uniquerecruitmentbackend2023-remake_database` which connect to `recruitment_backend ` , create volume. 

   ```bash
   docker run -p 5430:5432 --name unique_recruitment_psql_dev -v "your file path :/var/lib/postgresql/data"  --network uniquerecruitmentbackend2023-remake_database -e POSTGRES_PASSWORD=mysecretpassword postgres:latest`
   docker exec -it unique_recruitment_psql_dev env  # cat psql env
   docker exec -it unique_recruitment_psql_dev psql -U postgres # attach container
   create database 
   ```
   
2. Build compose

   ```bash
   docker-compose -f Docker-compose.local.yml build
   ```

3. Run service

   ```bash
   docker-compose -f Docker-compose.local.yml up
   ```


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