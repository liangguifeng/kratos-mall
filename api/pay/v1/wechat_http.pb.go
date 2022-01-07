// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type WechatHTTPServer interface {
	H5(context.Context, *H5Request) (*H5Response, error)
}

func RegisterWechatHTTPServer(s *http.Server, srv WechatHTTPServer) {
	r := s.Route("/")
	r.GET("/pay/v1/wechat/h5", _Wechat_H50_HTTP_Handler(srv))
}

func _Wechat_H50_HTTP_Handler(srv WechatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in H5Request
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.pay.v1.Wechat/H5")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.H5(ctx, req.(*H5Request))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*H5Response)
		return ctx.Result(200, reply)
	}
}

type WechatHTTPClient interface {
	H5(ctx context.Context, req *H5Request, opts ...http.CallOption) (rsp *H5Response, err error)
}

type WechatHTTPClientImpl struct {
	cc *http.Client
}

func NewWechatHTTPClient(client *http.Client) WechatHTTPClient {
	return &WechatHTTPClientImpl{client}
}

func (c *WechatHTTPClientImpl) H5(ctx context.Context, in *H5Request, opts ...http.CallOption) (*H5Response, error) {
	var out H5Response
	pattern := "/pay/v1/wechat/h5"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.pay.v1.Wechat/H5"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}