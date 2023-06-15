// Package factory
// @file:utility.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package factory

import (
	"context"
	"unsafe"
)

func GetKeyValueFromContext(ctx context.Context) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	getKeyValueFromContext(ctx, m)
	return m
}

func getKeyValueFromContext(ctx context.Context, m map[interface{}]interface{}) {
	ictx := *(*iface)(unsafe.Pointer(&ctx))
	if ictx.data == 0 || int(*(*emptyCtx)(unsafe.Pointer(ictx.data))) == 0 {
		return
	}
	valCtx := (*valueCtx)(unsafe.Pointer(ictx.data))
	if valCtx != nil && valCtx.key != nil {
		m[valCtx.key] = valCtx.val
	}
	getKeyValueFromContext(valCtx.Context, m)
}
