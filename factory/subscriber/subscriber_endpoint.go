// Package subscriber
// @file:subscriber_endpoint.go
// @description:
// @date: 02/24/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package subscriber

import (
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/pkg/errors"
	"math/rand"
	"sync"
)

type subscriberEndpoint struct {
	conf    conf.DiscoverClientConf
	base    []string
	rwMutex sync.RWMutex
}

func (s subscriberEndpoint) Start() {
}

func (s subscriberEndpoint) Stop() {
}

func (s subscriberEndpoint) GetHost() (string, error) {
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

func (s subscriberEndpoint) Scheme() string {
	if s.conf.TLS {
		return "https://"
	}
	return "http://"
}

func NewSubscriberEndpoint(conf conf.DiscoverClientConf) Subscriber {
	return &subscriberEndpoint{
		conf: conf,
		base: conf.Hosts,
	}
}
