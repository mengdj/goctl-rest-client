package world

import (
	"context"
	"fmt"
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
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req *types.HelloRequest) (*types.Response, error) {
	var (
		ret = &types.Response{
			Code: 0,
			Msg:  "",
		}
		err  error = nil
		resp *client.Response
	)
	resp, err = exa.MustClient(l.svcCtx.Config.HelloDiscoverConf).Hello(l.ctx, &client.HelloRequest{
		Msg: fmt.Sprintf("recive id:%d", req.ID),
	})
	if nil != err {
		ret.Code = 1
		ret.Msg = err.Error()
	} else {
		ret.Code = resp.Code
		ret.Msg = resp.Msg
	}
	return ret, nil
}
