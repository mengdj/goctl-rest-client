// Package world
// @file:client.go
// @version:v0.0.5
// @date:2023-03-19 15:34:42.1936773 +0800 CST m=+0.672860401
// Code generated by goctl-resty-client. DO NOT EDIT.
package world

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
		Hello(context.Context, *client.WorldRequest) (*client.Response, error)
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
func (cf *clientFactory) Hello(ctx context.Context, entity *client.WorldRequest) (resp *client.Response, err error) {
	resp = new(client.Response)
	err = cf.Invoke(ctx, "GET", "/api/v1/app/test/world/:id", entity, resp)
	if nil != err {
		return nil, err
	}
	return resp, nil
}
