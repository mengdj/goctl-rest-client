package rest

import "github.com/pkg/errors"

var (
	NotSupport     = errors.New("Not Support")
	ErrGetWithBody = errors.New("HTTP GET should not have body")
)
