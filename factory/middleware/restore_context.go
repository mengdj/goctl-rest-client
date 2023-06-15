// Package middleware
// @file:restory_context.go
// @description:
// @date: 15/06/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package middleware

import "net/http"

type restoreContext struct {
}

func GetRestoreContext() *restoreContext {
	return &restoreContext{}
}

func (rc *restoreContext) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		next(writer, request)
	}
}
