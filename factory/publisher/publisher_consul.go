// Package publisher
// @file:consul.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package publisher

import (
	"fmt"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/zeromicro/go-zero/core/netx"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type publisherConsul struct {
	cnf conf.DiscoverServerConf
}

func (r publisherConsul) Start() {
	host := r.cnf.Host
	if len(host) > 0 {
		if "0.0.0.0" == host {
			host = netx.InternalIp()
		}
	}
	consul.RegisterService(fmt.Sprintf("%s:%d", host, r.cnf.Port), r.cnf.Consul)
}

func (r publisherConsul) Stop() {
	//empty
}

func NewPublisherConsul(cnf conf.DiscoverServerConf) Publisher {
	return publisherConsul{
		cnf: cnf,
	}
}
