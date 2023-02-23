package exa

import (
	"context"

	"github.com/mengdj/goctl-rest-client/example/exa2/internal/svc"
	"github.com/mengdj/goctl-rest-client/example/exa2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req *types.HelloRequest) (resp *types.Response, err error) {
	return &types.Response{
		Code: 0,
		Msg:  "from exa2:" + req.Msg,
	}, nil
}
