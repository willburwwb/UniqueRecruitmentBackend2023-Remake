# Record

写下这个markdown的初衷在于记录开发hr系统中的一些问题，方便后续同学维护，后续会整理移交到飞书上。

## 部署

#### 本地部署：

先创建docker网络，用于连接数据库与后端容器`docker network create db `

##### postgres

- 建议先运行一个postgres容器，你可以命令行启动:

  ```bash
  # windows:
  docker run -d `
      --name db_postgres `
      -p 5432:5432 `
      -v ${your_local_file_path}:/var/lib/postgresql/data `
      -e POSTGRES_USER=postgres `
      -e POSTGRES_PASSWORD="" `
      -e POSTGRES_DB=recruitment_dev `
      --network db `
      postgres
  ```

  也可以选择`docker compose -f ./Docker-compose.withdb.yml up db_postgres`，这里出于方便我没有给db_postgres设置密码，假如出现postgres与backend容器出现密码认证的问题，可以将本地postgres 中`/var/lib/postgresql/data/pg_hda.conf`中最后一行改为为`host all all all trust`，即不通过密码认证（不够安全，仅适用于本地）

##### redis

redis容器同理

```bash
docker run -d `
    --name db_redis `
    -p 6379:6379 `
    -v D:\\service\\db_redis:/data `
    --restart always `
    --expose 6379 `
    --network db `
    redis:latest `
    redis-server --requirepass your_password
```

最后`docker compose up `运行backend容器

## SSO

hr系统与sso系统之间存在网络通信，hr系统从sso中获取用户信息和用户权限，目前hr与sso通过http协议进行通讯。

SSO RBAC机制：

- user：用户
- role：角色，其中每个用户可以拥有一个或多个角色
- object：用于表示一个具体权限，从属于object group
- object group：用于表示某些权限的集合，其中一个角色可以拥有一个或多个object group

hr系统只需要知道用户角色，即admin/member/candidate，后续开发hackday等其他系统时可以考虑增添其他角色。

**请求需要携带cookie**

> ​	POST /rbac/user/check_permission_by_role

```json
{
    "uid":"",
    "role":"admin",
}
```

用于检测用户是否具有权限

> GET /rbac/user?uid=

获取用户详细信息(包含权限)



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

