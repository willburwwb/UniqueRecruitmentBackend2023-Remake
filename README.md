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

### Docker

#### local:

1. Create  container `unique_recruitment_psql_dev`  by image `postgres:latest` , add password `mysecretpassword` , create network `uniquerecruitmentbackend2023-remake_database` which connect to `recruitment_backend ` , create volume. 

   ```bash
   docker run -p 5430:5432 --name unique_recruitment_psql_dev -v "your file path :/var/lib/postgresql/data"  --network uniquerecruitmentbackend2023-remake_database -e POSTGRES_PASSWORD=mysecretpassword postgres:latest`
   
   docker exec -it unique_recruitment_psql_dev env # cat psql env 
   docker exec -it unique_recruitment_psql_dev psql -U postgres # attach container
   ```

   

2. Build compose

   ```bash
   docker-compose -f Docker-compose.local.yml build
   ```

3. Run service

   ```bash
   docker-compose -f Docker-compose.local.yml up
   ```

   



### docker-compose

- local

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

