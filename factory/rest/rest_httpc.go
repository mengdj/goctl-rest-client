// Package rest
// @file:rest_httpc.go
// @description:
// @date: 02/24/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package rest

import (
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/rest/httpc"
	"net/http"
)

type (
	restHttpcOption = httpc.Option
	restHttpc       struct {
		httpc.Service
	}
)

func (rds restHttpc) Do(ctx context.Context, method, url string, data interface{}) (*http.Response, error) {
	var (
		payload, _ = data.(*RestPayload)
		resp       *http.Response
		err        error
	)
	if resp, err = rds.Service.Do(ctx, method, url, payload.Request); nil != err {
		return nil, err
	}
	if nil != payload.Response {
		if err = jsonx.UnmarshalFromReader(resp.Body, payload.Response); nil != err {
			return nil, err
		}
	}
	return resp, err
}

func (rds restHttpc) DoRequest(r *http.Request) (*http.Response, error) {
	return rds.Service.DoRequest(r)
}

func NewRestHttpc(name string, opts ...RestOption) httpc.Service {
	return &restHttpc{
		Service: httpc.NewServiceWithClient(name, http.DefaultClient),
	}
}
