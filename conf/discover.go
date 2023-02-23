package conf

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type (
	DiscoverClientConf struct {
		Etcd      discov.EtcdConf `json:",optional"` //etcd
		Consul    consul.Conf     //consul
		Endpoints []string        `json:",optional"` //直连
		TLS       bool            `json:",optional"`
		Resolver  string          `json:"Resolver"` //resolver
	}
	DiscoverServerConf struct {
		rest.RestConf
		DiscoverClientConf
	}
)
