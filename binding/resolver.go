package binding

import (
	"fmt"

	"github.com/zhiyunliu/xbinding"
)

var (
	bindingMap map[string]Binding
)

type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
}

type bindingResolver struct {
	name string
}

func (r *bindingResolver) Name() string {
	return r.name
}
func (r *bindingResolver) Resolve(opts *xbinding.Options) (xbinding.Codec, error) {
	return &bindingCodecWrap{
		binding: getBinding(opts),
	}, nil
}

type bindingCodecWrap struct {
	binding Binding
}

func (w *bindingCodecWrap) Marshal(v interface{}) ([]byte, error) {
	marshaler, ok := w.binding.(Marshaler)
	if !ok {
		return nil, fmt.Errorf("not implemented Marshaler,type[%T]", w.binding)
	}
	return marshaler.Marshal(v)
}

func (w *bindingCodecWrap) Unmarshal(reader xbinding.Reader, v interface{}) error {
	return w.binding.Bind(reader, v)
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
	if bindingObj, ok := bindingMap[opts.ContextType]; ok {
		return bindingObj
	}

	bindObj := Default(opts.Method, opts.ContextType)
	return bindObj
}
