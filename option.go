package xbinding

type Option func(o *Options)

type Options struct {
	Proto       string
	Method      string
	ContextType string
}

func NewOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func WithContextType(contextType string) Option {
	return func(o *Options) {
		o.ContextType = contextType
	}
}

func WithMethod(method string) Option {
	return func(o *Options) {
		o.Method = method
	}
}

func WithProto(proto string) Option {
	return func(o *Options) {
		o.Proto = proto
	}
}
