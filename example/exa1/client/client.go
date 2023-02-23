// Code generated by goctl-resty-client. DO NOT EDIT.
package client

import (
	"context"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/mengdj/goctl-rest-client/factory"
)

// begin
//
//go:generate fieldalignment -fix client.go
type (

	//aa
	Response struct {
		Msg  string `json:"msg"`
		Code uint32 `json:"code"`
	}
	//xxx
	HelloRequest struct {

		//xxx

		Msg string `json:"msg"` //回显消息
	}

	// Client
	Client interface {

		// Hello
		Hello(context.Context, *HelloRequest) (*Response, error)
		Invoke(context.Context, string, string, interface{}, interface{}) error
	}
	clientFactory struct {
		factory.Client
	}
)

//end

// MustClient
func MustClient(c conf.DiscoverClientConf, opts ...factory.RestDiscoverServiceOption) Client {
	return &clientFactory{
		Client: factory.NewRestDiscoverClient(c, opts...),
	}
}

func MustClientWithHttpService(c conf.DiscoverClientConf, opts ...factory.RestDiscoverServiceOption) Client {
	return &clientFactory{
		Client: factory.NewRestDiscoverClient(c, opts...),
	}
}

func (cf *clientFactory) Invoke(ctx context.Context, method string, path string, entity interface{}, resp interface{}) error {
	return cf.Client.Invoke(ctx, method, path, entity, resp)
}

// Hello "xx"
func (cf *clientFactory) Hello(ctx context.Context, entity *HelloRequest) (resp *Response, err error) {
	resp = new(Response)
	err = cf.Invoke(ctx, "post", "/api/v1/app/demo/hello", entity, resp)
	if nil != err {
		return nil, err
	}
	return resp, nil
}
