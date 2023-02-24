// Package rest
// Copyright ©2023 深圳市慢工坊智能家居有限公司 All Rights reserved.
// @file:rest_resty.go
// @description:
// @date: 02/24/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package rest

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/zeromicro/go-zero/rest/httpc"
	"net/http"
)

type (
	restResty struct {
		client *resty.Client
	}
)

func (rds restResty) Do(ctx context.Context, method, url string, data interface{}) (*http.Response, error) {
	return nil, nil
}

func (rds restResty) DoRequest(r *http.Request) (*http.Response, error) {
	return nil, nil
}

func NewRestResty(cnf conf.TransferConf, opts ...RestOption) httpc.Service {
	return &restResty{
		client: resty.New().SetDebug(cnf.Rety.Debug),
	}
}
