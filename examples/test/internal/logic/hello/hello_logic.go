package hello

import (
	"context"
	"github.com/mengdj/goctl-rest-client/examples/test/internal/svc"
	"github.com/mengdj/goctl-rest-client/examples/test/internal/types"

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

func (l *HelloLogic) Hello(req *types.HelloRequest) (*types.Response, error) {
	return nil, nil
}
