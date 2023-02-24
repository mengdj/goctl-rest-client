// Package rest
// Copyright ©2023 深圳市慢工坊智能家居有限公司 All Rights reserved.
// @file:rest_resty.go
// @description:
// @date: 02/24/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package rest

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mapping"
	"github.com/zeromicro/go-zero/rest/httpc"
	"net/http"
	nurl "net/url"
)

type (
	restResty struct {
		client *resty.Client
	}
)

const (
	pathKey   = "path"
	formKey   = "form"
	headerKey = "header"
	jsonKey   = "json"
	slash     = "/"
	colon     = ':'
)

var (
	NotSupport     = errors.New("Not Support")
	ErrGetWithBody = errors.New("HTTP GET should not have body")
)

func (rds *restResty) Do(ctx context.Context, method, url string, data interface{}) (*http.Response, error) {
	var (
		payload, _ = data.(RestPayload)
		rertyR     = rds.client.R().SetContext(ctx)
		rertyP     *resty.Response
		err        error
		purl       *nurl.URL
		val        map[string]map[string]interface{}
	)
	if purl, err = nurl.Parse(url); err != nil {
		//parse url
		return nil, err
	}
	if payload.Request != nil {
		//parse request
		if val, err = mapping.Marshal(payload.Request); err != nil {
			return nil, err
		}
		//path
		if pv, ok := val[pathKey]; ok {
			if err = fillPath(purl, pv); err != nil {
				return nil, err
			}
		}
		//form
		if fv := buildFormQuery(purl, val[formKey]); len(fv) > 0 {
			rertyR.SetFormDataFromValues(fv)
		}
		//body
		if jv, ok := val[jsonKey]; ok {
			if method == http.MethodGet {
				return nil, ErrGetWithBody
			}
			rertyR.SetBody(jv)
			rertyR.SetHeader(ContentType, JsonContentType)
		}
		//header
		if hv, ok := val[headerKey]; ok {
			for k, v := range hv {
				rertyR.SetHeader(k, fmt.Sprint(v))
			}
		}
		//result
		if nil != payload.Response {
			rertyR.SetResult(payload.Response)
		}
	}
	if rertyP, err = rertyR.Execute(method, purl.String()); nil != err {
		return nil, err
	}
	return rertyP.RawResponse, nil
}

func (rds *restResty) DoRequest(r *http.Request) (*http.Response, error) {
	return nil, NotSupport
}

func NewRestResty(cnf conf.TransferConf, opts ...RestOption) httpc.Service {
	client := resty.New().SetDebug(cnf.Rety.Debug).SetAllowGetMethodPayload(cnf.Rety.AllowGetMethodPayload)
	if cnf.Rety.Token != "" {
		client.SetAuthToken(cnf.Rety.Token)
	}
	return &restResty{
		client: client,
	}
}
