# UniqueRecruitmentBackend2023-Remake

Backend of recruitment system for Unique Studio 

------

### External packages in project 

- gin
- gorm
- go-redis
- zap + lumberjack
- 
  - log 
- swag : converts 
  - Go annotations to Swagger Documentation
- viper: 
- session: 
  - Use redis store the session.
  - First, the SSO system stores the session of the login users.
  - Second, the backend system get the session through middleware.

### docker 

- build image 
  - `docker build -t unique_backend2023 .`
- run container
  - `docker run -p 8080:3333 --name unique_backend_test unique_backend2023:latest  ` 



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
│   └── server
├── pkg
├── config.yaml
├── Docker-compose.yml
├── Dockerfile
└── main.go
```

