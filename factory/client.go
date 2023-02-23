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
	RestDiscoverClientOption func(*RestDiscoverClient)
	RestDiscoverClient       struct {
		config     conf.DiscoverClientConf
		protocol   string
		base       []string
		service    httpc.Service
		rwMutex    sync.RWMutex
		subscriber Subscriber
	}
)

func (f *RestDiscoverClient) Close() error {
	if nil != f.subscriber {
		f.subscriber.Stop()
	}
	return nil
}

func NewRestDiscoverClient(c conf.DiscoverClientConf, opts ...RestDiscoverServiceOption) Client {
	return NewRestDiscoverClientWithService(c, NewRestDiscoverService(c.Etcd.Key, opts...))
}

func NewRestDiscoverClientWithService(c conf.DiscoverClientConf, s httpc.Service) Client {
	var (
		sub      Subscriber = nil
		protocol            = "http://"
	)
	if c.TLS {
		protocol = "https://"
	}
	switch c.Resolver {
	case "etcd":
		sub = NewSubscriberEtcd(c)
		break
	case "consul":
		sub = NewSubscriberConsul(c)
		break
	default:
		break
	}
	//begin
	sub.Start()
	ret := &RestDiscoverClient{
		protocol:   protocol,
		config:     c,
		service:    s,
		subscriber: sub,
	}
	return ret
}

// 调用
func (f *RestDiscoverClient) Invoke(ctx context.Context, method string, path string, data interface{}, result interface{}) error {
	var (
		host     = ""
		response *http.Response
		err      error
		urls     strings.Builder
	)
	if host, err = f.subscriber.GetHost(); nil != err {
		return err
	}
	//buffer
	urls.WriteString(f.protocol)
	urls.WriteString(host)
	urls.WriteString(path)
	if response, err = f.service.Do(ctx, strings.ToUpper(method), urls.String(), data); nil != err {
		return err
	}
	//200~226
	if !(response.StatusCode >= http.StatusOK && response.StatusCode <= http.StatusIMUsed) {
		return errors.New(response.Status)
	}
	if nil != result {
		if err = jsonx.UnmarshalFromReader(response.Body, result); nil != err {
			return err
		}
	}
	return nil
}
