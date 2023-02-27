package rest

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/mengdj/goctl-rest-client/conf"
	"github.com/valyala/fasthttp"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/mapping"
	"net/http"
	nurl "net/url"
	"time"
)

type restFastHttp struct {
	name   string
	client *fasthttp.Client
	cnf    conf.FastHttpConf
}

func (r restFastHttp) Do(ctx context.Context, method, url string, req interface{}, resp interface{}) (*RestResponse, error) {
	var (
		request  = fasthttp.AcquireRequest()
		response = fasthttp.AcquireResponse()
		err      error
		purl     *nurl.URL
		val      map[string]map[string]interface{}
		body     []byte
	)
	defer func() {
		//clear
		request.Reset()
		response.Reset()
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()
	if purl, err = nurl.Parse(url); err != nil {
		return nil, err
	}
	//method
	request.Header.SetMethod(method)
	if len(r.cnf.Header) > 0 {
		for k, v := range r.cnf.Header {
			request.Header.Set(k, v)
		}
	}
	if req != nil {
		if val, err = mapping.Marshal(req); err != nil {
			return nil, err
		}
		if pv, ok := val[pathKey]; ok {
			if err = fillPath(purl, pv); err != nil {
				return nil, err
			}
		}
		if fv := buildFormQuery(purl, val[formKey]); len(fv) > 0 {
			args := request.PostArgs()
			for k, v := range fv {
				args.Set(k, fmt.Sprint(v))
			}

		}
		if jv, ok := val[jsonKey]; ok {
			if method == http.MethodGet {
				return nil, ErrGetWithBody
			}
			if body, err = jsonx.Marshal(jv); nil != err {
				return nil, err
			}
			request.SetBodyRaw(body)
			request.Header.SetContentType(JsonContentType)
		}
		if hv, ok := val[headerKey]; ok {
			for k, v := range hv {
				request.Header.Add(k, fmt.Sprint(v))
			}
		}
	}
	request.SetRequestURI(purl.String())
	if err = r.client.Do(request, response); nil != err {
		return nil, err
	}
	body = response.Body()
	if nil != resp {
		if err = jsonx.Unmarshal(body, resp); nil != err {
			return nil, err
		}
	}
	return &RestResponse{
		StatusCode: response.StatusCode(),
	}, nil
}

func (r restFastHttp) DoRequest(req *http.Request) (*http.Response, error) {
	return nil, NotSupport
}

func NewRestFastHttp(name string, cnf conf.TransferConf, opts ...RestOption) RestService {
	//init
	dial := &fasthttp.TCPDialer{
		Concurrency:      4096,      // 最大并发数，0表示无限制
		DNSCacheDuration: time.Hour, // 将 DNS 缓存时间从默认分钟增加到一小时
	}
	r := &restFastHttp{
		cnf:  cnf.Fasthttp,
		name: name,
		client: &fasthttp.Client{
			Name:      name,
			Dial:      dial.Dial,
			TLSConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
