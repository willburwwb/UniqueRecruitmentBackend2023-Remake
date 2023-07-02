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

------

### Postgresql

##### How to export database schema from postgresql 

- ssh connect to remote server
- `docker exec -it db_postgres bash       `
- `pg_dump -U postgres -s recruitment`  
  - dump the postgres database detail (tables,types,indexs...)  

- then get the SQL file about recruitment
- `psql -d recruitment_dev -U postgres -f filepath`  
  - import SQL file to database


##### Delete table and its dependences


- `drop table applications cascade;`
  ​	

