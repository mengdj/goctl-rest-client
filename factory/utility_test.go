// Package factory
// @file:utility.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package factory

import (
	"context"
	"testing"
)

func TestGetKeyValueFromContext(t *testing.T) {
	t.Log(GetKeyValueFromContext(context.WithValue(context.WithValue(context.TODO(), "aa", 22), "y", struct {
		Age  int
		Name string
	}{
		Age:  22,
		Name: "dj",
	})))
}
