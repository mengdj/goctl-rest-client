package factory

import (
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/zeromicro/go-zero/rest"
)

type RestDiscoverServer struct {
	*rest.Server
	config    conf.DiscoverServerConf
	publisher Publisher
}

func MustNewServer(c conf.DiscoverServerConf, opts ...rest.RunOption) *RestDiscoverServer {
	var pub Publisher = nil
	switch c.Resolver {
	case "etcd":
		pub = NewPublisherEtcd(c)
		break
	case "consul":
		pub = NewPublisherConsul(c)
		break
	default:
		break
	}
	return &RestDiscoverServer{
		Server:    rest.MustNewServer(c.RestConf, opts...),
		config:    c,
		publisher: pub,
	}
}

func (r *RestDiscoverServer) Start() {
	if nil != r.publisher {
		r.publisher.Start()
	}
	r.Server.Start()
}

// Stop stops the Server.
func (r *RestDiscoverServer) Stop() {
	if nil != r.publisher {
		r.publisher.Stop()
	}
	r.Server.Stop()
}
