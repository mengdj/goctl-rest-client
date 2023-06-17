// Package exa
// @file:client.go
// @version:v0.1.0
// @date:2023-06-17 16:07:12.263680664 +0800 CST m=+0.014465604
// Code generated by goctl-resty-client. DO NOT EDIT.
package exa

import (
	"context"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/mengdj/goctl-rest-client/examples/exa2/client"
	"github.com/mengdj/goctl-rest-client/factory"
	"github.com/mengdj/goctl-rest-client/factory/rest"
)

//go:generate fieldalignment -fix client.go
//begin
type (
	// Client
	Client interface {
		// Hello
		Hello(context.Context, *client.HelloRequest) (*client.Response, error)
		Invoke(context.Context, string, string, interface{}, interface{}) error
	}
	clientFactory struct {
		factory.Client
	}
)

// MustClient
func MustClient(c conf.DiscoverClientConf, opts ...rest.RestOption) Client {
	return &clientFactory{
		Client: factory.NewRestDiscoverClient("exa_api", c, opts...),
	}
}

//Invoke extend
func (cf *clientFactory) Invoke(ctx context.Context, method string, path string, entity interface{}, resp interface{}) error {
	return cf.Client.Invoke(ctx, method, path, entity, resp)
}

// Hello
func (cf *clientFactory) Hello(ctx context.Context, entity *client.HelloRequest) (resp *client.Response, err error) {
	resp = new(client.Response)
	err = cf.Invoke(ctx, "POST", "/api/v1/app/demo/hello", entity, resp)
	if nil != err {
		return nil, err
	}
	return resp, nil
}
