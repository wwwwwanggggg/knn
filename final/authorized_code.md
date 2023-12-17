## 1 authorization_code

### 1-1 获取授权code

**请求方式**

`GET` `/oauth2/authorize`

**参数说明**  

|参数|类型|说明|
|-|-|-|
|client_id|string|在oauth2 server注册的client_id,见配置文件`oauth2.client.id`|
|response_type|string|固定值:`code`|
|scope|string|权限范围,如:`str1,str2,str3`,str为配置文件中`oauth2.client.scope.id`的值 |
|state|string|表示客户端的当前状态,可以指定任意值,认证服务器会原封不动地返回这个值|
|redirect_uri|string|回调uri,会在后面添加query参数`?code=xxx&state=xxx`,发放的code就在其中|

**请求示例**

```sh
# 浏览器请求
http://gitea.com/oauth2/authorize?client_id=test_client_1&response_type=code&scope=all&state=xyz&redirect_uri=http://localhost:9093/cb

# 302跳转,返回code
http://localhost:9093/callback?code=XUNKO4OPPROWAPFKEWNZWA&state=xyz
```

### 1-2 使用`code`交换`token`

**请求方式**

`POST` `/oauth2/token`

**请求头 Authorization**

- basic auth
- username: `client_id`
- password: `client_secret`

**Header**  
`Content-Type: application/x-www-form-urlencoded`

**Body参数说明**  

|参数|类型|说明|
|-|-|-|
|grant_type|string|固定值`authorization_code`|
|code|string| 1-1 发放的code|
|redirect_uri|string| 1-1 填写的redirect_uri|

**Response返回示例**  

```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIyMjIyMjIiLCJleHAiOjE1ODU3MTU1NTksInN1YiI6InRlc3QifQ.ZMgIDQMW7FGxbF1V8zWOmEkmB7aLH1suGYjhDdrT7aCYMEudWUoiCkWHSvBmJahGm0RDXa3IyDoGFxeMfzlDNQ",
    "expires_in": 7200,
    "refresh_token": "JG7_WGLWXUOW2KV2VLJKSG",
    "scope": "all",
    "token_type": "Bearer"
}
```
