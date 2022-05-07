# Based on [THIS](https://github.com/victorsteven/Go-JWT-Postgres-Mysql-Restful-API)

# 使用方法

go build main.go

然后 ./main 运行

# 接口说明

## Login

①登录 

url: /login

method: POST

```
{
    "email": "email",
    "password": "password"
}
```

## Users

① 创建用户

url: /users

method: POST

```
{
    "email": "email",
    "password": "password"
}
```

