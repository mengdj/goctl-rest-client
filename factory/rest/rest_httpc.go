// Package rest
// @file:rest_httpc.go
// @description:
// @date: 02/24/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package rest

import (
	"context"
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
	return rds.Service.Do(ctx, method, url, data)
}

func (rds restHttpc) DoRequest(r *http.Request) (*http.Response, error) {
	return rds.Service.DoRequest(r)
}

func NewRestHttpc(name string, opts ...RestOption) httpc.Service {
	return &restHttpc{
		Service: httpc.NewServiceWithClient(name, http.DefaultClient),
	}
}
