// Package publisher
// @file:publisher_endpoint.go
// @description:
// @date: 02/24/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package publisher

import (
	"github.com/mengdj/goctl-rest-client/conf"
)

type publisherEndpoint struct {
	cnf conf.DiscoverServerConf
}

func (r publisherEndpoint) Start() {
}

func (r publisherEndpoint) Stop() {
	//empty
}

func NewPublisherEndpoint(cnf conf.DiscoverServerConf) Publisher {
	return publisherConsul{
		cnf: cnf,
	}
}
