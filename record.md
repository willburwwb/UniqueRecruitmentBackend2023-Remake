# Record

写下这个markdown的初衷在于记录开发hr系统中的一些问题，方便后续同学维护，后续会整理移交到飞书上。

### GPRC

sso 

``` 
protoc --go_out="./internal" --go_opt=paths=source_relative \          
--go-grpc_out="./internal" --go-grpc_opt=paths=source_relative \          
proto/sso/sso.proto    
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