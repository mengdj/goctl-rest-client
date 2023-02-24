// Package subscriber
// @file:subscriber.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package subscriber

type Subscriber interface {
	// Start service
	Start()
	// Stop service
	Stop()
	// GetHost get host and port
	GetHost() (string, error)
	// Scheme getscheme example https:// or http
	Scheme() string
}
