# goctl-rest-discover

* api文件描述

```api
syntax = "v1"

info(
    title: "type title here"
    desc: "type desc here"
    author: "type author here"
    email: "type email here"
    version: "type version here"
)

type (
    Response {
        Code uint32 `json:"code"`
        Msg string `json:"msg"`
    }
    HelloRequest {
        Msg string `json:"msg"`
    }
)

@server(
    group: exa
    timeout: 30s
)
service demo_api {
    @doc "xx"
    @handler Hello
    post /api/v1/app/demo/hello (HelloRequest) returns (Response)
}
```

* 生成基础框架

```shell
goctl api go -api exa.api -dir . --style go_zero
```

* 为api生成调用客户端(插件形式)

```shell
goctl api plugin -p goctl-rest-discover="rest-discover" -api exa.api -dir .
```

* 调用(通过服务名称调用)

```go
resp, err = client.MustClient(l.svcCtx.Config.HelloDiscoverConf).Hello(l.ctx, &client.HelloRequest{
Msg: "hello,rest",
})
```

调用api像rpc一样简单

