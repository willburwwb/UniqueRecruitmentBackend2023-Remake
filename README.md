# UniqueRecruitmentBackend2023-Remake

This repository houses the source code for the backend of the Unique Studio recruitment system

### 📚 External Packages

- gin - A web framework written in Go
- gorm - A powerful ORM library for handling relational databases
- gRPC - A high performance, open source universal RPC framework
- swag - Automatically generates Swagger documentation from Go annotations
- zapx - A wrapper of zap , get from https://github.com/xylonx/zapx.git


------

### 🗂️ Directory Structure

```bash
uniqueRecruitmentBackend2023-Remake
├── configs
├── docs
├── global
├── internal
│   ├── cmd
│   ├── common
│   ├── controllers
│   ├── middlewares
│   ├── models
│   ├── router
│   ├── tracer
│   └── utils
├── pkg
│   ├── grpc
│   ├── logger
│   ├── proto
│   ├── sms
│   ├── sso
│   ├── constants.go
│   └── type.go
├── config.local.yaml
├── docker-compose.yml
├── Dockerfile
└── main.go
```



------

###  📝**Todo list:** 

- [X] Add swagger annotations to the interfaces
- [ ] Implement a new interface for updating user information on SSO
- [ ] Collaborate with the frontend for alignment
- [ ] Solve psql time zone problem
- [ ] Connect loki service to recruitment

------

### 🔑**Note:** 

