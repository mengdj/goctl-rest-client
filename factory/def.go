// Package factory
// @file:def.go
// @description:
// @date: 15/06/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package factory

import "context"

const PrefixRestClientHeader = "rest-client-"

type iface struct {
	itab, data uintptr
}

type emptyCtx int

type valueCtx struct {
	context.Context
	key, val interface{}
}

type EnableContextTransfer struct{}
