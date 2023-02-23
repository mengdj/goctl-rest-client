// Package registry
// Copyright ©2023 深圳市慢工坊智能家居有限公司 All Rights reserved.
// @file:registry.go
// @description:
// @date: 02/23/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package factory

type Publisher interface {
	Start()
	Stop()
}
