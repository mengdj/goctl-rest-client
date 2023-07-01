package hello

import (
	"context"
	"github.com/mengdj/goctl-rest-client/examples/exa1/client"
	"github.com/mengdj/goctl-rest-client/examples/exa1/client/exa"
	"github.com/mengdj/goctl-rest-client/examples/test/internal/svc"
	"github.com/mengdj/goctl-rest-client/examples/test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	exac   exa.Client
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		exac:   exa.MustClient(svcCtx.Config.HelloDiscoverConf),
	}
}

func (l *HelloLogic) Hello(req *types.HelloRequest) (*types.Response, error) {
	resp, err := l.exac.Hello(l.ctx, &client.HelloRequest{
		Msg: "hello",
	})
	if nil != err {
		logx.Error("请求错误", err)
	} else {
		logx.Info(resp)
	}
	return nil, nil
}
