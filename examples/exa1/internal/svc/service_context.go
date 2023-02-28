package svc

import (
	"github.com/mengdj/goctl-rest-client/examples/exa1/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
