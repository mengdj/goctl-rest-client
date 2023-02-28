package exa

import (
	"net/http"

	"github.com/mengdj/goctl-rest-client/examples/exa1/internal/logic/exa"
	"github.com/mengdj/goctl-rest-client/examples/exa1/internal/svc"
	"github.com/mengdj/goctl-rest-client/examples/exa1/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := exa.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
