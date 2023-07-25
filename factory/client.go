// Package factory
// @file:client.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package factory

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/mengdj/goctl-rest-client/factory/rest"
	subscriber2 "github.com/mengdj/goctl-rest-client/factory/subscriber"
	"strings"
	"sync"
)

type Client interface {
	// Invoke invoke
	Invoke(ctx context.Context, method string, path string, data interface{}, result interface{}) error
	// Close release resource
	Close() error
}

type (
	restDiscoverClient struct {
		config      conf.DiscoverClientConf
		base        []string
		service     rest.RestService
		rwMutex     sync.RWMutex
		subscriber  subscriber2.Subscriber
		destination string
		contextPath string
	}
)

var (
	restDiscoverClientOnce     = sync.Once{}
	restDiscoverClientInstance *restDiscoverClient
)

func (f *restDiscoverClient) Close() error {
	if nil != f.subscriber {
		f.subscriber.Stop()
	}
	return nil
}

// WithErrorHandler 错误处理
func WithErrorHandler(fn func(status int, body []byte) error) rest.RestOption {
	return func(v interface{}) {
		if target, ok := v.(rest.RestService); ok {
			target.SetErrorHandler(fn)
		}
	}
}

func NewRestDiscoverClient(destination string, c conf.DiscoverClientConf, opts ...rest.RestOption) Client {
	var (
		transfer rest.RestService = nil
	)
	if c.Name == "" {
		//default
		c.Name = uuid.NewString()
	}
	if c.Transfer.Type == "resty" {
		transfer = rest.NewRestResty(c.Name, c.Transfer, opts...)
	} else if c.Transfer.Type == "fasthttp" {
		transfer = rest.NewRestFastHttp(c.Name, c.Transfer, opts...)
	} else {
		//default
		transfer = rest.NewRestHttpc(c.Name, opts...)
	}

	return NewRestDiscoverClientWithService(destination, c, transfer)
}

func NewRestDiscoverClientWithService(destination string, c conf.DiscoverClientConf, s rest.RestService) Client {
	restDiscoverClientOnce.Do(func() {
		//just once
		restDiscoverClientInstance = &restDiscoverClient{
			config:      c,
			service:     s,
			subscriber:  nil,
			destination: destination,
			contextPath: c.ContextPath, //add
		}
		switch c.Resolver {
		case "etcd":
			if "" == c.Etcd.Key {
				c.Etcd.Key = destination
			}
			restDiscoverClientInstance.subscriber = subscriber2.NewSubscriberEtcd(c)
			break
		case "consul":
			if "" == c.Consul.Key {
				c.Consul.Key = destination
			}
			restDiscoverClientInstance.subscriber = subscriber2.NewSubscriberConsul(c)
			break
		case "endpoint":
			if 0 == len(c.Hosts) {
				c.Hosts = []string{
					destination,
				}
			}
			restDiscoverClientInstance.subscriber = subscriber2.NewSubscriberEndpoint(c)
			break
		}
		if nil != restDiscoverClientInstance.subscriber {
			restDiscoverClientInstance.subscriber.Start()
		}
	})
	return restDiscoverClientInstance
}

// Invoke invoke method
func (f *restDiscoverClient) Invoke(ctx context.Context, method string, path string, data interface{}, result interface{}) error {
	var (
		host = ""
		err  error
		urls strings.Builder
		tar  string
	)
	if host, err = f.subscriber.GetHost(); nil == err {
		urls.WriteString(f.subscriber.Scheme())
		urls.WriteString(host)
		urls.WriteString(f.contextPath)
		urls.WriteString(path)
	} else {
		//default
		urls.WriteString(f.destination)
		urls.WriteString(f.contextPath)
		urls.WriteString(path)
	}
	tar = urls.String()
	if _, err = f.service.Do(
		ctx,
		method,
		tar,
		data,
		result,
	); nil != err {
		return fmt.Errorf("%s(%s)", tar, err.Error())
	}
	return nil
}
