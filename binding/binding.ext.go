package binding

import "github.com/zhiyunliu/xbinding"

var (
	Plain = textCodec{}
	TOML  = tomlBinding{}
)

const (
	MIMETOML = "application/toml"
)

// Binding describes the interface which needs to be implemented for binding the
// data present in the request such as JSON request body, query parameters or
// the form POST.
type Binding interface {
	Name() string
	Bind(xbinding.Reader, interface{}) error
	ContentType() string
	Marshal(v interface{}) ([]byte, error)
}
