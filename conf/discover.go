package conf

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type (
	DiscoverClientConf struct {
		Name     string          `json:",optional"`
		Hosts    []string        `json:",optional"`
		Etcd     discov.EtcdConf `json:",optional"` //etcd
		Consul   consul.Conf     `json:",optional"` //consul
		Resolver string          `json:"Resolver"`  //resolver[etcd,consul,endpoint]
		Transfer TransferConf    `json:",optional"` //transfer[httpc,resty]
		TLS      bool            `json:",default=false"`
	}
	DiscoverServerConf struct {
		rest.RestConf
		DiscoverClientConf
	}
	TransferConf struct {
		Type  string    `json:",default=httpc"` //httpc,resty
		Rety  RertyConf `json:",optional"`
		Httpc HttpcConf `json:",optional"`
	}

	HttpcConf struct {
	}

	RertyConf struct {
		Agent                 string            `json:",optional"`
		AllowGetMethodPayload bool              `json:",default=false"`
		Token                 string            `json:",optional"`
		Debug                 bool              `json:",default=false"`
		Timeout               int64             `json:",default=0"`
		Header                map[string]string `json:",optional"`
	}
)
