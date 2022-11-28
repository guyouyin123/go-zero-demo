### 1. "获取用户信息"

1. route definition

- Url: /user/info
- Method: POST
- Request: `UserInfoRes`
- Response: `UserInfoReq`

2. request definition



```golang
type UserInfoRes struct {
	UserId int `json:"userId"`
}
```


3. response definition



```golang
type UserInfoReq struct {
	UserId int `json:"userId"`
	Name string `json:"name"`
}
```

