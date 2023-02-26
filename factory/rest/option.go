// Package rest
// @file:option.go
// @description:
// @date: 02/24/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package rest

import "github.com/pkg/errors"

type (
	// RestOption any
	RestOption func(v interface{})
)

var (
	NotSupport     = errors.New("Not Support")
	ErrGetWithBody = errors.New("HTTP GET should not have body")
)
