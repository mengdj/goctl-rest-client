package world

import (
	"net/http"

	"github.com/mengdj/goctl-rest-client/example/test/internal/logic/world"
	"github.com/mengdj/goctl-rest-client/example/test/internal/svc"
	"github.com/mengdj/goctl-rest-client/example/test/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := world.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
