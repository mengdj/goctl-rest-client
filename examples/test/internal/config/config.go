package config

import (
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	HelloDiscoverConf conf.DiscoverClientConf
}
