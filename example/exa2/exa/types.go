// Package exa
// @file:types.go
// @version:v0.0.2
// Code generated by goctl-resty-client. DO NOT EDIT.
package exa

// begin
//
//go:generate fieldalignment -fix types.go
type (
	Response struct {
		Code uint32 `json:"code"`
		Msg  string `json:"msg"`
	}

	HelloRequest struct {
		Msg string `json:"msg"`
	}
)

func (t *Response) Reset() {
	*t = Response{}
}

func (t *Response) SetCode(v uint32) *Response {
	t.Code = v
	return t
}

func (t *Response) GetCode() uint32 {
	return t.Code
}

func (t *Response) SetMsg(v string) *Response {
	t.Msg = v
	return t
}

func (t *Response) GetMsg() string {
	return t.Msg
}

func (t *HelloRequest) Reset() {
	*t = HelloRequest{}
}

func (t *HelloRequest) SetMsg(v string) *HelloRequest {
	t.Msg = v
	return t
}

func (t *HelloRequest) GetMsg() string {
	return t.Msg
}
