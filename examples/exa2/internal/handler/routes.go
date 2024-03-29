// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	exa "github.com/mengdj/goctl-rest-client/examples/exa2/internal/handler/exa"
	"github.com/mengdj/goctl-rest-client/examples/exa2/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/app/demo/hello",
				Handler: exa.HelloHandler(serverCtx),
			},
		},
		rest.WithTimeout(30000*time.Millisecond),
	)
}
