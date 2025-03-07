package binding

import (
	"net/textproto"
	"reflect"

	"github.com/zhiyunliu/xbinding"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(reader xbinding.Reader, obj interface{}) error {
	dataObj, err := reader.ReadObject()
	if err != nil {
		return err
	}

	sourceData, err := transferMapArrayData(dataObj)
	if err != nil {
		return err
	}

	if err := mapHeader(obj, sourceData); err != nil {
		return err
	}

	return validate(obj)
}

func mapHeader(ptr interface{}, h map[string][]string) error {
	return mappingByPtr(ptr, headerSource(h), "header")
}

type headerSource map[string][]string

var _ setter = headerSource(nil)

func (hs headerSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt setOptions) (bool, error) {
	return setByForm(value, field, hs, textproto.CanonicalMIMEHeaderKey(tagValue), opt)
}
