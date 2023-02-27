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
	name    string
	client  *fasthttp.Client
	cnf     conf.FastHttpConf
	retryIf fasthttp.RetryIfFunc
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
		//release
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
		//parse result
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

func WithRetryIf(fn fasthttp.RetryIfFunc) RestOption {
	return func(v interface{}) {
		if vv, ok := v.(*restFastHttp); ok {
			vv.retryIf = fn
		}
	}
}
func NewRestFastHttp(name string, cnf conf.TransferConf, opts ...RestOption) RestService {
	//init
	dial := &fasthttp.TCPDialer{
		Concurrency:      cnf.Fasthttp.TCPDialer.Concurrency,                                   //
		DNSCacheDuration: time.Duration(cnf.Fasthttp.TCPDialer.DNSCacheDuration) * time.Second, //
	}
	r := &restFastHttp{
		cnf:  cnf.Fasthttp,
		name: name,
	}
	r.client = &fasthttp.Client{
		Name:                name,
		Dial:                dial.Dial,
		ReadTimeout:         time.Duration(cnf.Fasthttp.ReadTimeout) * time.Second,
		MaxConnWaitTimeout:  time.Duration(cnf.Fasthttp.MaxConnWaitTimeout) * time.Second,
		WriteTimeout:        time.Duration(cnf.Fasthttp.WriteTimeout) * time.Second,
		MaxConnDuration:     time.Duration(cnf.Fasthttp.MaxConnDuration) * time.Second,
		MaxIdleConnDuration: time.Duration(cnf.Fasthttp.MaxIdleConnDuration) * time.Second,
		TLSConfig:           &tls.Config{InsecureSkipVerify: true},
		RetryIf: func(request *fasthttp.Request) bool {
			if nil != r.retryIf {
				return r.retryIf(request)
			}
			//default
			return false
		},
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
