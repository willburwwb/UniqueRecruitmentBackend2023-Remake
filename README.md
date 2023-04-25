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



### How to import data from postgresql

#### example:

- ssh connect to remote server

- `docker exec -it db_postgres bash       `
- `pg_dump -U postgres -s recruitment`//dump the postgres database detail (tables,types,indexs...)  
- `psql -d recruitment_dev -U postgres -f filepath`//import sql file to database
- then get the sql file about applications


- delete table and its dependences
- `drop table applications cascade;`
​	
