package conf

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

//go:generate fieldalignment -fix discover.go

type (
	DiscoverClientConf struct {
		Transfer TransferConf    `json:",optional"`
		Consul   consul.Conf     `json:",optional"`
		Name     string          `json:",optional"`
		Resolver string          `json:"Resolver"`             //resolver[etcd,consul,endpoint]
		Balancer string          `json:",default=round-robin"` //round-robin,random,power of 2 random choice,consistent hash,consistent hash with bounded,ip-hash,least-load
		Etcd     discov.EtcdConf `json:",optional"`
		Hosts    []string        `json:",optional"`
		TLS      bool            `json:",default=false"`
	}
	DiscoverServerConf struct {
		DiscoverClientConf
		rest.RestConf
	}
	TransferConf struct {
		Fasthttp FastHttpConf `json:",optional"`
		Resty    RestyConf    `json:",optional"`
		Type     string       `json:",default=fasthttp"`
	}

	HttpcConf struct {
	}
	FastHttpConf struct {
		Header              map[string]string `json:",optional"`
		ReadTimeout         int64             `json:",default=10"` //seconds
		MaxConnWaitTimeout  int64             `json:",default=0"`
		WriteTimeout        int               `json:",default=0"`
		MaxConnDuration     int               `json:",default=0"`
		MaxIdleConnDuration int               `json:",default=0"`
		TCPDialer           struct {
			Concurrency      int   `json:",default=64"`   //Concurrency
			DNSCacheDuration int64 `json:",default=1800"` //DNSCacheDuration
		} `json:",optional"`
	}
	RestyConf struct {
		Header                map[string]string `json:",optional"`
		Agent                 string            `json:",optional"`
		Token                 string            `json:",optional"`
		Timeout               int64             `json:",default=0"`
		AllowGetMethodPayload bool              `json:",default=false"`
		Debug                 bool              `json:",default=false"`
		Trace                 bool              `json:",default=false"`
	}
)
