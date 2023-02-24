// Package publisher
// @file:etcd.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package publisher

import (
	"fmt"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/netx"
)

type publisherEtcd struct {
	cnf       conf.DiscoverServerConf
	publisher *discov.Publisher
}

func (r publisherEtcd) Start() {
	if err := r.publisher.KeepAlive(); nil != err {
		logx.Errorf("keepalive error:%s", err.Error())
	}
}

func (r publisherEtcd) Stop() {
	r.publisher.Stop()
}

func NewPublisherEtcd(c conf.DiscoverServerConf) Publisher {
	host := c.Host
	if len(host) > 0 {
		if "0.0.0.0" == host {
			host = netx.InternalIp()
		}
	}
	return publisherEtcd{
		publisher: discov.NewPublisher(c.Etcd.Hosts, c.Etcd.Key, fmt.Sprintf("%s:%d", host, c.Port)),
	}
}
