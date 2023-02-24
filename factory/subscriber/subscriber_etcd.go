// Package subscriber
// @file:subscriber_etcd.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package subscriber

import (
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/discov"
	"math/rand"
	"sync"
)

type subscriberEtcd struct {
	conf     conf.DiscoverClientConf
	protocol string
	base     []string
	rwMutex  sync.RWMutex
}

func (s *subscriberEtcd) Scheme() string {
	if s.conf.TLS {
		return "https://"
	}
	return "http://"
}

func (s *subscriberEtcd) Start() {
	if len(s.conf.Etcd.Hosts) == 0 {
		return
	}
	//discover
	subOpts := make([]discov.SubOption, 0)
	if "" != s.conf.Etcd.User {
		subOpts = append(subOpts, discov.WithSubEtcdAccount(s.conf.Etcd.User, s.conf.Etcd.Pass))
	}
	if sub, err := discov.NewSubscriber(s.conf.Etcd.Hosts, s.conf.Etcd.Key, subOpts...); nil == err {
		update := func() {
			s.rwMutex.Lock()
			defer func() {
				s.rwMutex.Unlock()
			}()
			if values := sub.Values(); len(values) > 0 {
				s.base = append([]string{}, values...)
			}
		}
		sub.AddListener(update)
		update()
	}
}

func (s *subscriberEtcd) Stop() {
}

func (s *subscriberEtcd) GetHost() (string, error) {
	s.rwMutex.RLock()
	defer func() {
		s.rwMutex.RUnlock()
	}()
	if bs := len(s.base); bs > 0 {
		if 1 != bs {
			rand.Shuffle(len(s.base), func(i, j int) {
				s.base[i], s.base[j] = s.base[j], s.base[i]
			})
		}
		return s.base[0], nil
	}
	return "", errors.New("host can't empty")
}

func NewSubscriberEtcd(conf conf.DiscoverClientConf) Subscriber {
	return &subscriberEtcd{
		conf: conf,
		base: []string{},
	}
}
