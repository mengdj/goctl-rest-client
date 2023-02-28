// Package hello
// @file:client.go
// @version:v0.0.3
// @date:2023-02-27 09:28:38.83772833 +0800 CST m=+0.010425701
// Code generated by goctl-resty-client. DO NOT EDIT.
package hello

import (
	"context"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/mengdj/goctl-rest-client/examples/test/client"
	"github.com/mengdj/goctl-rest-client/factory"
	"github.com/mengdj/goctl-rest-client/factory/rest"
)

// begin
//
//go:generate fieldalignment -fix client.go
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

// Invoke extend
func (cf *clientFactory) Invoke(ctx context.Context, method string, path string, entity interface{}, resp interface{}) error {
	return cf.Client.Invoke(ctx, method, path, entity, resp)
}

// Hello
func (cf *clientFactory) Hello(ctx context.Context, entity *client.HelloRequest) (resp *client.Response, err error) {
	resp = new(client.Response)
	err = cf.Invoke(ctx, "GET", "/api/v1/app/test/hello/:id", entity, resp)
	if nil != err {
		return nil, err
	}
	return resp, nil
}