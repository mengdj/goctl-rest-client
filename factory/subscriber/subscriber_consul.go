// Package subscriber
// @file:subscriber_consul.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package subscriber

import (
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

type (
	//consul
	subscriberConsulOnce = sync.Once
	subscriberConsul     struct {
		conf      conf.DiscoverClientConf
		rwMutex   sync.RWMutex
		watchPlan *watch.Plan
		base      []string
	}
)

func (s *subscriberConsul) Scheme() string {
	if s.conf.TLS {
		return "https://"
	}
	return "http://"
}

func (s *subscriberConsul) Start() {
	if "" == s.conf.Consul.Host {
		return
	}
	if client, err := api.NewClient(&api.Config{Scheme: "http", Address: s.conf.Consul.Host, Token: s.conf.Consul.Token}); nil == err {
		//init
		if services, servicesErr := client.Agent().Services(); nil == servicesErr {
			for k, v := range services {
				logx.Debugf("found service:%s", k)
				addr := strings.Builder{}
				base := make([]string, 0, len(services))
				if v.Service == s.conf.Consul.Key {
					addr.Reset()
					addr.WriteString(v.Address)
					addr.WriteString(":")
					addr.WriteString(strconv.Itoa(v.Port))
					base = append(base, addr.String())
				}
				if bs := len(base); bs > 0 {
					s.base = s.base[0:0]
					s.base = base[0:bs]
				}
			}
		}
		//watch
		if s.watchPlan, err = watch.Parse(map[string]interface{}{
			"type":    "service",
			"service": s.conf.Consul.Key,
		}); nil == err {
			//monitor
			s.watchPlan.Handler = func(u uint64, i interface{}) {
				if vv, ok := i.([]*api.ServiceEntry); ok {
					if vsize := len(vv); vsize > 0 {
						//lock
						s.rwMutex.Lock()
						defer func() {
							s.rwMutex.Unlock()
						}()
						addr := strings.Builder{}
						base := make([]string, 0, vsize)
						s.base = s.base[0:0]
						for _, v := range vv {
							//check health
							if api.HealthPassing == v.Checks.AggregatedStatus() {
								addr.Reset()
								addr.WriteString(v.Service.Address)
								addr.WriteString(":")
								addr.WriteString(strconv.Itoa(v.Service.Port))
								base = append(base, addr.String())
							}
						}
						//clone
						if bs := len(base); bs > 0 {
							s.base = s.base[0:0]
							s.base = base[0:bs]
						}
					}
				}
			}
			threading.GoSafe(func() {
				if err = s.watchPlan.RunWithClientAndHclog(client, nil); nil != err {
					logx.Errorf("consul watch error:%v", err)
				}
			})
		}
		logx.Info("watch success")
	}
}

func (s *subscriberConsul) Stop() {
	if nil != s.watchPlan {
		if !s.watchPlan.IsStopped() {
			s.watchPlan.Stop()
		}
	}
}

func (s *subscriberConsul) GetHost() (string, error) {
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

func NewSubscriberConsul(conf conf.DiscoverClientConf) Subscriber {
	return &subscriberConsul{
		conf:      conf,
		watchPlan: nil,
		base:      []string{},
	}
}
