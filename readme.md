# xinding 说明
## 介绍
xbinding是基于gin框架的binding模块的扩展，主要从项目gin中分离处理binding的部分内容，进行改造处理适配除http以外的其他数据绑定过程。并且兼容http.request的绑定方式。

## 安装

Requires Go 1.21 or above.
Install with `go install github.com/zhiyunliu/xbinding@latest`.
   
## 示例 

```go

    type Config struct {
        SpanKind   trace.SpanKind `json:"span_kind" yaml:"span_kind"`
        Provider   string         `json:"provider" yaml:"provider"`
        Propagator string         `json:"propagator" yaml:"propagator"`
    }

    traceCfg := &Config{}

    dataType := "json" //yaml , application/json, application/x-yaml
	codec, err := xbinding.GetCodec(xbinding.WithContentType(dataType))
	if err != nil {
		return nil, fmt.Errorf("xbinding.GetCodec err:%w", err)
	}

	if err = codec.Bind(xbinding.BytesReader(data.Data), traceCfg); err != nil {
		return nil, err
	}

```

## 扩展接口--帮助快速类型转换

辅助数据类型定义
```go
type BytesReader []byte
type MapReader map[string][]string
type SMapReader map[string]string
type XMapReader map[string]any
type ReaderWrapper struct {
	Data any
}

```

扩展实现,实现以下两个接口可自定义实现Codec的注入

```go
type Resolver interface {
	Name() string
	Resolve(opts *Options) (Codec, error)
}

type Codec interface {
	ContentType() string
	Marshal(v interface{}) ([]byte, error)
	Bind(reader Reader, v interface{}) error
}

```

注入到引擎中

```go

var custom struct{}

func(c custom)	Name() string{return "custom"}
func(c custom) Resolve(opts *Options) (Codec, error){
    return nil,fmt.Errorf("not implement")
}

xbinding.Register(&custom{})

//获取自定义的codec json解析处理
xbinding.GetCodec(xbinding.WithProto("custom"),xbinding.WithContentType("json"))


```


## 默认类型支持

 | 类型                              | 描述 | 是否支持 |
 | --------------------------------- | ---- | -------- |
 | json                              | json | 是       |
 | yaml                              | yaml | 是       |
 | xml                               | xml  | 是       |
 | text                              | text | 是       |
 | application/json                  |      |          |
 | text/html                         |      |          |
 | application/xml                   |      |          |
 | text/xml                          |      |          |
 | text/plain                        |      |          |
 | application/x-www-form-urlencoded |      |          |
 | multipart/form-data               |      |          |
 | application/x-protobuf            |      |          |
 | application/x-msgpack             |      |          |
 | application/msgpack               |      |          |
 | application/x-yaml                |      |          |

## 注意事项




## 版本变更记录

* v0.1.0 2024-11-19 初始化
* v0.1.1 重写绑定reader实现
* v0.1.2 优化绑定，处理go版本依赖问题
