## Quick Start

### 1. Setup Basic Dependence

```shell
docker-compose up
```

### 2. Run User RPC Server

```shell
cd cmd/user
bash build.sh
bash output/bootstrap.sh
```

### 3. Run API Server

```shell
cd cmd/user
bash build.sh
bash output/bootstrap.sh
```

### 4. Jaeger

`http://127.0.0.1:16686/`

## APT requests

### Register

```shell
# 正常注册
curl --location --request POST '127.0.0.1:8080/douyin/user/register' --header 'Content-Type: application/json' --data-raw '{"username":"Stone","password":"123456"}'
# 用户已存在
curl --location --request POST '127.0.0.1:8080/douyin/user/register' --header 'Content-Type: application/json' --data-raw '{"username":"Stone","password":"222222"}'
# 用户名或密码为空
curl --location --request POST '127.0.0.1:8080/douyin/user/register' --header 'Content-Type: application/json' --data-raw '{"username":"bravozyz","password":""}'
```

**response**

```text
# 正常注册
{"status_code":0,"status_msg":"Success","user_id":6,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NiwiQXV0aG9yaXR5SWQiOjB9.9yviB30k0NWgxyyp4wt7wGoqRc36Ea9tnAc3ajn0V2w"}
# 用户已存在
{"status_code":10005,"status_msg":"User already exists","user_id":-1,"token":""}
# 账号或密码为空
{"status_code":10002,"status_msg":"Wrong Parameter has been given","user_id":-1,"token":""}{"status_code":10002,"status_msg":"Wrong Parameter has been given","user_id":-1,"token":""}
```

### Login

```shell
# 正确登录
curl --location --request POST '127.0.0.1:8080/douyin/user/login' --header 'Content-Type: application/json' --data-raw '{"username":"Stone","password":"123456"}'
# 账号或密码为空
curl --location --request POST '127.0.0.1:8080/douyin/user/login' --header 'Content-Type: application/json' --data-raw '{"username":"Stone","password":""}'
# 用户不存在
curl --location --request POST '127.0.0.1:8080/douyin/user/login' --header 'Content-Type: application/json' --data-raw '{"username":"Rock","password":"123456"}'
# 密码错误
curl --location --request POST '127.0.0.1:8080/douyin/user/login' --header 'Content-Type: application/json' --data-raw '{"username":"Stone","password":"222222"}'
```

**response**

```text
# 正确登录
{"status_code":0,"status_msg":"Success","user_id":6,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NiwiQXV0aG9yaXR5SWQiOjB9.9yviB30k0NWgxyyp4wt7wGoqRc36Ea9tnAc3ajn0V2w"}
# 账号或密码为空
{"status_code":10002,"status_msg":"Wrong Parameter has been given","user_id":-1,"token":""}{"status_code":10002,"status_msg":"Wrong Parameter has been given","user_id":-1,"token":""}
# 用户不存在
{"status_code":10004,"status_msg":"User does not exists","user_id":-1,"token":""}
# 密码错误
{"status_code":10003,"status_msg":"Wrong username or password","user_id":-1,"token":""}
```

### GetUserInfo

```shell
curl --location --request GET '127.0.0.1:8080/douyin/user/' --header 'Content-Type: application/json' --data-raw '{"user_id":"6","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NiwiQXV0aG9yaXR5SWQiOjB9.9yviB30k0NWgxyyp4wt7wGoqRc36Ea9tnAc3ajn0V2w"}'
```