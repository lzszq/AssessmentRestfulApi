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

return：
```
{
    "token": "token"
}
```

## Users

① 创建用户

url: /users

method: POST

```
{
    "nickname": "nickname",
    "email": "email",
    "password": "password",
    "avatar_url": "avatar_url"
}
```

return：
```
{
    "id": id,
    "nickname": "nickname",
    "email": "email",
    "password": "",
    "avatar_url": "avatar_url",
    "created_at": "created_at",
    "updated_at": "updated_at"
}
```

② 获取用户列表

url: /users

method: GET

return：
```
[
    {
        "id": id,
        "nickname": "nickname",
        "email": "email",
        "password": "",
        "avatar_url": "avatar_url",
        "created_at": "created_at",
        "updated_at": "updated_at"
    },
    {
        "id": id,
        "nickname": "nickname",
        "email": "email",
        "password": "",
        "avatar_url": "avatar_url",
        "created_at": "created_at",
        "updated_at": "updated_at"
    }
]
```

③ 获取单用户信息

url: /users/{id}

method: GET

return：
```
{
    "id": id,
    "nickname": "nickname",
    "email": "email",
    "password": "",
    "avatar_url": "avatar_url",
    "created_at": "created_at",
    "updated_at": "updated_at"
}
```


④ 更新用户信息

header 须带 Bearer Token

url: /users/{id}

method: PUT

```
{
    "nickname": "nickname",
    "email": "email",
    "password": "password",
    "avatar_url": "avatar_url"
}
```

return：
```
{
    "id": id,
    "nickname": "nickname",
    "email": "email",
    "password": "",
    "avatar_url": "avatar_url",
    "created_at": "created_at",
    "updated_at": "updated_at"
}
```

⑤ 删除对应用户

header 须带 Bearer Token

url: /users/{id}

method: DELETE


## Posts

① 创建文章

header 须带 Bearer Token

url: /posts

method: POST

```
{
    "title": "title",
    "author_id": author_id,
    "content": "content"
}
```

return：
```
{
    "id": id,
    "title": "title",
    "content": "content",
    "author": {
        "id": id,
        "nickname": "nickname",
        "email": "email",
        "password": "",
        "avatar_url": "avatar_url",
        "created_at": "created_at",
        "updated_at": "updated_at"
    },
    "author_id": author_id,
    "created_at": "created_at",
    "updated_at": "updated_at"
}
```

② 获取文章列表

url: /posts

method: GET

return：
```
[
    {
        "id": id,
        "title": "title",
        "content": "content",
        "author": {
            "id": id,
            "nickname": "nickname",
            "email": "email",
            "password": "",
            "avatar_url": "avatar_url",
            "created_at": "created_at",
            "updated_at": "updated_at"
        },
        "author_id": author_id,
        "created_at": "created_at",
        "updated_at": "updated_at"
    },
    {
        "id": id,
        "title": "title",
        "content": "content",
        "author": {
            "id": id,
            "nickname": "nickname",
            "email": "email",
            "password": "",
            "avatar_url": "avatar_url",
            "created_at": "created_at",
            "updated_at": "updated_at"
        },
        "author_id": author_id,
        "created_at": "created_at",
        "updated_at": "updated_at"
    }
]
```

③ 获取单文章信息

url: /posts/{id}

method: GET

return：
```
{
    "id": id,
    "title": "title",
    "content": "content",
    "author": {
        "id": id,
        "nickname": "nickname",
        "email": "email",
        "password": "",
        "avatar_url": "avatar_url",
        "created_at": "created_at",
        "updated_at": "updated_at"
    },
    "author_id": author_id,
    "created_at": "created_at",
    "updated_at": "updated_at"
}
```


④ 更新文章信息

url: /posts/{id}

method: PUT

```
{
    "title": "title",
    "author_id": author_id,
    "content": "content"
}
```

return：
```
{
    "id": id,
    "title": "title",
    "content": "content",
    "author": {
        "id": id,
        "nickname": "nickname",
        "email": "email",
        "password": "",
        "avatar_url": "avatar_url",
        "created_at": "created_at",
        "updated_at": "updated_at"
    },
    "author_id": author_id,
    "created_at": "created_at",
    "updated_at": "updated_at"
}
```

⑤ 删除对应文章

header 须带 Bearer Token

url: /posts/{id}

method: DELETE


## Images

① 上传图片

url: /upload

method: POST

```
<form method="post" enctype="multipart/form-data">
    <input type="file" name="image" />
    <input type="submit" />
</form>
```

return：
```
{
    "url": "/uploaded/{filename}"
}
```

② 获取图片

url: /uploaded/{filename}

method: GET

return：
```
image
```