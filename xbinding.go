package xbinding

import (
	"fmt"
)

var ErrUnsupportedContentType = fmt.Errorf("xbinding: 不支持的Content-Type")

const DefaultProto = "binding"

// 解析器
type Resolver interface {
	Name() string
	Resolve(opts *Options) (Codec, error)
}

type Codec interface {
	ContentType() string
	Marshal(v interface{}) ([]byte, error)
	Bind(reader Reader, v interface{}) error
}

var resolvers = make(map[string]Resolver)

func Register(resolver Resolver) {
	proto := resolver.Name()
	if _, ok := resolvers[proto]; ok {
		panic(fmt.Errorf("xbinding: 不能重复注册:%s", proto))
	}
	resolvers[proto] = resolver
}

// GetCodec 根据适配器名称及参数返回配置处理器
func GetCodec(opts ...Option) (Codec, error) {
	botps := NewOptions(opts...)
	//默认的绑定适配器
	if botps.Proto == "" {
		botps.Proto = DefaultProto
	}
	resolver, ok := resolvers[botps.Proto]
	if !ok {
		return nil, fmt.Errorf("xbinding: 未知的协议类型:%s", botps.Proto)
	}
	return resolver.Resolve(botps)
}
