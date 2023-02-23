// Package subscriber
// Copyright ©2023 深圳市慢工坊智能家居有限公司 All Rights reserved.
// @file:subscriber_etcd.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package factory

import (
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"sync"
)

type subscriberEtcd struct {
	conf     conf.DiscoverClientConf
	protocol string
	base     []string
	rwMutex  sync.RWMutex
}

func (s subscriberEtcd) Start() {
	subOpts := make([]discov.SubOption, 0)
	if "" != s.conf.Etcd.User {
		subOpts = append(subOpts, discov.WithSubEtcdAccount(s.conf.Etcd.User, s.conf.Etcd.Pass))
	}
	sub, err := discov.NewSubscriber(s.conf.Etcd.Hosts, s.conf.Etcd.Key, subOpts...)
	if nil != err {
		panic(err)
	}
	update := func() {
		s.rwMutex.Lock()
		defer func() {
			s.rwMutex.Unlock()
		}()
		if values := sub.Values(); len(values) > 0 {
			logx.Debugw("found values", logx.Field("values", values))
			s.base = values
		}
	}
	sub.AddListener(update)
	update()
}

func (s subscriberEtcd) Stop() {
}

func (s subscriberEtcd) GetHost() (string, error) {
	s.rwMutex.RLock()
	defer func() {
		s.rwMutex.RUnlock()
	}()
	if len(s.base) > 0 {
		rand.Shuffle(len(s.base), func(i, j int) {
			s.base[i], s.base[j] = s.base[j], s.base[i]
		})
		return s.base[0], nil
	}
	return "", errors.New("host can't empty")
}

func NewSubscriberEtcd(conf conf.DiscoverClientConf) Subscriber {
	return subscriberEtcd{
		conf: conf,
		base: conf.Endpoints,
	}
}
