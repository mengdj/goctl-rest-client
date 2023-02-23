// Package subscriber
// Copyright ©2023 深圳市慢工坊智能家居有限公司 All Rights reserved.
// @file:subscriber_consul.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package factory

import (
	"github.com/hashicorp/consul/api"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type subscriberConsul struct {
	conf    conf.DiscoverClientConf
	client  *api.Client
	base    []string
	rwMutex sync.RWMutex
}

func (s subscriberConsul) Start() {
	ttlTicker := time.Duration(s.conf.Consul.TTL-1) * time.Second
	if ttlTicker < time.Second {
		ttlTicker = time.Second
	}
	// routine to update ttl
	go func() {
		ticker := time.NewTicker(ttlTicker)
		defer ticker.Stop()
		for {
			<-ticker.C
			if services, err := s.client.Agent().Services(); nil != err {
				logx.Error("query service err", err)
			} else {
				s.rwMutex.Lock()
				s.base = []string{}
				for _, serv := range services {
					if serv.Service == s.conf.Consul.Key {
						addr := strings.Builder{}
						addr.Reset()
						addr.WriteString(serv.Address)
						addr.WriteString(":")
						addr.WriteString(strconv.Itoa(serv.Port))
						s.base = append(s.base, addr.String())
					}
				}
				logx.Info(s.base)
				s.rwMutex.Unlock()
			}
		}
	}()
}

func (s subscriberConsul) Stop() {
}

func (s subscriberConsul) GetHost() (string, error) {
	s.rwMutex.Lock()
	defer func() {
		s.rwMutex.Unlock()
	}()
	logx.Info(len(s.base))
	if len(s.base) > 0 {
		rand.Shuffle(len(s.base), func(i, j int) {
			s.base[i], s.base[j] = s.base[j], s.base[i]
		})
		return s.base[0], nil
	}
	return "", errors.New("host can't empty")
}

func NewSubscriberConsul(conf conf.DiscoverClientConf) Subscriber {
	client, _ := api.NewClient(&api.Config{Scheme: "http", Address: conf.Consul.Host, Token: conf.Consul.Token})
	return subscriberConsul{
		conf:   conf,
		client: client,
		base:   conf.Endpoints,
	}
}
