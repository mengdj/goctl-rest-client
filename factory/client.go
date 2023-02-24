// Package factory
// @file:client.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package factory

import (
	"context"
	"github.com/mengdj/goctl-rest-client/conf"
	subscriber2 "github.com/mengdj/goctl-rest-client/factory/subscriber"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/rest/httpc"
	"net/http"
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
	restDiscoverClientOption func(*restDiscoverClient)
	restDiscoverClient       struct {
		config      conf.DiscoverClientConf
		base        []string
		service     httpc.Service
		rwMutex     sync.RWMutex
		subscriber  subscriber2.Subscriber
		destination string
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

func NewRestDiscoverClient(destination string, c conf.DiscoverClientConf, opts ...RestDiscoverServiceOption) Client {
	return NewRestDiscoverClientWithService(destination, c, NewRestDiscoverService(c.Etcd.Key, opts...))
}

func NewRestDiscoverClientWithService(destination string, c conf.DiscoverClientConf, s httpc.Service) Client {
	restDiscoverClientOnce.Do(func() {
		restDiscoverClientInstance = &restDiscoverClient{
			config:      c,
			service:     s,
			subscriber:  nil,
			destination: destination,
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
			//start service
			restDiscoverClientInstance.subscriber.Start()
		}
	})
	return restDiscoverClientInstance
}

// Invoke invoke method
func (f *restDiscoverClient) Invoke(ctx context.Context, method string, path string, data interface{}, result interface{}) error {
	var (
		host     = ""
		response *http.Response
		err      error
		urls     strings.Builder
	)
	if host, err = f.subscriber.GetHost(); nil == err {
		urls.WriteString(f.subscriber.Scheme())
		urls.WriteString(host)
		urls.WriteString(path)
	} else {
		//use param
		urls.WriteString(f.destination)
		urls.WriteString(path)
	}
	if response, err = f.service.Do(ctx, strings.ToUpper(method), urls.String(), data); nil != err {
		return err
	}
	//200~226
	if !(response.StatusCode >= http.StatusOK && response.StatusCode <= http.StatusIMUsed) {
		return errors.New(response.Status)
	}
	if nil != result {
		return jsonx.UnmarshalFromReader(response.Body, result)
	}
	return nil
}
