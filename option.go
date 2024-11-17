package xbinding

import "github.com/zhiyunliu/glue/context"

type Option func(o *Options)

type Options struct {
	method  string
	headers context.Header
}

func NewOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func WithHeaders(headers context.Header) Option {
	return func(o *Options) {
		o.headers = headers
	}
}

func WithMethod(method string) Option {
	return func(o *Options) {
		o.method = method
	}
}
