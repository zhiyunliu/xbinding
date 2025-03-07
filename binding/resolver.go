package binding

import (
	"github.com/zhiyunliu/xbinding"
)

var (
	bindingMap map[string]Binding
)

type bindingResolver struct {
	name string
}

func (r *bindingResolver) Name() string {
	return r.name
}
func (r *bindingResolver) Resolve(opts *xbinding.Options) (codec xbinding.Codec, err error) {
	codec = getBinding(opts)
	if codec == nil {
		return nil, xbinding.ErrUnsupportedContentType
	}
	return codec, nil
}

func init() {
	bindingMap = make(map[string]Binding)
	xbinding.Register(&bindingResolver{name: "binding"})

	bindingMap[Plain.Name()] = Plain
	bindingMap[JSON.Name()] = JSON
	bindingMap[XML.Name()] = XML
	bindingMap[Form.Name()] = Form
	bindingMap[Query.Name()] = Query
	bindingMap[FormPost.Name()] = FormPost
	bindingMap[FormMultipart.Name()] = FormMultipart
	bindingMap[ProtoBuf.Name()] = ProtoBuf
	bindingMap[YAML.Name()] = YAML
	bindingMap[Uri.Name()] = Uri
	bindingMap[Header.Name()] = Header
	bindingMap[TOML.Name()] = TOML
}

func getBinding(opts *xbinding.Options) Binding {
	if bindingObj, ok := bindingMap[opts.ContentType]; ok {
		return bindingObj
	}

	return Default(opts.Method, opts.ContentType)
}
