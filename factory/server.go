package factory

import (
	"github.com/mengdj/goctl-rest-client/conf"
	publisher2 "github.com/mengdj/goctl-rest-client/factory/publisher"
	"github.com/zeromicro/go-zero/rest"
)

type RestDiscoverServer struct {
	*rest.Server
	config    conf.DiscoverServerConf
	publisher publisher2.Publisher
}

func MustNewServer(c conf.DiscoverServerConf, opts ...rest.RunOption) *RestDiscoverServer {
	r := &RestDiscoverServer{
		Server:    rest.MustNewServer(c.RestConf, opts...),
		config:    c,
		publisher: nil,
	}
	switch c.Resolver {
	case "etcd":
		r.publisher = publisher2.NewPublisherEtcd(c)
		break
	case "consul":
		r.publisher = publisher2.NewPublisherConsul(c)
	case "endpoint":
		r.publisher = publisher2.NewPublisherConsul(c)
		break
	}
	return r
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
