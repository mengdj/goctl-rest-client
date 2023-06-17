// Package middleware
// @file:restory_context.go
// @description:
// @date: 15/06/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package middleware

import (
	"context"
	"github.com/mengdj/goctl-rest-client/factory/utility"
	"net/http"
	"strings"
)

type restoreContext struct {
}

func GetRestoreContext() *restoreContext {
	return &restoreContext{}
}

func (rc *restoreContext) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		for k, v := range request.Header {
			if strings.HasPrefix(k, utility.PrefixRestClientHeader) {
				if len(v) == 1 {
					//prefix rest-client-
					ctx = context.WithValue(ctx, k, v[0])
				}
			}
		}
		next(writer, request.WithContext(ctx))
	}
}
