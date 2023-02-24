package conf

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type (
	DiscoverClientConf struct {
		Hosts    []string        `json:",optional"`
		Etcd     discov.EtcdConf `json:",optional"` //etcd
		Consul   consul.Conf     `json:",optional"` //consul
		Resolver string          `json:"Resolver"`  //resolver[etcd,consul,endpoint]
		TLS      bool            `json:",default=false"`
	}
	DiscoverServerConf struct {
		rest.RestConf
		DiscoverClientConf
	}
)
