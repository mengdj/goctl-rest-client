// Package rest
// @file:payload.go
// @description:
// @date: 02/24/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package rest

//
//go:generate requestgen -type RestPayload -tags json -output payload_requestgen.go
type RestPayload struct {
	//request object
	Request interface{} `json:"request" param:"request"`
	//response object
	Response interface{} `json:"response" param:"response"`
}

func (r *RestPayload) Reset() {
	r.Request = nil
	r.Response = nil
}

func (r *RestPayload) SetRequest(req interface{}) *RestPayload {
	r.Request = req
	return r
}

func (r *RestPayload) SetResponse(resp interface{}) *RestPayload {
	r.Response = resp
	return r
}
