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
	HttpcBeforeRequest func(ctx context.Context, url string, data interface{}) (interface{}, error)
	restHttpc          struct {
		httpc.Service
		beforeRequest HttpcBeforeRequest
	}
)

func (rds *restHttpc) Do(ctx context.Context, method, url string, data interface{}) (*http.Response, error) {
	var (
		payload, _ = data.(*RestPayload)
		resp       *http.Response
		err        error
	)
	if nil != rds.Service {
		//before
		if payload.Request, err = rds.beforeRequest(ctx, url, data); nil != err {
			return nil, err
		}
	}
	if resp, err = rds.Service.Do(ctx, method, url, payload.Request); nil != err {
		return nil, err
	}
	if nil != payload.Response {
		//json
		if err = jsonx.UnmarshalFromReader(resp.Body, payload.Response); nil != err {
			return nil, err
		}
	}
	return resp, err
}

func (rds *restHttpc) DoRequest(r *http.Request) (*http.Response, error) {
	return rds.Service.DoRequest(r)
}

func WithHttpcBeforeRequest(fn HttpcBeforeRequest) RestOption {
	return func(v interface{}) {
		if target, ok := v.(*restHttpc); ok {
			target.beforeRequest = fn
		}
	}
}
func NewRestHttpc(name string, opts ...RestOption) httpc.Service {
	r := &restHttpc{
		Service:       httpc.NewServiceWithClient(name, http.DefaultClient),
		beforeRequest: nil,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
