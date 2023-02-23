// Package factory
// @file:service.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package factory

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpc"
	"net/http"
)

type (
	RestDiscoverServiceOption = httpc.Option
	RestDiscoverService       struct {
		target httpc.Service
	}
)

func (rds RestDiscoverService) Do(ctx context.Context, method, url string, data interface{}) (*http.Response, error) {
	return rds.target.Do(ctx, method, url, data)
}

func (rds RestDiscoverService) DoRequest(r *http.Request) (*http.Response, error) {
	return rds.target.DoRequest(r)
}

func NewRestDiscoverService(name string, opts ...RestDiscoverServiceOption) httpc.Service {
	return &RestDiscoverService{
		target: httpc.NewService(name, opts...),
	}
}
