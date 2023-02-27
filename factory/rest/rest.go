// Package rest
// @file:payload.go
// @description:
// @date: 02/24/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package rest

import (
	"context"
)

//
//go:generate requestgen -type RestPayload -tags json -output payload_requestgen.go
type (
	RestResponse struct {
		Status     string // e.g. "200 OK"
		StatusCode int    // e.g. 200
	}

	RestService interface {
		Do(ctx context.Context, method, url string, req interface{}, resp interface{}) (*RestResponse, error)
	}
)
