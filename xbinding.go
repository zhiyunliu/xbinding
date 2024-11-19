package xbinding

import (
	"fmt"
)

// 解析器
type Resolver interface {
	Name() string
	Resolve(opts *Options) (Codec, error)
}

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(reader Reader, v interface{}) error
}

var resolvers = make(map[string]Resolver)

func Register(resolver Resolver) {
	proto := resolver.Name()
	if _, ok := resolvers[proto]; ok {
		panic(fmt.Errorf("xbinding: 不能重复注册:%s", proto))
	}
	resolvers[proto] = resolver
}

// NewBinding 根据适配器名称及参数返回配置处理器
func New(opts ...Option) (Codec, error) {
	botps := NewOptions(opts...)
	//默认的绑定适配器
	proto := "binding"
	if botps.Proto != "" {
		proto = botps.Proto
	}
	resolver, ok := resolvers[proto]
	if !ok {
		return nil, fmt.Errorf("xbinding: 未知的协议类型:%s", proto)
	}
	return resolver.Resolve(botps)
}
