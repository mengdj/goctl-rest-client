// Package exa
// @file:client.go
// @version:v0.0.3
// @date:2023-02-26 12:28:21.1831437 +0800 CST m=+0.969367301
// Code generated by goctl-resty-client. DO NOT EDIT.
package exa

import (
	"context"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/mengdj/goctl-rest-client/example/exa1/client"
	"github.com/mengdj/goctl-rest-client/factory"
	"github.com/mengdj/goctl-rest-client/factory/rest"
)

// begin
//
//go:generate fieldalignment -fix client.go
type (
	// Client
	Client interface {
		// Hello xx
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

// Hello xx
func (cf *clientFactory) Hello(ctx context.Context, entity *client.HelloRequest) (resp *client.Response, err error) {
	resp = new(client.Response)
	err = cf.Invoke(ctx, "post", "/api/v1/app/demo/hello", entity, resp)
	if nil != err {
		return nil, err
	}
	return resp, nil
}
