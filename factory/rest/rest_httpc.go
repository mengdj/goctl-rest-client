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
	//before hook
	HttpcBeforeRequest func(ctx context.Context, url string, req interface{}) (interface{}, error)
	restHttpc          struct {
		httpc.Service
		beforeRequest HttpcBeforeRequest
	}
)

func (rds *restHttpc) Do(ctx context.Context, method, url string, req interface{}, resp interface{}) (*RestResponse, error) {
	var (
		response *http.Response
		err      error
	)
	if nil != rds.Service {
		//before
		if req, err = rds.beforeRequest(ctx, url, req); nil != err {
			return nil, err
		}
	}
	if response, err = rds.Service.Do(ctx, method, url, req); nil != err {
		return nil, err
	}
	if nil != resp {
		//json
		if err = jsonx.UnmarshalFromReader(response.Body, resp); nil != err {
			return nil, err
		}
	}
	return &RestResponse{
		StatusCode: response.StatusCode,
		Status:     response.Status,
	}, err
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
func NewRestHttpc(name string, opts ...RestOption) RestService {
	r := &restHttpc{
		Service:       httpc.NewServiceWithClient(name, http.DefaultClient),
		beforeRequest: nil,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
