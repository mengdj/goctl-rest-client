// Package rest
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
	"github.com/mengdj/goctl-rest-client/factory"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/breaker"
	"github.com/zeromicro/go-zero/core/mapping"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"net/http"
	nurl "net/url"
	"time"
)

type (
	restResty struct {
		client *resty.Client
		name   string
		trace  bool
	}
)

func (rds *restResty) Do(ctx context.Context, method, url string, req interface{}, resp interface{}) (*RestResponse, error) {
	var (
		rertyR    = rds.client.R().SetContext(ctx)
		rertyP    *resty.Response
		err, errx error
		purl      *nurl.URL
		val       map[string]map[string]interface{}
		span      oteltrace.Span
	)
	if purl, err = nurl.Parse(url); err != nil {
		return nil, err
	}
	// 0.0.9 新增context传递参数,key必须为string
	if ctx.Value(factory.EnableContextTransfer{}) != nil {
		if ctxkv := factory.GetKeyValueFromContext(ctx); len(ctxkv) > 0 {
			for k, v := range ctxkv {
				if name, ok := k.(string); ok {
					rertyR.SetHeader(factory.PrefixRestClientHeader+name, cast.ToString(v))
				}
			}
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
			rertyR.SetFormDataFromValues(fv)
		}
		if jv, ok := val[jsonKey]; ok {
			if method == http.MethodGet {
				return nil, ErrGetWithBody
			}
			rertyR.SetBody(jv)
			rertyR.SetHeader(ContentType, JsonContentType)
		}
		if hv, ok := val[headerKey]; ok {
			for k, v := range hv {
				rertyR.SetHeader(k, fmt.Sprint(v))
			}
		}
		if nil != resp {
			rertyR.SetResult(resp)
		}
	}
	if rds.trace {
		//链路跟踪
		ctx, span = otel.Tracer(traceName).Start(
			ctx,
			url,
			oteltrace.WithSpanKind(oteltrace.SpanKindClient),
		)
		defer span.End()
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(rertyR.Header))
	}
	//熔断
	if errx = breaker.GetBreaker(rds.name).DoWithAcceptable(func() error {
		if rertyP, err = rertyR.Execute(method, purl.String()); nil != err {
			return err
		}
		return nil
	}, func(err error) bool {
		if nil != err {
			return false
		}
		return rertyP.StatusCode() < http.StatusInternalServerError
	}); nil != errx {
		if rds.trace {
			span.RecordError(errx)
			span.SetStatus(codes.Error, errx.Error())
		}
		return nil, errx
	}
	if rds.trace {
		span.SetAttributes(semconv.HTTPAttributesFromHTTPStatusCode(rertyP.StatusCode())...)
		span.SetStatus(semconv.SpanStatusFromHTTPStatusCodeAndSpanKind(rertyP.StatusCode(), oteltrace.SpanKindClient))
	}
	return &RestResponse{
		StatusCode: rertyP.StatusCode(),
		Status:     rertyP.Status(),
	}, nil
}

func (rds *restResty) DoRequest(r *http.Request) (*http.Response, error) {
	return nil, NotSupport
}

func NewRestResty(name string, cnf conf.TransferConf, opts ...RestOption) RestService {
	r := &restResty{
		client: resty.New().SetDebug(cnf.Resty.Debug).SetAllowGetMethodPayload(cnf.Resty.AllowGetMethodPayload),
		name:   name,
	}
	//init
	if cnf.Resty.Token != "" {
		r.client.SetAuthToken(cnf.Resty.Token)
	}
	if cnf.Resty.Timeout != 0 {
		r.client.SetTimeout(time.Duration(cnf.Resty.Timeout))
	}
	if len(cnf.Resty.Header) > 0 {
		r.client.SetHeaders(cnf.Resty.Header)
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func WithRestyErrorHook(hook resty.ErrorHook) RestOption {
	return func(v interface{}) {
		if target, ok := v.(*restResty); ok {
			target.client.OnError(hook)
		}
	}
}

func WithRestyBeforeRequest(req resty.RequestMiddleware) RestOption {
	return func(v interface{}) {
		if target, ok := v.(*restResty); ok {
			target.client.OnBeforeRequest(req)
		}
	}
}

func WithDisableWarn(dis bool) RestOption {
	return func(v interface{}) {
		if target, ok := v.(*restResty); ok {
			target.client.SetDisableWarn(dis)
		}
	}
}

func WithTrace(trace bool) RestOption {
	return func(v interface{}) {
		if target, ok := v.(*restResty); ok {
			target.trace = trace
		}
	}
}

func WithRestyAfterResponse(resp resty.ResponseMiddleware) RestOption {
	return func(v interface{}) {
		if target, ok := v.(*restResty); ok {
			target.client.OnAfterResponse(resp)
		}
	}
}
